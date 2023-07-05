# Build a OpenIM release.  This will build the binaries, create the Docker
# images and other build artifacts.

set -o errexit
set -o nounset
set -o pipefail

OPENIM_ROOT=$(dirname "${BASH_SOURCE[0]}")/..
source "${OPENIM_ROOT}/scripts/common.sh"
source "${OPENIM_ROOT}/scripts/lib/release.sh"

OPENIM_RELEASE_RUN_TESTS=${OPENIM_RELEASE_RUN_TESTS-y}

openim::golang::setup_env
openim::build::verify_prereqs
openim::release::verify_prereqs
#openim::build::build_image
openim::build::build_command
openim::release::package_tarballs
openim::release::updload_tarballs
git push origin ${VERSION}
#openim::release::github_release
#openim::release::generate_changelog
