#! /bin/zsh

cd ${PKM_GIT_ROOT}
# Fetch PR review requests that are open
gh search prs --review-requested=@me --state=open --sort updated --json repository,number,title,updatedAt > ${PKM_DATA_DIR}/gh_reviews.json
