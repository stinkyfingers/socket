package run_handlers

import (
	"log"

	"github.com/stinkyfingers/socket/server/game"
	"golang.org/x/net/websocket"
	"gopkg.in/mgo.v2/bson"
)

// Works, but is the simpler, cruder version of the ws handler
// Use game_ws_handler

type Hub struct {
	Clients map[string][]Client   // game id: client
	Games   map[string]*game.Game // game id: game
}

type Client struct {
	WS     *websocket.Conn
	GameID string
}

func Hubify() *Hub {
	return &Hub{
		Clients: make(map[string][]Client),
		Games:   make(map[string]*game.Game),
	}
}

func Handle(ws *websocket.Conn, h *Hub) {
	id := ws.Request().URL.Query().Get("id")
	cli := Client{
		WS:     ws,
		GameID: id,
	}
	h.Clients[id] = append(h.Clients[id], cli)

	g := &game.Game{ID: bson.ObjectIdHex(id)}
	err := g.Get()
	if err != nil {
		log.Print("ERR GETTING GAME ", err)
		return
	}
	h.Games[id] = g

	err = websocket.JSON.Send(ws, h.Games[id])
	if err != nil {
		log.Print("MOUNT ERR", err)
		return
	}

	for {
		// receive
		var play game.Play
		err = websocket.JSON.Receive(ws, &play)
		if err != nil {
			log.Print("REC ERR", err)
			break
		}

		// process
		h.Games[id] = process(*h.Games[id], play)

		// send
		log.Print("NUM CLIENTS ", len(h.Clients[id]))

		for i, c := range h.Clients[id] {

			var g game.Game
			g = *h.Games[id]

			err = websocket.JSON.Send(c.WS, g)
			if err != nil {
				h.Clients[id] = append(h.Clients[id][:i], h.Clients[id][i+1:]...)
				log.Print("SEND ERR ", err, " REMOVED CLIENT: ", h.Clients[id][i])
				continue
			}
		}
	}
}

func process(g game.Game, p game.Play) *game.Game {
	switch p.PlayType {
	case "play":
		g.Round.Plays[p.Player.ID.Hex()] = p
	case "vote":
		g.Round.Votes[p.Player.ID.Hex()] = p
	default:
		log.Print("type not supported")
	}
	log.Print("P ", p.Card.Phrase)

	if p.PlayType == "play" {
		g.Round.Plays[p.Player.ID.Hex()] = p
	}
	if p.PlayType == "vote" {
		g.Round.Votes[p.Player.ID.Hex()] = p
	}

	if len(g.Round.Plays) >= len(g.Players) {
		err := g.UpdatePlays()
		if err != nil {
			log.Print("ERR 1 ", err)
			return &g
		}

	}
	if len(g.Round.Votes) >= len(g.Players) {
		err := g.UpdateVotes()
		if err != nil {
			log.Print("ERR 2 ", err)
			return &g
		}
	}
	return &g
}

// ORIGINAL TEST CODE------------->

// type Client struct {
// 	ws *websocket.Conn
// 	IP string
// }

// var clients map[string][]Client
// var games map[string]game.Game

// func handler(ws *websocket.Conn) {
// 	id := ws.Request().URL.Query().Get("id")

// 	var g game.Game
// 	g.ID = bson.ObjectIdHex(id)
// 	err := g.Get()
// 	if err != nil {
// 		log.Print(err)
// 		return
// 	}
// 	games[g.ID.Hex()] = g

// 	// client return
// 	client := Client{
// 		ws,
// 		ws.Request().RemoteAddr,
// 	}
// 	clients[id] = append(clients[id], client)

// 	for {

// 		for _, c := range clients[id] {
// 			err = websocket.JSON.Send(c.ws, games[id])
// 			if err != nil {
// 				log.Print("WS client connection error: ", err)
// 				// break
// 			}
// 		}

// 		var p game.Play
// 		err = websocket.JSON.Receive(ws, &p)
// 		if err != nil {
// 			log.Print(err)
// 			break //TODO - handle errors in WS
// 		}

// 		switch p.PlayType {
// 		case "play":
// 			games[id].Round.Plays[p.Player.ID.Hex()] = p // TODO - switch to id.hex
// 		case "vote":
// 			games[id].Round.Votes[p.Player.ID.Hex()] = p
// 		default:
// 			log.Print("type not supported")
// 		}

// 		ch := make(chan game.Game)
// 		go func() {
// 			if p.PlayType == "play" && len(games[id].Round.Plays) > 1 { //TODO -len equal to players
// 				ga := games[id]
// 				(&ga).UpdatePlays()
// 				ch <- ga
// 			} else if p.PlayType == "vote" && len(games[id].Round.Votes) > 1 {
// 				ga := games[id]
// 				(&ga).UpdateVotes()
// 				ch <- ga
// 			} else {
// 				ch <- games[id]
// 			}
// 		}()

// 		ga := <-ch
// 		games[id] = ga
// 	}
// }
