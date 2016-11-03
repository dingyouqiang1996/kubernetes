#!/usr/bin/env bash
# Copyright 2016 The Kubernetes Authors.
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

# This command is used by bazel as the workspace_status_command
# to implement build stamping with git information.

set -o errexit
set -o nounset
set -o pipefail

export KUBE_ROOT=$(dirname "${BASH_SOURCE}")/..

source hack/lib/version.sh
kube::version::get_version_vars

cat <<EOF
BUILD_GIT_COMMIT ${KUBE_GIT_COMMIT-}
BUILD_SCM_STATUS ${KUBE_GIT_TREE_STATE-}
BUILD_SCM_REVISION ${KUBE_GIT_VERSION-}
BUILD_MAJOR_VERSION ${KUBE_GIT_MAJOR-}
BUILD_MINOR_VERSION ${KUBE_GIT_MINOR-}
EOF
