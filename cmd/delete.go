package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"gopkg.in/kyokomi/emoji.v1"
)

var deleteCommand = &cobra.Command{
	Use:   "del",
	Short: "Delete a note",
	Long: "Deletes a note from the terminal. Use `notes del Note ID` to delete from the Default" +
		"notebook or `notes del NotebookName Note ID` to delete from a specific notebook",
	Run: func(cmd *cobra.Command, args []string) {
		db := setupDatabase()

		switch len(args) {
		case 0:
			emoji.Println(" :warning: You need to specify a note to delete ")
		case 1:
			noteID := os.Args[2]
			noteid, err := strconv.ParseInt(noteID, 10, 64)
			if err != nil {
				fmt.Println("ERROR")
			}
			err = db.DeleteNote("Default", uint64(noteid))
			fmt.Println(err)
			emoji.Println(" :pencil2: Note deleted from Default")

		}

	},
}

func init() {
	root.AddCommand(deleteCommand)
}
