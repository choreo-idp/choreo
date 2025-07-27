// Copyright 2025 The OpenChoreo Authors
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Repository defines the source repository configuration
type Repository struct {
	// URL is the repository URL
	URL string `json:"url"`

	// Revision specifies the revision to build from
	Revision Revision `json:"revision"`

	// AppPath is the path to the application within the repository
	AppPath string `json:"appPath"`
}

// Revision defines the revision specification
type Revision struct {
	// Branch specifies the branch to build from
	Branch string `json:"branch,omitempty"`

	// Commit specifies the commit hash to build from
	Commit string `json:"commit,omitempty"`
}

// TemplateRef defines the build template reference
type TemplateRef struct {
	// Engine specifies the build engine
	Engine string `json:"engine,omitempty"`

	// Name is the template name
	Name string `json:"name"`

	// Parameters contains the template parameters
	Parameters []Parameter `json:"parameters,omitempty"`
}

// Parameter defines a template parameter
type Parameter struct {
	// Name is the parameter name
	Name string `json:"name"`

	// Value is the parameter value
	Value string `json:"value"`
}

type BuildOwner struct {
	// +kubebuilder:validation:MinLength=1
	ProjectName string `json:"projectName"`
	// +kubebuilder:validation:MinLength=1
	ComponentName string `json:"componentName"`
}

// BuildSpec defines the desired state of Build.
type BuildSpec struct {
	Owner BuildOwner `json:"owner"`

	// Repository contains the source repository configuration
	Repository Repository `json:"repository"`

	// TemplateRef contains the build template reference and parameters
	TemplateRef TemplateRef `json:"templateRef"`
}

// BuildStatus defines the observed state of Build.
type BuildStatus struct {
	// Conditions represent the latest available observations of the build's current state
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// ImageStatus contains information about the built image
	ImageStatus Image `json:"imageStatus,omitempty"`
}

type Image struct {
	Image string `json:"image"`
}

// GetConditions returns the conditions slice
func (b *Build) GetConditions() []metav1.Condition {
	return b.Status.Conditions
}

// SetConditions sets the conditions slice
func (b *Build) SetConditions(conditions []metav1.Condition) {
	b.Status.Conditions = conditions
}

type DockerConfiguration struct {
	// Context specifies the build context path
	Context string `json:"context"`
	// DockerfilePath specifies the path to the Dockerfile
	DockerfilePath string `json:"dockerfilePath"`
}

type BuildpackConfiguration struct {
	Name    BuildpackName `json:"name"`
	Version string        `json:"version,omitempty"`
}

// BuildConfiguration specifies the build configuration details
type BuildConfiguration struct {
	// Docker specifies the Docker-specific build configuration
	Docker *DockerConfiguration `json:"docker,omitempty"`
	// Buildpack specifies the buildpack to use
	Buildpack *BuildpackConfiguration `json:"buildpack,omitempty"`
}

type BuildpackName string

const (
	BuildpackReact     BuildpackName = "React"
	BuildpackGo        BuildpackName = "Go"
	BuildpackBallerina BuildpackName = "Ballerina"
	BuildpackNodeJS    BuildpackName = "Node.js"
	BuildpackPython    BuildpackName = "Python"
	BuildpackRuby      BuildpackName = "Ruby"
	BuildpackPHP       BuildpackName = "PHP"
	// BuildpackJava      BuildpackName = "Java"
)

// SupportedVersions maps each buildpack to its supported versions.
// Refer (builder:google-22): https://cloud.google.com/docs/buildpacks/builders
var SupportedVersions = map[BuildpackName][]string{
	BuildpackReact:     {"18.20.6", "19.9.0", "20.18.3", "21.7.3", "22.14.0", "23.7.0"},
	BuildpackGo:        {"1.x"},
	BuildpackBallerina: {"2201.7.5", "2201.8.9", "2201.9.6", "2201.10.4", "2201.11.0"},
	BuildpackNodeJS:    {"12.x.x", "14.x.x", "16.x.x", "18.x.x", "20.x.x", "22.x.x"},
	BuildpackPython:    {"3.10.x", "3.11.x", "3.12.x"},
	BuildpackRuby:      {"3.1.x", "3.2.x", "3.3.x"},
	BuildpackPHP:       {"8.1.x", "8.2.x", "8.3.x"},
	// BuildpackJava:      {"8", "11", "17", "18", "21"},
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Build is the Schema for the builds API.
type Build struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BuildSpec   `json:"spec,omitempty"`
	Status BuildStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BuildList contains a list of Build.
type BuildList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Build `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Build{}, &BuildList{})
}
