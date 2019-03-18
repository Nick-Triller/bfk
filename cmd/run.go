package cmd

import (
	"bfk/interpreter"
	"fmt"
	"io/ioutil"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run <file>",
	Short: "Run a program.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		in, err := ioutil.ReadFile(args[0])
		if err != nil {
			fmt.Println("Error reading file")
			fmt.Println(err)
			return
		}

		input := string(in)

		interpreter.Execute(input)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
