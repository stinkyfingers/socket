package game

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/stinkyfingers/socket/server/db"
)

func TestDeal(t *testing.T) {
	defer cleanup("test_game")

	err := db.NewSession()
	if err != nil {
		t.Error(err)
	}
	db.DB = "test_game"

	players, err := TwoUsers()
	if err != nil {
		t.Error(err)
	}

	err = setupDecks()
	if err != nil {
		t.Error(err)
	}

	deck, err := GetAllCards()
	if err != nil {
		t.Error(err)
	}
	dealerDeck, err := GetAllDealerCards()
	if err != nil {
		t.Error(err)
	}

	g := Game{
		Deck:       deck,
		DealerDeck: dealerDeck,
		Players:    players,
	}
	err = g.Create()

	err = g.Deal()
	if err != nil {
		t.Error(err)
	}

}

func setupDecks() error {
	//cards
	var cards []Card
	b, err := ioutil.ReadFile("../cards.json")
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, &cards)
	if err != nil {
		return err
	}

	for _, card := range cards {
		err = card.Create()
		if err != nil {
			return err
		}
	}

	var dealerCards []DealerCard
	b, err = ioutil.ReadFile("../dealerCards.json")
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, &dealerCards)
	if err != nil {
		return err
	}

	for _, card := range dealerCards {
		err = card.Create()
		if err != nil {
			return err
		}
	}

	return err
}

func TwoUsers() ([]Player, error) {
	err := db.NewSession()
	if err != nil {
		return nil, err
	}
	db.DB = "test_game"
	var users []Player
	p1 := Player{
		Name:     "Zach",
		Password: "foo",
	}
	p2 := Player{
		Name:     "Zed",
		Password: "foo",
	}
	err = p1.Create()
	if err != nil {
		return users, err
	}

	err = p2.Create()
	if err != nil {
		return users, err
	}
	users = []Player{p1, p2}
	return users, err
}
