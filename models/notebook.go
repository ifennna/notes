package models

import (
	"encoding/json"
	"github.com/boltdb/bolt"
	"fmt"
	"bytes"
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
		bucket := tx.Bucket([]byte("Notebook")).Cursor()

		prefix := []byte(notebookTitle)
		for k, v := bucket.Seek(prefix); k != nil && bytes.HasPrefix(k, prefix); k, v = bucket.Next() {
			return json.Unmarshal(v, &notebook)
		}
		return nil
	})
	return notebook, err
}
