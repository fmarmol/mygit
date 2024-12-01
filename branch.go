package main

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/spf13/cobra"
)

func FindBranch(r *git.Repository, branch string) (*plumbing.Reference, error) {
	branches, err := r.Branches()
	if err != nil {
		return nil, err
	}
	var refFound *plumbing.Reference
	branches.ForEach(func(ref *plumbing.Reference) error {
		if strings.Contains(ref.Name().String(), branch) {
			refFound = ref
			return nil
		}
		if strings.Contains(ref.Hash().String(), branch) {
			refFound = ref
			return nil
		}
		return nil
	})
	if refFound == nil {
		return nil, fmt.Errorf("branch %v not found", branch)
	}
	ref, err := r.Reference(refFound.Name(), true)
	if err != nil {
		return nil, err
	}
	return ref, nil

}

var cmdDeleteBranch = &cobra.Command{
	Use: "delete",
	RunE: func(cmd *cobra.Command, args []string) error {
		r, err := NewRepo()
		if err != nil {
			return err
		}
		branch, err := FindBranch(r, args[0])
		if err != nil {
			return err
		}
		return r.Storer.RemoveReference(branch.Name())
	},
}

var cmdCreateBranch = &cobra.Command{
	Use: "create",
	RunE: func(cmd *cobra.Command, args []string) error {
		r, err := NewRepo()
		if err != nil {
			return err
		}
		tree, err := r.Worktree()
		if err != nil {
			return err
		}
		head, err := r.Head()
		if err != nil {
			return err
		}

		branchName := func() string {
			b := args[0]
			if strings.HasPrefix(b, "refs/heads/") {
				return b
			} else {
				return filepath.Join("refs/heads", b)
			}
		}()

		err = tree.Checkout(&git.CheckoutOptions{
			Hash:   head.Hash(),
			Branch: plumbing.ReferenceName(branchName),
			Create: true,
			Keep:   true,
		})
		return err
	},
}

var cmdBranch = &cobra.Command{
	Use: "branch",
	RunE: func(cmd *cobra.Command, args []string) error {
		r, err := NewRepo()
		if err != nil {
			return err
		}
		t, err := r.Worktree()
		if err != nil {
			return err
		}
		b, err := FindBranch(r, args[0])
		if err != nil {
			return err
		}
		return t.Checkout(&git.CheckoutOptions{Branch: b.Name()})
	},
}

var cmdBranches = &cobra.Command{
	Use:   "branches",
	Short: "list branches",
	RunE: func(cmd *cobra.Command, args []string) error {
		r, err := NewRepo()
		if err != nil {
			return err
		}
		branches, err := r.Branches()
		if err != nil {
			return err
		}
		branches.ForEach(func(ref *plumbing.Reference) error {
			fmt.Println(ref.Hash(), ref.Name())
			return nil
		})
		remotesFlag, err := cmd.Flags().GetBool("remotes")
		if err != nil {
			return err
		}
		if !remotesFlag {
			return nil
		}
		remotes, err := r.Remotes()
		if err != nil {
			return err
		}
		for _, remote := range remotes {
			refs, err := remote.List(&git.ListOptions{})
			if err != nil {
				return err
			}
			for _, ref := range refs {
				if !ref.Hash().IsZero() {
					fmt.Println(remote.Config().Name, ref.Hash(), ref.Name())
				}
			}
		}
		return nil
	},
}
