package cmd

import (
	"log"
	"os/user"
	"path"

	"github.com/noculture/notes/models"
)

// TODO: configure a global path here because as it is,
// a new db is created when the app is run from a diff folder
func setupDatabase() models.Datastore {
	var database models.Datastore
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	notebook := path.Join(usr.HomeDir, ".notebooks.db")

	database, err = models.NewDB(notebook)
	if err != nil {
		log.Panic(err)
	}
	return database
}
