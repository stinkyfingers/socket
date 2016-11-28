package handlers

import (
	"errors"

	"golang.org/x/net/websocket"
)

type ChatHub struct {
	ClientMap  map[string][]ChatClient
	Broadcast  chan *Message
	Register   chan *ChatClient
	Unregister chan *ChatClient
}

type ChatClient struct {
	GameID  string
	send    chan *Message
	Conn    websocket.Conn
	ChatHub *ChatHub
}

type Message struct {
	ID         string `json:"id"`
	Text       string `json:"text"`
	PlayerName string `json:"playerName"`
}

func NewChatHub() *ChatHub {
	return &ChatHub{
		Broadcast:  make(chan *Message),
		Register:   make(chan *ChatClient),
		Unregister: make(chan *ChatClient),
		ClientMap:  make(map[string][]ChatClient),
	}
}

func (h *ChatHub) Run() {
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
			for _, client := range h.ClientMap[message.ID] {
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

func ChatHandler(ws *websocket.Conn, h *ChatHub) {
	id := ws.Request().URL.Query().Get("id")
	if id == "" {
		return
	}

	cli := ChatClient{Conn: *ws, GameID: id, ChatHub: h, send: make(chan *Message)}
	h.Register <- &cli
	go cli.writeMessage()
	err := cli.readMessage()
	if err != nil {
		websocket.JSON.Send(ws, HttpError{Error: err, Message: "Error in Websocket read loop", Status: 500})
		return
	}
}

func (wsc *ChatClient) writeMessage() {
	for {
		select {
		case msg, ok := <-wsc.send:
			if !ok {
				websocket.JSON.Send(&wsc.Conn, HttpError{Error: errors.New("channel is closed"), Message: "Error sending on Websocket connection"})
				// return
				continue
			}
			err := websocket.JSON.Send(&wsc.Conn, msg)
			if err != nil {
				// remove conn
				wsc.ChatHub.Unregister <- wsc

			}
		}

	}
}

func (wsc *ChatClient) readMessage() error {
	defer func() {
		wsc.ChatHub.Unregister <- wsc
		wsc.Conn.Close()
	}()
	for {
		var msg Message
		err := websocket.JSON.Receive(&wsc.Conn, &msg)
		if err != nil {
			return err
		}
		wsc.ChatHub.Broadcast <- &msg
	}
}
