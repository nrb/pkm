#!/bin/zsh

PKM_DIR=${HOME}/.pkm
PKM_DATA_DIR=${PKM_DIR}/data
PKM_SCRIPTS_DIR=${PKM_DIR}/scripts

# Create our working directory
if [[ ! -d ${PKM_DIR} ]]; then
  mkdir -p ${PKM_DATA_DIR}
  mkdir -p ${PKM_SCRIPTS_DIR}
  # TODO: maybe bundle these into the go binary for installation?
  cp scripts/*.sh ${PKM_SCRIPTS_DIR}
fi

# TODO: install launchd plists on macOS
