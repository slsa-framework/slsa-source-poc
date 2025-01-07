#!/bin/sh
set -e

# Determine's the source level of a GitHub repository/commit/branch.
# Usage:
# ./determine_source_level_gh.sh <COMMIT> <REPO> <BRANCH>

# TODO: Commit is currently ignored, do we really need it
COMMIT=$1
REPO=$2
BRANCH=$3

# TODO: Replace with scripts/rulesets.sh from https://github.com/slsa-framework/slsa-source-poc/pull/6 ?

# TODO: Add GitHub token for non-public repos?
GITHUB_RULESET=$(curl -L \
    -H "Accept: application/vnd.github+json" \
    -H "X-GitHub-Api-Version: 2022-11-28" \
    https://api.github.com/repos/${REPO}/rules/branches/${BRANCH})

# Check continuity requirement
# We'll assume it meets this requirement if the branch prevents deletions
# and force pushes.
# TODO: Should other things be checked too?
NO_DELETION=$(echo $GITHUB_RULESET | jq '.[] | select(.type=="deletion") | any')
NO_FORCE_PUSH=$(echo $GITHUB_RULESET | jq '.[] | select(.type=="non_fast_forward") | any')

SOURCE_LEVEL="SLSA_SOURCE_LEVEL_1"

if [ "$NO_DELETION" == "true" ] && [ "$NO_FORCE_PUSH" == "true" ]; then
    SOURCE_LEVEL="SLSA_SOURCE_LEVEL_2"
fi
echo $SOURCE_LEVEL