package models

import (
	"bytes"
	"encoding/json"
	"strconv"

	"github.com/boltdb/bolt"
)

/**
 * DTO for a Note within a Notebook
 */
type Note struct {
	// TODO: explore allowing 'naming' notes within a notebook
	//Title   string `json:"title"`
	Id      uint64 `json:"id"`
	Content string `json:"content"`
}

/**
 * Returns whether or not note with a given id exists
 * in the given notebook or not
 */
func (db *DB) NoteExists(notebookName string, reqNoteId uint64) (bool, error) {
	noteExists := false
	err := db.View(func(tx *bolt.Tx) error {
		reqNoteIdBytes := []byte(strconv.FormatUint(reqNoteId, 10))
		notebookBucket := tx.Bucket([]byte("Notebook")).Bucket([]byte(notebookName))

		foundNoteIdBytes, _ := notebookBucket.Cursor().Seek(reqNoteIdBytes)
		if foundNoteIdBytes != nil && bytes.Equal(reqNoteIdBytes, foundNoteIdBytes) {
			noteExists = true
		}

		return nil
	})
	return noteExists, err
}

/**
 * Adds notes in the given notebook
 * notes' auto-increment 'Id' are generated and stored in the db by this method itself
 * param: string notebookName
 * param: ...Note notes
 * return: error
 */
func (db *DB) AddNote(notebookName string, notes ...Note) error {
	// create a bolt-db transaction with deferred-rollback
	tx, err := db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// create or retrieve (2nd order) bucket with given notebookName
	notebook, err := tx.Bucket([]byte("Notebook")).CreateBucketIfNotExists([]byte(notebookName))

	// for each noteContent to be added
	for _, note := range notes {
		// gereate noteId
		noteID, err := notebook.NextSequence()
		if err != nil {
			return err
		}
		note.Id = noteID
		// put JSON-marshalled noteContent into bolt-db bucket (of given Notebook) with noteId as key
		if encodedNote, err := json.Marshal(note); err != nil {
			return err
		} else if err := notebook.Put([]byte(strconv.FormatUint(noteID, 10)), encodedNote); err != nil {
			return err
		}
	}

	// Commit the transaction.
	if err := tx.Commit(); err != nil {
		return err
	}
	return err
}

/**
 * Deletes notes with given ids from the given notebook
 * param: string notebookName
 * param: ...uint64 noteIds
 * return: error
 */
func (db *DB) DeleteNote(notebookName string, noteIDs ...uint64) error {
	// TODO: try to remove code-duplication: txn creation & notebook notebookBucket retrieval logic can be extracted out
	// create a bolt-db transaction with deferred-rollback
	tx, err := db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// retrieve (2nd order) bucket with given notebookName
	bucket := tx.Bucket([]byte("Notebook")).Bucket([]byte(notebookName))

	// for each noteId supplied
	for _, noteID := range noteIDs {
		// delete the note with given noteId from notebook's bucket
		err = bucket.Delete([]byte(strconv.FormatUint(noteID, 10)))
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

/**
 * Retrives note with a given id
 * param: uint64 noteId
 * return: (Note, error)
 */
func (db *DB) GetNote(noteIndex uint64) (Note, error) {
	var note Note
	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("Notebook")).Cursor()

		prefix := []byte(strconv.FormatUint(noteIndex, 10))
		for key, value := bucket.Seek(prefix); key != nil && bytes.HasPrefix(key, prefix); key, value = bucket.Next() {
			return json.Unmarshal(value, &note)
		}
		return nil
	})
	return note, err
}
