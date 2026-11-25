package main

import (
	// "github.com/go-git/go-git/v5"
	// "github.com/go-git/go-git/v5/plumbing"

	"bytes"
	"fmt"
	"text/tabwriter"

	"github.com/go-git/go-git/v5"
	"github.com/spf13/cobra"
)

var cmdStatus = &cobra.Command{
	Use: "status",
	RunE: func(cmd *cobra.Command, args []string) error {
		r, err := NewRepo()
		if err != nil {
			return err
		}
		tree, err := r.Worktree()
		if err != nil {
			return err
		}
		status, err := tree.Status()
		if err != nil {
			return err
		}

		type FileStatus struct {
			Path   string
			Status *git.FileStatus
		}
		buff := bytes.NewBuffer(nil)
		twriter := tabwriter.NewWriter(buff, 0, 1, 1, ' ', 0)
		_, err = fmt.Fprintln(twriter, "filename\tstage\tworktree\textra")
		_, err = fmt.Fprintln(twriter, "--------\t-----\t--------\t-----")
		if err != nil {
			return err
		}
		for path, filestatus := range status {
			var sname string
			var wname string
			switch filestatus.Staging {
			case git.Unmodified:
				sname = "unmodified"
			case git.Untracked:
				sname = "untracked"
			case git.Modified:
				sname = "modified"
			case git.Added:
				sname = "added"
			case git.Deleted:
				sname = "deleted"
			case git.Renamed:
				sname = "renamed"
			case git.Copied:
				sname = "copied"
			case git.UpdatedButUnmerged:
				sname = "updated_but_unmerged"
			}
			switch filestatus.Worktree {
			case git.Unmodified:
				wname = "unmodified"
			case git.Untracked:
				wname = "untracked"
			case git.Modified:
				wname = "modified"
			case git.Added:
				wname = "added"
			case git.Deleted:
				wname = "deleted"
			case git.Renamed:
				wname = "renamed"
			case git.Copied:
				wname = "copied"
			case git.UpdatedButUnmerged:
				wname = "updated_but_unmerged"
			}
			_, err = fmt.Fprintf(twriter, "%s\t%s\t%s\t%s\n", path, sname, wname, filestatus.Extra)
			if err != nil {
				return err
			}
			// fmt.Println(path, wname, sname, filestatus.Extra)
		}
		err = twriter.Flush()
		if err != nil {
			return err
		}
		fmt.Println(buff.String())

		return err
	},
}
