// SPDX-FileCopyrightText: Copyright 2026 The SLSA Authors
// SPDX-License-Identifier: Apache-2.0

package github

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"path"
	"regexp"
	"slices"
	"strings"

	"github.com/google/go-github/v88/github"

	"github.com/slsa-framework/source-tool/pkg/slsa"
	"github.com/slsa-framework/source-tool/pkg/sourcetool/models"
)

// workflowsDir is the directory where GitHub looks for actions workflows
const workflowsDir = ".github/workflows"

// legacyActionsRepos lists the repositories that hosted the SLSA source
// actions before they moved to ActionsOrg/ActionsRepo. Workflows calling
// actions from these locations need to be updated.
var legacyActionsRepos = []string{
	"slsa-framework/source-actions",
	"slsa-framework/slsa-source-poc",
}

// actionsUsesRegexp matches the `uses:` lines of a workflow calling an action
// or a reusable workflow hosted in the current or any of the legacy SLSA
// actions repositories. The submatches capture the line prefix (up to and
// including any opening quote), the repository, the path of the action in the
// repository, the git reference, the closing quote and any trailing comment.
var actionsUsesRegexp = buildActionsUsesRegexp()

func buildActionsUsesRegexp() *regexp.Regexp {
	repos := make([]string, 0, len(legacyActionsRepos)+1)
	repos = append(repos, regexp.QuoteMeta(ActionsOrg+"/"+ActionsRepo))
	for _, r := range legacyActionsRepos {
		repos = append(repos, regexp.QuoteMeta(r))
	}
	return regexp.MustCompile(
		`^(\s*(?:-\s+)?uses:\s*["']?)(` + strings.Join(repos, "|") + `)/([^@\s"']+)@([^\s"'#]+)(["']?)(\s*#.*)?\s*$`,
	)
}

// actionsReference is a reference to a SLSA action or reusable workflow
// found in a workflow file.
type actionsReference struct {
	// Line is the 1-based line number where the reference was found
	Line int
	// Repo is the GitHub repository (owner/name) hosting the action
	Repo string
	// Path of the action or reusable workflow inside the repository
	Path string
	// Ref is the git reference (branch, tag or digest) the action is pinned to
	Ref string
}

// IsLegacy returns true when the reference calls the action from a
// repository that no longer hosts the SLSA actions.
func (ar *actionsReference) IsLegacy() bool {
	return slices.Contains(legacyActionsRepos, ar.Repo)
}

// findActionsReferences scans the contents of a workflow file and returns all
// the references to the SLSA actions it finds, current or legacy.
func findActionsReferences(content string) []*actionsReference {
	refs := []*actionsReference{}
	for i, line := range strings.Split(content, "\n") {
		m := actionsUsesRegexp.FindStringSubmatch(line)
		if m == nil {
			continue
		}
		refs = append(refs, &actionsReference{
			Line: i + 1,
			Repo: m[2],
			Path: m[3],
			Ref:  m[4],
		})
	}
	return refs
}

// workflowFile is a GitHub actions workflow read from a repository
type workflowFile struct {
	// Path of the workflow file, relative to the repository root
	Path string
	// Content is the raw workflow file
	Content string
}

// readWorkflowFiles returns all the workflow files found in the repository
// at the specified git reference. If ref is empty, the repository default
// branch is read. Repositories without workflows return an empty list.
func readWorkflowFiles(ctx context.Context, client *github.Client, owner, repoName, ref string) ([]*workflowFile, error) {
	opts := &github.RepositoryContentGetOptions{Ref: ref}
	_, entries, resp, err := client.Repositories.GetContents(ctx, owner, repoName, workflowsDir, opts)
	if err != nil {
		if resp != nil && resp.StatusCode == http.StatusNotFound {
			return nil, nil
		}
		return nil, fmt.Errorf("listing %s: %w", workflowsDir, err)
	}

	files := []*workflowFile{}
	for _, entry := range entries {
		if entry.GetType() != "file" {
			continue
		}
		if ext := strings.ToLower(path.Ext(entry.GetName())); ext != ".yml" && ext != ".yaml" {
			continue
		}

		content, _, _, err := client.Repositories.GetContents(ctx, owner, repoName, entry.GetPath(), opts)
		if err != nil {
			return nil, fmt.Errorf("reading %s: %w", entry.GetPath(), err)
		}
		if content == nil {
			return nil, fmt.Errorf("reading %s: no file contents returned", entry.GetPath())
		}
		data, err := content.GetContent()
		if err != nil {
			return nil, fmt.Errorf("decoding %s: %w", entry.GetPath(), err)
		}
		files = append(files, &workflowFile{
			Path:    entry.GetPath(),
			Content: data,
		})
	}
	return files, nil
}

// provenanceWorkflow is a workflow file calling the SLSA source actions
type provenanceWorkflow struct {
	*workflowFile
	// References are the calls to the SLSA actions found in the workflow
	References []*actionsReference
}

// LegacyRepos returns the deduplicated list of legacy repositories the
// workflow calls actions from, in the order they appear in the file.
func (pw *provenanceWorkflow) LegacyRepos() []string {
	repos := []string{}
	for _, ref := range pw.References {
		if ref.IsLegacy() && !slices.Contains(repos, ref.Repo) {
			repos = append(repos, ref.Repo)
		}
	}
	return repos
}

// toModel returns the public representation of the workflow, including the
// recommended action to update it when it calls actions from a legacy repo.
func (pw *provenanceWorkflow) toModel(repo *models.Repository) *models.ProvenanceWorkflow {
	res := &models.ProvenanceWorkflow{
		Path:               pw.Path,
		LegacyActionsRepos: pw.LegacyRepos(),
	}
	if res.IsLegacy() {
		res.RecommendedAction = &slsa.ControlRecommendedAction{
			Message: fmt.Sprintf(
				"Update %s to call the SLSA actions from %s/%s", pw.Path, ActionsOrg, ActionsRepo,
			),
			Command: fmt.Sprintf(
				"sourcetool setup controls --config=%s %s", models.CONFIG_GEN_PROVENANCE, repo.Path,
			),
		}
	}
	return res
}

// findProvenanceWorkflows reads the workflows of a repository at the
// specified git reference and returns those calling the SLSA source actions.
func findProvenanceWorkflows(ctx context.Context, client *github.Client, owner, repoName, ref string) ([]*provenanceWorkflow, error) {
	files, err := readWorkflowFiles(ctx, client, owner, repoName, ref)
	if err != nil {
		return nil, fmt.Errorf("reading repository workflows: %w", err)
	}

	workflows := []*provenanceWorkflow{}
	for _, f := range files {
		refs := findActionsReferences(f.Content)
		if len(refs) == 0 {
			continue
		}
		workflows = append(workflows, &provenanceWorkflow{
			workflowFile: f,
			References:   refs,
		})
	}
	return workflows, nil
}

// FindProvenanceWorkflows returns the workflows in the branch that call the
// SLSA source actions, flagging those still calling them from a legacy
// repository.
func (b *Backend) FindProvenanceWorkflows(ctx context.Context, branch *models.Branch) ([]*models.ProvenanceWorkflow, error) {
	if branch == nil || branch.Repository == nil {
		return nil, errors.New("branch has no repository")
	}

	owner, repoName, err := branch.Repository.PathAsGitHubOwnerName()
	if err != nil {
		return nil, err
	}

	client, err := b.authenticator.GetGitHubClient()
	if err != nil {
		return nil, fmt.Errorf("getting GitHub client: %w", err)
	}

	workflows, err := findProvenanceWorkflows(ctx, client, owner, repoName, branch.GetName())
	if err != nil {
		return nil, err
	}

	res := make([]*models.ProvenanceWorkflow, 0, len(workflows))
	for _, wf := range workflows {
		res = append(res, wf.toModel(branch.Repository))
	}
	return res, nil
}
