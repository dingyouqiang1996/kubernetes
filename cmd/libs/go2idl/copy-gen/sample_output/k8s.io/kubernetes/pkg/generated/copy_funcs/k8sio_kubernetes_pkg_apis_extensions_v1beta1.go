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
	v1 "k8s.io/kubernetes/pkg/api/v1"
	v1beta1 "k8s.io/kubernetes/pkg/apis/extensions/v1beta1"
)

func copy_ApisExtensionsV1beta1APIVersion(in, out *v1beta1.APIVersion) error {
	*out = *in

}

func copy_ApisExtensionsV1beta1DaemonSetStatus(in, out *v1beta1.DaemonSetStatus) error {
	*out = *in

}

func copy_ApisExtensionsV1beta1DeploymentStatus(in, out *v1beta1.DeploymentStatus) error {
	*out = *in

}

func copy_ApisExtensionsV1beta1IngressStatus(in, out *v1beta1.IngressStatus) error {
	*out = *in
	{
		in, out := &(*in).LoadBalancer, &(*out).LoadBalancer
		*out = *in
		{
			in, out := &(*in).Ingress, &(*out).Ingress
			*out = make([]v1.LoadBalancerIngress, len(*in))
			for i := range *in {
				{
					in, out := &(*in[i]), &(*out[i])
					*out = *in
				}
			}
		}
	}

}

func copy_ApisExtensionsV1beta1ReplicationControllerDummy(in, out *v1beta1.ReplicationControllerDummy) error {
	*out = *in
	{
		in, out := &(*in).TypeMeta, &(*out).TypeMeta
		*out = *in
	}

}

func copy_ApisExtensionsV1beta1ScaleSpec(in, out *v1beta1.ScaleSpec) error {
	*out = *in

}

func copy_ApisExtensionsV1beta1ScaleStatus(in, out *v1beta1.ScaleStatus) error {
	*out = *in
	{
		in, out := &(*in).Selector, &(*out).Selector
		*out = make(map[string]string, len(*in))
		for k, v := range *in {
			var k2 string
			{
				in, out := &k, &k2
				*in = *out
			}
			var v2 string
			{
				in, out := &v, &v2
				*in = *out
			}
			(*out)[k2] = v2
		}
	}

}

func copy_ApisExtensionsV1beta1SubresourceReference(in, out *v1beta1.SubresourceReference) error {
	*out = *in

}
