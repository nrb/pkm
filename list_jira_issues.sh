#!/bin/zsh

# Find all Jira issues assigned to me along with their associated metadata.
jira issue list -s~Closed -a$(jira me) -q "Project in (OCPBUGS,OCPCLOUD)" --raw > ${HOME}/.pkm_cache/jira_issues.json
