package main

import (
	"fmt"
	"log"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/spf13/cobra"
)

var cmdPush = &cobra.Command{
	Use: "push",
	RunE: func(cmd *cobra.Command, args []string) error {
		r, err := NewRepo()
		if err != nil {
			return err
		}
		return nil
		refSpec := fmt.Sprintf("refs/heads/%s", "main") // or any branch name
		err = r.Push(&git.PushOptions{
			RemoteName: "origin",
			RefSpecs:   []config.RefSpec{config.RefSpec(refSpec)},
			Auth:       nil, // Optional: Use for authentication, e.g., SSH keys or credentials.
		})
		if err != nil {
			log.Fatalf("Failed to push to remote: %v", err)
		}
		return nil
	},
}
