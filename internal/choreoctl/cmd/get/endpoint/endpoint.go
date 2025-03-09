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

package endpoint

import (
	"fmt"
	"os"
	"text/tabwriter"

	corev1 "github.com/choreo-idp/choreo/api/v1"
	"github.com/choreo-idp/choreo/internal/choreoctl/util"
	"github.com/choreo-idp/choreo/pkg/cli/common/constants"
	"github.com/choreo-idp/choreo/pkg/cli/types/api"
)

type ListEndpointImpl struct {
	config constants.CRDConfig
}

func NewListEndpointImpl(config constants.CRDConfig) *ListEndpointImpl {
	return &ListEndpointImpl{
		config: config,
	}
}

func (i *ListEndpointImpl) ListEndpoint(params api.ListEndpointParams) error {
	if params.Interactive {
		return listEndpointInteractive(i.config)
	}

	if err := util.ValidateParams(util.CmdGet, util.ResourceEndpoint, params); err != nil {
		return err
	}

	return listEndpoints(params, i.config)
}

func listEndpoints(params api.ListEndpointParams, config constants.CRDConfig) error {
	var endpoints []corev1.Endpoint

	if params.Name != "" {
		// Get specific endpoint
		endpoint, err := util.GetEndpoint(
			params.Organization,
			params.Project,
			params.Component,
			params.Environment,
			params.Name,
		)
		if err != nil {
			return err
		}
		endpoints = []corev1.Endpoint{*endpoint}
	} else {
		// List all endpoints
		endpointList, err := util.GetAllEndpoints(
			params.Organization,
			params.Project,
			params.Component,
			params.Environment,
		)
		if err != nil {
			return err
		}
		endpoints = endpointList.Items
	}

	if len(endpoints) == 0 {
		fmt.Printf("No endpoints found for organization: %s, project: %s, component: %s\n",
			params.Organization, params.Project, params.Component)
		return nil
	}

	if params.OutputFormat == constants.OutputFormatYAML {
		return printEndpointYAML(endpoints, params.Organization, config)
	}
	return printEndpointTable(endpoints, params.Organization)
}

func printEndpointYAML(endpoints []corev1.Endpoint, orgName string, config constants.CRDConfig) error {
	for _, endpoint := range endpoints {
		yamlStr, err := util.GetK8sObjectYAMLFromCRD(
			config.Group,
			string(config.Version),
			config.Kind,
			endpoint.Name,
			orgName,
		)
		if err != nil {
			return err
		}
		fmt.Printf("---\n%s\n", yamlStr)
	}
	return nil
}

func printEndpointTable(endpoints []corev1.Endpoint, orgName string) error {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	fmt.Fprintln(w, "NAME\tTYPE\tPORT\tBASE PATH\tVISIBILITY\tAGE")

	for _, endpoint := range endpoints {
		age := util.FormatAge(endpoint.CreationTimestamp.Time)
		visibility := ""
		if endpoint.Spec.NetworkVisibilities.Public.Enable {
			visibility = "Public"
		}
		if endpoint.Spec.NetworkVisibilities.Organization.Enable {
			if visibility != "" {
				visibility += ","
			}
			visibility += "Organization"
		}

		fmt.Fprintf(w, "%s\t%s\t%d\t%s\t%s\t%s\t%s\n",
			endpoint.Name,
			endpoint.Spec.Type,
			endpoint.Spec.Service.Port,
			endpoint.Spec.Service.BasePath,
			visibility,
			age, orgName)
	}

	return w.Flush()
}
