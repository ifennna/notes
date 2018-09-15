package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"fmt"
)

var lsCommand = &cobra.Command{
	Use: "ls",
	Short: "List stuff",
	Long: "Show a list of notes or notebooks",
	Run: func(cmd *cobra.Command, args []string) {
		db := setupDatabase()
		notebooks, err := db.GetAllNotebooks()
		if err != nil{
			log.Panic()
		}
		for _, n := range notebooks {
			fmt.Printf(n.Name + "\n")
		}
		//db.Dump()
	},
}

func init()  {
	root.AddCommand(lsCommand)
}

