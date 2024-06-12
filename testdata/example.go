package main

import (
	"github.com/spf13/cobra"
)

var help bool

func SomeCommand() {
	myCmd := &cobra.Command{
		Use:   "fake",
		Short: "this is a fake command",
		Run: func(cmd *cobra.Command, _ []string) {
			cmd.Println("Hello, World!")
		},
	}

	myCmd.Flags().BoolP("verbose", "V", false, "verbose output")
	myCmd.Flags().BoolVarP(&help, "help", "H", false, "show help")
	myCmd.Flags().StringP("file", "f", "a_file.txt", "some filename")
	myCmd.Flags().IntP("version-number", "v", 1, "some version number")
}
