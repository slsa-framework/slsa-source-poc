// SPDX-FileCopyrightText: Copyright 2025 The SLSA Authors
// SPDX-License-Identifier: Apache-2.0

package ghcontrol

var defaultOptions = Options{
	AllowMergeCommits: false,
}

type Options struct {
	// AllowMergeCommits causes the GitHub connector to reject merge
	// commits when set to false.
	AllowMergeCommits bool

	// accessToken is the token we will use to connect to the GitHub API
	accessToken string
}
