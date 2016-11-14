package main

import (
	"log"
	"net/http"

	"github.com/stinkyfingers/socket/server/db"
	"github.com/stinkyfingers/socket/server/game"
	"github.com/stinkyfingers/socket/server/handlers"
	"golang.org/x/net/websocket"
	"gopkg.in/mgo.v2/bson"
)

// Server needs to:
// 1. Send each client Lead Cards and their Player Hand
// 2. Receive each client's play
// 3. Send each client Played Cards
// 4. Receive each client's vote

func main() {
	err := db.NewSession()
	if err != nil {
		log.Fatal(err)
	}

	clients = make(map[string][]Client)
	games = make(map[string]game.Game)

	http.Handle("/", websocket.Handler(handler))
	http.HandleFunc("/test", handlers.HandleTestSetup)
	log.Fatal(http.ListenAndServe(":7000", nil))
}

type Client struct {
	ws *websocket.Conn
	IP string
}

var clients map[string][]Client
var games map[string]game.Game

func handler(ws *websocket.Conn) {
	id := ws.Request().URL.Query().Get("id")

	var g game.Game
	g.ID = bson.ObjectIdHex(id)
	err := g.Get()
	if err != nil {
		log.Print(err)
		return
	}
	games[g.ID.Hex()] = g

	// client return
	client := Client{
		ws,
		ws.Request().RemoteAddr,
	}
	clients[id] = append(clients[id], client)

	for {

		for _, c := range clients[id] {
			err = websocket.JSON.Send(c.ws, games[id])
			if err != nil {
				log.Print("WS client connection error: ", err)
				// break
			}
		}

		var p game.Play
		err = websocket.JSON.Receive(ws, &p)
		if err != nil {
			log.Print(err)
			break //TODO - handle errors in WS
		}

		switch p.PlayType {
		case "play":
			games[id].Round.Plays[p.Player.ID.Hex()] = p // TODO - switch to id.hex
		case "vote":
			games[id].Round.Votes[p.Player.ID.Hex()] = p
		default:
			log.Print("type not supported")
		}

		ch := make(chan game.Game)
		go func() {
			if p.PlayType == "play" && len(games[id].Round.Plays) > 1 { //TODO -len equal to players
				ga := games[id]
				(&ga).UpdatePlays()
				ch <- ga
			} else if p.PlayType == "vote" && len(games[id].Round.Votes) > 1 {
				ga := games[id]
				(&ga).UpdateVotes()
				ch <- ga
			} else {
				ch <- games[id]
			}
		}()

		ga := <-ch
		games[id] = ga
	}
}
