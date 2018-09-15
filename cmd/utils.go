package cmd

import (
	"github.com/uncultured/notes/models"
	"log"
)


func setupDatabase() models.Datastore {
	var database models.Datastore
	database, err := models.NewDB("notebooks.db")
	if err != nil {
		log.Panic(err)
	}
	return database
}