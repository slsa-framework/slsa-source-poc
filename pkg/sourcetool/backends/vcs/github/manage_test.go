// SPDX-FileCopyrightText: Copyright 2026 The SLSA Authors
// SPDX-License-Identifier: Apache-2.0

package github

import (
	"net/http"
	"testing"

	"github.com/google/go-github/v88/github"
	"github.com/migueleliasweb/go-github-mock/src/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/slsa-framework/source-tool/pkg/sourcetool/models"
)

func testTag(name, sha string) *github.RepositoryTag {
	return &github.RepositoryTag{
		Name:   github.Ptr(name),
		Commit: &github.Commit{SHA: github.Ptr(sha)},
	}
}

func TestBuildWorkflowYAML(t *testing.T) {
	t.Parallel()
	yaml := buildWorkflowYAML(
		[]*models.Branch{{Name: "main"}, {Name: "release-1.0"}},
		"v0.1.0", "dea965cdca5e0cb422bf7b2653c9d15f678ad01c",
	)

	assert.Contains(t, yaml, `branches: [ "main", "release-1.0" ]`)
	assert.Contains(
		t, yaml,
		"uses: slsa-framework/actions/.github/workflows/compute_slsa_source.yml@dea965cdca5e0cb422bf7b2653c9d15f678ad01c # v0.1.0",
	)
	assert.NotContains(t, yaml, "@main")
}

func TestLatestReleaseTag(t *testing.T) {
	t.Parallel()
	for _, tc := range []struct {
		name   string
		tags   []*github.RepositoryTag
		expect string
	}{
		{"empty", nil, ""},
		{"single", []*github.RepositoryTag{testTag("v0.1.0", "a")}, "v0.1.0"},
		{
			// Sorting must be numeric, not lexicographic
			"numeric-order",
			[]*github.RepositoryTag{
				testTag("v0.9.0", "a"), testTag("v0.10.0", "b"), testTag("v0.1.0", "c"),
			},
			"v0.10.0",
		},
		{
			"skip-prereleases",
			[]*github.RepositoryTag{testTag("v0.1.0", "a"), testTag("v0.2.0-rc.1", "b")},
			"v0.1.0",
		},
		{
			"skip-floating-tags",
			[]*github.RepositoryTag{testTag("v1", "a"), testTag("v1.0", "b"), testTag("v0.5.0", "c")},
			"v0.5.0",
		},
		{
			"skip-non-semver",
			[]*github.RepositoryTag{
				testTag("latest", "a"), testTag("1.2.3", "b"), testTag("v0.3.0", "c"), testTag("", "d"),
			},
			"v0.3.0",
		},
		{"only-prereleases", []*github.RepositoryTag{testTag("v1.0.0-alpha", "a")}, ""},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			latest := latestReleaseTag(tc.tags)
			if tc.expect == "" {
				assert.Nil(t, latest)
				return
			}
			require.NotNil(t, latest)
			assert.Equal(t, tc.expect, latest.GetName())
		})
	}
}

func TestSearchPullRequestsByTitle(t *testing.T) {
	t.Parallel()
	openPRs := []*github.PullRequest{
		{Number: github.Ptr(1), Title: github.Ptr("Bump something")},
		{Number: github.Ptr(2), Title: github.Ptr(workflowUpdateCommitMessage)},
		{Number: github.Ptr(3), Title: github.Ptr(workflowCommitMessage)},
	}
	for _, tc := range []struct {
		name         string
		prs          []*github.PullRequest
		queries      []string
		expectNumber int
	}{
		{"no-prs", []*github.PullRequest{}, []string{workflowCommitMessage}, 0},
		{"no-match", openPRs, []string{"Something else"}, 0},
		{"add-pr", openPRs, []string{workflowCommitMessage}, 3},
		{"update-pr", openPRs, []string{workflowUpdateCommitMessage}, 2},
		{"any-of", openPRs, []string{workflowCommitMessage, workflowUpdateCommitMessage}, 2},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			client, err := github.NewClient(github.WithHTTPClient(mock.NewMockedHTTPClient(
				mock.WithRequestMatch(mock.GetReposPullsByOwnerByRepo, tc.prs),
			)))
			require.NoError(t, err)

			pr, err := searchPullRequestsByTitle(t.Context(), client, "owner", "repo", tc.queries...)
			require.NoError(t, err)
			if tc.expectNumber == 0 {
				assert.Nil(t, pr)
				return
			}
			require.NotNil(t, pr)
			assert.Equal(t, tc.expectNumber, pr.GetNumber())
		})
	}

	t.Run("api-error", func(t *testing.T) {
		t.Parallel()
		client, err := github.NewClient(github.WithHTTPClient(mock.NewMockedHTTPClient(
			mock.WithRequestMatchHandler(
				mock.GetReposPullsByOwnerByRepo,
				http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
					mock.WriteError(w, http.StatusInternalServerError, "boom")
				}),
			),
		)))
		require.NoError(t, err)
		_, err = searchPullRequestsByTitle(t.Context(), client, "owner", "repo", workflowCommitMessage)
		require.Error(t, err)
	})
}

func TestLatestActionsTag(t *testing.T) {
	t.Parallel()
	for _, tc := range []struct {
		name         string
		mockOption   mock.MockBackendOption
		expectTag    string
		expectDigest string
		mustErr      bool
	}{
		{
			// The latest tag lives in the second page, ensuring all pages are read
			name: "paginated",
			mockOption: mock.WithRequestMatchPages(
				mock.GetReposTagsByOwnerByRepo,
				[]*github.RepositoryTag{testTag("v0.1.0", "aaa"), testTag("v0.2.0-rc.1", "bbb")},
				[]*github.RepositoryTag{testTag("v0.1.1", "ccc"), testTag("v1", "ddd")},
			),
			expectTag:    "v0.1.1",
			expectDigest: "ccc",
		},
		{
			name: "no-release-tags",
			mockOption: mock.WithRequestMatch(
				mock.GetReposTagsByOwnerByRepo,
				[]*github.RepositoryTag{testTag("main", "aaa"), testTag("v1.0.0-rc.1", "bbb")},
			),
			mustErr: true,
		},
		{
			name:       "no-tags",
			mockOption: mock.WithRequestMatch(mock.GetReposTagsByOwnerByRepo, []*github.RepositoryTag{}),
			mustErr:    true,
		},
		{
			name: "missing-digest",
			mockOption: mock.WithRequestMatch(
				mock.GetReposTagsByOwnerByRepo,
				[]*github.RepositoryTag{{Name: github.Ptr("v0.1.0")}},
			),
			mustErr: true,
		},
		{
			name: "api-error",
			mockOption: mock.WithRequestMatchHandler(
				mock.GetReposTagsByOwnerByRepo,
				http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
					mock.WriteError(w, http.StatusInternalServerError, "boom")
				}),
			),
			mustErr: true,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			client, err := github.NewClient(github.WithHTTPClient(mock.NewMockedHTTPClient(tc.mockOption)))
			require.NoError(t, err)

			tag, digest, err := latestActionsTag(t.Context(), client)
			if tc.mustErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tc.expectTag, tag)
			assert.Equal(t, tc.expectDigest, digest)
		})
	}
}
