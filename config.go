package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var cmdConfig = &cobra.Command{
	Use: "config",
}

var cmdSetConfig = &cobra.Command{
	Use: "set",
}

var cmdSetConfigName = &cobra.Command{
	Use: "name",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("command config set name expects one argument")
		}
		r, err := NewRepo()
		if err != nil {
			return err
		}
		cfg, err := r.Config()
		if err != nil {
			return nil
		}
		cfg.User.Name = args[0]
		return r.SetConfig(cfg)
	},
}
var cmdSetConfigEmail = &cobra.Command{
	Use: "email",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("command config set email expects one argument")
		}
		r, err := NewRepo()
		if err != nil {
			return err
		}
		cfg, err := r.Config()
		if err != nil {
			return nil
		}
		cfg.User.Email = args[0]
		return r.SetConfig(cfg)
	},
}
