package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/stinkyfingers/socket/server/game"
)

func HandleAuthenticate(w http.ResponseWriter, r *http.Request) {
	var p game.Player
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		HttpError{err, "error decoding player", 500, w}.HandleErr()
		return
	}
	err = p.Authenticate()
	if err != nil {
		HttpError{err, "error authenticating player", 500, w}.HandleErr()
		return
	}
	sendJson(w, p)
}
