package handlers

import (
	"encoding/csv"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/stinkyfingers/socket/server/game"
)

func HandleExportDealerDeck(w http.ResponseWriter, r *http.Request) {
	file := r.URL.Query().Get("file")
	dealerDeck, err := game.GetAllDealerCards()
	if err != nil {
		HttpError{Error: err, Message: "Error getting deck", Status: 500, W: w}.HandleErr()
		return
	}

	f, err := os.Create(file)
	if err != nil {
		HttpError{Error: err, Message: "Error creating file", Status: 500, W: w}.HandleErr()
		return
	}
	writer := csv.NewWriter(f)
	var lines [][]string
	for _, card := range dealerDeck {
		line := []string{card.ID.Hex(), card.Phrase, card.CreatedBy.Hex(), strconv.FormatBool(card.CorporateApproved)}
		lines = append(lines, line)
	}
	writer.WriteAll(lines)
	writer.Flush()
	w.Header().Set("Content-Disposition", "attachment; filename=WHATEVER_YOU_WANT")
	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
	io.Copy(w, f)
}

func HandleExportDeck(w http.ResponseWriter, r *http.Request) {
	file := r.URL.Query().Get("file")
	deck, err := game.GetAllCards()
	if err != nil {
		HttpError{Error: err, Message: "Error getting deck", Status: 500, W: w}.HandleErr()
		return
	}

	f, err := os.Create(file)
	if err != nil {
		HttpError{Error: err, Message: "Error creating file", Status: 500, W: w}.HandleErr()
		return
	}
	writer := csv.NewWriter(f)
	var lines [][]string
	for _, card := range deck {
		line := []string{card.ID.Hex(), card.Phrase, card.CreatedBy.Hex(), strconv.FormatBool(card.CorporateApproved)}
		lines = append(lines, line)
	}
	writer.WriteAll(lines)
	writer.Flush()
	w.Header().Set("Content-Disposition", "attachment; filename=WHATEVER_YOU_WANT")
	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
	io.Copy(w, f)
}
