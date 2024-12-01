package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var cmdAddFiles = &cobra.Command{
	Use: "add",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("add expect at least one file")
		}
		r, err := NewRepo()
		if err != nil {
			return err
		}
		t, err := r.Worktree()
		if err != nil {
			return err
		}
		for _, arg := range args {
			err := t.AddGlob(arg)
			if err != nil {
				return err
			}
		}
		return nil
	},
}
