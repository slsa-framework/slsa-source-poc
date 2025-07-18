package options

import (
	"context"
	"errors"
	"fmt"

	"github.com/slsa-framework/slsa-source-poc/sourcetool/pkg/ghcontrol"
	"github.com/slsa-framework/slsa-source-poc/sourcetool/pkg/policy"
)

type Options struct {
	Repo   string
	Owner  string
	Branch string
	Commit string

	// Organization to look for slsa and user forks
	UserForkOrg string
	Enforce     bool
	UseSSH      bool
	UpdateRepo  bool

	// PolicyRepo is the repository where the policies are stored
	PolicyRepo string
}

// DefaultOptions holds the default options the tool initializes with
var Default = Options{
	PolicyRepo: fmt.Sprintf("%s/%s", policy.SourcePolicyRepoOwner, policy.SourcePolicyRepo),
	UseSSH:     true,
}

// GetGitHubConnection creates a new github connection to the repository
// defined in the options set.
func (o *Options) GetGitHubConnection() (*ghcontrol.GitHubConnection, error) {
	if o.Owner == "" {
		return nil, errors.New("owner not set")
	}
	if o.Repo == "" {
		return nil, errors.New("repository not set")
	}

	return ghcontrol.NewGhConnection(o.Owner, o.Repo, ghcontrol.BranchToFullRef(o.Branch)), nil
}

// EnsureCommit checks the options have a commit sha defined. If not, then
// the latest commit from the loaded branch is read from the GitHub API.
func (o *Options) EnsureCommit() error {
	if o.Commit != "" {
		return nil
	}

	if o.Branch == "" {
		return errors.New("unable to fetch latest commit, no branch defined")
	}

	ghc, err := o.GetGitHubConnection()
	if err != nil {
		return fmt.Errorf("getting GitHub connection: %w", err)
	}

	commit, err := ghc.GetLatestCommit(context.Background(), o.Branch)
	if err != nil {
		return fmt.Errorf("fetching latest commit: %w", err)
	}
	o.Commit = commit
	return nil
}

// EnsureBranch checks that the options set has a branch set and if not, it
// reads the repository's default branch from the GitHub API.
func (o *Options) EnsureBranch() error {
	if o.Branch != "" {
		return nil
	}

	ghc, err := o.GetGitHubConnection()
	if err != nil {
		return fmt.Errorf("getting GitHub connection: %w", err)
	}

	branch, err := ghc.GetDefaultBranch(context.Background())
	if err != nil {
		return fmt.Errorf("fetching default branch: %w", err)
	}
	o.Branch = branch
	return nil
}

type Fn func(*Options) error
