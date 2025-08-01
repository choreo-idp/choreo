// Copyright 2025 The OpenChoreo Authors
// SPDX-License-Identifier: Apache-2.0

package kinds

import (
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	openchoreov1alpha1 "github.com/openchoreo/openchoreo/api/v1alpha1"
	"github.com/openchoreo/openchoreo/internal/choreoctl/resources"
	"github.com/openchoreo/openchoreo/pkg/cli/common/constants"
	"github.com/openchoreo/openchoreo/pkg/cli/types/api"
)

// DeploymentTrackResource provides operations for DeploymentTrack CRs.
type DeploymentTrackResource struct {
	*resources.BaseResource[*openchoreov1alpha1.DeploymentTrack, *openchoreov1alpha1.DeploymentTrackList]
}

// NewDeploymentTrackResource constructs a DeploymentTrackResource with CRDConfig and optionally sets organization, project, and component.
func NewDeploymentTrackResource(cfg constants.CRDConfig, org string, project string, component string) (*DeploymentTrackResource, error) {
	cli, err := resources.GetClient()
	if err != nil {
		return nil, fmt.Errorf(ErrCreateKubeClient, err)
	}

	options := []resources.ResourceOption[*openchoreov1alpha1.DeploymentTrack, *openchoreov1alpha1.DeploymentTrackList]{
		resources.WithClient[*openchoreov1alpha1.DeploymentTrack, *openchoreov1alpha1.DeploymentTrackList](cli),
		resources.WithConfig[*openchoreov1alpha1.DeploymentTrack, *openchoreov1alpha1.DeploymentTrackList](cfg),
	}

	// Add organization namespace if provided
	if org != "" {
		options = append(options, resources.WithNamespace[*openchoreov1alpha1.DeploymentTrack, *openchoreov1alpha1.DeploymentTrackList](org))
	}

	// Create labels for filtering
	labels := map[string]string{}

	// Add project label if provided
	if project != "" {
		labels[constants.LabelProject] = project
	}

	// Add component label if provided
	if component != "" {
		labels[constants.LabelComponent] = component
	}

	// Add labels if any were set
	if len(labels) > 0 {
		options = append(options, resources.WithLabels[*openchoreov1alpha1.DeploymentTrack, *openchoreov1alpha1.DeploymentTrackList](labels))
	}

	return &DeploymentTrackResource{
		BaseResource: resources.NewBaseResource[*openchoreov1alpha1.DeploymentTrack, *openchoreov1alpha1.DeploymentTrackList](options...),
	}, nil
}

// WithNamespace sets the namespace for the deployment track resource (usually the organization name)
func (d *DeploymentTrackResource) WithNamespace(namespace string) {
	d.BaseResource.WithNamespace(namespace)
}

// GetStatus returns the status of a DeploymentTrack with detailed information.
func (d *DeploymentTrackResource) GetStatus(track *openchoreov1alpha1.DeploymentTrack) string {
	// DeploymentTrack typically uses Created or Ready conditions
	priorityConditions := []string{ConditionTypeCreated, ConditionTypeReady}

	return resources.GetResourceStatus(
		track.Status.Conditions,
		priorityConditions,
		StatusPending,
		StatusReady,
		StatusNotReady,
	)
}

// GetAge returns the age of a DeploymentTrack.
func (d *DeploymentTrackResource) GetAge(track *openchoreov1alpha1.DeploymentTrack) string {
	return resources.FormatAge(track.GetCreationTimestamp().Time)
}

// GetAutoDeploy returns whether a DeploymentTrack has auto-deployment enabled.
func (d *DeploymentTrackResource) GetAutoDeploy(track *openchoreov1alpha1.DeploymentTrack) string {
	autoDeploy := track.GetAnnotations()[AutoDeployAnnotation]
	if autoDeploy == "true" {
		return "Yes"
	}
	return "No"
}

// PrintTableItems formats deployment tracks into a table
func (d *DeploymentTrackResource) PrintTableItems(tracks []resources.ResourceWrapper[*openchoreov1alpha1.DeploymentTrack]) error {
	if len(tracks) == 0 {
		// Provide a more descriptive message
		namespaceName := d.GetNamespace()
		labels := d.GetLabels()

		message := "No deployment tracks found"

		if namespaceName != "" {
			message += " in organization " + namespaceName
		}

		if project, ok := labels[constants.LabelProject]; ok {
			message += ", project " + project
		}

		if component, ok := labels[constants.LabelComponent]; ok {
			message += ", component " + component
		}

		fmt.Println(message)
		return nil
	}

	rows := make([][]string, 0, len(tracks))

	for _, wrapper := range tracks {
		track := wrapper.Resource
		rows = append(rows, []string{
			wrapper.LogicalName,
			resources.FormatValueOrPlaceholder("v1.0"),
			resources.FormatBoolAsYesNo(track.GetAnnotations()[AutoDeployAnnotation] == "true"),
			resources.FormatAge(track.GetCreationTimestamp().Time),
			track.GetLabels()[constants.LabelComponent],
			track.GetLabels()[constants.LabelProject],
			track.GetLabels()[constants.LabelOrganization],
		})
	}
	return resources.PrintTable(HeadersDeploymentTrack, rows)
}

// Print overrides the base Print method to ensure our custom PrintTableItems is called
func (d *DeploymentTrackResource) Print(format resources.OutputFormat, filter *resources.ResourceFilter) error {
	// List resources
	deploymentTracks, err := d.List()
	if err != nil {
		return err
	}

	// Apply name filter if specified
	if filter != nil && filter.Name != "" {
		filtered, err := resources.FilterByName(deploymentTracks, filter.Name)
		if err != nil {
			return err
		}
		deploymentTracks = filtered
	}

	// Call the appropriate print method based on format
	switch format {
	case resources.OutputFormatTable:
		return d.PrintTableItems(deploymentTracks)
	case resources.OutputFormatYAML:
		return d.BaseResource.PrintItems(deploymentTracks, format)
	default:
		return fmt.Errorf(ErrFormatUnsupported, format)
	}
}

// CreateDeploymentTrack creates a new DeploymentTrack CR.
func (d *DeploymentTrackResource) CreateDeploymentTrack(params api.CreateDeploymentTrackParams) error {
	// Generate a K8s-compliant name for the deployment track
	k8sName := resources.GenerateResourceName(params.Organization, params.Project, params.Component, params.Name)

	// Create the DeploymentTrack resource
	deploymentTrack := &openchoreov1alpha1.DeploymentTrack{
		ObjectMeta: metav1.ObjectMeta{
			Name:      k8sName,
			Namespace: params.Organization,
			Annotations: map[string]string{
				constants.AnnotationDisplayName: resources.DefaultIfEmpty(params.DisplayName, params.Name),
				constants.AnnotationDescription: params.Description,
				AutoDeployAnnotation:            fmt.Sprintf("%t", params.AutoDeploy),
			},
			Labels: map[string]string{
				constants.LabelName:         params.Name,
				constants.LabelOrganization: params.Organization,
				constants.LabelProject:      params.Project,
				constants.LabelComponent:    params.Component,
			},
		},
		Spec: openchoreov1alpha1.DeploymentTrackSpec{
			BuildTemplateSpec: params.BuildTemplateSpec,
		},
	}

	// Create the deployment track using the base create method
	if err := d.Create(deploymentTrack); err != nil {
		return fmt.Errorf(ErrCreateDeploymentTrack, err)
	}

	fmt.Printf(FmtDeploymentTrackSuccess,
		params.Name, params.Component, params.Project, params.Organization)
	return nil
}

// GetDeploymentTracksForComponent returns deployment tracks filtered by component
func (d *DeploymentTrackResource) GetDeploymentTracksForComponent(componentName string) ([]resources.ResourceWrapper[*openchoreov1alpha1.DeploymentTrack], error) {
	// List all deployment tracks in the namespace
	allDeploymentTracks, err := d.List()
	if err != nil {
		return nil, err
	}

	// Filter by component
	var deploymentTracks []resources.ResourceWrapper[*openchoreov1alpha1.DeploymentTrack]
	for _, wrapper := range allDeploymentTracks {
		if wrapper.Resource.GetLabels()[constants.LabelComponent] == componentName {
			deploymentTracks = append(deploymentTracks, wrapper)
		}
	}

	return deploymentTracks, nil
}
