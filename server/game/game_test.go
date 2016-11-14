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
	t.Log(rounds)

	if !reflect.DeepEqual(rounds, []Round{r3, r2, r1}) {
		t.Error("Expected to find three rounds")
	}
}
