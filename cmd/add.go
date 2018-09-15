package cmd

import (
	"github.com/spf13/cobra"
		"github.com/uncultured/notes/models"
	"log"
	"fmt"
)

var addCommand = &cobra.Command{
	Use: "add",
	Short: "Add a note",
	Long: "Load up your favourite editor and jot something down",
	Run: func(cmd *cobra.Command, args []string) {
		db := setupDatabase()
		var err error

		switch len(args) {
		case 0:
			fmt.Print("You need to add a note")
		case 1:
			err = db.AddNote("Default", models.Note{Content:args[0]})
		default:
			err = db.AddNote(args[0], models.Note{Content:args[1]})
		}
		if err != nil{
			log.Panic()
		}
	},
}

func init()  {
	root.AddCommand(addCommand)
}
