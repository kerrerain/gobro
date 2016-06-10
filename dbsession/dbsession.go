package dbsession

import (
	"gopkg.in/mgo.v2"
	"time"
	"log"
)

// As seen in http://blog.mongodb.org/post/80579086742/running-mongodb-queries-concurrently-with-go
func Create() *mgo.Session {
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{"localhost:27017"},
		Timeout:  60 * time.Second,
		Database: "expenses-analyzer"}
	// Create a session which maintains a pool of socket connections
	// to our MongoDB.
	mongoSession, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		log.Fatalf("CreateSession: %s\n", err)
	}
	return mongoSession
}
