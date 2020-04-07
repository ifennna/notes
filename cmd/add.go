package cmd

import (
	"log"

	"github.com/noculture/notes/models"
	"github.com/spf13/cobra"
	"gopkg.in/kyokomi/emoji.v1"
)

var addCommand = &cobra.Command{
	Use:   "add",
	Short: "Add a note",
	Long: "Add a note to a notebook from the terminal. Use `notes add \"text\"` to jot in the default notebook or" +
		"`notes add NotebookName \"text\"` to add notes to other notebooks",
	Run: func(cmd *cobra.Command, args []string) {
		db := setupDatabase()
		var err error

		switch len(args) {
		case 0:
			emoji.Println(" :warning: You need to add some text")
		case 1:
			err = db.AddNote("Default", models.Note{Content: args[0]})
			emoji.Println(" :pencil2: Note added")
		default:
			var notes []models.Note
			for _, note := range args[1:] {
				notes = append(notes, models.Note{Content: note})
			}
			err = db.AddNote(args[0], notes...)
			emoji.Println(" :pencil2: Note(s) added")
		}
		if err != nil {
			log.Panic()
		}

	},
}

func init() {
	root.AddCommand(addCommand)
}
