package database

import (
	"gopkg.in/mgo.v2"
)

type Datastore struct {
	Session Session
}

func (self *Datastore) BindSession() {
	self.Session = GetSession().Clone()
}

func (self *Datastore) CloseSession() {
	if self.Session != nil {
		self.Session.Close()
	}
}

// Clones the main database session and executes the given function.
// The session is bound to the Datastore.Session attribute.
// It is automatically closed after the execution of the given function.
func (self *Datastore) ExecuteInSession(fn func()) {
	self.BindSession()
	defer self.CloseSession()
	fn()
}

func (self *Datastore) Collection(name string) *mgo.Collection {
	return self.Session.DefaultSchema().Collection(name)
}
