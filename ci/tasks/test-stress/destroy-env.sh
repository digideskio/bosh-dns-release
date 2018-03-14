#!/bin/bash
main() {
  source $PWD/bosh-dns-release/ci/assets/utils.sh
  OUTPUT_DIR="cleanup-bbl-state/"
  trap "commit_bbl_state_dir ${PWD} 'Remove bbl state dir' ${OUTPUT_DIR}" EXIT

  source_bbl_env bbl-state/${BBL_STATE_DIR}
  clean_up_director docker

  pushd bbl-state/${BBL_STATE_DIR}
    bbl version

    bbl --debug destroy --no-confirm
  popd
}

main
