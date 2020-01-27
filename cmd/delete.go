package cmd

import (
	"github.com/spf13/cobra"
	"gopkg.in/kyokomi/emoji.v1"
	"log"
	"strconv"
)

var deleteCommand = &cobra.Command{
	Use:   "del",
	Short: "Delete a note",
	Long: "Deletes a note from the terminal. Use `notes del Note ID` to delete from the Default" +
	"notebook or `notes del NotebookName Note ID` to delete from a specific notebook",
	Run: func(cmd *cobra.Command, args []string) {
		db := setupDatabase()
		var err error

		switch len(args) {
		case 0:
			emoji.Println(" :warning: You need to specify a note to delete ")
		case 1:
			noteID, err := strconv.ParseInt(args[1], 10, 64)
			if err != nil {
				emoji.Println(" :warning: Please pass valid noteID")
			}
			err = db.DeleteNote("Default", uint64(noteID))
			emoji.Println(" :pencil2: Note deleted from Default")
		default:
			notebookName := args[0]
			noteID, err := strconv.ParseInt(args[1], 10, 64)
			if err != nil {
				emoji.Println(":warning: Please pass valid noteID")
			}
			err = db.DeleteNote(notebookName, uint64(noteID))
			emoji.Println(" :pencil2: Note deleted from ", notebookName)
		}
		if err != nil {
			log.Panic()
		}

	},
}

func init() {
	root.AddCommand(deleteCommand)
}
