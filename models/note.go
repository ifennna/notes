package models

import (
	"bytes"
	"encoding/json"
	"strconv"

	"github.com/boltdb/bolt"
)

type Note struct {
	// TODO: explore allowing 'naming' notes within a notebook
	//Title   string `json:"title"`
	Id      uint64 `json:"id"`
	Content string `json:"content"`
}

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

func (db *DB) DeleteNote(notebookName string, noteIDs ...uint64) error {
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
