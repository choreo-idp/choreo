// Copyright 2025 The OpenChoreo Authors
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// DockerConfiguration specifies the Docker build configuration

// BuildTemplateSpec defines the build template configuration
type BuildTemplateSpec struct {
	// Branch specifies the Git branch to use
	Branch string `json:"branch"`
	// Path specifies the repository path to use
	Path string `json:"path"`
	// BuildConfiguration specifies the build settings
	BuildConfiguration *BuildConfiguration `json:"buildConfiguration,omitempty"`
}

// DeploymentTrackSpec defines the desired state of DeploymentTrack.
type DeploymentTrackSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// AutoDeploy defines whether deployment should be triggered automatically
	AutoDeploy bool `json:"autoDeploy,omitempty"`

	// BuildTemplateSpec defines the build template configuration
	BuildTemplateSpec *BuildTemplateSpec `json:"buildTemplateSpec,omitempty"`
}

// DeploymentTrackStatus defines the observed state of DeploymentTrack.
type DeploymentTrackStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	ObservedGeneration int64              `json:"observedGeneration,omitempty"`
	Conditions         []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Namespaced,shortName=deptrack;deptracks

// DeploymentTrack is the Schema for the deploymenttracks API.
type DeploymentTrack struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DeploymentTrackSpec   `json:"spec,omitempty"`
	Status DeploymentTrackStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DeploymentTrackList contains a list of DeploymentTrack.
type DeploymentTrackList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DeploymentTrack `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DeploymentTrack{}, &DeploymentTrackList{})
}

func (p *DeploymentTrack) GetConditions() []metav1.Condition {
	return p.Status.Conditions
}

func (p *DeploymentTrack) SetConditions(conditions []metav1.Condition) {
	p.Status.Conditions = conditions
}
