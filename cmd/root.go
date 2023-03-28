package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var root = &cobra.Command{
	Use:           "notes",
	Short:         "Jot things down quickly from the command line",
	SilenceErrors: true,
	SilenceUsage:  true,
}

// Execute runs the main command
func Execute() {
	if err := root.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
