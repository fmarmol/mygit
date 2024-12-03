package main

import (
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
		head, err := r.Head()
		if err != nil {
			return err
		}
		refSpec := head.Name().String() + ":" + head.Name().String()
		err = r.Push(&git.PushOptions{
			RemoteName: "origin",
			RefSpecs:   []config.RefSpec{config.RefSpec(refSpec)},
			Auth:       nil,
		})
		if err != nil {
			return err
		}
		return nil
	},
}
