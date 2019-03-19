package cmd

import (
	"bfk/interpreter"
	"fmt"
	"io/ioutil"

	"github.com/spf13/cobra"
)

var minifyCmd = &cobra.Command{
	Use:   "minify <file>",
	Short: "Minifies a program.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		in, err := ioutil.ReadFile(args[0])
		if err != nil {
			fmt.Println("Error reading file")
			fmt.Println(err)
			return
		}

		input := string(in)

		fmt.Println(string(interpreter.Tokenize(input)))
	},
}

func init() {
	rootCmd.AddCommand(minifyCmd)
}
