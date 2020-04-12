package cmd

import (
	"fmt"
	"github.com/noculture/notes/models"
	"github.com/noculture/notes/utils"
	"github.com/spf13/cobra"
	"gopkg.in/kyokomi/emoji.v1"
	"log"
)

var deleteCommand = &cobra.Command{
	Use:   "del",
	Short: "Delete notes",
	Long: "Deletes notes from the terminal. Use `notes del noteId` to delete a note from the Default notebook" +
		"`notes del NotebookName noteId-1 noteId-2 ..` to delete notes from a specific notebook",
	Run: func(cmd *cobra.Command, args []string) {
		db := setupDatabase()

		switch len(args) {
		case 0:
			emoji.Println(" :warning: You need to specify a note to delete ")
		case 1:
			usNoteId, _ := utils.ParseUInt64(args[0])
			deleteNotesIfExist(db, "Default", usNoteId)
		default:
			notebookName := args[0]
			usNoteIds, _ := utils.ParseUInt64Slice(args[1:])
			deleteNotesIfExist(db, notebookName, usNoteIds...)
		}
	},
}

/**
 * Deletes notes with given ids from the given notebook if they exist
 * Displays appropriate messages whether or not note exists
 *
 * param: models.Datastore db
 * param: string           notebookName
 * param: ...uint64        noteIds
 */
func deleteNotesIfExist(db models.Datastore, notebookName string, noteIds ...uint64) {
	for _, noteId := range noteIds {
		noteExists, _ := db.NoteExists(notebookName, noteId)
		if noteExists {
			err := db.DeleteNotes(notebookName, noteId)
			if err != nil {
				log.Panic()
			} else {
				emoji.Println(fmt.Sprintf(" :pencil2: Note with noteId %d deleted from notebook %s", noteId, notebookName))
			}
		} else {
			emoji.Println(fmt.Sprintf(" :pencil2: Note with noteId %d does not exist in notebook %s", noteId, notebookName))
		}
	}
}

func init() {
	root.AddCommand(deleteCommand)
}
