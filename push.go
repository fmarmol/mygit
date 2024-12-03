package main

import (
	"fmt"
	"path/filepath"

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
		// remote, err := r.Remote("origin")
		// if err != nil {
		// 	return nil
		// }
		// for _, ref := range remote.Config().Fetch {
		// 	if ref.Match(head.Name()) {
		// 		fmt.Println(ref, head.Name().String())
		// 	}

		// }
		// return nil
		// fmt.Println("HEAD TARGET:", filepath.Base(head.Name().String()), "END")
		branchName := filepath.Base(head.Name().String())
		_ = branchName
		// refSpec := head.Name().String() + ":refs/remotes/origin/" + filepath.Base(head.Name().String())
		refSpec := head.Name().String() + ":" + head.Name().String()
		// _ = refSpec
		fmt.Println(refSpec)
		err = r.Push(&git.PushOptions{
			RemoteName: "origin",
			RefSpecs:   []config.RefSpec{config.RefSpec(refSpec)},
			Auth:       nil, // Optional: Use for authentication, e.g., SSH keys or credentials.
		})
		if err != nil {
			return err
		}
		return nil
	},
}
