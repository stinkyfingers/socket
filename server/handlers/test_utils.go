package handlers

import (
	"encoding/json"
	"io/ioutil"

	"github.com/stinkyfingers/socket/server/db"
	"github.com/stinkyfingers/socket/server/game"
	"gopkg.in/mgo.v2/bson"
)

func Setup(path string) ([]game.Card, []game.DealerCard, error) {
	decksUser := struct {
		Deck       []game.Card       `json:"deck"`
		Player     game.Player       `json:"player"`
		DealerDeck []game.DealerCard `json:"dealerDeck"`
	}{}

	b, err := ioutil.ReadFile(path)
	if err != nil {
		return []game.Card{}, []game.DealerCard{}, err
	}

	err = json.Unmarshal(b, &decksUser)
	if err != nil {
		return []game.Card{}, []game.DealerCard{}, err
	}

	decksUser.Player.ID = bson.NewObjectId()

	return decksUser.Deck, decksUser.DealerDeck, nil
}

func SetupPlayers() ([]game.Player, error) {
	players := []game.Player{
		{
			Name:     "Alice",
			Password: "test",
		}, {
			Name:     "Bob",
			Password: "test",
		}, {
			Name:     "Carl",
			Password: "test",
		},
	}
	for i := range players {
		err := players[i].Create()
		if err != nil {
			return players, err
		}
	}
	return players, nil
}

func Cleanup() error {
	return db.Session.DB(db.DB).DropDatabase()
}
