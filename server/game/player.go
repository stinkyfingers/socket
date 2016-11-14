package game

import (
	"errors"

	"github.com/stinkyfingers/socket/server/db"
	"gopkg.in/mgo.v2/bson"
)

var playerCollection = "difference-between-players"

func (p *Player) Create() error {
	var err error
	query := bson.M{
		"name": p.Name,
	}

	count, err := db.Session.DB(db.DB).C(playerCollection).Find(query).Count()
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("user name already exists")
	}

	p.ID = bson.NewObjectId()
	// u.Created = time.Now()
	// u.EncryptedPassword, err = bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	// if err != nil {
	// 	return err
	// }
	return db.Session.DB(db.DB).C(playerCollection).Insert(&p)
}
