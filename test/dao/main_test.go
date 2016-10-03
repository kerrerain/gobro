package dao_test

import (
	"github.com/magleff/gobro/database"
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	impl := &database.MgoDatabaseTest{}

	log.Println("Tests on DAO, creating mongodb container...")

	database.InitDatabaseWithImpl(impl)

	log.Println("Container successfully created. Running tests.")

	// Run tests
	result := m.Run()

	// Close the database's main session after running the command
	database.GetSession().Close()

	log.Println("Deleting mongodb container...")

	// Clean up the docker container
	impl.Container.KillRemove()

	log.Println("Container successfully deleted.")

	// Exit tests
	os.Exit(result)
}
