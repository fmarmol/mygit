package main

import (
	"log"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/spf13/cobra"
)

func raise(err error) {
	if err != nil {
		panic(err)
	}
}
func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func NewRepo() (*git.Repository, error) {
	path, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	r, err := git.PlainOpen(path)
	if err != nil {
		return nil, err
	}
	return r, nil
}

var cmdRoot = &cobra.Command{}

func main() {
	cmdBranch.AddCommand(cmdCreateBranch, cmdDeleteBranch)
	cmdRemote.AddCommand(cmdAddRemote)

	//CONFIG
	cmdSetConfig.AddCommand(cmdSetConfigName)
	cmdSetConfig.AddCommand(cmdSetConfigEmail)
	cmdConfig.AddCommand(cmdSetConfig)

	flags := cmdBranches.Flags()
	flags.BoolP("remotes", "r", false, "display remotes branches")

	cmdRoot.AddCommand(cmdPush)
	cmdRoot.AddCommand(cmdConfig)
	cmdRoot.AddCommand(cmdCommit)
	cmdRoot.AddCommand(cmdAddFiles)
	cmdRoot.AddCommand(cmdInit)
	cmdRoot.AddCommand(cmdBranches)
	cmdRoot.AddCommand(cmdWorkTree)
	cmdRoot.AddCommand(cmdBranch)
	cmdRoot.AddCommand(cmdRemote)
	err := cmdRoot.Execute()
	if err != nil {
		os.Exit(1)
	}
}
