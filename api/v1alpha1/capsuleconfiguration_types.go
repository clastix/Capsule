/*
Copyright 2020 Clastix Labs.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CapsuleConfigurationSpec defines the Capsule configuration
// nolint:maligned
type CapsuleConfigurationSpec struct {
	// Names of the groups for Capsule users.
	// +kubebuilder:default={capsule.clastix.io}
	UserGroups []string `json:"userGroups,omitempty"`
	// Enforces the Tenant owner, during Namespace creation, to name it using the selected Tenant name as prefix,
	// separated by a dash. This is useful to avoid Namespace name collision in a public CaaS environment.
	// +kubebuilder:default=false
	ForceTenantPrefix bool `json:"forceTenantPrefix,omitempty"`
	// Disallow creation of namespaces, whose name matches this regexp
	ProtectedNamespaceRegexpString string `json:"protectedNamespaceRegex,omitempty"`
	// When defining the exact match for allowed Ingress hostnames at Tenant level, a collision is not allowed.
	// Toggling this, Capsule will not check if a hostname collision is in place, allowing the creation of
	// two or more Tenant resources although sharing the same allowed hostname(s).
	//
	// The JSON path of the resource is: /spec/ingressHostnames/allowed
	AllowTenantIngressHostnamesCollision bool `json:"allowTenantIngressHostnamesCollision,omitempty"`
	// Allow the collision of Ingress resource hostnames across all the Tenants.
	// +kubebuilder:default=true
	AllowIngressHostnameCollision bool `json:"allowIngressHostnameCollision,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Cluster

// CapsuleConfiguration is the Schema for the Capsule configuration API
type CapsuleConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec CapsuleConfigurationSpec `json:"spec,omitempty"`
}

// +kubebuilder:object:root=true

// CapsuleConfigurationList contains a list of CapsuleConfiguration
type CapsuleConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CapsuleConfiguration `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CapsuleConfiguration{}, &CapsuleConfigurationList{})
}
