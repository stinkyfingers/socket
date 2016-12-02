package handlers

import (
	"errors"

	"github.com/stinkyfingers/socket/server/game"
	"golang.org/x/net/websocket"
	"gopkg.in/mgo.v2/bson"
)

type Hub struct {
	ClientMap  map[string][]WSClient
	Broadcast  chan *game.Game
	Register   chan *WSClient
	Unregister chan *WSClient
	Games      map[string]*game.Game
}

type WSClient struct {
	GameID string
	send   chan *game.Game
	Conn   websocket.Conn
	Hub    *Hub
}

func NewHub() *Hub {
	return &Hub{
		Broadcast:  make(chan *game.Game),
		Register:   make(chan *WSClient),
		Unregister: make(chan *WSClient),
		ClientMap:  make(map[string][]WSClient),
		Games:      make(map[string]*game.Game),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.ClientMap[client.GameID] = append(h.ClientMap[client.GameID], *client)

		case client := <-h.Unregister:
			if _, ok := h.ClientMap[client.GameID]; ok {
				cm := h.ClientMap[client.GameID]
				for i := range h.ClientMap[client.GameID] {
					if &h.ClientMap[client.GameID][i] == client {
						cm = append(h.ClientMap[client.GameID][:i], h.ClientMap[client.GameID][i+1:]...)
					}
				}
				h.ClientMap[client.GameID] = cm
			}

		case message := <-h.Broadcast:
			for _, client := range h.ClientMap[message.ID.Hex()] {
				select {
				case client.send <- message:
				default:
					cm := h.ClientMap[client.GameID]
					for i := range h.ClientMap[client.GameID] {
						if h.ClientMap[client.GameID][i] == client {
							cm = append(h.ClientMap[client.GameID][:i], h.ClientMap[client.GameID][i+1:]...)
						}
					}
					h.ClientMap[client.GameID] = cm
				}
			}
		}
	}
}

func ServeWS(ws *websocket.Conn, h *Hub) {
	id := ws.Request().URL.Query().Get("id")
	if id == "" {
		return
	}

	var g game.Game
	g.ID = bson.ObjectIdHex(id)
	err := g.Get()
	if err != nil {
		websocket.JSON.Send(ws, HttpError{Error: err, Message: "Error getting game", Status: 500})
		return
	}

	if h.Games[g.ID.Hex()] == nil || !h.Games[g.ID.Hex()].Initialized {
		h.Games[g.ID.Hex()] = &g
	}

	cli := WSClient{Conn: *ws, GameID: id, Hub: h, send: make(chan *game.Game)}
	h.Register <- &cli
	go cli.write()
	cli.Hub.Broadcast <- h.Games[g.ID.Hex()]
	err = cli.read()
	if err != nil {
		websocket.JSON.Send(ws, HttpError{Error: err, Message: "Error in Websocket read loop", Status: 500})
		return
	}
}

func (wsc *WSClient) write() {
	for {
		select {
		case g, ok := <-wsc.send:
			if !ok {
				websocket.JSON.Send(&wsc.Conn, HttpError{Error: errors.New("channel is closed"), Message: "Error sending on Websocket connection"})
				// return
				continue
			}
			err := websocket.JSON.Send(&wsc.Conn, g)
			if err != nil {
				// remove conn
				wsc.Hub.Unregister <- wsc

			}
		}

	}
}

func (wsc *WSClient) read() error {
	defer func() {
		wsc.Hub.Unregister <- wsc
		wsc.Conn.Close()
	}()
	for {
		var p game.Play
		err := websocket.JSON.Receive(&wsc.Conn, &p)
		if err != nil {
			return err
		}

		switch p.PlayType {
		case "play":
			wsc.Hub.Games[wsc.GameID].Round.Plays[p.Player.ID.Hex()] = p
		case "vote":
			wsc.Hub.Games[wsc.GameID].Round.Votes[p.Player.ID.Hex()] = p
		default:
			return errors.New("Play type not supported")
		}

		ch := make(chan *game.Game)

		go func() {
			if p.PlayType == "play" && len(wsc.Hub.Games[wsc.GameID].Round.Plays) >= len(wsc.Hub.Games[wsc.GameID].Players) { //TODO -len equal to players
				ga := wsc.Hub.Games[wsc.GameID]
				ga.UpdatePlays()
				ch <- ga
			} else if p.PlayType == "vote" && len(wsc.Hub.Games[wsc.GameID].Round.Votes) >= len(wsc.Hub.Games[wsc.GameID].Players) {
				ga := wsc.Hub.Games[wsc.GameID]
				ga.UpdateVotes()
				ch <- ga
			} else {
				ch <- wsc.Hub.Games[wsc.GameID]
			}
		}()
		ga := <-ch
		wsc.Hub.Broadcast <- ga
	}
}
