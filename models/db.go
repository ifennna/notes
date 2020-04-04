package models

import (
	"fmt"

	"github.com/boltdb/bolt"
)

type Datastore interface {
	AddNotebook(notebook Notebook) error
	GetNotebook(notebookTitle string) (Notebook, error)
	GetAllNotebooks() ([]Notebook, error)
	AddNote(notebookTitle string, note ...Note) error
	DeleteNote(notebookName string, noteId ...uint64) error
	GetNote(noteIndex uint64) (Note, error)
	Dump()
}

type DB struct {
	*bolt.DB
}

func NewDB(dbFileName string) (*DB, error) {
	db, err := bolt.Open(dbFileName, 0600, nil)
	if err != nil {
		return nil, err
	}
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("Notebook"))
		if err != nil {
			return fmt.Errorf("could not create root bucket: %v", err)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &DB{db}, nil
}
