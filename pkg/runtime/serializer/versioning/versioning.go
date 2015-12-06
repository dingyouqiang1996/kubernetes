/*
Copyright 2014 The Kubernetes Authors All rights reserved.

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

package versioning

import (
	"errors"
	"fmt"
	"io"

	"k8s.io/kubernetes/pkg/api/unversioned"
	"k8s.io/kubernetes/pkg/conversion"
	"k8s.io/kubernetes/pkg/runtime"
)

// ExtensibleConvertor is an object that can specialize an ObjectConvertor for a particular
// context.
type ExtensibleConvertor interface {
	WithConversions(fns *conversion.ConversionFuncs) runtime.ObjectConvertor
}

// NewCodecForScheme is a convenience method for callers that are using a scheme.
func NewCodecForScheme(
	// TODO: I should be a scheme interface?
	scheme *runtime.Scheme,
	serializer runtime.Serializer,
	encodeVersion []unversioned.GroupVersion,
	decodeVersion []unversioned.GroupVersion,
) runtime.Codec {
	return NewCodec(scheme, serializer, scheme, runtime.ObjectTyperToTyper(scheme), encodeVersion, decodeVersion)
}

// NewCodec takes objects in their internal versions and converts them to external versions before
// serializing them. It assumes the serializer provided to it only deals with external versions.
// This class is also a serializer, but is generally used with a specific version.
func NewCodec(
	extensibleConvertor ExtensibleConvertor,
	serializer runtime.Serializer,
	creater runtime.ObjectCreater,
	typer runtime.Typer,
	encodeVersion []unversioned.GroupVersion,
	decodeVersion []unversioned.GroupVersion,
) runtime.Codec {
	internal := &codec{
		serializer: serializer,
		creater:    creater,
		typer:      typer,
	}
	// when we perform conversions of runtime.Object -> []byte, use this serializer to encode and decode the
	// objects
	internal.convertor = extensibleConvertor.WithConversions(runtime.NewSerializedConversions(internal))
	if encodeVersion != nil {
		internal.encodeVersion = make(map[string]unversioned.GroupVersion)
		for _, v := range encodeVersion {
			internal.encodeVersion[v.Group] = v
		}
	}
	if decodeVersion != nil {
		internal.decodeVersion = make(map[string]unversioned.GroupVersion)
		for _, v := range decodeVersion {
			internal.decodeVersion[v.Group] = v
		}
	}

	return internal
}

type codec struct {
	serializer runtime.Serializer
	convertor  runtime.ObjectConvertor
	creater    runtime.ObjectCreater
	typer      runtime.Typer

	encodeVersion map[string]unversioned.GroupVersion
	decodeVersion map[string]unversioned.GroupVersion
}

// Decode attempts a decode of the object, then tries to convert it to the internal version.
func (c *codec) Decode(data []byte, defaultGVK *unversioned.GroupVersionKind, into runtime.Object) (runtime.Object, *unversioned.GroupVersionKind, error) {
	obj, gvk, err := c.serializer.Decode(data, defaultGVK, into)
	if err != nil {
		return nil, gvk, err
	}

	// if we specify a target, use generic conversion.
	if into != nil {
		if err := c.convertor.Convert(obj, into); err != nil {
			return nil, gvk, err
		}
		return into, gvk, nil
	}

	// invoke a version conversion
	group := gvk.Group
	if defaultGVK != nil {
		group = defaultGVK.Group
	}
	var targetGV unversioned.GroupVersion
	if c.decodeVersion == nil {
		// convert to internal by default
		targetGV.Group = group
		targetGV.Version = runtime.APIVersionInternal
	} else {
		gv, ok := c.decodeVersion[group]
		if !ok {
			// unknown objects are left in their original version
			return obj, gvk, nil
		}
		targetGV = gv
	}

	if gvk.GroupVersion() == targetGV {
		return obj, gvk, nil
	}

	// Convert if needed.
	out, err := c.convertor.ConvertToVersion(obj, targetGV.String())
	if err != nil {
		return nil, gvk, err
	}
	return out, gvk, nil
}

// EncodeToVersionStream
func (c *codec) EncodeToStream(obj runtime.Object, w io.Writer) error {
	gvk, err := c.typer.ObjectVersionAndKind(obj)
	if err != nil {
		return err
	}

	if c.encodeVersion == nil {
		return c.serializer.EncodeToStream(obj, w)
	}

	targetGV, ok := c.encodeVersion[gvk.Group]
	if !ok {
		return fmt.Errorf("the codec does not recognize group %s for kind %s and cannot encode it", gvk.Group, gvk.Kind)
	}

	// Perform a conversion if necessary.
	if gvk.GroupVersion() != targetGV {
		out, err := c.convertor.ConvertToVersion(obj, targetGV.String())
		if err != nil {
			return err
		}
		obj = out
	}

	return c.serializer.EncodeToStream(obj, w)
}

func NewEnforcingDecoder(codec runtime.Codec) runtime.Codec {
	return enforcingDecoder{Codec: codec}
}

type enforcingDecoder struct {
	runtime.Codec
}

func (c enforcingDecoder) Decode(data []byte, requestedGVK *unversioned.GroupVersionKind, into runtime.Object) (runtime.Object, *unversioned.GroupVersionKind, error) {
	out, gvk, err := c.Codec.Decode(data, requestedGVK, into)
	if err != nil {
		return nil, gvk, err
	}
	if requestedGVK != nil {
		if (len(requestedGVK.Group) > 0 || len(requestedGVK.Version) > 0) && gvk.GroupVersion() != requestedGVK.GroupVersion() {
			return nil, gvk, errors.New(fmt.Sprintf("the API version in the data (%s) does not match the expected API version (%v)", gvk.Kind, requestedGVK.GroupVersion()))
		}
		if len(requestedGVK.Kind) > 0 && (gvk.Kind != requestedGVK.Kind) {
			return nil, gvk, errors.New(fmt.Sprintf("the kind in the data (%s) does not match the expected kind (%v)", gvk.Kind, requestedGVK))
		}
	}
	return out, gvk, nil
}

// DefaultGroupVersionKindForObject calculates the expected outcome type for an object.
func DefaultGroupVersionKindForObject(typer runtime.Typer, obj runtime.Object, defaults ...unversioned.GroupVersionKind) (*unversioned.GroupVersionKind, error) {
	gvk, err := typer.ObjectVersionAndKind(obj)
	if err != nil {
		return gvk, err
	}
	for _, d := range defaults {
		if len(gvk.Kind) == 0 {
			// Assume objects with unset Kind fields are being unmarshalled into the
			// correct type.
			gvk.Kind = d.Kind
		}
		if len(gvk.Version) == 0 && len(gvk.Group) == 0 {
			// Assume objects with unset Version fields are being unmarshalled into the
			// correct type.
			gvk.Version = d.Version
			gvk.Group = d.Group
		}
		if len(gvk.Kind) > 0 && len(gvk.Version) > 0 {
			break
		}
	}
	return gvk, nil
}
