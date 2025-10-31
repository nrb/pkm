#!/bin/zsh

source env.sh

$(pwd)/get_todays_commits.sh > ${PKM_CACHE_DIR}/todays_commits.json
