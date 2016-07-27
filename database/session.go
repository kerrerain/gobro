package database

import (
	"gopkg.in/mgo.v2"
)

type Session interface {
	Clone() Session
	Close()
	Schema(string) Schema
	DefaultSchema() Schema
}

/*
	Implementation
*/

type MgoSession struct {
	Session *mgo.Session
}

func (self MgoSession) Clone() Session {
	return MgoSession{self.Session.Clone()}
}

func (self MgoSession) Close() {
	if self.Session != nil {
		self.Session.Close()
	}
}

func (self MgoSession) Schema(name string) Schema {
	return MgoSchema{self.Session.DB(name)}
}

func (self MgoSession) DefaultSchema() Schema {
	return MgoSchema{self.Session.DB("")}
}
