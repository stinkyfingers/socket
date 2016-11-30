package game

import (
	"testing"

	"github.com/stinkyfingers/socket/server/db"
	"gopkg.in/mgo.v2/bson"
)

func TestCardCRUD(t *testing.T) {
	db.NewSession()
	cardCollection = "tempCardCollection"
	defer cleanup(cardCollection)

	c := Card{
		Phrase:    "test 1",
		CreatedBy: bson.NewObjectId(),
	}

	// Create
	err := c.Create()
	if err != nil {
		t.Error(err)
	}

	// Find
	card, err := FindCard(c.Phrase)
	if err != nil {
		t.Error(err)
	}
	if card.ID != c.ID {
		t.Error("Expected found card to match")
	}

	// Audit
	newCard1 := Card{Phrase: "funk"}
	err = newCard1.Audit()
	if err != nil {
		t.Error(err)
	}
	if newCard1.Phrase != "Funk" {
		t.Error("Expected card to be capitalized")
	}

	newCard2 := Card{Phrase: "Test 1"}
	err = newCard2.Audit()
	if err == nil {
		t.Error("Expected card to already exist")
	}

	//Get All
	cards, err := GetAllCards()
	if err != nil {
		t.Error(err)
	}
	if len(cards) < 1 {
		t.Error("Expected to find dealer cards")
	}

	// Delete
	err = c.Delete()
	if err != nil {
		t.Error(err)
	}
}

func TestDealerCardCRUD(t *testing.T) {
	db.NewSession()
	dealerCardCollection = "tempDealerCollection"
	defer cleanup(dealerCardCollection)
	c := DealerCard{
		Phrase:    "dealertest 1",
		CreatedBy: bson.NewObjectId(),
	}
	// Create
	err := c.Create()
	if err != nil {
		t.Error(err)
	}

	// Fine
	card, err := FindDealerCard(c.Phrase)
	if err != nil {
		t.Error(err)
	}

	if card.ID != c.ID {
		t.Error("Expected found card to match")
	}

	// Audit
	newCard1 := DealerCard{Phrase: "funk"}
	err = newCard1.Audit()
	if err != nil {
		t.Error(err)
	}
	if newCard1.Phrase != "Funk" {
		t.Error("Expected card to be capitalized")
	}

	newCard2 := DealerCard{Phrase: "Dealertest 1"}
	err = newCard2.Audit()
	if err == nil {
		t.Error("Expected card to already exist")
	}

	//Get All
	cards, err := GetAllDealerCards()
	if err != nil {
		t.Error(err)
	}
	if len(cards) < 1 {
		t.Error("Expected to find dealer cards")
	}

	// Delete
	err = c.Delete()
	if err != nil {
		t.Error(err)
	}
}

// remove test db by name
func cleanup(name string) {
	db.Session.DB(name).DropDatabase()
}
