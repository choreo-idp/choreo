// Copyright 2025 The OpenChoreo Authors
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// WebApplicationBindingSpec defines the desired state of WebApplicationBinding.
type WebApplicationBindingSpec struct {
	// Owner defines the component and project that owns this web application binding
	Owner WebApplicationOwner `json:"owner"`

	// Environment is the target environment for this binding
	// +kubebuilder:validation:MinLength=1
	Environment string `json:"environment"`

	// ClassName is the name of the web application class that provides the web application-specific deployment configuration.
	// +kubebuilder:default=default
	ClassName string `json:"className"`

	// WorkloadSpec contains the copied workload specification for this environment-specific binding
	WorkloadSpec WorkloadTemplateSpec `json:"workloadSpec"`

	// Overrides contains web application-specific overrides for this binding
	Overrides map[string]bool `json:"overrides,omitempty"`

	// ReleaseState controls the state of the Release created by this binding.
	// Active: Resources are deployed normally
	// Suspend: Resources are suspended (scaled to zero or paused)
	// Undeploy: Resources are removed from the data plane
	// +kubebuilder:default=Active
	// +kubebuilder:validation:Enum=Active;Suspend;Undeploy
	// +optional
	ReleaseState ReleaseState `json:"releaseState,omitempty"`
}

// WebApplicationBindingStatus defines the observed state of WebApplicationBinding.
type WebApplicationBindingStatus struct {
	// Conditions represent the latest available observations of the WebApplicationBinding's current state.
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// Endpoints contain the status of each endpoint
	// +optional
	Endpoints []EndpointStatus `json:"endpoints,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// WebApplicationBinding is the Schema for the webapplicationbindings API.
type WebApplicationBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WebApplicationBindingSpec   `json:"spec,omitempty"`
	Status WebApplicationBindingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WebApplicationBindingList contains a list of WebApplicationBinding.
type WebApplicationBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WebApplicationBinding `json:"items"`
}

// GetConditions returns the conditions from the status
func (wab *WebApplicationBinding) GetConditions() []metav1.Condition {
	return wab.Status.Conditions
}

// SetConditions sets the conditions in the status
func (wab *WebApplicationBinding) SetConditions(conditions []metav1.Condition) {
	wab.Status.Conditions = conditions
}

func init() {
	SchemeBuilder.Register(&WebApplicationBinding{}, &WebApplicationBindingList{})
}
