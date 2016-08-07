package database

import (
	"gopkg.in/mgo.v2"
	"log"
	"sync"
	"time"
)

var session Session
var database Database
var once sync.Once

type Database interface {
	DialDatabase() Session
}

// Gets the singleton session of the database.
// The dabatase is dialed only once: calling this method twice
// won't create another sessions.
func GetSession() Session {
	once.Do(func() {
		session = database.DialDatabase()
	})
	return session
}

// Inits the database with a default implementation.
// The default is MGO
func InitDatabase() {
	InitDatabaseWithImpl(MgoDatabase{})
}

func InitDatabaseWithImpl(impl Database) {
	database = impl
}

/*
	Implementation
*/

type MgoDatabase struct{}

func (self MgoDatabase) DialDatabase() Session {
	var err error
	var mgoSession *mgo.Session

	info := &mgo.DialInfo{
		Addrs:    []string{"localhost:27017"},
		Database: "gobro",
		Timeout:  60 * time.Second}

	mgoSession, err = mgo.DialWithInfo(info)
	if err != nil {
		log.Fatalf("CreateSession: %s\n", err)
	}

	return MgoSession{mgoSession}
}
