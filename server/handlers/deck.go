package handlers

import (
	"encoding/csv"
	"io"
	"net/http"
	"os"

	"github.com/stinkyfingers/socket/server/game"
)

func HandleExportDealerDeck(w http.ResponseWriter, r *http.Request) {
	file := r.URL.Query().Get("file")
	dealerDeck, err := game.GetDealerDeck()
	if err != nil {
		HttpError{Error: err, Message: "Error getting deck", Status: 500, w: w}.HandleErr()
		return
	}

	f, err := os.Create(file)
	if err != nil {
		HttpError{Error: err, Message: "Error creating file", Status: 500, w: w}.HandleErr()
		return
	}
	writer := csv.NewWriter(f)
	var lines [][]string
	for _, card := range dealerDeck.Cards {
		line := []string{card.Phrase}
		lines = append(lines, line)
	}
	writer.WriteAll(lines)
	writer.Flush()
	w.Header().Set("Content-Disposition", "attachment; filename=WHATEVER_YOU_WANT")
	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
	io.Copy(w, f)
}

func HandleImportDealerDeck(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Query().Get("file")
	f, err := os.Open(path)
	if err != nil {
		HttpError{Error: err, Message: "Error opening file", Status: 500, w: w}.HandleErr()
		return
	}
	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		HttpError{Error: err, Message: "Error reading file", Status: 500, w: w}.HandleErr()
		return
	}
	deck, err := game.GetDealerDeck()
	if err != nil {
		HttpError{Error: err, Message: "Error getting dealer deck", Status: 500, w: w}.HandleErr()
		return
	}

	for _, line := range records {
		deck.Cards = append(deck.Cards, game.DealerCard{Phrase: line[0]})
	}
	err = deck.Update()
	if err != nil {
		HttpError{Error: err, Message: "Error updating deck", Status: 500, w: w}.HandleErr()
		return
	}
	sendJson(w, deck)
}
