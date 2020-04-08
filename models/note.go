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
 * Adds a note in the given notebook
 *  - the passed note is expected to contain only 'Content'
 *  - note's auto-increment 'Id' is generated and stored in the db by this method itself
 * // TODO: accept only 'Content' of Note
 * param: string notebookName
 * param: Note   note
 * return: error
 */
func (db *DB) AddNote(notebookName string, notes ...Note) error {
	tx, err := db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	notebook, err := tx.Bucket([]byte("Notebook")).CreateBucketIfNotExists([]byte(notebookName))

	for _, note := range notes {
		noteID, err := notebook.NextSequence()
		if err != nil {
			return err
		}
		note.Id = noteID
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
 * Deletes note with given id from the given notebook
 * param: string notebookName
 * param: uint64 noteID
 * return: error
 */
func (db *DB) DeleteNote(notebookName string, noteIDs ...uint64) error {
	// TODO: try to remove code-duplication: txn creation & notebook bucket retrieval logic can be extracted out
	tx, err := db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	bucket := tx.Bucket([]byte("Notebook")).Bucket([]byte(notebookName))
	for _, noteID := range noteIDs {
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
 * // TODO: also accept notebookName in input
 * param: uint64 noteId
 * return: (Note, error)
 */
func (db *DB) GetNote(notebookName string, reqNoteId uint64) (Note, error) {
	var note Note
	err := db.View(func(tx *bolt.Tx) error {
		reqNoteIdBytes := []byte(strconv.FormatUint(reqNoteId, 10))

		notebookBucket := tx.Bucket([]byte("Notebook")).Bucket([]byte(notebookName))
		foundNoteIdBytes, foundNoteContentBytes := notebookBucket.Cursor().Seek(reqNoteIdBytes)
		if foundNoteIdBytes != nil && bytes.Equal(reqNoteIdBytes, foundNoteIdBytes) {
			return json.Unmarshal(foundNoteContentBytes, &note)
		}

		return nil
	})
	return note, err
}
