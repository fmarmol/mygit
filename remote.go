package main

import (
	"fmt"

	"github.com/go-git/go-git/v5/config"
	"github.com/spf13/cobra"
)

var cmdRemote = &cobra.Command{
	Use: "remote",
}

var cmdAddRemote = &cobra.Command{
	Use: "add",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("remote add expects an url to set the remote origin")
		}
		url := args[0]
		r, err := NewRepo()
		if err != nil {
			return err
		}
		_, err = r.CreateRemote(&config.RemoteConfig{Name: "origin", URLs: []string{url}})
		return err
	},
}
