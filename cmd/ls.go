package cmd

import (
	"github.com/noculture/notes/models"
	"github.com/spf13/cobra"
	"gopkg.in/kyokomi/emoji.v1"
	"log"
	"strconv"
)

var lsCommand = &cobra.Command{
	Use:   "ls <notebook>",
	Short: "List stuff",
	Long: "Show a list of notes or notebooks. `notes ls` will show a list of your notebooks and `notes ls NotebookName` " +
		"will show a list of notes in the notebook you've specified",
	Run: func(cmd *cobra.Command, args []string) {
		db := setupDatabase()

		switch len(args) {
		case 0:
			printAllNotebooks(db)
		default:
			getSpecificNotebook(db, args[0])
		}

		//db.Dump()
	},
}

func getSpecificNotebook(db models.Datastore, notebookTitle string) {
	notebook, err := db.GetNotebook(notebookTitle)
	if err != nil {
		log.Panic()
	}
	emoji.Println(notebook.Name)
	for _, note := range notebook.Notes {
		emoji.Println(" " + strconv.FormatUint(note.Id, 10) + "	" + note.Content)
	}
}

func printAllNotebooks(db models.Datastore) {
	notebookNames, err := db.GetAllNotebookNames()
	if err != nil {
		log.Panic()
	}
	for _, notebookName := range notebookNames {
		emoji.Println(" :notebook_with_decorative_cover: " + notebookName)
	}
}

func init() {
	root.AddCommand(lsCommand)
}
