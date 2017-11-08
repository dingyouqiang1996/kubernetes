// +build !ignore_autogenerated

/*
Copyright 2017 The Kubernetes Authors.

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

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package kubeadm

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	reflect "reflect"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
//
// Deprecated: deepcopy registration will go away when static deepcopy is fully implemented.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*API).DeepCopyInto(out.(*API))
			return nil
		}, InType: reflect.TypeOf(&API{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Etcd).DeepCopyInto(out.(*Etcd))
			return nil
		}, InType: reflect.TypeOf(&Etcd{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*HostPathMount).DeepCopyInto(out.(*HostPathMount))
			return nil
		}, InType: reflect.TypeOf(&HostPathMount{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*MasterConfiguration).DeepCopyInto(out.(*MasterConfiguration))
			return nil
		}, InType: reflect.TypeOf(&MasterConfiguration{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Networking).DeepCopyInto(out.(*Networking))
			return nil
		}, InType: reflect.TypeOf(&Networking{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*NodeConfiguration).DeepCopyInto(out.(*NodeConfiguration))
			return nil
		}, InType: reflect.TypeOf(&NodeConfiguration{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*SelfHostedEtcd).DeepCopyInto(out.(*SelfHostedEtcd))
			return nil
		}, InType: reflect.TypeOf(&SelfHostedEtcd{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*TokenDiscovery).DeepCopyInto(out.(*TokenDiscovery))
			return nil
		}, InType: reflect.TypeOf(&TokenDiscovery{})},
	)
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *API) DeepCopyInto(out *API) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new API.
func (in *API) DeepCopy() *API {
	if in == nil {
		return nil
	}
	out := new(API)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Etcd) DeepCopyInto(out *Etcd) {
	*out = *in
	if in.Endpoints != nil {
		in, out := &in.Endpoints, &out.Endpoints
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ExtraArgs != nil {
		in, out := &in.ExtraArgs, &out.ExtraArgs
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.SelfHosted != nil {
		in, out := &in.SelfHosted, &out.SelfHosted
		if *in == nil {
			*out = nil
		} else {
			*out = new(SelfHostedEtcd)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Etcd.
func (in *Etcd) DeepCopy() *Etcd {
	if in == nil {
		return nil
	}
	out := new(Etcd)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostPathMount) DeepCopyInto(out *HostPathMount) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostPathMount.
func (in *HostPathMount) DeepCopy() *HostPathMount {
	if in == nil {
		return nil
	}
	out := new(HostPathMount)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MasterConfiguration) DeepCopyInto(out *MasterConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.API = in.API
	in.Etcd.DeepCopyInto(&out.Etcd)
	out.Networking = in.Networking
	if in.AuthorizationModes != nil {
		in, out := &in.AuthorizationModes, &out.AuthorizationModes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.TokenTTL != nil {
		in, out := &in.TokenTTL, &out.TokenTTL
		if *in == nil {
			*out = nil
		} else {
			*out = new(v1.Duration)
			**out = **in
		}
	}
	if in.APIServerExtraArgs != nil {
		in, out := &in.APIServerExtraArgs, &out.APIServerExtraArgs
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ControllerManagerExtraArgs != nil {
		in, out := &in.ControllerManagerExtraArgs, &out.ControllerManagerExtraArgs
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.SchedulerExtraArgs != nil {
		in, out := &in.SchedulerExtraArgs, &out.SchedulerExtraArgs
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.APIServerExtraVolumes != nil {
		in, out := &in.APIServerExtraVolumes, &out.APIServerExtraVolumes
		*out = make([]HostPathMount, len(*in))
		copy(*out, *in)
	}
	if in.ControllerManagerExtraVolumes != nil {
		in, out := &in.ControllerManagerExtraVolumes, &out.ControllerManagerExtraVolumes
		*out = make([]HostPathMount, len(*in))
		copy(*out, *in)
	}
	if in.SchedulerExtraVolumes != nil {
		in, out := &in.SchedulerExtraVolumes, &out.SchedulerExtraVolumes
		*out = make([]HostPathMount, len(*in))
		copy(*out, *in)
	}
	if in.APIServerCertSANs != nil {
		in, out := &in.APIServerCertSANs, &out.APIServerCertSANs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.FeatureGates != nil {
		in, out := &in.FeatureGates, &out.FeatureGates
		*out = make(map[string]bool, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MasterConfiguration.
func (in *MasterConfiguration) DeepCopy() *MasterConfiguration {
	if in == nil {
		return nil
	}
	out := new(MasterConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MasterConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Networking) DeepCopyInto(out *Networking) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Networking.
func (in *Networking) DeepCopy() *Networking {
	if in == nil {
		return nil
	}
	out := new(Networking)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeConfiguration) DeepCopyInto(out *NodeConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.DiscoveryTokenAPIServers != nil {
		in, out := &in.DiscoveryTokenAPIServers, &out.DiscoveryTokenAPIServers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.DiscoveryTokenCACertHashes != nil {
		in, out := &in.DiscoveryTokenCACertHashes, &out.DiscoveryTokenCACertHashes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeConfiguration.
func (in *NodeConfiguration) DeepCopy() *NodeConfiguration {
	if in == nil {
		return nil
	}
	out := new(NodeConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SelfHostedEtcd) DeepCopyInto(out *SelfHostedEtcd) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SelfHostedEtcd.
func (in *SelfHostedEtcd) DeepCopy() *SelfHostedEtcd {
	if in == nil {
		return nil
	}
	out := new(SelfHostedEtcd)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TokenDiscovery) DeepCopyInto(out *TokenDiscovery) {
	*out = *in
	if in.Addresses != nil {
		in, out := &in.Addresses, &out.Addresses
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TokenDiscovery.
func (in *TokenDiscovery) DeepCopy() *TokenDiscovery {
	if in == nil {
		return nil
	}
	out := new(TokenDiscovery)
	in.DeepCopyInto(out)
	return out
}
