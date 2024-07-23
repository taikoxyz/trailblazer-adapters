package main

import (
	"github.com/spf13/cobra"
	"github.com/taikoxyz/trailblazer-adapters/cmd"
)

func main() {
	var rootCmd = &cobra.Command{Use: "trailblazer-adapters"}
	rootCmd.AddCommand(cmd.ProcessCmd)
	rootCmd.Execute()
}
