package models

import (
	"encoding/json"
	"github.com/boltdb/bolt"
		"bytes"
	"strconv"
	)

type Note struct {
	//Title   string `json:"title"`
	Id      uint64 `json:"id"`
	Content string `json:"content"`
}

func (db *DB) AddNote(notebookName string, note Note) (error) {
	tx, err := db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	notebook, err := tx.Bucket([]byte("Notebook")).CreateBucketIfNotExists([]byte(notebookName))

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