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
func (in *Widget) DeepCopyInto(out *Widget) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Widget.
func (in *Widget) DeepCopy() *Widget {
	if in == nil {
		return nil
	}
	out := new(Widget)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Widget) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WidgetInitParameters) DeepCopyInto(out *WidgetInitParameters) {
	*out = *in
	if in.AccountID != nil {
		in, out := &in.AccountID, &out.AccountID
		*out = new(string)
		**out = **in
	}
	if in.BotFightMode != nil {
		in, out := &in.BotFightMode, &out.BotFightMode
		*out = new(bool)
		**out = **in
	}
	if in.ClearanceLevel != nil {
		in, out := &in.ClearanceLevel, &out.ClearanceLevel
		*out = new(string)
		**out = **in
	}
	if in.Domains != nil {
		in, out := &in.Domains, &out.Domains
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.EphemeralID != nil {
		in, out := &in.EphemeralID, &out.EphemeralID
		*out = new(bool)
		**out = **in
	}
	if in.Mode != nil {
		in, out := &in.Mode, &out.Mode
		*out = new(string)
		**out = **in
	}
	if in.Offlabel != nil {
		in, out := &in.Offlabel, &out.Offlabel
		*out = new(bool)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WidgetInitParameters.
func (in *WidgetInitParameters) DeepCopy() *WidgetInitParameters {
	if in == nil {
		return nil
	}
	out := new(WidgetInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WidgetList) DeepCopyInto(out *WidgetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Widget, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WidgetList.
func (in *WidgetList) DeepCopy() *WidgetList {
	if in == nil {
		return nil
	}
	out := new(WidgetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WidgetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WidgetObservation) DeepCopyInto(out *WidgetObservation) {
	*out = *in
	if in.AccountID != nil {
		in, out := &in.AccountID, &out.AccountID
		*out = new(string)
		**out = **in
	}
	if in.BotFightMode != nil {
		in, out := &in.BotFightMode, &out.BotFightMode
		*out = new(bool)
		**out = **in
	}
	if in.ClearanceLevel != nil {
		in, out := &in.ClearanceLevel, &out.ClearanceLevel
		*out = new(string)
		**out = **in
	}
	if in.CreatedOn != nil {
		in, out := &in.CreatedOn, &out.CreatedOn
		*out = new(string)
		**out = **in
	}
	if in.Domains != nil {
		in, out := &in.Domains, &out.Domains
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.EphemeralID != nil {
		in, out := &in.EphemeralID, &out.EphemeralID
		*out = new(bool)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Mode != nil {
		in, out := &in.Mode, &out.Mode
		*out = new(string)
		**out = **in
	}
	if in.ModifiedOn != nil {
		in, out := &in.ModifiedOn, &out.ModifiedOn
		*out = new(string)
		**out = **in
	}
	if in.Offlabel != nil {
		in, out := &in.Offlabel, &out.Offlabel
		*out = new(bool)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.Secret != nil {
		in, out := &in.Secret, &out.Secret
		*out = new(string)
		**out = **in
	}
	if in.Sitekey != nil {
		in, out := &in.Sitekey, &out.Sitekey
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WidgetObservation.
func (in *WidgetObservation) DeepCopy() *WidgetObservation {
	if in == nil {
		return nil
	}
	out := new(WidgetObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WidgetParameters) DeepCopyInto(out *WidgetParameters) {
	*out = *in
	if in.AccountID != nil {
		in, out := &in.AccountID, &out.AccountID
		*out = new(string)
		**out = **in
	}
	if in.BotFightMode != nil {
		in, out := &in.BotFightMode, &out.BotFightMode
		*out = new(bool)
		**out = **in
	}
	if in.ClearanceLevel != nil {
		in, out := &in.ClearanceLevel, &out.ClearanceLevel
		*out = new(string)
		**out = **in
	}
	if in.Domains != nil {
		in, out := &in.Domains, &out.Domains
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.EphemeralID != nil {
		in, out := &in.EphemeralID, &out.EphemeralID
		*out = new(bool)
		**out = **in
	}
	if in.Mode != nil {
		in, out := &in.Mode, &out.Mode
		*out = new(string)
		**out = **in
	}
	if in.Offlabel != nil {
		in, out := &in.Offlabel, &out.Offlabel
		*out = new(bool)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WidgetParameters.
func (in *WidgetParameters) DeepCopy() *WidgetParameters {
	if in == nil {
		return nil
	}
	out := new(WidgetParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WidgetSpec) DeepCopyInto(out *WidgetSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WidgetSpec.
func (in *WidgetSpec) DeepCopy() *WidgetSpec {
	if in == nil {
		return nil
	}
	out := new(WidgetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WidgetStatus) DeepCopyInto(out *WidgetStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WidgetStatus.
func (in *WidgetStatus) DeepCopy() *WidgetStatus {
	if in == nil {
		return nil
	}
	out := new(WidgetStatus)
	in.DeepCopyInto(out)
	return out
}
