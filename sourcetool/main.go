package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/go-github/v68/github"
)

const (
	SlsaSourceLevel1      = "SLSA_SOURCE_LEVEL_1"
	SlsaSourceLevel2      = "SLSA_SOURCE_LEVEL_2"
	SourcePolicyRepoOwner = "slsa-framework"
	SourcePolicyRepo      = "slsa-source-poc"
)

type activity struct {
	Id           int
	Before       string
	After        string
	Ref          string
	Timestamp    time.Time
	ActivityType string `json:"activity_type"`
}

type protectedBranch struct {
	Name                  string
	Since                 time.Time
	TargetSlsaSourceLevel string `json:"target_slsa_source_level"`
}
type repoPolicy struct {
	CanonicalRepo     string            `json:"canonical_repo"`
	ProtectedBranches []protectedBranch `json:"protected_branches"`
}

func getBranchPolicy(ctx context.Context, gh_client *github.Client, owner string, repo string, branch string) (*protectedBranch, error) {
	path := fmt.Sprintf("../policy/github.com/%s/%s/source-policy.json", owner, repo)

	policyContents, _, _, err := gh_client.Repositories.GetContents(ctx, SourcePolicyRepoOwner, SourcePolicyRepo, path, nil)
	if err != nil {
		return nil, err
	}

	var p repoPolicy
	err = json.Unmarshal(policyContents.GetContents(), &p)
	if err != nil {
		return nil, err
	}

	for _, pb := range p.ProtectedBranches {
		if pb.Name == branch {
			return &pb, nil
		}
	}

	return nil, errors.New(fmt.Sprintf("Could not find rule for branch %s", branch))
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
	for _, activity := range result {
		if activity.ActivityType != "push" && activity.ActivityType != "force_push" {
			continue
		}
		if activity.After == commit && activity.Ref == targetRef {
			// Found it
			return activity.Timestamp, nil
		}
	}

	return time.Time{}, errors.New(fmt.Sprintf("Could not find repo activity for commit %s", commit))
}

func determineSourceLevel(ctx context.Context, gh_client *github.Client, commit string, owner string, repo string, branch string, minDays int) (string, error) {
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
		return SlsaSourceLevel1, nil
	}

	// We want to know when this commit was pushed to ensure the rules were active _then_.
	pushTime, err := commitPushTime(ctx, gh_client, commit, owner, repo, branch)
	if err != nil {
		return "", err
	}

	// We want to check to ensure the repo hasn't enabled/disabled the rules since
	// setting the 'since' field in their policy.
	branchPolicy, err := getBranchPolicy(ctx, gh_client, owner, repo, branch)
	if err != nil {
		return "", err
	}

	if pushTime.Before(branchPolicy.Since) {
		// This commit was pushed before they had an explicit policy.
		return SlsaSourceLevel1, nil
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
		return SlsaSourceLevel2, nil
	}

	return SlsaSourceLevel1, nil
}

// Determines the source level of a repo.
func main() {
	var commit, owner, repo, branch, outputVsa string
	var minDays int
	flag.StringVar(&commit, "commit", "", "The commit to check.")
	flag.StringVar(&owner, "owner", "", "The GitHub repository owner - required.")
	flag.StringVar(&repo, "repo", "", "The GitHub repository name - required.")
	flag.StringVar(&branch, "branch", "", "The branch within the repository - required.")
	flag.IntVar(&minDays, "min_days", 1, "The minimum duration that the rules need to have been enabled for.")
	flag.StringVar(&outputVsa, "output_vsa", "", "The path to write a signed VSA with the determined level.")
	flag.Parse()

	if commit == "" || owner == "" || repo == "" || branch == "" {
		log.Fatal("Must set commit, owner, repo, and branch flags.")
	}

	gh_client := github.NewClient(nil)
	ctx := context.Background()

	sourceLevel, err := determineSourceLevel(ctx, gh_client, commit, owner, repo, branch, minDays)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(sourceLevel)

	if outputVsa != "" {
		// This will output in the sigstore bundle format.
		signedVsa, err := createSignedSourceVsa(owner, repo, commit, sourceLevel)
		if err != nil {
			log.Fatal(err)
		}
		err = os.WriteFile(outputVsa, []byte(signedVsa), 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
}
