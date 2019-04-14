package main

import (
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use: "recluse",
	}
	rootCmd.AddCommand(cmdRun)
	rootCmd.AddCommand(cmdBuild)
	rootCmd.Execute()
}
