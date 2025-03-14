//go:build !ignore_autogenerated

// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OneRequest) DeepCopyInto(out *OneRequest) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OneRequest.
func (in *OneRequest) DeepCopy() *OneRequest {
	if in == nil {
		return nil
	}
	out := new(OneRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OneRequest) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OneRequestAsset) DeepCopyInto(out *OneRequestAsset) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OneRequestAsset.
func (in *OneRequestAsset) DeepCopy() *OneRequestAsset {
	if in == nil {
		return nil
	}
	out := new(OneRequestAsset)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OneRequestAsset) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OneRequestAssetInitParameters) DeepCopyInto(out *OneRequestAssetInitParameters) {
	*out = *in
	if in.AccountIdentifier != nil {
		in, out := &in.AccountIdentifier, &out.AccountIdentifier
		*out = new(string)
		**out = **in
	}
	if in.Page != nil {
		in, out := &in.Page, &out.Page
		*out = new(float64)
		**out = **in
	}
	if in.PerPage != nil {
		in, out := &in.PerPage, &out.PerPage
		*out = new(float64)
		**out = **in
	}
	if in.RequestIdentifier != nil {
		in, out := &in.RequestIdentifier, &out.RequestIdentifier
		*out = new(string)
		**out = **in
	}
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OneRequestAssetInitParameters.
func (in *OneRequestAssetInitParameters) DeepCopy() *OneRequestAssetInitParameters {
	if in == nil {
		return nil
	}
	out := new(OneRequestAssetInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OneRequestAssetList) DeepCopyInto(out *OneRequestAssetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OneRequestAsset, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OneRequestAssetList.
func (in *OneRequestAssetList) DeepCopy() *OneRequestAssetList {
	if in == nil {
		return nil
	}
	out := new(OneRequestAssetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OneRequestAssetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OneRequestAssetObservation) DeepCopyInto(out *OneRequestAssetObservation) {
	*out = *in
	if in.AccountIdentifier != nil {
		in, out := &in.AccountIdentifier, &out.AccountIdentifier
		*out = new(string)
		**out = **in
	}
	if in.Created != nil {
		in, out := &in.Created, &out.Created
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.FileType != nil {
		in, out := &in.FileType, &out.FileType
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Page != nil {
		in, out := &in.Page, &out.Page
		*out = new(float64)
		**out = **in
	}
	if in.PerPage != nil {
		in, out := &in.PerPage, &out.PerPage
		*out = new(float64)
		**out = **in
	}
	if in.RequestIdentifier != nil {
		in, out := &in.RequestIdentifier, &out.RequestIdentifier
		*out = new(string)
		**out = **in
	}
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OneRequestAssetObservation.
func (in *OneRequestAssetObservation) DeepCopy() *OneRequestAssetObservation {
	if in == nil {
		return nil
	}
	out := new(OneRequestAssetObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OneRequestAssetParameters) DeepCopyInto(out *OneRequestAssetParameters) {
	*out = *in
	if in.AccountIdentifier != nil {
		in, out := &in.AccountIdentifier, &out.AccountIdentifier
		*out = new(string)
		**out = **in
	}
	if in.Page != nil {
		in, out := &in.Page, &out.Page
		*out = new(float64)
		**out = **in
	}
	if in.PerPage != nil {
		in, out := &in.PerPage, &out.PerPage
		*out = new(float64)
		**out = **in
	}
	if in.RequestIdentifier != nil {
		in, out := &in.RequestIdentifier, &out.RequestIdentifier
		*out = new(string)
		**out = **in
	}
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OneRequestAssetParameters.
func (in *OneRequestAssetParameters) DeepCopy() *OneRequestAssetParameters {
	if in == nil {
		return nil
	}
	out := new(OneRequestAssetParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OneRequestAssetSpec) DeepCopyInto(out *OneRequestAssetSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OneRequestAssetSpec.
func (in *OneRequestAssetSpec) DeepCopy() *OneRequestAssetSpec {
	if in == nil {
		return nil
	}
	out := new(OneRequestAssetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OneRequestAssetStatus) DeepCopyInto(out *OneRequestAssetStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OneRequestAssetStatus.
func (in *OneRequestAssetStatus) DeepCopy() *OneRequestAssetStatus {
	if in == nil {
		return nil
	}
	out := new(OneRequestAssetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OneRequestInitParameters) DeepCopyInto(out *OneRequestInitParameters) {
	*out = *in
	if in.AccountIdentifier != nil {
		in, out := &in.AccountIdentifier, &out.AccountIdentifier
		*out = new(string)
		**out = **in
	}
	if in.Content != nil {
		in, out := &in.Content, &out.Content
		*out = new(string)
		**out = **in
	}
	if in.Priority != nil {
		in, out := &in.Priority, &out.Priority
		*out = new(string)
		**out = **in
	}
	if in.RequestType != nil {
		in, out := &in.RequestType, &out.RequestType
		*out = new(string)
		**out = **in
	}
	if in.Summary != nil {
		in, out := &in.Summary, &out.Summary
		*out = new(string)
		**out = **in
	}
	if in.Tlp != nil {
		in, out := &in.Tlp, &out.Tlp
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OneRequestInitParameters.
func (in *OneRequestInitParameters) DeepCopy() *OneRequestInitParameters {
	if in == nil {
		return nil
	}
	out := new(OneRequestInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OneRequestList) DeepCopyInto(out *OneRequestList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OneRequest, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OneRequestList.
func (in *OneRequestList) DeepCopy() *OneRequestList {
	if in == nil {
		return nil
	}
	out := new(OneRequestList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OneRequestList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OneRequestMessage) DeepCopyInto(out *OneRequestMessage) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OneRequestMessage.
func (in *OneRequestMessage) DeepCopy() *OneRequestMessage {
	if in == nil {
		return nil
	}
	out := new(OneRequestMessage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OneRequestMessage) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OneRequestMessageInitParameters) DeepCopyInto(out *OneRequestMessageInitParameters) {
	*out = *in
	if in.AccountIdentifier != nil {
		in, out := &in.AccountIdentifier, &out.AccountIdentifier
		*out = new(string)
		**out = **in
	}
	if in.Content != nil {
		in, out := &in.Content, &out.Content
		*out = new(string)
		**out = **in
	}
	if in.RequestIdentifier != nil {
		in, out := &in.RequestIdentifier, &out.RequestIdentifier
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OneRequestMessageInitParameters.
func (in *OneRequestMessageInitParameters) DeepCopy() *OneRequestMessageInitParameters {
	if in == nil {
		return nil
	}
	out := new(OneRequestMessageInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OneRequestMessageList) DeepCopyInto(out *OneRequestMessageList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OneRequestMessage, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OneRequestMessageList.
func (in *OneRequestMessageList) DeepCopy() *OneRequestMessageList {
	if in == nil {
		return nil
	}
	out := new(OneRequestMessageList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OneRequestMessageList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OneRequestMessageObservation) DeepCopyInto(out *OneRequestMessageObservation) {
	*out = *in
	if in.AccountIdentifier != nil {
		in, out := &in.AccountIdentifier, &out.AccountIdentifier
		*out = new(string)
		**out = **in
	}
	if in.Author != nil {
		in, out := &in.Author, &out.Author
		*out = new(string)
		**out = **in
	}
	if in.Content != nil {
		in, out := &in.Content, &out.Content
		*out = new(string)
		**out = **in
	}
	if in.Created != nil {
		in, out := &in.Created, &out.Created
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.IsFollowOnRequest != nil {
		in, out := &in.IsFollowOnRequest, &out.IsFollowOnRequest
		*out = new(bool)
		**out = **in
	}
	if in.RequestIdentifier != nil {
		in, out := &in.RequestIdentifier, &out.RequestIdentifier
		*out = new(string)
		**out = **in
	}
	if in.Updated != nil {
		in, out := &in.Updated, &out.Updated
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OneRequestMessageObservation.
func (in *OneRequestMessageObservation) DeepCopy() *OneRequestMessageObservation {
	if in == nil {
		return nil
	}
	out := new(OneRequestMessageObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OneRequestMessageParameters) DeepCopyInto(out *OneRequestMessageParameters) {
	*out = *in
	if in.AccountIdentifier != nil {
		in, out := &in.AccountIdentifier, &out.AccountIdentifier
		*out = new(string)
		**out = **in
	}
	if in.Content != nil {
		in, out := &in.Content, &out.Content
		*out = new(string)
		**out = **in
	}
	if in.RequestIdentifier != nil {
		in, out := &in.RequestIdentifier, &out.RequestIdentifier
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OneRequestMessageParameters.
func (in *OneRequestMessageParameters) DeepCopy() *OneRequestMessageParameters {
	if in == nil {
		return nil
	}
	out := new(OneRequestMessageParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OneRequestMessageSpec) DeepCopyInto(out *OneRequestMessageSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OneRequestMessageSpec.
func (in *OneRequestMessageSpec) DeepCopy() *OneRequestMessageSpec {
	if in == nil {
		return nil
	}
	out := new(OneRequestMessageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OneRequestMessageStatus) DeepCopyInto(out *OneRequestMessageStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OneRequestMessageStatus.
func (in *OneRequestMessageStatus) DeepCopy() *OneRequestMessageStatus {
	if in == nil {
		return nil
	}
	out := new(OneRequestMessageStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OneRequestObservation) DeepCopyInto(out *OneRequestObservation) {
	*out = *in
	if in.AccountIdentifier != nil {
		in, out := &in.AccountIdentifier, &out.AccountIdentifier
		*out = new(string)
		**out = **in
	}
	if in.Completed != nil {
		in, out := &in.Completed, &out.Completed
		*out = new(string)
		**out = **in
	}
	if in.Content != nil {
		in, out := &in.Content, &out.Content
		*out = new(string)
		**out = **in
	}
	if in.Created != nil {
		in, out := &in.Created, &out.Created
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.MessageTokens != nil {
		in, out := &in.MessageTokens, &out.MessageTokens
		*out = new(float64)
		**out = **in
	}
	if in.Priority != nil {
		in, out := &in.Priority, &out.Priority
		*out = new(string)
		**out = **in
	}
	if in.ReadableID != nil {
		in, out := &in.ReadableID, &out.ReadableID
		*out = new(string)
		**out = **in
	}
	if in.Request != nil {
		in, out := &in.Request, &out.Request
		*out = new(string)
		**out = **in
	}
	if in.RequestType != nil {
		in, out := &in.RequestType, &out.RequestType
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.Summary != nil {
		in, out := &in.Summary, &out.Summary
		*out = new(string)
		**out = **in
	}
	if in.Tlp != nil {
		in, out := &in.Tlp, &out.Tlp
		*out = new(string)
		**out = **in
	}
	if in.Tokens != nil {
		in, out := &in.Tokens, &out.Tokens
		*out = new(float64)
		**out = **in
	}
	if in.Updated != nil {
		in, out := &in.Updated, &out.Updated
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OneRequestObservation.
func (in *OneRequestObservation) DeepCopy() *OneRequestObservation {
	if in == nil {
		return nil
	}
	out := new(OneRequestObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OneRequestParameters) DeepCopyInto(out *OneRequestParameters) {
	*out = *in
	if in.AccountIdentifier != nil {
		in, out := &in.AccountIdentifier, &out.AccountIdentifier
		*out = new(string)
		**out = **in
	}
	if in.Content != nil {
		in, out := &in.Content, &out.Content
		*out = new(string)
		**out = **in
	}
	if in.Priority != nil {
		in, out := &in.Priority, &out.Priority
		*out = new(string)
		**out = **in
	}
	if in.RequestType != nil {
		in, out := &in.RequestType, &out.RequestType
		*out = new(string)
		**out = **in
	}
	if in.Summary != nil {
		in, out := &in.Summary, &out.Summary
		*out = new(string)
		**out = **in
	}
	if in.Tlp != nil {
		in, out := &in.Tlp, &out.Tlp
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OneRequestParameters.
func (in *OneRequestParameters) DeepCopy() *OneRequestParameters {
	if in == nil {
		return nil
	}
	out := new(OneRequestParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OneRequestPriority) DeepCopyInto(out *OneRequestPriority) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OneRequestPriority.
func (in *OneRequestPriority) DeepCopy() *OneRequestPriority {
	if in == nil {
		return nil
	}
	out := new(OneRequestPriority)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OneRequestPriority) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OneRequestPriorityInitParameters) DeepCopyInto(out *OneRequestPriorityInitParameters) {
	*out = *in
	if in.AccountIdentifier != nil {
		in, out := &in.AccountIdentifier, &out.AccountIdentifier
		*out = new(string)
		**out = **in
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Priority != nil {
		in, out := &in.Priority, &out.Priority
		*out = new(float64)
		**out = **in
	}
	if in.Requirement != nil {
		in, out := &in.Requirement, &out.Requirement
		*out = new(string)
		**out = **in
	}
	if in.Tlp != nil {
		in, out := &in.Tlp, &out.Tlp
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OneRequestPriorityInitParameters.
func (in *OneRequestPriorityInitParameters) DeepCopy() *OneRequestPriorityInitParameters {
	if in == nil {
		return nil
	}
	out := new(OneRequestPriorityInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OneRequestPriorityList) DeepCopyInto(out *OneRequestPriorityList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OneRequestPriority, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OneRequestPriorityList.
func (in *OneRequestPriorityList) DeepCopy() *OneRequestPriorityList {
	if in == nil {
		return nil
	}
	out := new(OneRequestPriorityList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OneRequestPriorityList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OneRequestPriorityObservation) DeepCopyInto(out *OneRequestPriorityObservation) {
	*out = *in
	if in.AccountIdentifier != nil {
		in, out := &in.AccountIdentifier, &out.AccountIdentifier
		*out = new(string)
		**out = **in
	}
	if in.Completed != nil {
		in, out := &in.Completed, &out.Completed
		*out = new(string)
		**out = **in
	}
	if in.Content != nil {
		in, out := &in.Content, &out.Content
		*out = new(string)
		**out = **in
	}
	if in.Created != nil {
		in, out := &in.Created, &out.Created
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.MessageTokens != nil {
		in, out := &in.MessageTokens, &out.MessageTokens
		*out = new(float64)
		**out = **in
	}
	if in.Priority != nil {
		in, out := &in.Priority, &out.Priority
		*out = new(float64)
		**out = **in
	}
	if in.ReadableID != nil {
		in, out := &in.ReadableID, &out.ReadableID
		*out = new(string)
		**out = **in
	}
	if in.Request != nil {
		in, out := &in.Request, &out.Request
		*out = new(string)
		**out = **in
	}
	if in.Requirement != nil {
		in, out := &in.Requirement, &out.Requirement
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.Summary != nil {
		in, out := &in.Summary, &out.Summary
		*out = new(string)
		**out = **in
	}
	if in.Tlp != nil {
		in, out := &in.Tlp, &out.Tlp
		*out = new(string)
		**out = **in
	}
	if in.Tokens != nil {
		in, out := &in.Tokens, &out.Tokens
		*out = new(float64)
		**out = **in
	}
	if in.Updated != nil {
		in, out := &in.Updated, &out.Updated
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OneRequestPriorityObservation.
func (in *OneRequestPriorityObservation) DeepCopy() *OneRequestPriorityObservation {
	if in == nil {
		return nil
	}
	out := new(OneRequestPriorityObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OneRequestPriorityParameters) DeepCopyInto(out *OneRequestPriorityParameters) {
	*out = *in
	if in.AccountIdentifier != nil {
		in, out := &in.AccountIdentifier, &out.AccountIdentifier
		*out = new(string)
		**out = **in
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Priority != nil {
		in, out := &in.Priority, &out.Priority
		*out = new(float64)
		**out = **in
	}
	if in.Requirement != nil {
		in, out := &in.Requirement, &out.Requirement
		*out = new(string)
		**out = **in
	}
	if in.Tlp != nil {
		in, out := &in.Tlp, &out.Tlp
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OneRequestPriorityParameters.
func (in *OneRequestPriorityParameters) DeepCopy() *OneRequestPriorityParameters {
	if in == nil {
		return nil
	}
	out := new(OneRequestPriorityParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OneRequestPrioritySpec) DeepCopyInto(out *OneRequestPrioritySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OneRequestPrioritySpec.
func (in *OneRequestPrioritySpec) DeepCopy() *OneRequestPrioritySpec {
	if in == nil {
		return nil
	}
	out := new(OneRequestPrioritySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OneRequestPriorityStatus) DeepCopyInto(out *OneRequestPriorityStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OneRequestPriorityStatus.
func (in *OneRequestPriorityStatus) DeepCopy() *OneRequestPriorityStatus {
	if in == nil {
		return nil
	}
	out := new(OneRequestPriorityStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OneRequestSpec) DeepCopyInto(out *OneRequestSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OneRequestSpec.
func (in *OneRequestSpec) DeepCopy() *OneRequestSpec {
	if in == nil {
		return nil
	}
	out := new(OneRequestSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OneRequestStatus) DeepCopyInto(out *OneRequestStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OneRequestStatus.
func (in *OneRequestStatus) DeepCopy() *OneRequestStatus {
	if in == nil {
		return nil
	}
	out := new(OneRequestStatus)
	in.DeepCopyInto(out)
	return out
}
