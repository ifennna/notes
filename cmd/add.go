package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"gopkg.in/kyokomi/emoji.v1"
)

var addCommand = &cobra.Command{
	Use:   "add",
	Short: "Adds notes",
	Long: "Adds notes to a notebook from the terminal. Use `notes add \"text\"` to jot in the default notebook or" +
		"`notes add NotebookName \"text-1\" \"text-2\" ..` to add notes to other notebooks",
	Run: func(cmd *cobra.Command, args []string) {
		db := setupDatabase()
		var err error

		switch len(args) {
		case 0:
			emoji.Println(" :warning: You need to add some text")
		case 1:
			err = db.AddNotes("Default", args[0])
			emoji.Println(" :pencil2: Note added to 'Default' Notebook")
		default:
			err = db.AddNotes(args[0], args[1:]...)
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
