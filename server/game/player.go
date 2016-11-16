package game

import (
	"errors"
	"math/rand"

	"github.com/stinkyfingers/socket/server/db"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

type Player struct {
	ID   bson.ObjectId `bson:"_id" json:"_id"`
	Name string        `bson:"name" json:"name"`
	Hand []Card        `bson:"hand" json:"hand"`

	Password          string `bson:"-" json:"password"`
	EncryptedPassword []byte `bson:"encryptedPassword" json:"-"`
	Email             string `bson:"email" json:"email"`
}

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
	p.EncryptedPassword, err = bcrypt.GenerateFromPassword([]byte(p.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	return db.Session.DB(db.DB).C(playerCollection).Insert(&p)
}

func (p *Player) Get() error {
	return db.Session.DB(db.DB).C(playerCollection).FindId(p.ID).One(&p)
}

func (p *Player) Update() error {
	var err error
	p.EncryptedPassword, err = bcrypt.GenerateFromPassword([]byte(p.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	return db.Session.DB(db.DB).C(playerCollection).UpdateId(p.ID, p)
}

func (p *Player) Delete() error {
	return db.Session.DB(db.DB).C(playerCollection).RemoveId(p.ID)
}

func (p *Player) Authenticate() error {
	query := bson.M{
		"name": p.Name,
	}
	password := p.Password

	count, err := db.Session.DB(db.DB).C(playerCollection).Find(query).Count()
	if err != nil {
		return err
	}
	if count > 1 {
		return errors.New("too many users with that username.")
	}
	if count == 0 {
		return errors.New("no users with that username")
	}
	err = db.Session.DB(db.DB).C(playerCollection).Find(query).One(&p)
	if err != nil {
		return err
	}
	return bcrypt.CompareHashAndPassword(p.EncryptedPassword, []byte(password))
}

func (p *Player) ResetPassword() error {
	n := 8
	var letterRunes = []rune("abcdefghijkmnpqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ23456789")

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	p.Password = string(b)
	return p.Update()
}
