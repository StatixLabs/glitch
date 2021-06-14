package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "glitch",
	Short: "Glitch is a multi use tool that lets you do stuff.",
	Long:  `Glitch is a tool created to make development easier. Its a cli that does random things.`,
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
