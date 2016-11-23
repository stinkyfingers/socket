package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/stinkyfingers/socket/server/game"
	"golang.org/x/net/websocket"
	"gopkg.in/mgo.v2/bson"
)

type Client struct {
	ws *websocket.Conn
	IP string
}

var Clients map[string][]Client
var Games map[string]game.Game

func Game(ws *websocket.Conn) {
	id := ws.Request().URL.Query().Get("id")
	if id == "" {
		return
	}

	var g game.Game
	g.ID = bson.ObjectIdHex(id)
	err := g.Get()
	if err != nil {
		ws.Write([]byte("game not found"))
		return
	}

	Games[g.ID.Hex()] = g

	// client return
	client := Client{
		ws,
		ws.Request().RemoteAddr,
	}
	if clients, ok := Clients[id]; ok {
		var found bool
		for _, cli := range clients {
			if cli == client {
				found = true
			}
		}
		if !found {
			Clients[id] = append(Clients[id], client)
		}
	} else {
		Clients[id] = append(Clients[id], client)
	}

	for {

		for i, c := range Clients[id] {
			err = websocket.JSON.Send(c.ws, Games[id])
			if err != nil {
				log.Print("WS client connection error: ", err)
				Clients[id] = append(Clients[id][:i], Clients[id][i+1:]...)
			}
		}

		ch := make(chan game.Game)
		go func() {
			var p game.Play
			err = websocket.JSON.Receive(ws, &p)
			if err != nil {
				return // break w/o returning game
			}

			switch p.PlayType {
			case "play":
				Games[id].Round.Plays[p.Player.ID.Hex()] = p // TODO - switch to id.hex
			case "vote":
				Games[id].Round.Votes[p.Player.ID.Hex()] = p
			default:
				log.Print("type not supported")
			}

			if p.PlayType == "play" && len(Games[id].Round.Plays) == len(Games[id].Players) {
				ga := Games[id]
				(&ga).UpdatePlays()
				ch <- ga
			} else if p.PlayType == "vote" && len(Games[id].Round.Votes) == len(Games[id].Players) {
				ga := Games[id]
				(&ga).UpdateVotes()
				ch <- ga
			} else {
				ch <- Games[id]
			}
		}()

		ga := <-ch
		Games[id] = ga
	}
}

func HandleNewGame(w http.ResponseWriter, r *http.Request) {
	var player game.Player
	err := json.NewDecoder(r.Body).Decode(&player)
	if err != nil {
		HttpError{err, "Player required to start game", 500, w}.HandleErr()
		return
	}
	deck, err := game.GetDeck()
	if err != nil {
		HttpError{err, "error getting deck", 500, w}.HandleErr()
		return
	}
	dealerDeck, err := game.GetDealerDeck()
	if err != nil {
		HttpError{err, "error getting dealer deck", 500, w}.HandleErr()
		return
	}
	g := game.Game{
		Deck:       deck.Cards,
		DealerDeck: dealerDeck.Cards,
		Players:    []game.Player{player},
	}
	err = g.Create()
	if err != nil {
		HttpError{err, "error creating game", 500, w}.HandleErr()
		return
	}
	sendJson(w, g)
}

func HandleAddPlayer(w http.ResponseWriter, r *http.Request) {
	var player game.Player
	err := json.NewDecoder(r.Body).Decode(&player)
	if err != nil {
		HttpError{err, "Player required to join game", 500, w}.HandleErr()
		return
	}
	id := r.URL.Query().Get("id")
	if !bson.IsObjectIdHex(id) {
		HttpError{err, "error parsing game id", 500, w}.HandleErr()
		return
	}
	g := game.Game{
		ID: bson.ObjectIdHex(id),
	}
	err = g.AddPlayer(player)
	if err != nil {
		HttpError{err, "error adding player", 500, w}.HandleErr()
		return
	}
	sendJson(w, g)
}

func HandleStartGame(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if !bson.IsObjectIdHex(id) {
		HttpError{nil, "error parsing game id", 500, w}.HandleErr()
		return
	}
	g := game.Game{
		ID: bson.ObjectIdHex(id),
	}
	err := g.Get()
	if err != nil {
		HttpError{err, "error retrieving game", 500, w}.HandleErr()
		return
	}
	err = g.Initialize()
	if err != nil {
		HttpError{err, "error initializing game", 500, w}.HandleErr()
		return
	}
	sendJson(w, g)
}

func HandleGetGame(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if !bson.IsObjectIdHex(id) {
		HttpError{nil, "error parsing game id", 500, w}.HandleErr()
		return
	}
	g := game.Game{
		ID: bson.ObjectIdHex(id),
	}
	err := g.Get()
	if err != nil {
		HttpError{err, "error retrieving game", 500, w}.HandleErr()
		return
	}
	sendJson(w, g)
}

func HandleExitGame(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if !bson.IsObjectIdHex(id) {
		HttpError{nil, "error parsing game id", 500, w}.HandleErr()
		return
	}
	g := game.Game{
		ID: bson.ObjectIdHex(id),
	}
	err := g.Get()
	if err != nil {
		HttpError{err, "error retrieving game", 500, w}.HandleErr()
		return
	}
	g.Initialized = false
	err = g.Update()
	if err != nil {
		HttpError{err, "error updating game", 500, w}.HandleErr()
		return
	}
	sendJson(w, g)
}
