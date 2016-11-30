package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/stinkyfingers/socket/server/game"
	"gopkg.in/mgo.v2/bson"
)

func HandleCreateCard(w http.ResponseWriter, r *http.Request) {
	var card game.Card
	err := json.NewDecoder(r.Body).Decode(&card)
	if err != nil {
		HttpError{err, "", 500, w}.HandleErr()
		return
	}
	err = card.Audit()
	if err != nil {
		HttpError{err, "Card audit error", 500, w}.HandleErr()
		return
	}
	err = card.Create()
	if err != nil {
		HttpError{err, "Error saving card", 500, w}.HandleErr()
		return
	}
	sendJson(w, card)
}

func HandleCreateDealerCard(w http.ResponseWriter, r *http.Request) {
	var card game.DealerCard
	err := json.NewDecoder(r.Body).Decode(&card)
	if err != nil {
		HttpError{err, "", 500, w}.HandleErr()
		return
	}
	err = card.Audit()
	if err != nil {
		HttpError{err, "Card audit error", 500, w}.HandleErr()
		return
	}
	err = card.Create()
	if err != nil {
		HttpError{err, "Error saving card", 500, w}.HandleErr()
		return
	}
	sendJson(w, card)
}

func HandleImportCards(w http.ResponseWriter, r *http.Request) {
	f, _, err := r.FormFile("file")
	if err != nil {
		HttpError{err, "Error uploading file", 500, w}.HandleErr()
		return
	}
	defer f.Close()

	idStr := r.URL.Query().Get("id")
	if !bson.IsObjectIdHex(idStr) {
		HttpError{nil, "No id provided", 500, w}.HandleErr()
		return
	}
	id := bson.ObjectIdHex(idStr)

	var cards []game.Card
	err = json.NewDecoder(f).Decode(&cards)
	if err != nil {
		HttpError{err, "Error decoding file", 500, w}.HandleErr()
		return
	}
	for _, card := range cards {
		card.CreatedBy = id
		err = card.Create()
		if err != nil {
			HttpError{err, "Error saving card: " + card.Phrase, 500, w}.HandleErr()
			return
		}
	}
	sendJson(w, cards)
}

func HandleImportDealerCards(w http.ResponseWriter, r *http.Request) {
	f, _, err := r.FormFile("file")
	if err != nil {
		HttpError{err, "Error uploading file", 500, w}.HandleErr()
		return
	}
	defer f.Close()

	idStr := r.URL.Query().Get("id")
	if !bson.IsObjectIdHex(idStr) {
		HttpError{nil, "No id provided", 500, w}.HandleErr()
		return
	}
	id := bson.ObjectIdHex(idStr)

	var cards []game.DealerCard
	err = json.NewDecoder(f).Decode(&cards)
	if err != nil {
		HttpError{err, "Error decoding file", 500, w}.HandleErr()
		return
	}
	for _, card := range cards {
		card.CreatedBy = id
		err = card.Create()
		if err != nil {
			HttpError{err, "Error saving card: " + card.Phrase, 500, w}.HandleErr()
			return
		}
	}
	sendJson(w, cards)
}
