package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/stinkyfingers/socket/server/game"
	"gopkg.in/mgo.v2/bson"
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

func HandleCreatePlayer(w http.ResponseWriter, r *http.Request) {
	var p game.Player
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		HttpError{err, "error decoding player", 500, w}.HandleErr()
		return
	}
	err = p.Create()
	if err != nil {
		HttpError{err, "error creating player", 500, w}.HandleErr()
		return
	}
	sendJson(w, p)
}

func HandleUpdatePlayer(w http.ResponseWriter, r *http.Request) {
	var p game.Player
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		HttpError{err, "error decoding player", 500, w}.HandleErr()
		return
	}
	err = p.Update()
	if err != nil {
		HttpError{err, "error updating player", 500, w}.HandleErr()
		return
	}
	sendJson(w, p)
}

func HandleGetPlayer(w http.ResponseWriter, r *http.Request) {
	var p game.Player
	id := r.URL.Query().Get("id")
	if !bson.IsObjectIdHex(id) {
		HttpError{nil, "id is not valid", 500, w}.HandleErr()
		return
	}
	p.ID = bson.ObjectIdHex(id)
	err := p.Get()
	if err != nil {
		HttpError{err, "error getting player", 500, w}.HandleErr()
		return
	}
	sendJson(w, p)
}

func HandleResetPassword(w http.ResponseWriter, r *http.Request) {
	add := r.URL.Query().Get("email")
	p := game.Player{
		Email: add,
	}
	err := p.FindByEmail()
	if err != nil {
		HttpError{err, "email not found", 500, w}.HandleErr()
		return
	}
	err = p.ResetPassword()
	if err != nil {
		HttpError{err, "error resetting password", 500, w}.HandleErr()
		return
	}
	err = p.PasswordEmail()
	if err != nil {
		HttpError{err, "error emailing password", 500, w}.HandleErr()
		return
	}
	p.Password = ""
	sendJson(w, p)
}
