package models

import (
	"encoding/json"
	"github.com/boltdb/bolt"
	"fmt"
	"bytes"
	"strings"
	"log"
)

type Notebook struct {
	Name  string  `json:"name"`
	Notes []Note  `json:"notes"`
}

func (db *DB) AddNotebook(notebook Notebook) (error) {
	encoded, err := json.Marshal(notebook)
	if err != nil {
		return err
	}
	err = db.Update(func(tx *bolt.Tx) error {
		err = tx.Bucket([]byte("Notebook")).Put([]byte(notebook.Name), encoded)
		if err != nil {
			return fmt.Errorf("could not set config: %v", err)
		}
		return nil
	})
	return err
}

func (db *DB) GetNotebook(notebookTitle string) (Notebook, error) {
	var notebook Notebook
	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("Notebook"))
		cursor := bucket.Cursor()
		prefix := []byte(notebookTitle)
		for key, _ := cursor.Seek(prefix); key != nil && bytes.HasPrefix(key, prefix); key, _ = cursor.Next() {
			var notes []Note
			notes = getNotesInNotebook(bucket, key, notes)
			notebook.Name = string(key)
			notebook.Notes = notes
		}
		return nil
	})
	return notebook, err
}

func (db *DB) GetAllNotebooks() ([]Notebook, error) {
	var notebooks []Notebook
	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("Notebook"))
		cursor := bucket.Cursor()
		notebooks = getNotebooksInRootBucket(cursor, bucket, notebooks)
		return nil
	})
	return notebooks, err
}

func getNotebooksInRootBucket(cursor *bolt.Cursor, bucket *bolt.Bucket, notebooks []Notebook) []Notebook {
	for key, _ := cursor.First(); key != nil; key, _ = cursor.Next() {
		var notebook Notebook
		var notes []Note
		notes = getNotesInNotebook(bucket, key, notes)
		notebook.Name = string(key)
		notebook.Notes = notes
		notebooks = append(notebooks, notebook)
	}
	return notebooks
}

func getNotesInNotebook(bucket *bolt.Bucket, key []byte, notes []Note) []Note {
	nestedBucketCursor := bucket.Bucket([]byte(key)).Cursor()
	for key, value := nestedBucketCursor.First(); key != nil; key, value = nestedBucketCursor.Next() {
		var note Note
		json.Unmarshal(value, &note)
		notes = append(notes, note)
	}
	return notes
}

func (db *DB) Dump() () {
	err := db.View(func(tx *bolt.Tx) error {
		c := tx.Cursor()
		dumpCursor(tx, c, 0)
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
}

func dumpCursor(tx *bolt.Tx, c *bolt.Cursor, indent int) {
	for k, v := c.First(); k != nil; k, v = c.Next() {
		if v == nil {
			fmt.Printf(strings.Repeat("\t", indent)+"[%s]\n", k)
			newBucket := c.Bucket().Bucket(k)
			if newBucket == nil {
				newBucket = tx.Bucket(k)
			}
			newCursor := newBucket.Cursor()
			dumpCursor(tx, newCursor, indent+1)
		} else {
			fmt.Printf(strings.Repeat("\t", indent)+"%s\n", k)
			fmt.Printf(strings.Repeat("\t", indent+1)+"%s\n", v)
		}
	}
}