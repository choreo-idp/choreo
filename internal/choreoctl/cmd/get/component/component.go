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

package component

import (
	"fmt"
	"os"
	"text/tabwriter"

	corev1 "github.com/choreo-idp/choreo/api/v1"
	"github.com/choreo-idp/choreo/internal/choreoctl/util"
	"github.com/choreo-idp/choreo/pkg/cli/common/constants"
	"github.com/choreo-idp/choreo/pkg/cli/types/api"
)

type ListCompImpl struct {
	config constants.CRDConfig
}

func NewListCompImpl(config constants.CRDConfig) *ListCompImpl {
	return &ListCompImpl{
		config: config,
	}
}

func (i *ListCompImpl) ListComponent(params api.ListComponentParams) error {
	if params.Interactive {
		return listComponentInteractive(i.config)
	}

	if err := util.ValidateParams(util.CmdGet, util.ResourceComponent, params); err != nil {
		return err
	}

	return listComponents(params, i.config)
}

func listComponents(params api.ListComponentParams, config constants.CRDConfig) error {
	var components []corev1.Component

	if params.Name != "" {
		component, err := util.GetComponent(params.Organization, params.Project, params.Name)
		if err != nil {
			return err
		}
		components = []corev1.Component{*component}
	} else {
		componentList, err := util.GetAllComponents(params.Organization, params.Project)
		if err != nil {
			return err
		}
		components = componentList.Items
	}

	if len(components) == 0 {
		fmt.Printf("No components found for organization: %s, project: %s\n", params.Organization, params.Project)
		return nil
	}

	if params.OutputFormat == constants.OutputFormatYAML {
		return printComponentYAML(components, params.Organization, config)
	}
	return printComponentTable(components, params.Organization, params.Project)
}
func printComponentYAML(components []corev1.Component, orgName string, config constants.CRDConfig) error {
	for _, component := range components {
		yamlStr, err := util.GetK8sObjectYAMLFromCRD(
			config.Group,
			string(config.Version),
			config.Kind,
			component.Name,
			orgName,
		)
		if err != nil {
			return err
		}
		fmt.Printf("---\n%s\n", yamlStr)
	}
	return nil
}

func printComponentTable(components []corev1.Component, orgName, projectName string) error {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	fmt.Fprintln(w, "NAME\tSTATUS\tAGE\tPROJECT\tORGANIZATION")

	for _, component := range components {
		age := util.FormatAge(component.CreationTimestamp.Time)
		status := util.GetStatus(component.Status.Conditions, "Created")
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n",
			component.Name,
			status,
			age,
			projectName,
			orgName)
	}

	return w.Flush()
}
