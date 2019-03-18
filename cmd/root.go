package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "bfk",
	Short:   "bfk is a brainfuck interpreter written in go.",
	Version: "0.0.1",
}

// Execute runs the CLI logic.
func Execute() {
	rootCmd.Execute()
}
