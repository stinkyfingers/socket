package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/stinkyfingers/socket/server/game"
	"gopkg.in/mgo.v2/bson"
)

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

func HandleUpdateGame(w http.ResponseWriter, r *http.Request) {
	var g game.Game
	err := json.NewDecoder(r.Body).Decode(&g)
	if err != nil {
		HttpError{err, "error decoding game", 500, w}.HandleErr()
		return
	}
	err = g.Update()
	if err != nil {
		HttpError{err, "error updating game", 500, w}.HandleErr()
		return
	}
	sendJson(w, g)
}
