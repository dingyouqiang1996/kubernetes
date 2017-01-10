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

// This file was autogenerated by conversion-gen. Do not edit it manually!

package v1beta1

import (
	storage "k8s.io/kubernetes/pkg/apis/storage"
	conversion "k8s.io/kubernetes/pkg/conversion"
	runtime "k8s.io/kubernetes/pkg/runtime"
	unsafe "unsafe"
)

func init() {
	SchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1beta1_StorageClass_To_storage_StorageClass,
		Convert_storage_StorageClass_To_v1beta1_StorageClass,
		Convert_v1beta1_StorageClassList_To_storage_StorageClassList,
		Convert_storage_StorageClassList_To_v1beta1_StorageClassList,
	)
}

func autoConvert_v1beta1_StorageClass_To_storage_StorageClass(in *StorageClass, out *storage.StorageClass, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Provisioner = in.Provisioner
	out.Parameters = *(*map[string]string)(unsafe.Pointer(&in.Parameters))
	return nil
}

func Convert_v1beta1_StorageClass_To_storage_StorageClass(in *StorageClass, out *storage.StorageClass, s conversion.Scope) error {
	return autoConvert_v1beta1_StorageClass_To_storage_StorageClass(in, out, s)
}

func autoConvert_storage_StorageClass_To_v1beta1_StorageClass(in *storage.StorageClass, out *StorageClass, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Provisioner = in.Provisioner
	out.Parameters = *(*map[string]string)(unsafe.Pointer(&in.Parameters))
	return nil
}

func Convert_storage_StorageClass_To_v1beta1_StorageClass(in *storage.StorageClass, out *StorageClass, s conversion.Scope) error {
	return autoConvert_storage_StorageClass_To_v1beta1_StorageClass(in, out, s)
}

func autoConvert_v1beta1_StorageClassList_To_storage_StorageClassList(in *StorageClassList, out *storage.StorageClassList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]storage.StorageClass)(unsafe.Pointer(&in.Items))
	return nil
}

func Convert_v1beta1_StorageClassList_To_storage_StorageClassList(in *StorageClassList, out *storage.StorageClassList, s conversion.Scope) error {
	return autoConvert_v1beta1_StorageClassList_To_storage_StorageClassList(in, out, s)
}

func autoConvert_storage_StorageClassList_To_v1beta1_StorageClassList(in *storage.StorageClassList, out *StorageClassList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]StorageClass)(unsafe.Pointer(&in.Items))
	return nil
}

func Convert_storage_StorageClassList_To_v1beta1_StorageClassList(in *storage.StorageClassList, out *StorageClassList, s conversion.Scope) error {
	return autoConvert_storage_StorageClassList_To_v1beta1_StorageClassList(in, out, s)
}
