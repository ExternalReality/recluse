package main

import (
	"github.com/spf13/cobra"
)

var cmdRun = &cobra.Command{
	Use:   "run [binary]",
	Short: "Run an executable under Hermitux",
	Args:  cobra.MinimumNArgs(1),
	Run:   run,
}

func run(cmd *cobra.Command, args []string) {

}
