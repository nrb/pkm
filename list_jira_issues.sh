#!/bin/zsh

source env.sh

# Find all Jira issues assigned to me along with their associated metadata.
jira issue list -s~Closed -a$(jira me) -q "Project in (OCPBUGS,OCPCLOUD)" --raw > ${PKM_CACHE_DIR}/jira_issues.json
