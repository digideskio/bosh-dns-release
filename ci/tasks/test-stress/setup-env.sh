#!/bin/bash
main() {
  source $PWD/bosh-dns-release/ci/assets/utils.sh
  local output_dir="$PWD/updated-bbl-state/"
  trap "commit_bbl_state_dir ${PWD} ${BBL_STATE_DIR} ${output_dir} 'Update bbl state dir'" EXIT

  export TEST_STRESS_ASSETS=$PWD/bosh-dns-release/ci/assets/test-stress
  export BOSH_DOCKER_CPI_RELEASE_TARBALL=$PWD/bosh-docker-cpi-release/*.tgz

  mkdir -p bbl-state/${BBL_STATE_DIR}

  pushd bbl-state/${BBL_STATE_DIR}
    bbl version
    bbl plan > bbl_plan.txt

    # Customize environment
    cp $TEST_STRESS_ASSETS/terraform/* terraform/
    cp $TEST_STRESS_ASSETS/director/*.sh .

    bbl --debug up
  popd
}

main
