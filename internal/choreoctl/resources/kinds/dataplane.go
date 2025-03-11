/*
 * Copyright (c) 2025, WSO2 Inc. (http://www.wso2.org) All Rights Reserved.
 *
 * WSO2 Inc. licenses this file to you under the Apache License,
 * Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package kinds

import (
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	choreov1 "github.com/choreo-idp/choreo/api/v1"
	"github.com/choreo-idp/choreo/internal/choreoctl/resources"
	"github.com/choreo-idp/choreo/pkg/cli/common/constants"
	"github.com/choreo-idp/choreo/pkg/cli/types/api"
)

// DataPlaneResource provides operations for DataPlane CRs.
type DataPlaneResource struct {
	*resources.BaseResource[*choreov1.DataPlane, *choreov1.DataPlaneList]
}

// NewDataPlaneResource constructs a DataPlaneResource with CRDConfig and optionally sets organization.
func NewDataPlaneResource(cfg constants.CRDConfig, org string) (*DataPlaneResource, error) {
	cli, err := resources.GetClient()
	if err != nil {
		return nil, fmt.Errorf(ErrCreateKubeClient, err)
	}

	options := []resources.ResourceOption[*choreov1.DataPlane, *choreov1.DataPlaneList]{
		resources.WithClient[*choreov1.DataPlane, *choreov1.DataPlaneList](cli),
		resources.WithConfig[*choreov1.DataPlane, *choreov1.DataPlaneList](cfg),
	}

	// Add organization namespace if provided
	if org != "" {
		options = append(options, resources.WithNamespace[*choreov1.DataPlane, *choreov1.DataPlaneList](org))
	}

	return &DataPlaneResource{
		BaseResource: resources.NewBaseResource(options...),
	}, nil
}

// WithNamespace sets the namespace for the dataplane resource (usually the organization name)
func (d *DataPlaneResource) WithNamespace(namespace string) {
	d.BaseResource.WithNamespace(namespace)
}

// GetStatus returns the status of a DataPlane with detailed information.
func (d *DataPlaneResource) GetStatus(dataPlane *choreov1.DataPlane) string {
	priorityConditions := []string{
		ConditionTypeReady,
		ConditionTypeAvailable,
		ConditionTypeConfigured,
	}

	return resources.GetResourceStatus(
		dataPlane.Status.Conditions,
		priorityConditions,
		StatusPending,
		StatusReady,
		StatusNotReady,
	)
}

// GetAge returns the age of a DataPlane.
func (d *DataPlaneResource) GetAge(dataPlane *choreov1.DataPlane) string {
	return resources.FormatAge(dataPlane.GetCreationTimestamp().Time)
}

// PrintTableItems formats dataplanes into a table
func (d *DataPlaneResource) PrintTableItems(dataPlanes []resources.ResourceWrapper[*choreov1.DataPlane]) error {
	if len(dataPlanes) == 0 {
		namespaceName := d.GetNamespace()

		message := "No data planes found"

		if namespaceName != "" {
			message += " in organization " + namespaceName
		}

		fmt.Println(message)
		return nil
	}

	rows := make([][]string, 0, len(dataPlanes))

	for _, wrapper := range dataPlanes {
		dataPlane := wrapper.Resource
		rows = append(rows, []string{
			wrapper.LogicalName,
			dataPlane.Spec.KubernetesCluster.Name,
			d.GetStatus(dataPlane),
			d.GetAge(dataPlane),
			dataPlane.GetLabels()[constants.LabelOrganization],
		})
	}
	return resources.PrintTable(HeadersDataPlane, rows)
}

// Print overrides the base Print method to ensure our custom PrintTableItems is called
func (d *DataPlaneResource) Print(format resources.OutputFormat, filter *resources.ResourceFilter) error {
	dataPlanes, err := d.List()
	if err != nil {
		return err
	}

	if filter != nil && filter.Name != "" {
		filtered, err := resources.FilterByName(dataPlanes, filter.Name)
		if err != nil {
			return err
		}
		dataPlanes = filtered
	}

	switch format {
	case resources.OutputFormatTable:
		return d.PrintTableItems(dataPlanes)
	case resources.OutputFormatYAML:
		return d.BaseResource.PrintItems(dataPlanes, format)
	default:
		return fmt.Errorf(ErrFormatUnsupported, format)
	}
}

// CreateDataPlane creates a new DataPlane CR.
func (d *DataPlaneResource) CreateDataPlane(params api.CreateDataPlaneParams) error {
	k8sName := resources.GenerateResourceName(params.Organization, params.Name)

	// Create the DataPlane resource
	dataPlane := &choreov1.DataPlane{
		ObjectMeta: metav1.ObjectMeta{
			Name:      k8sName,
			Namespace: params.Organization,
			Annotations: map[string]string{
				constants.AnnotationDisplayName: resources.DefaultIfEmpty(params.DisplayName, params.Name),
				constants.AnnotationDescription: params.Description,
			},
			Labels: map[string]string{
				constants.LabelName:         params.Name,
				constants.LabelOrganization: params.Organization,
			},
		},
		Spec: choreov1.DataPlaneSpec{
			KubernetesCluster: choreov1.KubernetesClusterSpec{
				Name:                params.KubernetesClusterName,
				ConnectionConfigRef: params.ConnectionConfigRef,
				FeatureFlags: choreov1.FeatureFlagsSpec{
					Cilium:      params.EnableCilium,
					ScaleToZero: params.EnableScaleToZero,
					GatewayType: params.GatewayType,
				},
			},
			Gateway: choreov1.GatewaySpec{
				PublicVirtualHost:       params.PublicVirtualHost,
				OrganizationVirtualHost: params.OrganizationVirtualHost,
			},
		},
	}

	// Create the dataplane using the base create method
	if err := d.Create(dataPlane); err != nil {
		return fmt.Errorf(ErrCreateDataPlane, err)
	}

	fmt.Printf(FmtDataPlaneCreateSuccess, params.Name, params.Organization)
	return nil
}

// GetDataPlanesForOrganization returns dataplanes filtered by organization
func (d *DataPlaneResource) GetDataPlanesForOrganization(orgName string) ([]resources.ResourceWrapper[*choreov1.DataPlane], error) {
	allDataPlanes, err := d.List()
	if err != nil {
		return nil, err
	}

	var dataPlanes []resources.ResourceWrapper[*choreov1.DataPlane]
	for _, wrapper := range allDataPlanes {
		if wrapper.Resource.GetLabels()[constants.LabelOrganization] == orgName {
			dataPlanes = append(dataPlanes, wrapper)
		}
	}

	return dataPlanes, nil
}
