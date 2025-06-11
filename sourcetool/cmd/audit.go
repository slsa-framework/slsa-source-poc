/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"errors"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/slsa-framework/slsa-source-poc/sourcetool/pkg/attest"
	"github.com/slsa-framework/slsa-source-poc/sourcetool/pkg/audit"
	"github.com/slsa-framework/slsa-source-poc/sourcetool/pkg/ghcontrol"
)

type AuditMode int

const (
	AuditModeBasic AuditMode = 1
	AuditModeFull  AuditMode = 2
)

// Enable audit mode enum
// String is used both by fmt.Print and by Cobra in help text
func (e *AuditMode) String() string {
	switch *e {
	case AuditModeBasic:
		return "basic"
	case AuditModeFull:
		return "full"
	}
	return "error"
}

// Set must have pointer receiver so it doesn't change the value of a copy
func (e *AuditMode) Set(v string) error {
	switch v {
	case "basic":
		*e = AuditModeBasic
		return nil
	case "full":
		*e = AuditModeFull
		return nil
	default:
		return errors.New(`must be one of "foo", "bar", or "moo"`)
	}
}

// Type is only used in help text
func (e *AuditMode) Type() string {
	return "AuditMode"
}

type AuditArgs struct {
	owner        string
	repo         string
	branch       string
	auditDepth   int
	endingCommit string
	auditMode    AuditMode
}

func (aa *AuditArgs) Validate() error {
	if aa.owner == "" || aa.repo == "" || aa.branch == "" {
		return errors.New("must set owner, repo, and branch flags")
	}
	return nil
}

var (
	auditArgs = &AuditArgs{}

	auditCmd = &cobra.Command{
		Use:   "audit",
		Short: "Audits the SLSA properties and controls of a repository",
		Long: `Checks the revisions on the specified branch within the repository.

Revisions 'pass'and audit if they have:
1. A corresponding VSA
2. Corresponding source provenance
3. The revision (commit) listed in the provenance matches the revision reported by GitHub

Future:
* Check the provenance to validate the verifiedLevels in the VSA match expectations
  (i.e. that the VSA was issued correctly)
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return doAudit(auditArgs)
		},
	}
)

func printResult(ghc *ghcontrol.GitHubConnection, ar *audit.AuditCommitResult, mode AuditMode) {
	good := ar.IsGood()
	status := "passed"
	if !good {
		status = "failed"
	}
	fmt.Printf("commit: %s - %v\n", ar.Commit, status)

	if good && AuditModeBasic == mode {
		return
	}

	if ar.VsaPred != nil {
		fmt.Printf("\tvsa: %v\n", ar.VsaPred.GetVerifiedLevels())
	} else {
		fmt.Printf("\tvsa: none\n")
	}
	if ar.ProvPred != nil {
		fmt.Print("\tprov:\n")
		fmt.Printf("\t\tcontrols: %v\n", ar.ProvPred.Controls)
		if ar.ProvPred.PrevCommit == ar.GhPriorCommit {
			fmt.Printf("\t\tPrevCommit matches GH commit: true\n")
		} else {
			fmt.Printf("\t\tPrevCommit matches GH commit: false: %s != %s\n", ar.ProvPred.PrevCommit, ar.GhPriorCommit)
		}
	} else {
		fmt.Printf("\tprov: none\n")
	}
	if ar.GhControlStatus != nil {
		fmt.Printf("\tgh controls: %v\n", ar.GhControlStatus.Controls)
	}

	fmt.Printf("\tlink: https://github.com/%s/%s/commit/%s\n", ghc.Owner(), ghc.Repo(), ar.GhPriorCommit)
}

func doAudit(auditArgs *AuditArgs) error {
	err := auditArgs.Validate()
	if err != nil {
		return err
	}

	ghc := ghcontrol.NewGhConnection(auditArgs.owner, auditArgs.repo, ghcontrol.BranchToFullRef(auditArgs.branch)).WithAuthToken(githubToken)
	ctx := context.Background()
	verifier := getVerifier()
	pa := attest.NewProvenanceAttestor(ghc, verifier)

	auditor := audit.NewAuditor(ghc, pa, verifier)

	latestCommit, err := ghc.GetLatestCommit(ctx, auditArgs.branch)
	if err != nil {
		return fmt.Errorf("could not get latest commit for %s", auditArgs.branch)
	}

	fmt.Printf("Auditing branch %s starting from revision %s\n", auditArgs.branch, latestCommit)

	count := 0
	for ar, err := range auditor.AuditBranch(ctx, auditArgs.branch) {
		if ar == nil {
			return err
		}
		if err != nil {
			fmt.Printf("\terror: %v\n", err)
		}
		printResult(ghc, ar, auditArgs.auditMode)
		if auditArgs.endingCommit != "" && auditArgs.endingCommit == ar.Commit {
			fmt.Printf("Found ending commit %s\n", auditArgs.endingCommit)
			return nil
		}
		if auditArgs.auditDepth > 0 && count >= auditArgs.auditDepth {
			fmt.Printf("Reached depth limit %d\n", auditArgs.auditDepth)
			return nil
		}
		count++
	}

	return nil
}

func init() {
	rootCmd.AddCommand(auditCmd)

	auditCmd.Flags().StringVar(&auditArgs.owner, "owner", "", "The GitHub repository owner - required.")
	auditCmd.Flags().StringVar(&auditArgs.repo, "repo", "", "The GitHub repository name - required.")
	auditCmd.Flags().StringVar(&auditArgs.branch, "branch", "", "The branch within the repository - required.")
	auditCmd.Flags().IntVar(&auditArgs.auditDepth, "depth", 0, "The max number of revisions to audit (depth <= audit all revisions).")
	auditCmd.Flags().StringVar(&auditArgs.endingCommit, "ending-commit", "", "The commit to stop auditing at.")
	auditArgs.auditMode = AuditModeBasic
	auditCmd.Flags().Var(&auditArgs.auditMode, "audit-mode", "'basic' for limited details (default), 'full' for all details")
}
