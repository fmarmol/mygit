package main

import (
	"fmt"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/spf13/cobra"
)

var cmdCommit = &cobra.Command{
	Use: "commit",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("commit command expect one commit message")
		}
		r, err := NewRepo()
		if err != nil {
			return err
		}
		t, err := r.Worktree()
		if err != nil {
			return err
		}
		config, err := r.Config()
		if err != nil {
			return err
		}
		if config.User.Name == "" {
			return fmt.Errorf("user's name is not defined")
		}
		if config.User.Email == "" {
			return fmt.Errorf("user's email is not defined")
		}

		hash, err := t.Commit(args[0], &git.CommitOptions{
			Author: &object.Signature{
				Name:  config.User.Name,
				Email: config.User.Email,
				When:  time.Now(),
			},
		})
		if err != nil {
			return err
		}
		fmt.Println("new commit:", hash)
		return nil
	},
}
