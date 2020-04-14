package models

import (
	"fmt"

	"github.com/boltdb/bolt"
)

/**
 * - interface defining operations of our 'notes' datastore
 * - notice the function prototypes here have one-to-one mapping with
 *   *almost* every command (plus it's subcommand / options)
 */
type Datastore interface {
	// notebook-related operations
	NotebookExists(notebookName string) (bool, error)
	GetNotebook(notebookTitle string) (Notebook, error)
	GetAllNotebooks() ([]Notebook, error)
	GetAllNotebookNames() ([]string, error)
	AddNotebook(notebook Notebook) error
	RmNotebook(notebookName string) error

	// note-related operations
	NoteExists(notebookName string, noteId uint64) (bool, error)
	GetNote(noteIndex uint64) (Note, error)
	AddNote(notebookTitle string, notes ...Note) error
	DeleteNote(notebookName string, noteIDs ...uint64) error
	// db-backup operation
	Dump()
}

/**
 * - structure that implements the above Datastore interface
 * - the concrete implementation of methods has been spread across 2 files
 *   1. notebook.go: Defines DTO (struct) 'Notebook'
 *     - notebook-related operations
 *     - db-backup operation
 *   2. note.go: Defines DTO (struct) 'Note'
 *     - note-related operations
 * - Looking at the implementation of following constructor 'GetOrCreateDB'
 *   it can be inferred that this is just a wrapper over BoltDb's DB struct
 */
type DB struct {
	*bolt.DB
}

/**
 * <Constructor for above DB struct>
 * Returns an instance of DB struct by either creating a new BoltDb
 * bucket for given `dbFileName` or using an existing one
 * @param dbFileName string The complete (path) qualified filename of for BoltDb file
 * @return (*DB, error) Tuple containing pointer to DB struct and optionally an error
 */
func GetOrCreateDB(dbFileName string) (*DB, error) {
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
