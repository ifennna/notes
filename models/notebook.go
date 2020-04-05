package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"strings"
)

/**
 * DTO for Notebook
 */
type Notebook struct {
	Name  string `json:"name"`
	Notes []Note `json:"notes"`
}

/**
 * Adds a notebook in db
 * - Puts key-value pair into 'Notebook' bucket of db
 *    - Key: Name of notebook
 *    - Value: marshalled JSON blob (bytes) of Notebook object
 */
func (db *DB) AddNotebook(notebook Notebook) error {
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

/**
 * Retrieves Notebook for given 'notebookTitle' name
 *  - uses cursor.Seek(..) to seek through keys (notebook-names) and find matching key
 *  - then uses getNotesInNotebook() call to retrieve notes of that notebook
 */
func (db *DB) GetNotebook(reqName string) (Notebook, error) {
	var notebook Notebook
	notebook.Name = reqName
	err := db.View(func(tx *bolt.Tx) error {
		// TODO: play with strings instead of bytes
		reqNameBytes := []byte(reqName)

		bucket := tx.Bucket([]byte("Notebook"))
		foundNameBytes, _ := bucket.Cursor().Seek(reqNameBytes)
		if foundNameBytes != nil && bytes.Equal(reqNameBytes, foundNameBytes) {
			notebook.Notes = getNotesInNotebook(bucket, reqNameBytes)
		}

		return nil
	})
	return notebook, err
}

/**
 * Retrieves all notebooks (along with their notes)
 *  - deletegates actual work to 'getNotebooksInRootBucket' function
 * two better alternatives
 *  - in current implementation GetNotebook() should use this call
 *  - (better) modify this method to 'GetAllNotebookNames'
 */
func (db *DB) GetAllNotebooks() ([]Notebook, error) {
	var notebooks []Notebook
	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("Notebook"))
		notebooks = getNotebooksInRootBucket(bucket.Cursor(), bucket)

		return nil
	})
	return notebooks, err
}

/**
 * Function wrapping the core logic of 'GetAllNotebooks'
 *  - iterates over notebooks in bucket
 *  - uses getNotesInNotebook() call to retrieve notes of notebooks
 * param: *bolt.Cursor cursor
 * param: *bolt.Bucket bucket
 * return: []Notebook
 */
func getNotebooksInRootBucket(cursor *bolt.Cursor, bucket *bolt.Bucket) []Notebook {
	var notebooks []Notebook
	for notebookNameBytes, _ := cursor.First(); notebookNameBytes != nil; notebookNameBytes, _ = cursor.Next() {
		var notebook Notebook
		notebook.Name = string(notebookNameBytes)
		// TODO: remove notes retrieval from here as an optimization
		notebook.Notes = getNotesInNotebook(bucket, notebookNameBytes)
		notebooks = append(notebooks, notebook)
	}
	return notebooks
}

/**
 * Retrieves notes in a given notebook
 * param: *bolt.Bucket bucket
 * param: []byte       notebookNameBytes
 * return: []Note
 */
func getNotesInNotebook(bucket *bolt.Bucket, notebookNameBytes []byte) []Note {
	var notes []Note
	nestedBucketCursor := bucket.Bucket([]byte(notebookNameBytes)).Cursor()
	// TODO: leverage this thing and allow 'naming' notes within a notebook
	for noteNameBytes, noteBodyBytes := nestedBucketCursor.First(); noteNameBytes != nil; noteNameBytes, noteBodyBytes = nestedBucketCursor.Next() {
		var note Note
		json.Unmarshal(noteBodyBytes, &note)
		notes = append(notes, note)
	}
	return notes
}

/**
 * Prints all data of all notebooks
 * Output can be piped into a file for persisting and sharing
 *  - delegates actual work to 'dumpCursor' method
 * TODO: add CLI interface to invoke this function
 */
func (db *DB) Dump() {
	err := db.View(func(tx *bolt.Tx) error {
		c := tx.Cursor()
		dumpCursor(tx, c, 0)
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
}

/**
 * Recursively prints all data of given cursor
 * param: *bolt.Tx     tx
 * param: *bolt.Cursor c
 * param: int          indent
 */
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
