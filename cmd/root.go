package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var root = &cobra.Command{
	Use:           "notes",
	Short:         "Jot things down quickly from the command line",
	SilenceErrors: true,
	SilenceUsage:  true,
}

// Register adds a new command
func Register(cmd *cobra.Command) {
	root.AddCommand(cmd)
}

// Execute runs the main command
func Execute() {
	if err := root.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
