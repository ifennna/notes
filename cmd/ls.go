package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"gopkg.in/kyokomi/emoji.v1"
	"github.com/uncultured/notes/models"
	"strconv"
)

var lsCommand = &cobra.Command{
	Use: "ls <notebook>",
	Short: "List stuff",
	Long: "Show a list of notes or notebooks",
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
	notebooks, err := db.GetAllNotebooks()
	if err != nil {
		log.Panic()
	}
	for _, n := range notebooks {
		emoji.Println(" :notebook_with_decorative_cover: " + n.Name)
	}
}

func init()  {
	root.AddCommand(lsCommand)
}

