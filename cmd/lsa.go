package cmd

import (
	"log"

	"github.com/noculture/notes/models"
	"github.com/spf13/cobra"
)

var lsaCommand = &cobra.Command{
	Use:   "lsa",
	Short: "List all notes in all notebooks",
	Long: "Show a list of notes and notebooks as a tree. `notes lsa` will show a list of your notes within all notebooks" +
		"represented as a tree",
	Run: func(cmd *cobra.Command, args []string) {
		db := setupDatabase()
		printAllNotesNotebooks(db)
		// db.Dump()
	},
}

func printAllNotesNotebooks(db models.Datastore) {
	notebookNames, err := db.GetAllNotebookNames()
	if err != nil {
		log.Panic()
	}
	for _, notebookName := range notebookNames {
		getSpecificNotebook(db, notebookName)
	}
}

func init() {
	root.AddCommand(lsaCommand)
}
