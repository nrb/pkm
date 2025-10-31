#! /bin/zh

# Fetch PR review requests that are open
gh search prs --review-requested=@me --state=open --sort updated --json repository,number,title,updatedAt > ${HOME}/.pkm_cache/gh_reviews.json
