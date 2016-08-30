package database

import (
	"gopkg.in/mgo.v2"
	dockertest "gopkg.in/ory-am/dockertest.v2"
	"log"
	"time"
)

type MgoDatabaseTest struct {
	Container dockertest.ContainerID
}

func (self *MgoDatabaseTest) DialDatabase() Session {
	var db *mgo.Session

	c, err := dockertest.ConnectToMongoDB(15, time.Millisecond*500, func(url string) bool {
		// This callback function checks if the image's process is responsive.
		// Sometimes, docker images are booted but the process (in this case MongoDB) is still doing maintenance
		// before being fully responsive which might cause issues like "TCP Connection reset by peer".
		var err error
		db, err = mgo.Dial(url)
		if err != nil {
			return false
		}

		// Sometimes, dialing the database is not enough because the port is already open but the process is not responsive.
		// Most database conenctors implement a ping function which can be used to test if the process is responsive.
		// Alternatively, you could execute a query to see if an error occurs or not.
		return db.Ping() == nil
	})

	if err != nil {
		log.Fatalf("Could not connect to database: %s", err)
	}

	self.Container = c

	return MgoSession{db}
}
