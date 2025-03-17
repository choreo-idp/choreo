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

package kubernetes

import (
	"fmt"
	"path"

	gatewayv1 "sigs.k8s.io/gateway-api/apis/v1"

	choreov1 "github.com/choreo-idp/choreo/api/v1"
	"github.com/choreo-idp/choreo/internal/controller/endpoint/integrations/kubernetes/visibility"
	"github.com/choreo-idp/choreo/internal/dataplane"
)

// makeHostname generates the hostname for an endpoint based on gateway type and component type
func makeHostname(epCtx *dataplane.EndpointContext, gwType visibility.GatewayType) gatewayv1.Hostname {
	var domain string
	switch gwType {
	case visibility.GatewayInternal:
		domain = epCtx.DataPlane.Spec.Gateway.OrganizationVirtualHost
	default:
		domain = epCtx.DataPlane.Spec.Gateway.PublicVirtualHost
	}
	if epCtx.Component.Spec.Type == choreov1.ComponentTypeWebApplication {
		return gatewayv1.Hostname(fmt.Sprintf("%s-%s.%s", epCtx.Component.Name, epCtx.Environment.Name, domain))
	}
	return gatewayv1.Hostname(fmt.Sprintf("%s.%s", epCtx.Environment.Spec.Gateway.DNSPrefix, domain))
}

// makePathPrefix returns the URL path prefix based on component type
func makePathPrefix(epCtx *dataplane.EndpointContext) string {
	if epCtx.Component.Spec.Type == choreov1.ComponentTypeWebApplication {
		return "/"
	}
	return path.Clean(path.Join("/", epCtx.Project.Name, epCtx.Component.Name))
}

// MakeAddress constructs the full HTTPS URL for an endpoint
func MakeAddress(epCtx *dataplane.EndpointContext, gwType visibility.GatewayType) string {
	host := makeHostname(epCtx, gwType)
	pathPrefix := makePathPrefix(epCtx)

	return fmt.Sprintf("https://%s%s", host, pathPrefix)
}
