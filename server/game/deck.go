package game

import (
	"github.com/stinkyfingers/socket/server/db"
	"gopkg.in/mgo.v2/bson"
)

type DealerDeck struct {
	ID    bson.ObjectId `json:"_id" bson:"_id"`
	Cards []DealerCard  `json:"cards" bson:"cards"`
}

type Deck struct {
	ID    bson.ObjectId `json:"_id" bson:"_id"`
	Cards []Card        `json:"cards" bson:"cards"`
}

var DeckCollection = "deck"
var DealerDeckCollection = "dealer-deck"

func (d *Deck) Create() error {
	d.ID = bson.NewObjectId()
	return db.Session.DB(db.DB).C(DeckCollection).Insert(&d)
}

func GetDeck() (Deck, error) {
	var d Deck
	err := db.Session.DB(db.DB).C(DeckCollection).Find(nil).One(&d)
	return d, err
}

func (d *Deck) Update() error {
	return db.Session.DB(db.DB).C(DeckCollection).UpdateId(d.ID, d)
}

func (d *DealerDeck) Create() error {
	d.ID = bson.NewObjectId()
	return db.Session.DB(db.DB).C(DealerDeckCollection).Insert(&d)
}

func GetDealerDeck() (DealerDeck, error) {
	var d DealerDeck
	err := db.Session.DB(db.DB).C(DealerDeckCollection).Find(nil).One(&d)
	return d, err
}

func (d *DealerDeck) Update() error {
	return db.Session.DB(db.DB).C(DealerDeckCollection).UpdateId(d.ID, d)
}
