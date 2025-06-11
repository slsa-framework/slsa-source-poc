/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/slsa-framework/slsa-source-poc/sourcetool/pkg/attest"
	"github.com/slsa-framework/slsa-source-poc/sourcetool/pkg/ghcontrol"
)

type ProvArgs struct {
	prevAttPath, commit, prevCommit, owner, repo, branch string
}

// provCmd represents the prov command
var (
	provArgs ProvArgs
	provCmd  = &cobra.Command{
		Use:   "prov",
		Short: "Creates provenance for the given commit, but does not check policy.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return doProv(provArgs.prevAttPath, provArgs.commit, provArgs.prevCommit, provArgs.owner, provArgs.repo, provArgs.branch)
		},
	}
)

func doProv(prevAttPath, commit, prevCommit, owner, repo, branch string) error {
	ghconnection := ghcontrol.NewGhConnection(owner, repo, ghcontrol.BranchToFullRef(branch)).WithAuthToken(githubToken)
	ctx := context.Background()
	pa := attest.NewProvenanceAttestor(ghconnection, getVerifier())
	newProv, err := pa.CreateSourceProvenance(ctx, prevAttPath, commit, prevCommit, ghconnection.GetFullRef())
	if err != nil {
		return err
	}
	provStr, err := protojson.Marshal(newProv)
	if err != nil {
		return err
	}
	fmt.Println(string(provStr))
	return nil
}

func init() {
	rootCmd.AddCommand(provCmd)

	provCmd.Flags().StringVar(&provArgs.prevAttPath, "prev_att_path", "", "Path to the file with the attestations for the previous commit (as an in-toto bundle).")
	provCmd.Flags().StringVar(&provArgs.commit, "commit", "", "The commit to check.")
	provCmd.Flags().StringVar(&provArgs.prevCommit, "prev_commit", "", "The commit prior to 'commit'.")
	provCmd.Flags().StringVar(&provArgs.owner, "owner", "", "The GitHub repository owner - required.")
	provCmd.Flags().StringVar(&provArgs.repo, "repo", "", "The GitHub repository name - required.")
	provCmd.Flags().StringVar(&provArgs.branch, "branch", "", "The branch within the repository - required.")
}
