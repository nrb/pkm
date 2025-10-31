#!/bin/zsh

# Create our working directory
if [[ ! -d ${HOME}/.pkm_cache ]]; then
  mkdir ${HOME}/.pkm_cache
fi

# TODO: install launchd plists on macOS
