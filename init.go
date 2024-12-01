package main

import (
	"github.com/go-git/go-git/v5"
	"github.com/spf13/cobra"
)

var cmdInit = &cobra.Command{
	Use: "init",
	RunE: func(cmd *cobra.Command, args []string) error {
		_, err := git.PlainInit(".", false)
		return err
	},
}
