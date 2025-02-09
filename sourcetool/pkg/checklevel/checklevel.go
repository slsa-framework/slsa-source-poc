package checklevel

import (
	"context"
	"errors"
	"fmt"
	"slices"
	"time"

	"github.com/google/go-github/v68/github"
	"github.com/slsa-framework/slsa-source-poc/sourcetool/pkg/policy"
)

type activity struct {
	Id           int
	Before       string
	After        string
	Ref          string
	Timestamp    time.Time
	ActivityType string `json:"activity_type"`
}

// Checks to see if the rule meets our requirements.
func checkRule(ctx context.Context, gh_client *github.Client, owner string, repo string, rule *github.RepositoryRule, minTime time.Time) (bool, error) {
	ruleset, _, err := gh_client.Repositories.GetRuleset(ctx, owner, repo, rule.RulesetID, false)
	if err != nil {
		return false, err
	}

	// We need rules to be 'active' and to have been updated no later than minTime.
	if ruleset.Enforcement != "active" {
		return false, nil
	}

	if minTime.Before(ruleset.UpdatedAt.Time) {
		return false, nil
	}

	return true, nil
}

func commitPushTime(ctx context.Context, gh_client *github.Client, commit string, owner string, repo string, branch string) (time.Time, error) {
	// Unfortunately the gh_client doesn't have native support for this...'
	reqUrl := fmt.Sprintf("repos/%s/%s/activity", owner, repo)
	req, err := gh_client.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return time.Time{}, err
	}

	var result []*activity
	_, err = gh_client.Do(ctx, req, &result)
	if err != nil {
		return time.Time{}, err
	}

	targetRef := fmt.Sprintf("refs/heads/%s", branch)
	monitoredTypes := []string{"push", "force_push", "pr_merge"}
	for _, activity := range result {
		if !slices.Contains(monitoredTypes, activity.ActivityType) {
			continue
		}
		if activity.After == commit && activity.Ref == targetRef {
			// Found it
			return activity.Timestamp, nil
		}
	}

	return time.Time{}, errors.New(fmt.Sprintf("Could not find repo activity for commit %s", commit))
}

// Determines the source level using GitHub's built in controls only.
// This is necessarily only as good as GitHub's controls and existing APIs.
// This is a useful demonstration on how SLSA Level 2 can be acheived with ~minimal effort.
//
// Returns the determined source level (level 2 max) or error.
func DetermineSourceLevelControlOnly(ctx context.Context, gh_client *github.Client, commit string, owner string, repo string, branch string) (string, error) {
	rules, _, err := gh_client.Repositories.GetRulesForBranch(ctx, owner, repo, branch)

	if err != nil {
		return "", err
	}

	var deletionRule *github.RepositoryRule
	var nonFastFowardRule *github.RepositoryRule
	for _, rule := range rules {
		switch rule.Type {
		case "deletion":
			deletionRule = rule
		case "non_fast_forward":
			nonFastFowardRule = rule
		default:
			// ignore
		}
	}

	if deletionRule == nil && nonFastFowardRule == nil {
		// For L2 we need deletion and non-fast-forward rules.
		return policy.SlsaSourceLevel1, nil
	}

	// We want to know when this commit was pushed to ensure the rules were active _then_.
	pushTime, err := commitPushTime(ctx, gh_client, commit, owner, repo, branch)
	if err != nil {
		return "", err
	}

	// We want to check to ensure the repo hasn't enabled/disabled the rules since
	// setting the 'since' field in their policy.
	branchPolicy, err := policy.GetBranchPolicy(ctx, gh_client, owner, repo, branch)
	if err != nil {
		return "", err
	}

	if pushTime.Before(branchPolicy.Since) {
		// This commit was pushed before they had an explicit policy.
		return policy.SlsaSourceLevel1, nil
	}

	deletionGood, err := checkRule(ctx, gh_client, owner, repo, deletionRule, branchPolicy.Since)
	if err != nil {
		return "", err
	}
	nonFFGood, err := checkRule(ctx, gh_client, owner, repo, nonFastFowardRule, branchPolicy.Since)
	if err != nil {
		return "", err
	}

	if deletionGood && nonFFGood {
		return policy.SlsaSourceLevel2, nil
	}

	return policy.SlsaSourceLevel1, nil
}

type SourceProvenanceProperty struct {
	// The name of the property.
	Name string
	// The time from which this property has been continuously enforced.
	Since time.Time
}
type SourceProvenance struct {
	// The commit this provenance documents.
	Commit string `json:"commit"`
	// The commit preceeding 'Commit' in the current context.
	PrevCommit string `json:"prev_commit"`
	// The properties observed for this commit.
	Properties []SourceProvenanceProperty `json:"properties"`
}

func DetermineSourceLevelProvenance() {

	// Compute current provenance based on existing controls only.
	//   Check to see if there's prov for the prior commit, if not return cur_prov
	//   Update 'since' for properties on cur_prov to be MIN(cur_prov[prop].since, prior_prov[prop].since)
	//   Ignore properties on prior_prov that don't exist in cur_prov.
	// Then check updated_cur_prov based on policy. Make sure updated_cur_prov[prop].since <= policy[prop].since.
}
