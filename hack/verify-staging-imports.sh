#!/bin/bash

# Copyright 2017 The Kubernetes Authors.
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

set -o errexit
set -o nounset
set -o pipefail

KUBE_ROOT=$(dirname "${BASH_SOURCE}")/..
source "${KUBE_ROOT}/hack/lib/init.sh"

kube::golang::setup_env


for dep in $(ls -1 ${KUBE_ROOT}/staging/src/k8s.io/); do
	if go list -f "{{.Deps}}" ./vendor/k8s.io/${dep}/... | tr " " '\n' | grep k8s.io/kubernetes | grep -v 'k8s.io/kubernetes/vendor' | LC_ALL=C sort -u | grep -qe "."; then
		echo "${dep} has a cyclical dependency:"
		echo
		go list -f '{{.ImportPath}} {{.Deps}}' ./vendor/k8s.io/${dep}/... |
			sed 's|^k8s.io/kubernetes/vendor/\([-a-zA-Z./0-9_]*\)|\1:|' | # remove vendor prefix of ImportPath
			sed 's| k8s.io/kubernetes/vendor/[-a-zA-Z./0-9_]*||g' |       # remove vendored packages
			grep k8s.io/kubernetes |                                      # only show packages with k8s.io/kubernetes deps
			sed 's|\(k8s.io/kubernetes/\)|§\1|g' |                        # mark kubernetes deps with §
			sed 's|\([[ ]\)[a-zA-Z][-a-zA-Z./0-9_]*|\1|g' |               # remove all other deps
			sed 's|§||g' |                                                # remove § mark
			sed 's|\]||g;s|\[||g' |                                       # remove [ and ]
			sed 's|  *| |g' |                                             # squash spaces
			sed $'s| |\\\n  |g' |                                         # put in newlines
			sed 's|^|  |'                                                 # indent everything two spaces
		exit 1
	fi
done

if grep -rq '// import "k8s.io/kubernetes/' 'staging/'; then
	echo 'file has "// import "k8s.io/kubernetes/"'
	exit 1
fi

exit 0