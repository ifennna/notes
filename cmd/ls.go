package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/noculture/notes/models"
	"github.com/spf13/cobra"
	"gopkg.in/kyokomi/emoji.v1"
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

func getSpecificNotebook(db models.Datastore, notebookName string) {
	notebookExists, _ := db.NotebookExists(notebookName)
	if notebookExists {
		notebook, err := db.GetNotebook(notebookName)
		if err != nil {
			log.Panic()
		}
		emoji.Println(" :notebook_with_decorative_cover: " + notebook.Name)
		for _, note := range notebook.Notes {
			emoji.Println("     " + strconv.FormatUint(note.ID, 10) + "| " + note.Content)
		}
	} else {
		emoji.Println(fmt.Sprintf(" :warning: Noteebook '%s' doesn't exist", notebookName))
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
