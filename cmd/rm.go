package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"gopkg.in/kyokomi/emoji.v1"
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
			err := db.RmNotebook(args[0])
			if err != nil {
				log.Panic()
			}
			emoji.Println(" :pencil2: Notebook deleted")
		}

	},
}

func init() {
	root.AddCommand(removeCommand)
}
