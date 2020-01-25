// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package image

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
	core "k8s.io/kubernetes/pkg/apis/core"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Image) DeepCopyInto(out *Image) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.DockerImageMetadata.DeepCopyInto(&out.DockerImageMetadata)
	if in.DockerImageLayers != nil {
		in, out := &in.DockerImageLayers, &out.DockerImageLayers
		*out = make([]ImageLayer, len(*in))
		copy(*out, *in)
	}
	if in.Signatures != nil {
		in, out := &in.Signatures, &out.Signatures
		*out = make([]ImageSignature, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DockerImageSignatures != nil {
		in, out := &in.DockerImageSignatures, &out.DockerImageSignatures
		*out = make([][]byte, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = make([]byte, len(*in))
				copy(*out, *in)
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Image.
func (in *Image) DeepCopy() *Image {
	if in == nil {
		return nil
	}
	out := new(Image)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Image) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageBlobReferences) DeepCopyInto(out *ImageBlobReferences) {
	*out = *in
	if in.Layers != nil {
		in, out := &in.Layers, &out.Layers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageBlobReferences.
func (in *ImageBlobReferences) DeepCopy() *ImageBlobReferences {
	if in == nil {
		return nil
	}
	out := new(ImageBlobReferences)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageImportSpec) DeepCopyInto(out *ImageImportSpec) {
	*out = *in
	out.From = in.From
	if in.To != nil {
		in, out := &in.To, &out.To
		*out = new(core.LocalObjectReference)
		**out = **in
	}
	out.ImportPolicy = in.ImportPolicy
	out.ReferencePolicy = in.ReferencePolicy
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageImportSpec.
func (in *ImageImportSpec) DeepCopy() *ImageImportSpec {
	if in == nil {
		return nil
	}
	out := new(ImageImportSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageImportStatus) DeepCopyInto(out *ImageImportStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(Image)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageImportStatus.
func (in *ImageImportStatus) DeepCopy() *ImageImportStatus {
	if in == nil {
		return nil
	}
	out := new(ImageImportStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageLayer) DeepCopyInto(out *ImageLayer) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageLayer.
func (in *ImageLayer) DeepCopy() *ImageLayer {
	if in == nil {
		return nil
	}
	out := new(ImageLayer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageLayerData) DeepCopyInto(out *ImageLayerData) {
	*out = *in
	if in.LayerSize != nil {
		in, out := &in.LayerSize, &out.LayerSize
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageLayerData.
func (in *ImageLayerData) DeepCopy() *ImageLayerData {
	if in == nil {
		return nil
	}
	out := new(ImageLayerData)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageList) DeepCopyInto(out *ImageList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Image, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageList.
func (in *ImageList) DeepCopy() *ImageList {
	if in == nil {
		return nil
	}
	out := new(ImageList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ImageList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageLookupPolicy) DeepCopyInto(out *ImageLookupPolicy) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageLookupPolicy.
func (in *ImageLookupPolicy) DeepCopy() *ImageLookupPolicy {
	if in == nil {
		return nil
	}
	out := new(ImageLookupPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageSignature) DeepCopyInto(out *ImageSignature) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Content != nil {
		in, out := &in.Content, &out.Content
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]SignatureCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SignedClaims != nil {
		in, out := &in.SignedClaims, &out.SignedClaims
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Created != nil {
		in, out := &in.Created, &out.Created
		*out = (*in).DeepCopy()
	}
	if in.IssuedBy != nil {
		in, out := &in.IssuedBy, &out.IssuedBy
		*out = new(SignatureIssuer)
		**out = **in
	}
	if in.IssuedTo != nil {
		in, out := &in.IssuedTo, &out.IssuedTo
		*out = new(SignatureSubject)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageSignature.
func (in *ImageSignature) DeepCopy() *ImageSignature {
	if in == nil {
		return nil
	}
	out := new(ImageSignature)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ImageSignature) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageStream) DeepCopyInto(out *ImageStream) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageStream.
func (in *ImageStream) DeepCopy() *ImageStream {
	if in == nil {
		return nil
	}
	out := new(ImageStream)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ImageStream) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageStreamImage) DeepCopyInto(out *ImageStreamImage) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Image.DeepCopyInto(&out.Image)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageStreamImage.
func (in *ImageStreamImage) DeepCopy() *ImageStreamImage {
	if in == nil {
		return nil
	}
	out := new(ImageStreamImage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ImageStreamImage) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageStreamImport) DeepCopyInto(out *ImageStreamImport) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageStreamImport.
func (in *ImageStreamImport) DeepCopy() *ImageStreamImport {
	if in == nil {
		return nil
	}
	out := new(ImageStreamImport)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ImageStreamImport) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageStreamImportSpec) DeepCopyInto(out *ImageStreamImportSpec) {
	*out = *in
	if in.Repository != nil {
		in, out := &in.Repository, &out.Repository
		*out = new(RepositoryImportSpec)
		**out = **in
	}
	if in.Images != nil {
		in, out := &in.Images, &out.Images
		*out = make([]ImageImportSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageStreamImportSpec.
func (in *ImageStreamImportSpec) DeepCopy() *ImageStreamImportSpec {
	if in == nil {
		return nil
	}
	out := new(ImageStreamImportSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageStreamImportStatus) DeepCopyInto(out *ImageStreamImportStatus) {
	*out = *in
	if in.Import != nil {
		in, out := &in.Import, &out.Import
		*out = new(ImageStream)
		(*in).DeepCopyInto(*out)
	}
	if in.Repository != nil {
		in, out := &in.Repository, &out.Repository
		*out = new(RepositoryImportStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.Images != nil {
		in, out := &in.Images, &out.Images
		*out = make([]ImageImportStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageStreamImportStatus.
func (in *ImageStreamImportStatus) DeepCopy() *ImageStreamImportStatus {
	if in == nil {
		return nil
	}
	out := new(ImageStreamImportStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageStreamLayers) DeepCopyInto(out *ImageStreamLayers) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Blobs != nil {
		in, out := &in.Blobs, &out.Blobs
		*out = make(map[string]ImageLayerData, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.Images != nil {
		in, out := &in.Images, &out.Images
		*out = make(map[string]ImageBlobReferences, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageStreamLayers.
func (in *ImageStreamLayers) DeepCopy() *ImageStreamLayers {
	if in == nil {
		return nil
	}
	out := new(ImageStreamLayers)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ImageStreamLayers) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageStreamList) DeepCopyInto(out *ImageStreamList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ImageStream, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageStreamList.
func (in *ImageStreamList) DeepCopy() *ImageStreamList {
	if in == nil {
		return nil
	}
	out := new(ImageStreamList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ImageStreamList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageStreamMapping) DeepCopyInto(out *ImageStreamMapping) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Image.DeepCopyInto(&out.Image)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageStreamMapping.
func (in *ImageStreamMapping) DeepCopy() *ImageStreamMapping {
	if in == nil {
		return nil
	}
	out := new(ImageStreamMapping)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ImageStreamMapping) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageStreamSpec) DeepCopyInto(out *ImageStreamSpec) {
	*out = *in
	out.LookupPolicy = in.LookupPolicy
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]TagReference, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageStreamSpec.
func (in *ImageStreamSpec) DeepCopy() *ImageStreamSpec {
	if in == nil {
		return nil
	}
	out := new(ImageStreamSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageStreamStatus) DeepCopyInto(out *ImageStreamStatus) {
	*out = *in
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]TagEventList, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageStreamStatus.
func (in *ImageStreamStatus) DeepCopy() *ImageStreamStatus {
	if in == nil {
		return nil
	}
	out := new(ImageStreamStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageStreamTag) DeepCopyInto(out *ImageStreamTag) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Tag != nil {
		in, out := &in.Tag, &out.Tag
		*out = new(TagReference)
		(*in).DeepCopyInto(*out)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]TagEventCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.LookupPolicy = in.LookupPolicy
	in.Image.DeepCopyInto(&out.Image)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageStreamTag.
func (in *ImageStreamTag) DeepCopy() *ImageStreamTag {
	if in == nil {
		return nil
	}
	out := new(ImageStreamTag)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ImageStreamTag) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageStreamTagList) DeepCopyInto(out *ImageStreamTagList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ImageStreamTag, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageStreamTagList.
func (in *ImageStreamTagList) DeepCopy() *ImageStreamTagList {
	if in == nil {
		return nil
	}
	out := new(ImageStreamTagList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ImageStreamTagList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageTag) DeepCopyInto(out *ImageTag) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Spec != nil {
		in, out := &in.Spec, &out.Spec
		*out = new(TagReference)
		(*in).DeepCopyInto(*out)
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(NamedTagEventList)
		(*in).DeepCopyInto(*out)
	}
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(Image)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageTag.
func (in *ImageTag) DeepCopy() *ImageTag {
	if in == nil {
		return nil
	}
	out := new(ImageTag)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ImageTag) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageTagList) DeepCopyInto(out *ImageTagList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ImageTag, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageTagList.
func (in *ImageTagList) DeepCopy() *ImageTagList {
	if in == nil {
		return nil
	}
	out := new(ImageTagList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ImageTagList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamedTagEventList) DeepCopyInto(out *NamedTagEventList) {
	*out = *in
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TagEvent, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]TagEventCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamedTagEventList.
func (in *NamedTagEventList) DeepCopy() *NamedTagEventList {
	if in == nil {
		return nil
	}
	out := new(NamedTagEventList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RepositoryImportSpec) DeepCopyInto(out *RepositoryImportSpec) {
	*out = *in
	out.From = in.From
	out.ImportPolicy = in.ImportPolicy
	out.ReferencePolicy = in.ReferencePolicy
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RepositoryImportSpec.
func (in *RepositoryImportSpec) DeepCopy() *RepositoryImportSpec {
	if in == nil {
		return nil
	}
	out := new(RepositoryImportSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RepositoryImportStatus) DeepCopyInto(out *RepositoryImportStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	if in.Images != nil {
		in, out := &in.Images, &out.Images
		*out = make([]ImageImportStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AdditionalTags != nil {
		in, out := &in.AdditionalTags, &out.AdditionalTags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RepositoryImportStatus.
func (in *RepositoryImportStatus) DeepCopy() *RepositoryImportStatus {
	if in == nil {
		return nil
	}
	out := new(RepositoryImportStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SignatureCondition) DeepCopyInto(out *SignatureCondition) {
	*out = *in
	in.LastProbeTime.DeepCopyInto(&out.LastProbeTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SignatureCondition.
func (in *SignatureCondition) DeepCopy() *SignatureCondition {
	if in == nil {
		return nil
	}
	out := new(SignatureCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SignatureGenericEntity) DeepCopyInto(out *SignatureGenericEntity) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SignatureGenericEntity.
func (in *SignatureGenericEntity) DeepCopy() *SignatureGenericEntity {
	if in == nil {
		return nil
	}
	out := new(SignatureGenericEntity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SignatureIssuer) DeepCopyInto(out *SignatureIssuer) {
	*out = *in
	out.SignatureGenericEntity = in.SignatureGenericEntity
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SignatureIssuer.
func (in *SignatureIssuer) DeepCopy() *SignatureIssuer {
	if in == nil {
		return nil
	}
	out := new(SignatureIssuer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SignatureSubject) DeepCopyInto(out *SignatureSubject) {
	*out = *in
	out.SignatureGenericEntity = in.SignatureGenericEntity
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SignatureSubject.
func (in *SignatureSubject) DeepCopy() *SignatureSubject {
	if in == nil {
		return nil
	}
	out := new(SignatureSubject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TagEvent) DeepCopyInto(out *TagEvent) {
	*out = *in
	in.Created.DeepCopyInto(&out.Created)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TagEvent.
func (in *TagEvent) DeepCopy() *TagEvent {
	if in == nil {
		return nil
	}
	out := new(TagEvent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TagEventCondition) DeepCopyInto(out *TagEventCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TagEventCondition.
func (in *TagEventCondition) DeepCopy() *TagEventCondition {
	if in == nil {
		return nil
	}
	out := new(TagEventCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TagEventList) DeepCopyInto(out *TagEventList) {
	*out = *in
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TagEvent, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]TagEventCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TagEventList.
func (in *TagEventList) DeepCopy() *TagEventList {
	if in == nil {
		return nil
	}
	out := new(TagEventList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TagImportPolicy) DeepCopyInto(out *TagImportPolicy) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TagImportPolicy.
func (in *TagImportPolicy) DeepCopy() *TagImportPolicy {
	if in == nil {
		return nil
	}
	out := new(TagImportPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TagReference) DeepCopyInto(out *TagReference) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.From != nil {
		in, out := &in.From, &out.From
		*out = new(core.ObjectReference)
		**out = **in
	}
	if in.Generation != nil {
		in, out := &in.Generation, &out.Generation
		*out = new(int64)
		**out = **in
	}
	out.ImportPolicy = in.ImportPolicy
	out.ReferencePolicy = in.ReferencePolicy
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TagReference.
func (in *TagReference) DeepCopy() *TagReference {
	if in == nil {
		return nil
	}
	out := new(TagReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TagReferencePolicy) DeepCopyInto(out *TagReferencePolicy) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TagReferencePolicy.
func (in *TagReferencePolicy) DeepCopy() *TagReferencePolicy {
	if in == nil {
		return nil
	}
	out := new(TagReferencePolicy)
	in.DeepCopyInto(out)
	return out
}
