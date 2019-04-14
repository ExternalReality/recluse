package main

import (
	"github.com/spf13/cobra"
)

var cmdBuild = &cobra.Command{
	Use:   "build [binary]",
	Short: "Build a Hermitux kernel for a given binary",
	Args:  cobra.MinimumNArgs(1),
	Run:   run,
}

func build(cmd *cobra.Command, args []string) {

}
