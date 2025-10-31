#!/bin/zsh

# look through the ~/projects directory for git repos and find branches that have commits from today.

gitFormat='{%n "abbreviated_commit": "%h",%n  "branch": "%D",%n  "author": {%n    "name": "%aN",%n    "email": "%aE",%n    "date": "%aD"%n  },%n  "committer": {%n    "name": "%cN",%n    "email": "%cE",%n    "date": "%cD"%n  },%n  "subject": "%s" %n}'
author="Nolan Brubaker"

pushd ${HOME}/projects
for dir in */; do
  if [ -d "${dir}/.git" ]; then
    cd "${dir}"
    for branch in $(git branch --format='%(refname:short)'); do
      commits=$(git log "${branch}" --since="1 days ago" --author="${author}" --committer="${author}" --pretty=format:${gitFormat} 2>/dev/null)
      if [ -n "${commits}" ]; then
        echo "${commits}"
        echo ""
      fi
    done
    cd - > /dev/null
  fi
done
popd
