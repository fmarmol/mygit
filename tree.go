package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/spf13/cobra"
)

var cmdWorkTree = &cobra.Command{
	Use: "tree",
	RunE: func(cmd *cobra.Command, args []string) error {
		path, err := os.Getwd()
		if err != nil {
			return err
		}
		r, err := git.PlainOpen(path)
		if err != nil {
			return err
		}
		tree, err := r.Worktree()
		if err != nil {
			return err
		}
		log.Println("PATH:", path)
		infos, err := tree.Filesystem.ReadDir(".")

		if err != nil {
			return err
		}

		for _, info := range infos {
			fmt.Println(info.Name(), info.IsDir())
		}
		return nil
	},
}
