package database

type Session interface {
	Clone() Session
	Close()
	Schema(string) Schema
	DefaultSchema() Schema
}
