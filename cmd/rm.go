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
			notebookName := args[0]
			notebookExists, _ := db.NotebookExists(notebookName)
			if notebookExists {
				err := db.RmNotebook(notebookName)
				if err != nil {
					log.Panic()
				}
				emoji.Println(" :pencil2: Notebook deleted")
			} else {
				emoji.Println(fmt.Sprintf(" :warning: Notebook '%s' does not exist", notebookName))
			}
		}

	},
}

func init() {
	root.AddCommand(removeCommand)
}
