package game

import (
	"errors"
	"strings"

	"github.com/stinkyfingers/socket/server/db"
	"gopkg.in/mgo.v2/bson"
)

type Card struct {
	ID                bson.ObjectId `bson:"_id" json:"_id"`
	Phrase            string        `bson:"phrase" json:"phrase"`
	PlayerID          bson.ObjectId `bson:"playerId,omitempty" json:"playerId,omitempty"`
	CreatedBy         bson.ObjectId `bson:"createdBy,omitempty" json:"createdBy,omitempty"`
	CorporateApproved bool          `bson:"corporateApproved" json:"corporateApproved"`
	CorporateReviewed bool          `bson:"corporateReviewed" json:"corporateReviewed"`
}

type DealerCard struct {
	ID                bson.ObjectId `bson:"_id" json:"_id"`
	Phrase            string        `bson:"phrase" json:"phrase"`
	CreatedBy         bson.ObjectId `bson:"createdBy,omitempty" json:"createdBy,omitempty"`
	CorporateApproved bool          `bson:"corporateApproved" json:"corporateApproved"`
	CorporateReviewed bool          `bson:"corporateReviewed" json:"corporateReviewed"`
}

var (
	cardCollection       = "cards"
	dealerCardCollection = "dealerCards"
)

func (c *Card) Create() error {
	c.ID = bson.NewObjectId()
	return db.Session.DB(db.DB).C(cardCollection).Insert(c)
}

func (c *Card) Delete() error {
	return db.Session.DB(db.DB).C(cardCollection).RemoveId(c.ID)
}

func GetAllCards() ([]Card, error) {
	var cards []Card
	err := db.Session.DB(db.DB).C(cardCollection).Find(nil).All(&cards)
	return cards, err
}

// FindCard finds card by phrase, case insensitve
func FindCard(phrase string) (Card, error) {
	query := bson.M{
		"phrase": bson.M{
			"$regex": bson.RegEx{
				Pattern: phrase,
				Options: "i",
			},
		},
	}
	var card Card
	err := db.Session.DB(db.DB).C(cardCollection).Find(query).One(&card)
	return card, err
}

// Audit capitalized first word, returns error if card phrase is empty or already exists
func (c *Card) Audit() error {
	strArray := strings.Split(c.Phrase, " ")
	if len(strArray) < 1 {
		return errors.New("Card does not have text")
	}
	strArray[0] = strings.Title(strArray[0])
	(*c).Phrase = strings.Join(strArray, " ")

	_, err := FindCard(c.Phrase)
	if err == nil {
		return errors.New("Card already exists")
	}
	return nil
}

func (c *DealerCard) Create() error {
	c.ID = bson.NewObjectId()
	return db.Session.DB(db.DB).C(dealerCardCollection).Insert(c)
}

func (c *DealerCard) Delete() error {
	return db.Session.DB(db.DB).C(dealerCardCollection).RemoveId(c.ID)
}

func GetAllDealerCards() ([]DealerCard, error) {
	var cards []DealerCard
	err := db.Session.DB(db.DB).C(dealerCardCollection).Find(nil).All(&cards)
	return cards, err
}

// FindCard finds card by phrase, case insensitve
func FindDealerCard(phrase string) (DealerCard, error) {
	query := bson.M{
		"phrase": bson.M{
			"$regex": bson.RegEx{
				Pattern: phrase,
				Options: "i",
			},
		},
	}
	var card DealerCard
	err := db.Session.DB(db.DB).C(dealerCardCollection).Find(query).One(&card)
	return card, err
}

// Audit capitalized first word, returns error if card phrase is empty or already exists
func (c *DealerCard) Audit() error {
	strArray := strings.Split(c.Phrase, " ")
	if len(strArray) < 1 {
		return errors.New("Card does not have text")
	}
	strArray[0] = strings.Title(strArray[0])
	(*c).Phrase = strings.Join(strArray, " ")

	_, err := FindDealerCard(c.Phrase)
	if err == nil {
		return errors.New("Card already exists")
	}
	return nil
}

func GetUnreviewedCards() ([]DealerCard, []Card, error) {
	query := bson.M{
		"$or": []interface{}{
			bson.M{"corporateReviewed": false}, bson.M{"corporateReviewed": nil},
		},
	}
	var dealerCards []DealerCard
	var cards []Card
	err := db.Session.DB(db.DB).C(dealerCardCollection).Find(query).All(&dealerCards)
	if err != nil {
		return dealerCards, cards, err
	}
	err = db.Session.DB(db.DB).C(cardCollection).Find(query).All(&cards)
	if err != nil {
		return dealerCards, cards, err
	}
	return dealerCards, cards, err
}

func (c *Card) Update() error {
	update := bson.M{
		"corporateApproved": c.CorporateApproved,
		"corporateReviewed": true,
	}
	return db.Session.DB(db.DB).C(cardCollection).UpdateId(c.ID, update)
}

func (c *DealerCard) Update() error {
	update := bson.M{
		"corporateApproved": c.CorporateApproved,
		"corporateReviewed": true,
	}
	return db.Session.DB(db.DB).C(dealerCardCollection).UpdateId(c.ID, update)
}
