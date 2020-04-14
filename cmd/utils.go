package cmd

import (
	"log"
	"os/user"
	"path"

	"github.com/noculture/notes/models"
)

/**
 * Creates (or uses existing one) bolt-db database storage file called '.notebooks.db' in user's home directory
 * Returns models.Datastore object that implements interfaces which can be used for managing our datastore
 * (like add note, show etc.)
 * return: models.Datastore
 */
func setupDatabase() models.Datastore {
	var database models.Datastore
	// determine current user's home directory and build path for a '.notebooks.db' file there
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	notebook := path.Join(usr.HomeDir, ".notebooks.db")

	// create a bolt-db file (.notebooks.db) or use the existing one
	database, err = models.GetOrCreateDB(notebook)
	if err != nil {
		log.Panic(err)
	}
	return database
}
