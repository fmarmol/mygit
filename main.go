package main

import (
	"log"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/spf13/cobra"
)

func raise(err error) {
	if err != nil {
		panic(err)
	}
}

var cmdRoot = &cobra.Command{}

var cmdBranches = &cobra.Command{
	Use: "branches",
	RunE: func(cmd *cobra.Command, args []string) error {
		path, err := os.Getwd()
		if err != nil {
			return err
		}
		r, err := git.PlainOpen(path)
		if err != nil {
			return err
		}
		branches, err := r.Branches()
		if err != nil {
			return err
		}
		branches.ForEach(func(ref *plumbing.Reference) error {
			log.Println(ref)
			return nil
		})
		return nil
	},
}

func main() {
	cmdRoot.AddCommand(cmdBranches)
	err := cmdRoot.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
