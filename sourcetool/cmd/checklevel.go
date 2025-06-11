/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/slsa-framework/slsa-source-poc/sourcetool/pkg/attest"
	"github.com/slsa-framework/slsa-source-poc/sourcetool/pkg/ghcontrol"
	"github.com/slsa-framework/slsa-source-poc/sourcetool/pkg/policy"
)

type CheckLevelArgs struct {
	commit, owner, repo, branch, outputVsa, outputUnsignedVsa, useLocalPolicy string
	allowMergeCommits                                                         bool
}

func (cla *CheckLevelArgs) Validate() error {
	if cla.commit == "" || cla.owner == "" || cla.repo == "" || cla.branch == "" {
		return errors.New("must set commit, owner, repo, and branch flags")
	}
	return nil
}

// checklevelCmd represents the checklevel command
var (
	checkLevelArgs CheckLevelArgs

	checklevelCmd = &cobra.Command{
		Use:   "checklevel",
		Short: "Determines the SLSA Source Level of the repo",
		Long: `Determines the SLSA Source Level of the repo.

This is meant to be run within the corresponding GitHub Actions workflow.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return doCheckLevel(&checkLevelArgs)
		},
	}
)

func doCheckLevel(cla *CheckLevelArgs) error {
	if err := cla.Validate(); err != nil {
		return err
	}

	ghconnection := ghcontrol.NewGhConnection(cla.owner, cla.repo, ghcontrol.BranchToFullRef(cla.branch)).WithAuthToken(githubToken)
	ghconnection.Options.AllowMergeCommits = cla.allowMergeCommits

	ctx := context.Background()
	controlStatus, err := ghconnection.GetBranchControls(ctx, cla.commit, ghconnection.GetFullRef())
	if err != nil {
		return err
	}
	pe := policy.NewPolicyEvaluator()
	pe.UseLocalPolicy = checkLevelProvArgs.useLocalPolicy
	verifiedLevels, policyPath, err := pe.EvaluateControl(ctx, ghconnection, controlStatus)
	if err != nil {
		return err
	}
	fmt.Print(verifiedLevels)

	unsignedVsa, err := attest.CreateUnsignedSourceVsa(ghconnection.GetRepoUri(), ghconnection.GetFullRef(), cla.commit, verifiedLevels, policyPath)
	if err != nil {
		return err
	}
	if cla.outputUnsignedVsa != "" {
		if err := os.WriteFile(cla.outputUnsignedVsa, []byte(unsignedVsa), 0o644); err != nil { //nolint:gosec
			return err
		}
	}

	if cla.outputVsa != "" {
		// This will output in the sigstore bundle format.
		signedVsa, err := attest.Sign(unsignedVsa)
		if err != nil {
			return err
		}
		err = os.WriteFile(cla.outputVsa, []byte(signedVsa), 0o644) //nolint:gosec
		if err != nil {
			return err
		}
	}

	return nil
}

func init() {
	rootCmd.AddCommand(checklevelCmd)

	// Here you will define your flags and configuration settings.
	checklevelCmd.Flags().StringVar(&checkLevelArgs.commit, "commit", "", "The commit to check.")
	checklevelCmd.Flags().StringVar(&checkLevelArgs.owner, "owner", "", "The GitHub repository owner - required.")
	checklevelCmd.Flags().StringVar(&checkLevelArgs.repo, "repo", "", "The GitHub repository name - required.")
	checklevelCmd.Flags().StringVar(&checkLevelArgs.branch, "branch", "", "The branch within the repository - required.")
	checklevelCmd.Flags().StringVar(&checkLevelArgs.outputVsa, "output_vsa", "", "The path to write a signed VSA with the determined level.")
	checklevelCmd.Flags().StringVar(&checkLevelArgs.outputUnsignedVsa, "output_unsigned_vsa", "", "The path to write an unsigned vsa with the determined level.")
	checklevelCmd.Flags().StringVar(&checkLevelArgs.useLocalPolicy, "use_local_policy", "", "UNSAFE: Use the policy at this local path instead of the official one.")
	checklevelCmd.Flags().BoolVar(&checkLevelArgs.allowMergeCommits, "allow-merge-commits", false, "[EXPERIMENTAL] Allow merge commits in branch.")
}
