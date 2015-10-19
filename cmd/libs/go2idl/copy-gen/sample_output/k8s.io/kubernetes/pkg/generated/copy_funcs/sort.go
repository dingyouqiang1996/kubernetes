/*
Copyright 2015 The Kubernetes Authors All rights reserved.

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

// This file was autogenerated by the command:
// $ ./copy-gen -o sample_output
// Do not edit it manually!

package copy_funcs

import (
	sort "sort"
)

func copy_SortIntSlice(in, out *sort.IntSlice) error {
	*out = make(sort.IntSlice, len(*in))
	copy(*out, *in)

}

func copy_SortStringSlice(in, out *sort.StringSlice) error {
	*out = make(sort.StringSlice, len(*in))
	copy(*out, *in)

}
