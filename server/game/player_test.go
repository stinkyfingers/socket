package game

import (
	"testing"

	"github.com/stinkyfingers/socket/server/db"
)

func TestCRUD(t *testing.T) {
	db.NewSession()
	defer Cleanup()
	db.DB = "game-test"
	p := Player{
		Name:     "test",
		Password: "test_pass",
	}
	err := p.Create()
	if err != nil {
		t.Fatal(err)
	}

	err = p.Get()
	if err != nil {
		t.Error(err)
	}
	if !p.ID.Valid() {
		t.Error("expected ID  to be populated")
	}

	p.Password = "new_pass"
	err = p.Update()
	if err != nil {
		t.Error(err)
	}

	newPlayer := Player{
		Name:     "test",
		Password: "new_pass",
	}
	err = newPlayer.Authenticate()
	if err != nil {
		t.Error(err)
	}

	err = p.ResetPassword()
	if err != nil {
		t.Error(err)
	}
	if p.Password == "new_pass" {
		t.Error("expected password to be reset")
	}

	err = p.Delete()
	if err != nil {
		t.Error(err)
	}

	err = p.Get()
	if err == nil {
		t.Error("expected User to not exists")
	}
}

func Cleanup() error {
	return db.Session.DB(db.DB).DropDatabase()
}
