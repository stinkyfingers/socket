package game

import (
	"testing"

	"github.com/stinkyfingers/socket/server/db"
	"gopkg.in/mgo.v2/bson"
)

func TestMain(m *testing.M) {
	db.NewSession()
}

func TestGet(t *testing.T) {
	g := Game{
		ID: bson.ObjectIdHex("58291ddadd162c05e145c413"),
	}
	err := g.Get()
	if err != nil {
		t.Error(err)
	}
}
