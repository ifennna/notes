package models

import (
	"bytes"
	"encoding/json"
	"strconv"

	"github.com/boltdb/bolt"
)

//Note DTO for a Note within a Notebook
type Note struct {
	// TODO: explore allowing 'naming' notes within a notebook
	//Title   string `json:"title"`
	ID      uint64 `json:"id"`
	Content string `json:"content"`
}

//NoteExists returns whether or not note with a given id exists in the given notebook or not
func (db *DB) NoteExists(notebookName string, noteID uint64) (bool, error) {
	noteExists := false
	err := db.View(func(tx *bolt.Tx) error {
		reqNoteIDBytes := []byte(strconv.FormatUint(noteID, 10))
		notebookBucket := tx.Bucket([]byte("Notebook")).Bucket([]byte(notebookName))

		foundNoteIDBytes, _ := notebookBucket.Cursor().Seek(reqNoteIDBytes)
		if foundNoteIDBytes != nil && bytes.Equal(reqNoteIDBytes, foundNoteIDBytes) {
			noteExists = true
		}

		return nil
	})
	return noteExists, err
}

//GetNote retreives note with the given id
/*
 * param: uint64 noteId
 * return: (Note, error)
 */
func (db *DB) GetNote(notebookName string, noteID uint64) (Note, error) {
	var note Note
	err := db.View(func(tx *bolt.Tx) error {
		reqNoteIDBytes := []byte(strconv.FormatUint(noteID, 10))
		notebookBucket := tx.Bucket([]byte("Notebook")).Bucket([]byte(notebookName))

		foundNoteIDBytes, foundNoteContentBytes := notebookBucket.Cursor().Seek(reqNoteIDBytes)
		if foundNoteIDBytes != nil && bytes.Equal(reqNoteIDBytes, foundNoteIDBytes) {
			return json.Unmarshal(foundNoteContentBytes, &note)
		}

		return nil
	})
	return note, err
}

//AddNotes adds notes in the given notebook
/*
 * Adds notes in the given notebook
 * notes' auto-increment 'Id' are generated and stored in the db by this method itself
 * param: string notebookName
 * param: ...Note notes
 * return: error
 */
func (db *DB) AddNotes(notebookName string, noteContents ...string) error {
	// create a bolt-db transaction with deferred-rollback
	tx, err := db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// create or retrieve (2nd order) bucket with given notebookName
	notebookBucket, err := tx.Bucket([]byte("Notebook")).CreateBucketIfNotExists([]byte(notebookName))
	if err != nil {
		return err
	}

	// for each noteContent to be added
	for _, noteContent := range noteContents {
		// create Note object
		var note Note = Note{Content: noteContent}

		// gereate noteId
		noteID, err := notebookBucket.NextSequence()
		if err != nil {
			return err
		}
		note.ID = noteID

		// put JSON-marshalled noteContent into bolt-db bucket (of given Notebook) with noteId as key
		if encodedNote, err := json.Marshal(note); err != nil {
			return err
		} else if err := notebookBucket.Put([]byte(strconv.FormatUint(noteID, 10)), encodedNote); err != nil {
			return err
		}
	}

	// Commit the transaction.
	if err := tx.Commit(); err != nil {
		return err
	}

	return err
}

//DeleteNotes deletes notes with given ids from the given notebook
/*
 * param: string notebookName
 * param: ...uint64 noteIds
 * return: error
 */
func (db *DB) DeleteNotes(notebookName string, noteIds ...uint64) error {
	// TODO: try to remove code-duplication: txn creation & notebook notebookBucket retrieval logic can be extracted out
	// create a bolt-db transaction with deferred-rollback
	tx, err := db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// retrieve (2nd order, noteBook) bucket with given notebookName
	notebookBucket := tx.Bucket([]byte("Notebook")).Bucket([]byte(notebookName))

	// for each noteId supplied
	for _, noteID := range noteIds {
		// delete the note with given noteId from notebook's bucket
		err = notebookBucket.Delete([]byte(strconv.FormatUint(noteID, 10)))
		if err != nil {
			return err
		}
	}

	// Commit the transaction.
	if err := tx.Commit(); err != nil {
		return err
	}

	return err
}
