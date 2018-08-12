package models

import (
	"github.com/boltdb/bolt"
)

type Datastore interface {
	AddNote(note Note) (error)
}

type DB struct {
	*bolt.DB
}

func NewDB(dbFileName string) (*DB, error) {
	db, err := bolt.Open(dbFileName, 0600, nil)
	if err != nil {
		return nil, err
	}

	return &DB{db}, nil
}
