#! /bin/zsh

source env.sh

# Fetch PR review requests that are open
gh search prs --review-requested=@me --state=open --sort updated --json repository,number,title,updatedAt > ${PKM_CACHE_DIR}/gh_reviews.json
