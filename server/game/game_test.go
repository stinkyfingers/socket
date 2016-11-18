package game

import (
	"os"
	"reflect"
	"testing"

	"github.com/stinkyfingers/socket/server/db"
	"gopkg.in/mgo.v2/bson"
)

func TestMain(m *testing.M) {
	db.NewSession()
	os.Exit(m.Run())
}

func TestGet(t *testing.T) {
	t.Log("Test Get")
	g := Game{
		ID: bson.ObjectIdHex("58291ddadd162c05e145c413"),
	}
	err := g.Get()
	if err != nil {
		t.Error(err)
	}
}

func TestGetRounds(t *testing.T) {
	t.Log("Test GetRounds")
	r1 := Round{
		Score:    map[string][]Play{"1": []Play{{}}},
		Previous: nil,
	}
	r2 := Round{
		Score:    map[string][]Play{"2": []Play{{}}},
		Previous: &r1,
	}
	r3 := Round{
		Score:    map[string][]Play{"3": []Play{{}}},
		Previous: &r2,
	}
	g := Game{
		Round: r3,
	}
	rounds := g.GetRounds()

	if !reflect.DeepEqual(rounds, []Round{r3, r2, r1}) {
		t.Error("Expected to find three rounds")
	}
}

func TestUpdateVotes(t *testing.T) {
	id1 := bson.NewObjectId()
	id2 := bson.NewObjectId()
	g := Game{
		Players: []Player{
			{
				Name: "Jim",
				Hand: []Card{
					{Phrase: "test"},
					{Phrase: "test2"},
				},
				ID: id1,
			},
			{
				Name: "Fred",
				Hand: []Card{
					{Phrase: "test3"},
					{Phrase: "test4"},
				},
				ID: id2,
			},
		},
		Round: Round{
			Votes: map[string]Play{
				id1.Hex(): {
					Card: Card{
						Phrase: "test",
					},
					Player: Player{ID: id1},
				}, id2.Hex(): {
					Card: Card{
						Phrase: "test",
					},
					Player: Player{ID: id1},
				},
			},
		},
		ID:         bson.NewObjectId(),
		DealerDeck: []DealerCard{{Phrase: ""}, {Phrase: ""}, {Phrase: ""}, {Phrase: ""}, {Phrase: ""}, {Phrase: ""}, {Phrase: ""}},
		Deck:       []Card{{}, {}, {}, {}},
	}
	err := g.UpdateVotes()
	if err != nil {
		t.Error(err)
	}
}
