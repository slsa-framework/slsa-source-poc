# SLSA Source PoC - Requirements Mapping

This file describes how this tool comes to a conclusion about what SLSA Source Level a given commit meets.

It is currently written based on https://slsa.dev/spec/draft/source-requirements as of 2025-01-06

## Safe Expunging Process

Users of this tool can support the safe expunging process by using the
[bypass list](https://docs.github.com/en/repositories/configuring-branches-and-merges-in-your-repository/managing-rulesets/creating-rulesets-for-a-repository#granting-bypass-permissions-for-your-branch-or-tag-ruleset)
feature of rulesets.  Organizations should have the bypass list set to include _only_ a custom
'safe-expunging' role which maps to the trusted accounts that can perform this operation.

In addition the role must only be assigned to accounts that are used for this process and
_must not_ include accounts that are used for day-to-day development.

There is no technical enforcement of this requirement.

## Organization Requirements

Open question: does our tooling even need to care about these things?  Should VSA creation
take this into account (it seems like it could depending on how the open questions below
are answered).

### Use Modern Tools

Level 1+: git is one of the modern tools, since this tool only works with git this requirement is met.

### Canonical location

Open question: It's unclear how this should be checked by tooling. One thought is that this
requirement may be in the wrong spot.  Instead perhaps what we're looking for is that given
packages distributed further downstream should indicate which source repos are canonical.
They could do this using SLSA _build_ provenance.

TODO: Now enabling people to set a canonical location in a policy file.  Not yet required for Level 1.

### Distribute summary attestations

This tool stores summary attestations with other attestations in `git notes`, one attestation per line.

### Distribute provenance attestations

Level 1: N/A

Level 2: N/A

Level 3:

This tool stores provenance attestations with other attestations in `git notes`, one attestation per line.

## Source Control System Requirements

### Revisions are immutable and uniquely identifiable

Level 1+: This tool only attests to git revisions (commits), and those are inherently immutable and
uniquely identifiable.

### Repository IDs	

Level 1+: This tool currently only supports GitHub repositories and those are uniquely identified
by the repository URL (e.g. `https://github.com/slsa-framework/slsa-source-poc`).

### Revision IDs	

Level 1+: This tool only attests to git revisions (commits) and those have immutability inherently
enforced.

Open question: Is this duplicative of "Revisions are immutable and uniquely identifiable"

### Source Summary Attestations

Level 1+: This is the tool that is generating these summary attestations.

This tool stores summary attestations with other attestations in `git notes`, one attestation per line.

Open question: Is this duplicative of "Distribute summary attestations"

### Branches

Level 1: N/A

Level 2+: For a commit on a branch to qualify as Level 2+ it must be explicitly indicated in the corresponding policy file.
This is taken as an indication that it is meant for consumption.

### Continuity

Level 1: N/A

Level 2+: Repos are eligible for Level 2 if they have enabled the "Restrict Deletions" (`deletion`) and "Block force pushes" (`non_fast_forward`) rules for the branch in question.

### Identity Management

Level 1: N/A

Level 2+: This tool relies on GitHub's [built-in user management](https://docs.github.com/en/get-started/learning-about-github/types-of-github-accounts#user-accounts).

### Strong Authentication	

Level 1: N/A

Level 2: N/A

Level 3:

#### Multi Factor

GitHub requires all users who contribute to github.com repos to use multi-factor auth
[reference](https://docs.github.com/en/authentication/securing-your-account-with-two-factor-authentication-2fa/about-mandatory-two-factor-authentication).

#### Strong auth id in source provenance

The identity used by GitHub (with strong auth) is stored in the source provenance `actor` field.

### Source Provenance	

Level 1: N/A

Level 2: N/A

Level 3:
This tool creates 'source provenance' attestations for each push to a protected branch.  It records, at least

* The actor that did the push
* The current commit
* The previous commit
* The SLSA source level the current commit meets
* The time the provenance was created
* The time the branch began meeting that level's requirements

### Enforced change management process

Level 1: N/A

Level 2: N/A

Level 3:

#### Rules discoverable by authorized users of the repo

This tool leverages GitHub rulesets which are readable by anyone with repository access.

#### Cannot be bypassed except by the change management process.

This as outlined in [Safe Expunging Process](#safe-expunging-process) and elsewhere
this tool relies on 'active' enforced rulesets and advises the organization to
only allow bypass to accounts used when performing safe expunging.
