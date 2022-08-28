package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	fmt.Println("----", "root init add wordcmd")

	rootCmd.AddCommand(wordCmd)
}
