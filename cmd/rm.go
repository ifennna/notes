package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gopkg.in/kyokomi/emoji.v1"
	"log"
)

var removeCommand = &cobra.Command{
	Use:   "rm",
	Short: "Remove a notebook",
	Long:  "Removes a notebook from the terminal. Use `notes rm NotebookName` to delete the specific Notebook",
	Run: func(cmd *cobra.Command, args []string) {
		db := setupDatabase()

		switch len(args) {
		case 0:
			emoji.Println(" :warning: You need to specify a notebook to remove ")
		case 1:
			// retrieve notebook name from CLI args
			notebookName := args[0]
			// check if notebook exists
			notebookExists, _ := db.NotebookExists(notebookName)
			if notebookExists {
				// if notebook exists, try deleting it
				err := db.RmNotebook(notebookName)
				if err != nil {
					log.Panic()
				}
				emoji.Println(fmt.Sprintf(" :pencil2: Notebook '%s' deleted", notebookName))
			} else {
				// if notebook doesn't exist, display warning message
				emoji.Println(fmt.Sprintf(" :warning: Notebook '%s' does not exist", notebookName))
			}
		}

	},
}

func init() {
	root.AddCommand(removeCommand)
}
