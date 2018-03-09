#!/bin/bash
main() {
  source $PWD/bosh-dns-release/ci/assets/utils.sh

  export BBL_STATE_DIR=$PWD/bbl-state/${BBL_STATE_SUBDIRECTORY}
  source_bbl_env $BBL_STATE_DIR

  logs_dir=$PWD/errand_logs
  mkdir -p $logs_dir

  pushd bosh-dns-release/ci/assets/test-stress/bosh-workspace
    # Run test
    seq 1 "${DEPLOYMENTS_OF_100}" \
      | xargs -n1 -P"${PARALLEL_DEPLOYMENTS}" -I{} \
      -- bosh -d bosh-dns-{} run-errand dns-lookuper --download-logs --logs-dir=$logs_dir
  popd
}

main
