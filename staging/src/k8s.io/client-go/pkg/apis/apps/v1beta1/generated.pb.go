/*
Copyright 2016 The Kubernetes Authors.

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

// Code generated by protoc-gen-gogo.
// source: k8s.io/kubernetes/pkg/apis/apps/v1beta1/generated.proto
// DO NOT EDIT!

/*
	Package v1beta1 is a generated protocol buffer package.

	It is generated from these files:
		k8s.io/kubernetes/pkg/apis/apps/v1beta1/generated.proto

	It has these top-level messages:
		StatefulSet
		StatefulSetList
		StatefulSetSpec
		StatefulSetStatus
*/
package v1beta1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import k8s_io_kubernetes_pkg_api_unversioned "k8s.io/client-go/pkg/apis/meta/v1"
import k8s_io_kubernetes_pkg_api_v1 "k8s.io/client-go/pkg/api/v1"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.GoGoProtoPackageIsVersion1

func (m *StatefulSet) Reset()                    { *m = StatefulSet{} }
func (*StatefulSet) ProtoMessage()               {}
func (*StatefulSet) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{0} }

func (m *StatefulSetList) Reset()                    { *m = StatefulSetList{} }
func (*StatefulSetList) ProtoMessage()               {}
func (*StatefulSetList) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{1} }

func (m *StatefulSetSpec) Reset()                    { *m = StatefulSetSpec{} }
func (*StatefulSetSpec) ProtoMessage()               {}
func (*StatefulSetSpec) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{2} }

func (m *StatefulSetStatus) Reset()                    { *m = StatefulSetStatus{} }
func (*StatefulSetStatus) ProtoMessage()               {}
func (*StatefulSetStatus) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{3} }

func init() {
	proto.RegisterType((*StatefulSet)(nil), "k8s.io.client-go.pkg.apis.apps.v1beta1.StatefulSet")
	proto.RegisterType((*StatefulSetList)(nil), "k8s.io.client-go.pkg.apis.apps.v1beta1.StatefulSetList")
	proto.RegisterType((*StatefulSetSpec)(nil), "k8s.io.client-go.pkg.apis.apps.v1beta1.StatefulSetSpec")
	proto.RegisterType((*StatefulSetStatus)(nil), "k8s.io.client-go.pkg.apis.apps.v1beta1.StatefulSetStatus")
}
func (m *StatefulSet) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *StatefulSet) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintGenerated(data, i, uint64(m.ObjectMeta.Size()))
	n1, err := m.ObjectMeta.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	data[i] = 0x12
	i++
	i = encodeVarintGenerated(data, i, uint64(m.Spec.Size()))
	n2, err := m.Spec.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	data[i] = 0x1a
	i++
	i = encodeVarintGenerated(data, i, uint64(m.Status.Size()))
	n3, err := m.Status.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	return i, nil
}

func (m *StatefulSetList) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *StatefulSetList) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintGenerated(data, i, uint64(m.ListMeta.Size()))
	n4, err := m.ListMeta.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n4
	if len(m.Items) > 0 {
		for _, msg := range m.Items {
			data[i] = 0x12
			i++
			i = encodeVarintGenerated(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *StatefulSetSpec) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *StatefulSetSpec) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Replicas != nil {
		data[i] = 0x8
		i++
		i = encodeVarintGenerated(data, i, uint64(*m.Replicas))
	}
	if m.Selector != nil {
		data[i] = 0x12
		i++
		i = encodeVarintGenerated(data, i, uint64(m.Selector.Size()))
		n5, err := m.Selector.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	data[i] = 0x1a
	i++
	i = encodeVarintGenerated(data, i, uint64(m.Template.Size()))
	n6, err := m.Template.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n6
	if len(m.VolumeClaimTemplates) > 0 {
		for _, msg := range m.VolumeClaimTemplates {
			data[i] = 0x22
			i++
			i = encodeVarintGenerated(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	data[i] = 0x2a
	i++
	i = encodeVarintGenerated(data, i, uint64(len(m.ServiceName)))
	i += copy(data[i:], m.ServiceName)
	return i, nil
}

func (m *StatefulSetStatus) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *StatefulSetStatus) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ObservedGeneration != nil {
		data[i] = 0x8
		i++
		i = encodeVarintGenerated(data, i, uint64(*m.ObservedGeneration))
	}
	data[i] = 0x10
	i++
	i = encodeVarintGenerated(data, i, uint64(m.Replicas))
	return i, nil
}

func encodeFixed64Generated(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Generated(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintGenerated(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *StatefulSet) Size() (n int) {
	var l int
	_ = l
	l = m.ObjectMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Spec.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Status.Size()
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *StatefulSetList) Size() (n int) {
	var l int
	_ = l
	l = m.ListMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	if len(m.Items) > 0 {
		for _, e := range m.Items {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	return n
}

func (m *StatefulSetSpec) Size() (n int) {
	var l int
	_ = l
	if m.Replicas != nil {
		n += 1 + sovGenerated(uint64(*m.Replicas))
	}
	if m.Selector != nil {
		l = m.Selector.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	l = m.Template.Size()
	n += 1 + l + sovGenerated(uint64(l))
	if len(m.VolumeClaimTemplates) > 0 {
		for _, e := range m.VolumeClaimTemplates {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	l = len(m.ServiceName)
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *StatefulSetStatus) Size() (n int) {
	var l int
	_ = l
	if m.ObservedGeneration != nil {
		n += 1 + sovGenerated(uint64(*m.ObservedGeneration))
	}
	n += 1 + sovGenerated(uint64(m.Replicas))
	return n
}

func sovGenerated(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozGenerated(x uint64) (n int) {
	return sovGenerated(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *StatefulSet) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&StatefulSet{`,
		`ObjectMeta:` + strings.Replace(strings.Replace(this.ObjectMeta.String(), "ObjectMeta", "k8s_io_kubernetes_pkg_api_v1.ObjectMeta", 1), `&`, ``, 1) + `,`,
		`Spec:` + strings.Replace(strings.Replace(this.Spec.String(), "StatefulSetSpec", "StatefulSetSpec", 1), `&`, ``, 1) + `,`,
		`Status:` + strings.Replace(strings.Replace(this.Status.String(), "StatefulSetStatus", "StatefulSetStatus", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *StatefulSetList) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&StatefulSetList{`,
		`ListMeta:` + strings.Replace(strings.Replace(this.ListMeta.String(), "ListMeta", "k8s_io_kubernetes_pkg_api_unversioned.ListMeta", 1), `&`, ``, 1) + `,`,
		`Items:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.Items), "StatefulSet", "StatefulSet", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *StatefulSetSpec) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&StatefulSetSpec{`,
		`Replicas:` + valueToStringGenerated(this.Replicas) + `,`,
		`Selector:` + strings.Replace(fmt.Sprintf("%v", this.Selector), "LabelSelector", "k8s_io_kubernetes_pkg_api_unversioned.LabelSelector", 1) + `,`,
		`Template:` + strings.Replace(strings.Replace(this.Template.String(), "PodTemplateSpec", "k8s_io_kubernetes_pkg_api_v1.PodTemplateSpec", 1), `&`, ``, 1) + `,`,
		`VolumeClaimTemplates:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.VolumeClaimTemplates), "PersistentVolumeClaim", "k8s_io_kubernetes_pkg_api_v1.PersistentVolumeClaim", 1), `&`, ``, 1) + `,`,
		`ServiceName:` + fmt.Sprintf("%v", this.ServiceName) + `,`,
		`}`,
	}, "")
	return s
}
func (this *StatefulSetStatus) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&StatefulSetStatus{`,
		`ObservedGeneration:` + valueToStringGenerated(this.ObservedGeneration) + `,`,
		`Replicas:` + fmt.Sprintf("%v", this.Replicas) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringGenerated(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *StatefulSet) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StatefulSet: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StatefulSet: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjectMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ObjectMeta.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Spec", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Spec.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Status.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *StatefulSetList) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StatefulSetList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StatefulSetList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ListMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ListMeta.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Items", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Items = append(m.Items, StatefulSet{})
			if err := m.Items[len(m.Items)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *StatefulSetSpec) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StatefulSetSpec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StatefulSetSpec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Replicas", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Replicas = &v
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Selector", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Selector == nil {
				m.Selector = &k8s_io_kubernetes_pkg_api_unversioned.LabelSelector{}
			}
			if err := m.Selector.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Template", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Template.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VolumeClaimTemplates", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VolumeClaimTemplates = append(m.VolumeClaimTemplates, k8s_io_kubernetes_pkg_api_v1.PersistentVolumeClaim{})
			if err := m.VolumeClaimTemplates[len(m.VolumeClaimTemplates)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServiceName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ServiceName = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *StatefulSetStatus) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StatefulSetStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StatefulSetStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObservedGeneration", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.ObservedGeneration = &v
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Replicas", wireType)
			}
			m.Replicas = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Replicas |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGenerated(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthGenerated
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowGenerated
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipGenerated(data[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthGenerated = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenerated   = fmt.Errorf("proto: integer overflow")
)

var fileDescriptorGenerated = []byte{
	// 637 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x93, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0x86, 0xe3, 0xa4, 0xe9, 0x97, 0x6f, 0x52, 0xfe, 0x86, 0x0a, 0x45, 0x11, 0x72, 0xab, 0x6c,
	0x08, 0x52, 0x3b, 0x56, 0x4a, 0x2b, 0x2a, 0x96, 0x46, 0x02, 0x21, 0x01, 0x45, 0x0e, 0xaa, 0xa0,
	0x08, 0xa4, 0xb1, 0x73, 0x9a, 0x9a, 0xd8, 0x1e, 0xcb, 0x73, 0x9c, 0x35, 0x1b, 0x16, 0xec, 0xb8,
	0x0b, 0x2e, 0x81, 0x5b, 0xa8, 0xc4, 0xa6, 0x4b, 0x56, 0x15, 0x0d, 0x37, 0x82, 0x3c, 0x99, 0x24,
	0xa6, 0x4e, 0x4a, 0xd5, 0x9d, 0xcf, 0xcc, 0x79, 0x9f, 0xf3, 0x33, 0xaf, 0xc9, 0xc3, 0xc1, 0xae,
	0x64, 0xbe, 0xb0, 0x06, 0xa9, 0x0b, 0x49, 0x04, 0x08, 0xd2, 0x8a, 0x07, 0x7d, 0x8b, 0xc7, 0xbe,
	0xb4, 0x78, 0x1c, 0x4b, 0x6b, 0xd8, 0x71, 0x01, 0x79, 0xc7, 0xea, 0x43, 0x04, 0x09, 0x47, 0xe8,
	0xb1, 0x38, 0x11, 0x28, 0xe8, 0xbd, 0xb1, 0x90, 0xcd, 0x84, 0x2c, 0x1e, 0xf4, 0x59, 0x26, 0x64,
	0x99, 0x90, 0x69, 0x61, 0x73, 0xb3, 0xef, 0xe3, 0x51, 0xea, 0x32, 0x4f, 0x84, 0x56, 0x5f, 0xf4,
	0x85, 0xa5, 0xf4, 0x6e, 0x7a, 0xa8, 0x22, 0x15, 0xa8, 0xaf, 0x31, 0xb7, 0xb9, 0xb5, 0xb0, 0x21,
	0x2b, 0x01, 0x29, 0xd2, 0xc4, 0x83, 0xf3, 0xbd, 0x34, 0x77, 0x16, 0x6b, 0xd2, 0x68, 0x08, 0x89,
	0xf4, 0x45, 0x04, 0xbd, 0x82, 0x6c, 0x63, 0xb1, 0x6c, 0x58, 0x18, 0xb8, 0xb9, 0x39, 0x3f, 0x3b,
	0x49, 0x23, 0xf4, 0xc3, 0x62, 0x4f, 0xdb, 0x17, 0xa7, 0x4b, 0xef, 0x08, 0x42, 0x5e, 0x50, 0x75,
	0xe6, 0xab, 0x52, 0xf4, 0x03, 0xcb, 0x8f, 0x50, 0x62, 0x72, 0x5e, 0xd2, 0xfa, 0x56, 0x26, 0xf5,
	0x2e, 0x72, 0x84, 0xc3, 0x34, 0xe8, 0x02, 0xd2, 0x37, 0xa4, 0x16, 0x02, 0xf2, 0x1e, 0x47, 0xde,
	0x30, 0xd6, 0x8d, 0x76, 0x7d, 0xab, 0xcd, 0x16, 0xbe, 0x15, 0x1b, 0x76, 0xd8, 0x9e, 0xfb, 0x11,
	0x3c, 0x7c, 0x01, 0xc8, 0x6d, 0x7a, 0x7c, 0xba, 0x56, 0x1a, 0x9d, 0xae, 0x91, 0xd9, 0x99, 0x33,
	0xa5, 0xd1, 0x03, 0xb2, 0x24, 0x63, 0xf0, 0x1a, 0x65, 0x45, 0xdd, 0x65, 0x97, 0x74, 0x00, 0xcb,
	0x75, 0xd7, 0x8d, 0xc1, 0xb3, 0x57, 0x74, 0x95, 0xa5, 0x2c, 0x72, 0x14, 0x93, 0xba, 0x64, 0x59,
	0x22, 0xc7, 0x54, 0x36, 0x2a, 0x8a, 0xfe, 0xe8, 0x4a, 0x74, 0x45, 0xb0, 0xaf, 0x6b, 0xfe, 0xf2,
	0x38, 0x76, 0x34, 0xb9, 0xf5, 0xc3, 0x20, 0x37, 0x72, 0xd9, 0xcf, 0x7d, 0x89, 0xf4, 0x7d, 0x61,
	0x5b, 0xd6, 0x05, 0xdb, 0xca, 0xb9, 0x89, 0x65, 0x72, 0xb5, 0xb4, 0x9b, 0xba, 0x5c, 0x6d, 0x72,
	0x92, 0x5b, 0xd9, 0x5b, 0x52, 0xf5, 0x11, 0x42, 0xd9, 0x28, 0xaf, 0x57, 0xda, 0xf5, 0xad, 0xed,
	0xab, 0x4c, 0x65, 0x5f, 0xd3, 0x05, 0xaa, 0xcf, 0x32, 0x94, 0x33, 0x26, 0xb6, 0xbe, 0x57, 0xfe,
	0x9a, 0x26, 0xdb, 0x25, 0x6d, 0x93, 0x5a, 0x02, 0x71, 0xe0, 0x7b, 0x5c, 0xaa, 0x69, 0xaa, 0xf6,
	0x4a, 0xd6, 0x98, 0xa3, 0xcf, 0x9c, 0xe9, 0x2d, 0xfd, 0x40, 0x6a, 0x12, 0x02, 0xf0, 0x50, 0x24,
	0xfa, 0x3d, 0xb7, 0x2f, 0x3b, 0x37, 0x77, 0x21, 0xe8, 0x6a, 0xed, 0x98, 0x3f, 0x89, 0x9c, 0x29,
	0x93, 0xbe, 0x23, 0x35, 0x84, 0x30, 0x0e, 0x38, 0x82, 0x7e, 0xd1, 0xcd, 0x8b, 0x5d, 0xf8, 0x4a,
	0xf4, 0x5e, 0x6b, 0x81, 0x32, 0xc9, 0x74, 0xab, 0x93, 0x53, 0x67, 0x0a, 0xa4, 0x9f, 0x0d, 0xb2,
	0x3a, 0x14, 0x41, 0x1a, 0xc2, 0xe3, 0x80, 0xfb, 0xe1, 0x24, 0x43, 0x36, 0x96, 0xd4, 0x96, 0x1f,
	0xfc, 0xa3, 0x52, 0x36, 0x8a, 0x44, 0x88, 0x70, 0x7f, 0xc6, 0xb0, 0xef, 0xea, 0x7a, 0xab, 0xfb,
	0x73, 0xc0, 0xce, 0xdc, 0x72, 0x74, 0x87, 0xd4, 0x25, 0x24, 0x43, 0xdf, 0x83, 0x97, 0x3c, 0x84,
	0x46, 0x75, 0xdd, 0x68, 0xff, 0x6f, 0xdf, 0xd6, 0xa0, 0x7a, 0x77, 0x76, 0xe5, 0xe4, 0xf3, 0x5a,
	0x5f, 0x0c, 0x72, 0xab, 0xe0, 0x5a, 0xfa, 0x84, 0x50, 0xe1, 0x66, 0x69, 0xd0, 0x7b, 0x3a, 0xfe,
	0xc5, 0x7d, 0x11, 0xa9, 0x57, 0xac, 0xd8, 0x77, 0x46, 0xa7, 0x6b, 0x74, 0xaf, 0x70, 0xeb, 0xcc,
	0x51, 0xd0, 0x8d, 0x9c, 0x07, 0xca, 0xca, 0x03, 0xd3, 0x55, 0x16, 0x7d, 0x60, 0xdf, 0x3f, 0x3e,
	0x33, 0x4b, 0x27, 0x67, 0x66, 0xe9, 0xe7, 0x99, 0x59, 0xfa, 0x34, 0x32, 0x8d, 0xe3, 0x91, 0x69,
	0x9c, 0x8c, 0x4c, 0xe3, 0xd7, 0xc8, 0x34, 0xbe, 0xfe, 0x36, 0x4b, 0x07, 0xff, 0x69, 0x4b, 0xfe,
	0x09, 0x00, 0x00, 0xff, 0xff, 0x64, 0x32, 0x5a, 0xad, 0x2b, 0x06, 0x00, 0x00,
}
