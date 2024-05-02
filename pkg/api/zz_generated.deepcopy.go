//go:build !ignore_autogenerated

// Copyright 2020-2023 Project Capsule Authors.
// SPDX-License-Identifier: Apache-2.0

// Code generated by controller-gen. DO NOT EDIT.

package api

import (
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	"k8s.io/api/rbac/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdditionalMetadataSpec) DeepCopyInto(out *AdditionalMetadataSpec) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdditionalMetadataSpec.
func (in *AdditionalMetadataSpec) DeepCopy() *AdditionalMetadataSpec {
	if in == nil {
		return nil
	}
	out := new(AdditionalMetadataSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdditionalRoleBindingsSpec) DeepCopyInto(out *AdditionalRoleBindingsSpec) {
	*out = *in
	if in.Subjects != nil {
		in, out := &in.Subjects, &out.Subjects
		*out = make([]v1.Subject, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdditionalRoleBindingsSpec.
func (in *AdditionalRoleBindingsSpec) DeepCopy() *AdditionalRoleBindingsSpec {
	if in == nil {
		return nil
	}
	out := new(AdditionalRoleBindingsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllowedListSpec) DeepCopyInto(out *AllowedListSpec) {
	*out = *in
	if in.Exact != nil {
		in, out := &in.Exact, &out.Exact
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllowedListSpec.
func (in *AllowedListSpec) DeepCopy() *AllowedListSpec {
	if in == nil {
		return nil
	}
	out := new(AllowedListSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllowedServices) DeepCopyInto(out *AllowedServices) {
	*out = *in
	if in.NodePort != nil {
		in, out := &in.NodePort, &out.NodePort
		*out = new(bool)
		**out = **in
	}
	if in.ExternalName != nil {
		in, out := &in.ExternalName, &out.ExternalName
		*out = new(bool)
		**out = **in
	}
	if in.LoadBalancer != nil {
		in, out := &in.LoadBalancer, &out.LoadBalancer
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllowedServices.
func (in *AllowedServices) DeepCopy() *AllowedServices {
	if in == nil {
		return nil
	}
	out := new(AllowedServices)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultAllowedListSpec) DeepCopyInto(out *DefaultAllowedListSpec) {
	*out = *in
	in.SelectorAllowedListSpec.DeepCopyInto(&out.SelectorAllowedListSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultAllowedListSpec.
func (in *DefaultAllowedListSpec) DeepCopy() *DefaultAllowedListSpec {
	if in == nil {
		return nil
	}
	out := new(DefaultAllowedListSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalServiceIPsSpec) DeepCopyInto(out *ExternalServiceIPsSpec) {
	*out = *in
	if in.Allowed != nil {
		in, out := &in.Allowed, &out.Allowed
		*out = make([]AllowedIP, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalServiceIPsSpec.
func (in *ExternalServiceIPsSpec) DeepCopy() *ExternalServiceIPsSpec {
	if in == nil {
		return nil
	}
	out := new(ExternalServiceIPsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ForbiddenListSpec) DeepCopyInto(out *ForbiddenListSpec) {
	*out = *in
	if in.Exact != nil {
		in, out := &in.Exact, &out.Exact
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ForbiddenListSpec.
func (in *ForbiddenListSpec) DeepCopy() *ForbiddenListSpec {
	if in == nil {
		return nil
	}
	out := new(ForbiddenListSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LimitRangesSpec) DeepCopyInto(out *LimitRangesSpec) {
	*out = *in
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]corev1.LimitRangeSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LimitRangesSpec.
func (in *LimitRangesSpec) DeepCopy() *LimitRangesSpec {
	if in == nil {
		return nil
	}
	out := new(LimitRangesSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkPolicySpec) DeepCopyInto(out *NetworkPolicySpec) {
	*out = *in
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]networkingv1.NetworkPolicySpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkPolicySpec.
func (in *NetworkPolicySpec) DeepCopy() *NetworkPolicySpec {
	if in == nil {
		return nil
	}
	out := new(NetworkPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodOptions) DeepCopyInto(out *PodOptions) {
	*out = *in
	if in.AdditionalMetadata != nil {
		in, out := &in.AdditionalMetadata, &out.AdditionalMetadata
		*out = new(AdditionalMetadataSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodOptions.
func (in *PodOptions) DeepCopy() *PodOptions {
	if in == nil {
		return nil
	}
	out := new(PodOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceQuotaSpec) DeepCopyInto(out *ResourceQuotaSpec) {
	*out = *in
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]corev1.ResourceQuotaSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceQuotaSpec.
func (in *ResourceQuotaSpec) DeepCopy() *ResourceQuotaSpec {
	if in == nil {
		return nil
	}
	out := new(ResourceQuotaSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SelectorAllowedListSpec) DeepCopyInto(out *SelectorAllowedListSpec) {
	*out = *in
	in.AllowedListSpec.DeepCopyInto(&out.AllowedListSpec)
	in.LabelSelector.DeepCopyInto(&out.LabelSelector)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SelectorAllowedListSpec.
func (in *SelectorAllowedListSpec) DeepCopy() *SelectorAllowedListSpec {
	if in == nil {
		return nil
	}
	out := new(SelectorAllowedListSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceOptions) DeepCopyInto(out *ServiceOptions) {
	*out = *in
	if in.AdditionalMetadata != nil {
		in, out := &in.AdditionalMetadata, &out.AdditionalMetadata
		*out = new(AdditionalMetadataSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.AllowedServices != nil {
		in, out := &in.AllowedServices, &out.AllowedServices
		*out = new(AllowedServices)
		(*in).DeepCopyInto(*out)
	}
	if in.ExternalServiceIPs != nil {
		in, out := &in.ExternalServiceIPs, &out.ExternalServiceIPs
		*out = new(ExternalServiceIPsSpec)
		(*in).DeepCopyInto(*out)
	}
	in.ForbiddenLabels.DeepCopyInto(&out.ForbiddenLabels)
	in.ForbiddenAnnotations.DeepCopyInto(&out.ForbiddenAnnotations)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceOptions.
func (in *ServiceOptions) DeepCopy() *ServiceOptions {
	if in == nil {
		return nil
	}
	out := new(ServiceOptions)
	in.DeepCopyInto(out)
	return out
}
