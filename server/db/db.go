package db

import (
	"os"
	"time"

	"gopkg.in/mgo.v2"
)

var Session *mgo.Session

var DB string = "boardgame"

func NewSession() error {
	var err error
	Session, err = mgo.DialWithInfo(getDialInfo())
	return err
}

func getDialInfo() *mgo.DialInfo {
	cs := os.Getenv("MONGO_URL")
	if cs == "" {
		cs = "127.0.0.1"
	}
	user := os.Getenv("MONGO_USER")
	pass := os.Getenv("MONGO_PASS")
	admin := os.Getenv("MONGO_ADMIN")
	if admin == "" {
		admin = "admin"
	}
	db := os.Getenv("MONGO_DB")
	if db != "" {
		DB = db
	}

	return &mgo.DialInfo{
		Addrs:    []string{cs},
		Database: admin,
		Timeout:  time.Second,
		Source:   admin,
		Username: user,
		Password: pass,
	}
}
