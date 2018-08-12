package models

import (
	"github.com/boltdb/bolt"
	"fmt"
)

type Datastore interface {
	AddNotebook(notebook Notebook) (error)
	GetNotebook(notebookTitle string) (Notebook, error)
	AddNote(note Note) (error)
	GetNote(noteIndex uint64) (Note, error)
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
		root, err := tx.CreateBucketIfNotExists([]byte("Notebook"))
		if err != nil {
			return fmt.Errorf("could not create root bucket: %v", err)
		}

		_, err = root.CreateBucketIfNotExists([]byte("Note"))
		if err != nil {
			return fmt.Errorf("could not create note bucket: %v", err)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &DB{db}, nil
}
