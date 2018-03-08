#!/bin/bash
set -eu -o pipefail

kill_ssh() {
  # kill the ssh tunnel to jumpbox, set up by bbl env
  # (or task will hang forever)
  pkill ssh || true
}

source_bbl_env() {
  local bbl_state_dir=$1

  trap kill_ssh EXIT

  set +x
  eval "$(bbl print-env --state-dir=$bbl_state_dir)"
  set -x
}

commit_bbl_state_dir() {
  local root_dir
  root_dir="${1}"
  local commit_message
  commit_message="${2}"

  pushd "${root_dir}/bbl-state/${BBL_STATE_DIR}"
    if [[ -n $(git status --porcelain) ]]; then
      git config user.name "CI Bot"
      git config user.email "cf-release-integration@pivotal.io"
      git add --all .
      git commit -m "${commit_message}"
    fi
  popd

  pushd "${root_dir}"
    shopt -s dotglob
    cp -R "bbl-state/." "updated-bbl-state/"
  popd
}

clean_up_director() {
  local deployment=$1

  # Ensure the environment is clean
  if [[ -z "$deployment" ]]; then
    bosh deployments --column=name | xargs -n1 bosh delete-deployment --force -n -d
  else
    bosh delete-deployment -d $deployment -n --force
  fi

  # Clean-up old artifacts
  bosh -n clean-up --all
}
