package db

import (
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
	return &mgo.DialInfo{
		Addrs:    []string{"127.0.0.1"},
		Database: "admin",
		Timeout:  time.Second,
	}
}
