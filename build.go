package main

import (
	"os"
	"os/user"

	"github.com/spf13/cobra"
)

var cmdBuild = &cobra.Command{
	Use:   "build [binary]",
	Short: "Build a Hermitux kernel for a given binary",
	Args:  cobra.MinimumNArgs(1),
	Run:   build,
}

func build(cmd *cobra.Command, args []string) {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	dir := usr.HomeDir + "/.recluse"
	_, err = os.Stat(dir)
	if !os.IsNotExist(err) {
		panic(err)
	}
	err = os.MkdirAll(dir+"/bin", os.ModePerm)
	err = os.MkdirAll(dir+"/lib", os.ModePerm)
	if err != nil {
		panic(err)
	}
}
