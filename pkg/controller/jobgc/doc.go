/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package jobgc contains a very simple job "garbage collector" implementation,
// JobGCController, that runs in the controller manager. If the number of jobs
// in finished status (right now either Failed or Succeeded) surpasses a
// configurable threshold, the controller will delete jobs in finished state
// until the system reaches the allowed threshold again. The JobGCController
// prioritizes jobs to delete by sorting by creation timestamp and deleting the
// oldest objects first. The JobGCController will not delete active jobs.
package jobgc // import "k8s.io/kubernetes/pkg/controller/jobgc"
