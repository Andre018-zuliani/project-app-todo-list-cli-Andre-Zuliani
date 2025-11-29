package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "todo",
    Short: "A simple CLI TODO app using Cobra",
    Long:  "To-Do List CLI dengan fitur add, list, done, delete & search",
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
    }
}
