#!/bin/zsh

source env.sh

# Create our working directory
if [[ ! -d ${PKM_DIR} ]]; then
  mkdir -p ${PKM_CACHE_DIR}
  mkdir -p ${PKM_REPORT_DIR}
fi

# TODO: install launchd plists on macOS
