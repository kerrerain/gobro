package models

import (
	"github.com/magleff/gobro/database"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	impl := &database.MgoDatabaseTest{}

	database.InitDatabaseWithImpl(impl)

	// Run tests
	result := m.Run()

	// Close the database's main session after running the command
	database.GetSession().Close()

	// Clean up the docker container
	impl.Container.KillRemove()

	// Exit tests
	os.Exit(result)
}
