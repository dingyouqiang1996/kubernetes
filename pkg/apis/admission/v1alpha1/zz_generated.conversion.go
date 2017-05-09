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

package v1alpha1

import (
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	pkg_admission "k8s.io/apiserver/pkg/admission"
	admission "k8s.io/kubernetes/pkg/apis/admission"
)

func init() {
	SchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1alpha1_AdmittanceReview_To_admission_AdmittanceReview,
		Convert_admission_AdmittanceReview_To_v1alpha1_AdmittanceReview,
		Convert_v1alpha1_AdmittanceReviewSpec_To_admission_AdmittanceReviewSpec,
		Convert_admission_AdmittanceReviewSpec_To_v1alpha1_AdmittanceReviewSpec,
		Convert_v1alpha1_AdmittanceReviewStatus_To_admission_AdmittanceReviewStatus,
		Convert_admission_AdmittanceReviewStatus_To_v1alpha1_AdmittanceReviewStatus,
	)
}

func autoConvert_v1alpha1_AdmittanceReview_To_admission_AdmittanceReview(in *AdmittanceReview, out *admission.AdmittanceReview, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_AdmittanceReviewSpec_To_admission_AdmittanceReviewSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_AdmittanceReviewStatus_To_admission_AdmittanceReviewStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_AdmittanceReview_To_admission_AdmittanceReview is an autogenerated conversion function.
func Convert_v1alpha1_AdmittanceReview_To_admission_AdmittanceReview(in *AdmittanceReview, out *admission.AdmittanceReview, s conversion.Scope) error {
	return autoConvert_v1alpha1_AdmittanceReview_To_admission_AdmittanceReview(in, out, s)
}

func autoConvert_admission_AdmittanceReview_To_v1alpha1_AdmittanceReview(in *admission.AdmittanceReview, out *AdmittanceReview, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_admission_AdmittanceReviewSpec_To_v1alpha1_AdmittanceReviewSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_admission_AdmittanceReviewStatus_To_v1alpha1_AdmittanceReviewStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_admission_AdmittanceReview_To_v1alpha1_AdmittanceReview is an autogenerated conversion function.
func Convert_admission_AdmittanceReview_To_v1alpha1_AdmittanceReview(in *admission.AdmittanceReview, out *AdmittanceReview, s conversion.Scope) error {
	return autoConvert_admission_AdmittanceReview_To_v1alpha1_AdmittanceReview(in, out, s)
}

func autoConvert_v1alpha1_AdmittanceReviewSpec_To_admission_AdmittanceReviewSpec(in *AdmittanceReviewSpec, out *admission.AdmittanceReviewSpec, s conversion.Scope) error {
	out.Kind = in.Kind
	if err := runtime.Convert_runtime_RawExtension_To_runtime_Object(&in.Object, &out.Object, s); err != nil {
		return err
	}
	if err := runtime.Convert_runtime_RawExtension_To_runtime_Object(&in.OldObject, &out.OldObject, s); err != nil {
		return err
	}
	out.Operation = pkg_admission.Operation(in.Operation)
	out.Name = in.Name
	out.Namespace = in.Namespace
	out.Resource = in.Resource
	out.SubResource = in.SubResource
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.UserInfo, &out.UserInfo, 0); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_AdmittanceReviewSpec_To_admission_AdmittanceReviewSpec is an autogenerated conversion function.
func Convert_v1alpha1_AdmittanceReviewSpec_To_admission_AdmittanceReviewSpec(in *AdmittanceReviewSpec, out *admission.AdmittanceReviewSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_AdmittanceReviewSpec_To_admission_AdmittanceReviewSpec(in, out, s)
}

func autoConvert_admission_AdmittanceReviewSpec_To_v1alpha1_AdmittanceReviewSpec(in *admission.AdmittanceReviewSpec, out *AdmittanceReviewSpec, s conversion.Scope) error {
	out.Kind = in.Kind
	out.Name = in.Name
	out.Namespace = in.Namespace
	if err := runtime.Convert_runtime_Object_To_runtime_RawExtension(&in.Object, &out.Object, s); err != nil {
		return err
	}
	if err := runtime.Convert_runtime_Object_To_runtime_RawExtension(&in.OldObject, &out.OldObject, s); err != nil {
		return err
	}
	out.Operation = pkg_admission.Operation(in.Operation)
	out.Resource = in.Resource
	out.SubResource = in.SubResource
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.UserInfo, &out.UserInfo, 0); err != nil {
		return err
	}
	return nil
}

// Convert_admission_AdmittanceReviewSpec_To_v1alpha1_AdmittanceReviewSpec is an autogenerated conversion function.
func Convert_admission_AdmittanceReviewSpec_To_v1alpha1_AdmittanceReviewSpec(in *admission.AdmittanceReviewSpec, out *AdmittanceReviewSpec, s conversion.Scope) error {
	return autoConvert_admission_AdmittanceReviewSpec_To_v1alpha1_AdmittanceReviewSpec(in, out, s)
}

func autoConvert_v1alpha1_AdmittanceReviewStatus_To_admission_AdmittanceReviewStatus(in *AdmittanceReviewStatus, out *admission.AdmittanceReviewStatus, s conversion.Scope) error {
	out.Allowed = in.Allowed
	out.Reason = in.Reason
	return nil
}

// Convert_v1alpha1_AdmittanceReviewStatus_To_admission_AdmittanceReviewStatus is an autogenerated conversion function.
func Convert_v1alpha1_AdmittanceReviewStatus_To_admission_AdmittanceReviewStatus(in *AdmittanceReviewStatus, out *admission.AdmittanceReviewStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_AdmittanceReviewStatus_To_admission_AdmittanceReviewStatus(in, out, s)
}

func autoConvert_admission_AdmittanceReviewStatus_To_v1alpha1_AdmittanceReviewStatus(in *admission.AdmittanceReviewStatus, out *AdmittanceReviewStatus, s conversion.Scope) error {
	out.Allowed = in.Allowed
	out.Reason = in.Reason
	return nil
}

// Convert_admission_AdmittanceReviewStatus_To_v1alpha1_AdmittanceReviewStatus is an autogenerated conversion function.
func Convert_admission_AdmittanceReviewStatus_To_v1alpha1_AdmittanceReviewStatus(in *admission.AdmittanceReviewStatus, out *AdmittanceReviewStatus, s conversion.Scope) error {
	return autoConvert_admission_AdmittanceReviewStatus_To_v1alpha1_AdmittanceReviewStatus(in, out, s)
}
