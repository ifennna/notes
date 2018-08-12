package cmd

import "github.com/spf13/cobra"

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
func Execute() error {
	return root.Execute()
}
