package database

import (
	"sync"
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

func ExecuteInSession(fn func(session Session)) {
	session := GetSession().Clone()
	defer session.Close()
	fn(session)
}

// Inits the database with a default implementation. The default is MGO.
func InitDatabase() {
	InitDatabaseWithImpl(MgoDatabase{})
}

func InitDatabaseWithImpl(impl Database) {
	database = impl
}
