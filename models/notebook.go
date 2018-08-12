package models

type Notebook struct {
	Name string
	Notes []Note
}

func (db *DB) AddNote(note Note) (error) {
	// _ ,err := db.Query("INSERT INTO notes")
	//if err != nil {
	//	return err
	//}
	 return nil
}
