package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/stinkyfingers/socket/server/db"
	"github.com/stinkyfingers/socket/server/game"
)

func HandleDefault(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}

func sendJson(w http.ResponseWriter, i interface{}) {
	j, err := json.Marshal(i)
	if err != nil {
		HttpError{err, "json error", 500, w}.HandleErr()
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(j)
	return
}

func HandleTestSetup(w http.ResponseWriter, r *http.Request) {
	db.NewSession()
	Cleanup()
	//decks
	deck, dealerDeck, err := Setup("sample.json")
	if err != nil {
		HttpError{err, "json error", 500, w}.HandleErr()
		return
	}
	dd := game.DealerDeck{
		Cards: dealerDeck,
	}
	d := game.Deck{
		Cards: deck,
	}
	err = dd.Create()
	if err != nil {
		HttpError{err, "error creating dealer deck", 500, w}.HandleErr()
		return
	}
	err = d.Create()
	if err != nil {
		HttpError{err, "error creating deck", 500, w}.HandleErr()
		return
	}

	//players
	players, err := SetupPlayers()
	if err != nil {
		if err.Error() != "username already exists" {
			HttpError{err, "json error", 500, w}.HandleErr()
			return
		}
	}
	// game
	g := game.Game{
		Deck:       deck,
		DealerDeck: dealerDeck,
	}
	err = g.Create()
	if err != nil {
		HttpError{err, "json error", 500, w}.HandleErr()
		return
	}
	for _, player := range players {
		err = g.AddPlayer(player)
		if err != nil {
			HttpError{err, "json error", 500, w}.HandleErr()
			return
		}
	}

	err = g.Initialize()
	if err != nil {
		HttpError{err, "json error", 500, w}.HandleErr()
		return
	}

	sendJson(w, g)
}
