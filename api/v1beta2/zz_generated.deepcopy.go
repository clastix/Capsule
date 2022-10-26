//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright 2020-2021 Clastix Labs
// SPDX-License-Identifier: Apache-2.0

// Code generated by controller-gen. DO NOT EDIT.

package v1beta2

import (
	"github.com/clastix/capsule/pkg/api"
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	"k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
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
func (in ByKindAndName) DeepCopyInto(out *ByKindAndName) {
	{
		in := &in
		*out = make(ByKindAndName, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ByKindAndName.
func (in ByKindAndName) DeepCopy() ByKindAndName {
	if in == nil {
		return nil
	}
	out := new(ByKindAndName)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CapsuleConfiguration) DeepCopyInto(out *CapsuleConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CapsuleConfiguration.
func (in *CapsuleConfiguration) DeepCopy() *CapsuleConfiguration {
	if in == nil {
		return nil
	}
	out := new(CapsuleConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CapsuleConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CapsuleConfigurationList) DeepCopyInto(out *CapsuleConfigurationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CapsuleConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CapsuleConfigurationList.
func (in *CapsuleConfigurationList) DeepCopy() *CapsuleConfigurationList {
	if in == nil {
		return nil
	}
	out := new(CapsuleConfigurationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CapsuleConfigurationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CapsuleConfigurationSpec) DeepCopyInto(out *CapsuleConfigurationSpec) {
	*out = *in
	if in.UserGroups != nil {
		in, out := &in.UserGroups, &out.UserGroups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.CapsuleResources = in.CapsuleResources
	if in.NodeMetadata != nil {
		in, out := &in.NodeMetadata, &out.NodeMetadata
		*out = new(NodeMetadata)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CapsuleConfigurationSpec.
func (in *CapsuleConfigurationSpec) DeepCopy() *CapsuleConfigurationSpec {
	if in == nil {
		return nil
	}
	out := new(CapsuleConfigurationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CapsuleResources) DeepCopyInto(out *CapsuleResources) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CapsuleResources.
func (in *CapsuleResources) DeepCopy() *CapsuleResources {
	if in == nil {
		return nil
	}
	out := new(CapsuleResources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalTenantResource) DeepCopyInto(out *GlobalTenantResource) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalTenantResource.
func (in *GlobalTenantResource) DeepCopy() *GlobalTenantResource {
	if in == nil {
		return nil
	}
	out := new(GlobalTenantResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GlobalTenantResource) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalTenantResourceList) DeepCopyInto(out *GlobalTenantResourceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GlobalTenantResource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalTenantResourceList.
func (in *GlobalTenantResourceList) DeepCopy() *GlobalTenantResourceList {
	if in == nil {
		return nil
	}
	out := new(GlobalTenantResourceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GlobalTenantResourceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalTenantResourceSpec) DeepCopyInto(out *GlobalTenantResourceSpec) {
	*out = *in
	in.TenantSelector.DeepCopyInto(&out.TenantSelector)
	in.TenantResourceSpec.DeepCopyInto(&out.TenantResourceSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalTenantResourceSpec.
func (in *GlobalTenantResourceSpec) DeepCopy() *GlobalTenantResourceSpec {
	if in == nil {
		return nil
	}
	out := new(GlobalTenantResourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalTenantResourceStatus) DeepCopyInto(out *GlobalTenantResourceStatus) {
	*out = *in
	if in.SelectedTenants != nil {
		in, out := &in.SelectedTenants, &out.SelectedTenants
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ProcessedItems != nil {
		in, out := &in.ProcessedItems, &out.ProcessedItems
		*out = make(ProcessedItems, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalTenantResourceStatus.
func (in *GlobalTenantResourceStatus) DeepCopy() *GlobalTenantResourceStatus {
	if in == nil {
		return nil
	}
	out := new(GlobalTenantResourceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressOptions) DeepCopyInto(out *IngressOptions) {
	*out = *in
	if in.AllowedClasses != nil {
		in, out := &in.AllowedClasses, &out.AllowedClasses
		*out = new(api.AllowedListSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.AllowedHostnames != nil {
		in, out := &in.AllowedHostnames, &out.AllowedHostnames
		*out = new(api.AllowedListSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressOptions.
func (in *IngressOptions) DeepCopy() *IngressOptions {
	if in == nil {
		return nil
	}
	out := new(IngressOptions)
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
func (in *NamespaceOptions) DeepCopyInto(out *NamespaceOptions) {
	*out = *in
	if in.Quota != nil {
		in, out := &in.Quota, &out.Quota
		*out = new(int32)
		**out = **in
	}
	if in.AdditionalMetadata != nil {
		in, out := &in.AdditionalMetadata, &out.AdditionalMetadata
		*out = new(api.AdditionalMetadataSpec)
		(*in).DeepCopyInto(*out)
	}
	in.ForbiddenLabels.DeepCopyInto(&out.ForbiddenLabels)
	in.ForbiddenAnnotations.DeepCopyInto(&out.ForbiddenAnnotations)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceOptions.
func (in *NamespaceOptions) DeepCopy() *NamespaceOptions {
	if in == nil {
		return nil
	}
	out := new(NamespaceOptions)
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
func (in *NodeMetadata) DeepCopyInto(out *NodeMetadata) {
	*out = *in
	in.ForbiddenLabels.DeepCopyInto(&out.ForbiddenLabels)
	in.ForbiddenAnnotations.DeepCopyInto(&out.ForbiddenAnnotations)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeMetadata.
func (in *NodeMetadata) DeepCopy() *NodeMetadata {
	if in == nil {
		return nil
	}
	out := new(NodeMetadata)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NonLimitedResourceError) DeepCopyInto(out *NonLimitedResourceError) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NonLimitedResourceError.
func (in *NonLimitedResourceError) DeepCopy() *NonLimitedResourceError {
	if in == nil {
		return nil
	}
	out := new(NonLimitedResourceError)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectReference) DeepCopyInto(out *ObjectReference) {
	*out = *in
	out.ObjectReferenceAbstract = in.ObjectReferenceAbstract
	in.Selector.DeepCopyInto(&out.Selector)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectReference.
func (in *ObjectReference) DeepCopy() *ObjectReference {
	if in == nil {
		return nil
	}
	out := new(ObjectReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectReferenceAbstract) DeepCopyInto(out *ObjectReferenceAbstract) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectReferenceAbstract.
func (in *ObjectReferenceAbstract) DeepCopy() *ObjectReferenceAbstract {
	if in == nil {
		return nil
	}
	out := new(ObjectReferenceAbstract)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectReferenceStatus) DeepCopyInto(out *ObjectReferenceStatus) {
	*out = *in
	out.ObjectReferenceAbstract = in.ObjectReferenceAbstract
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectReferenceStatus.
func (in *ObjectReferenceStatus) DeepCopy() *ObjectReferenceStatus {
	if in == nil {
		return nil
	}
	out := new(ObjectReferenceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in OwnerListSpec) DeepCopyInto(out *OwnerListSpec) {
	{
		in := &in
		*out = make(OwnerListSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OwnerListSpec.
func (in OwnerListSpec) DeepCopy() OwnerListSpec {
	if in == nil {
		return nil
	}
	out := new(OwnerListSpec)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OwnerSpec) DeepCopyInto(out *OwnerSpec) {
	*out = *in
	if in.ClusterRoles != nil {
		in, out := &in.ClusterRoles, &out.ClusterRoles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ProxyOperations != nil {
		in, out := &in.ProxyOperations, &out.ProxyOperations
		*out = make([]ProxySettings, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OwnerSpec.
func (in *OwnerSpec) DeepCopy() *OwnerSpec {
	if in == nil {
		return nil
	}
	out := new(OwnerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in ProcessedItems) DeepCopyInto(out *ProcessedItems) {
	{
		in := &in
		*out = make(ProcessedItems, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProcessedItems.
func (in ProcessedItems) DeepCopy() ProcessedItems {
	if in == nil {
		return nil
	}
	out := new(ProcessedItems)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProxySettings) DeepCopyInto(out *ProxySettings) {
	*out = *in
	if in.Operations != nil {
		in, out := &in.Operations, &out.Operations
		*out = make([]ProxyOperation, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProxySettings.
func (in *ProxySettings) DeepCopy() *ProxySettings {
	if in == nil {
		return nil
	}
	out := new(ProxySettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RawExtension) DeepCopyInto(out *RawExtension) {
	*out = *in
	in.RawExtension.DeepCopyInto(&out.RawExtension)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RawExtension.
func (in *RawExtension) DeepCopy() *RawExtension {
	if in == nil {
		return nil
	}
	out := new(RawExtension)
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
func (in *ResourceSpec) DeepCopyInto(out *ResourceSpec) {
	*out = *in
	if in.NamespaceSelector != nil {
		in, out := &in.NamespaceSelector, &out.NamespaceSelector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.NamespacedItems != nil {
		in, out := &in.NamespacedItems, &out.NamespacedItems
		*out = make([]ObjectReference, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RawItems != nil {
		in, out := &in.RawItems, &out.RawItems
		*out = make([]RawExtension, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AdditionalMetadata != nil {
		in, out := &in.AdditionalMetadata, &out.AdditionalMetadata
		*out = new(AdditionalMetadataSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceSpec.
func (in *ResourceSpec) DeepCopy() *ResourceSpec {
	if in == nil {
		return nil
	}
	out := new(ResourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Tenant) DeepCopyInto(out *Tenant) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Tenant.
func (in *Tenant) DeepCopy() *Tenant {
	if in == nil {
		return nil
	}
	out := new(Tenant)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Tenant) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TenantList) DeepCopyInto(out *TenantList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Tenant, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TenantList.
func (in *TenantList) DeepCopy() *TenantList {
	if in == nil {
		return nil
	}
	out := new(TenantList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TenantList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TenantResource) DeepCopyInto(out *TenantResource) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TenantResource.
func (in *TenantResource) DeepCopy() *TenantResource {
	if in == nil {
		return nil
	}
	out := new(TenantResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TenantResource) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TenantResourceList) DeepCopyInto(out *TenantResourceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TenantResource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TenantResourceList.
func (in *TenantResourceList) DeepCopy() *TenantResourceList {
	if in == nil {
		return nil
	}
	out := new(TenantResourceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TenantResourceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TenantResourceSpec) DeepCopyInto(out *TenantResourceSpec) {
	*out = *in
	out.ResyncPeriod = in.ResyncPeriod
	if in.PruningOnDelete != nil {
		in, out := &in.PruningOnDelete, &out.PruningOnDelete
		*out = new(bool)
		**out = **in
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]ResourceSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TenantResourceSpec.
func (in *TenantResourceSpec) DeepCopy() *TenantResourceSpec {
	if in == nil {
		return nil
	}
	out := new(TenantResourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TenantResourceStatus) DeepCopyInto(out *TenantResourceStatus) {
	*out = *in
	if in.ProcessedItems != nil {
		in, out := &in.ProcessedItems, &out.ProcessedItems
		*out = make(ProcessedItems, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TenantResourceStatus.
func (in *TenantResourceStatus) DeepCopy() *TenantResourceStatus {
	if in == nil {
		return nil
	}
	out := new(TenantResourceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TenantSpec) DeepCopyInto(out *TenantSpec) {
	*out = *in
	if in.Owners != nil {
		in, out := &in.Owners, &out.Owners
		*out = make(OwnerListSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NamespaceOptions != nil {
		in, out := &in.NamespaceOptions, &out.NamespaceOptions
		*out = new(NamespaceOptions)
		(*in).DeepCopyInto(*out)
	}
	if in.ServiceOptions != nil {
		in, out := &in.ServiceOptions, &out.ServiceOptions
		*out = new(api.ServiceOptions)
		(*in).DeepCopyInto(*out)
	}
	if in.StorageClasses != nil {
		in, out := &in.StorageClasses, &out.StorageClasses
		*out = new(api.AllowedListSpec)
		(*in).DeepCopyInto(*out)
	}
	in.IngressOptions.DeepCopyInto(&out.IngressOptions)
	if in.ContainerRegistries != nil {
		in, out := &in.ContainerRegistries, &out.ContainerRegistries
		*out = new(api.AllowedListSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.NetworkPolicies.DeepCopyInto(&out.NetworkPolicies)
	in.LimitRanges.DeepCopyInto(&out.LimitRanges)
	in.ResourceQuota.DeepCopyInto(&out.ResourceQuota)
	if in.AdditionalRoleBindings != nil {
		in, out := &in.AdditionalRoleBindings, &out.AdditionalRoleBindings
		*out = make([]AdditionalRoleBindingsSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ImagePullPolicies != nil {
		in, out := &in.ImagePullPolicies, &out.ImagePullPolicies
		*out = make([]api.ImagePullPolicySpec, len(*in))
		copy(*out, *in)
	}
	if in.PriorityClasses != nil {
		in, out := &in.PriorityClasses, &out.PriorityClasses
		*out = new(api.AllowedListSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TenantSpec.
func (in *TenantSpec) DeepCopy() *TenantSpec {
	if in == nil {
		return nil
	}
	out := new(TenantSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TenantStatus) DeepCopyInto(out *TenantStatus) {
	*out = *in
	if in.Namespaces != nil {
		in, out := &in.Namespaces, &out.Namespaces
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TenantStatus.
func (in *TenantStatus) DeepCopy() *TenantStatus {
	if in == nil {
		return nil
	}
	out := new(TenantStatus)
	in.DeepCopyInto(out)
	return out
}
