# Copyright © 2023 OpenIM open source community. All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# Build a OpenIM release.  This will build the binaries, create the Docker
# images and other build artifacts.

set -o errexit         # Exit the script if any command returns a non-zero exit status
set -o nounset         # Exit the script if any unset variables are encountered
set -o pipefail        # Exit the script if any command in a pipeline fails

OPENIM_ROOT=$(dirname "${BASH_SOURCE[0]}")/..
source "${OPENIM_ROOT}/scripts/common.sh"          # Source the common.sh script
source "${OPENIM_ROOT}/scripts/lib/release.sh"     # Source the release.sh script

OPENIM_RELEASE_RUN_TESTS=${OPENIM_RELEASE_RUN_TESTS-y}   # Set the OPENIM_RELEASE_RUN_TESTS variable

check_github_actions_failures
openim::build::verify_prereqs       # Verify prerequisites for building
check_github_actions_failures

# Build the image
#openim::build::build_image

openim::build::build_command
openim::release::package_tarballs   # Package the tarballs
openim::release::updload_tarballs   # Upload the tarballs
git push origin ${VERSION}           # Push the code to the remote repository with the specified VERSION

# GitHub release (commented out)
#openim::release::github_release

# Generate changelog (commented out)
#openim::release::generate_changelog
