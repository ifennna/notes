package cmd

import (
	"fmt"
	"log"
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
			noteID := args[0]
			noteid, err := strconv.ParseInt(noteID, 10, 64)
			if err != nil {
				fmt.Println("ERROR")
			}
			err = db.DeleteNote("Default", uint64(noteid))
			if err != nil {
				log.Panic()
			}
			emoji.Println(" :pencil2: Note deleted from Default")
		default:
			noteIDs := args[1:]
			var noteids []uint64
			for _, noteID := range noteIDs {
				noteid, err := strconv.ParseInt(noteID, 10, 64)
				if err != nil {
					fmt.Println("ERROR")
				}
				noteids = append(noteids, uint64(noteid))
			}
			err := db.DeleteNote(args[0], noteids...)
			if err != nil {
				log.Panic()
			}
			emoji.Println(" :pencil2: Note(s) deleted from " + args[0])
		}

	},
}

func init() {
	root.AddCommand(deleteCommand)
}
