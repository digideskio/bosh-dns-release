#!/bin/bash
main() {
  source $PWD/bosh-dns-release/ci/assets/utils.sh
  local output_dir="$PWD/cleanup-bbl-state/"
  trap "commit_bbl_state_dir ${PWD} ${BBL_STATE_DIR} ${output_dir} 'Remove bbl state dir'" EXIT

  (
    source_bbl_env bbl-state/${BBL_STATE_DIR}
    clean_up_director docker
  )

  pushd bbl-state/${BBL_STATE_DIR}
    bbl version

    bbl --debug destroy --no-confirm
  popd
}

main
