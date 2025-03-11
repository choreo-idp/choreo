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

package choreoctl

import (
	"github.com/choreo-idp/choreo/internal/choreoctl/cmd/apply"
	"github.com/choreo-idp/choreo/internal/choreoctl/cmd/config"
	"github.com/choreo-idp/choreo/internal/choreoctl/cmd/create/build"
	"github.com/choreo-idp/choreo/internal/choreoctl/cmd/create/component"
	"github.com/choreo-idp/choreo/internal/choreoctl/cmd/create/dataplane"
	"github.com/choreo-idp/choreo/internal/choreoctl/cmd/create/deployableartifact"
	"github.com/choreo-idp/choreo/internal/choreoctl/cmd/create/deployment"
	"github.com/choreo-idp/choreo/internal/choreoctl/cmd/create/deploymenttrack"
	"github.com/choreo-idp/choreo/internal/choreoctl/cmd/create/environment"
	"github.com/choreo-idp/choreo/internal/choreoctl/cmd/create/organization"
	"github.com/choreo-idp/choreo/internal/choreoctl/cmd/create/project"
	getbuild "github.com/choreo-idp/choreo/internal/choreoctl/cmd/get/build"
	getcomponent "github.com/choreo-idp/choreo/internal/choreoctl/cmd/get/component"
	getdataplane "github.com/choreo-idp/choreo/internal/choreoctl/cmd/get/dataplane"
	getdeployartifcat "github.com/choreo-idp/choreo/internal/choreoctl/cmd/get/deployableartifact"
	getdeploy "github.com/choreo-idp/choreo/internal/choreoctl/cmd/get/deployment"
	getdeploymenttrack "github.com/choreo-idp/choreo/internal/choreoctl/cmd/get/deploymenttrack"
	getendpoint "github.com/choreo-idp/choreo/internal/choreoctl/cmd/get/endpoint"
	getenv "github.com/choreo-idp/choreo/internal/choreoctl/cmd/get/environment"
	getorganization "github.com/choreo-idp/choreo/internal/choreoctl/cmd/get/organization"
	getproject "github.com/choreo-idp/choreo/internal/choreoctl/cmd/get/project"
	"github.com/choreo-idp/choreo/internal/choreoctl/cmd/login"
	"github.com/choreo-idp/choreo/internal/choreoctl/cmd/logout"
	"github.com/choreo-idp/choreo/internal/choreoctl/cmd/logs"
	"github.com/choreo-idp/choreo/pkg/cli/common/constants"
	"github.com/choreo-idp/choreo/pkg/cli/types/api"
)

type CommandImplementation struct{}

var _ api.CommandImplementationInterface = &CommandImplementation{}

func NewCommandImplementation() *CommandImplementation {
	return &CommandImplementation{}
}

// Get Operations

func (c *CommandImplementation) GetOrganization(params api.GetParams) error {
	orgImpl := getorganization.NewGetOrgImpl(constants.OrganizationV1Config)
	return orgImpl.GetOrganization(params)
}

func (c *CommandImplementation) GetProject(params api.GetProjectParams) error {
	projImpl := getproject.NewGetProjImpl(constants.ProjectV1Config)
	return projImpl.GetProject(params)
}

func (c *CommandImplementation) GetComponent(params api.GetComponentParams) error {
	compImpl := getcomponent.NewGetCompImpl(constants.ComponentV1Config)
	return compImpl.GetComponent(params)
}

func (c *CommandImplementation) GetBuild(params api.GetBuildParams) error {
	buildImpl := getbuild.NewGetBuildImpl(constants.BuildV1Config)
	return buildImpl.GetBuild(params)
}

func (c *CommandImplementation) GetDeployableArtifact(params api.GetDeployableArtifactParams) error {
	deployableArtifactImpl := getdeployartifcat.NewGetDeployableArtifactImpl(constants.DeployableArtifactV1Config)
	return deployableArtifactImpl.GetDeployableArtifact(params)
}

func (c *CommandImplementation) GetDeployment(params api.GetDeploymentParams) error {
	deploymentImpl := getdeploy.NewGetDeploymentImpl(constants.DeploymentV1Config)
	return deploymentImpl.GetDeployment(params)
}

func (c *CommandImplementation) GetEnvironment(params api.GetEnvironmentParams) error {
	envImpl := getenv.NewGetEnvironmentImpl(constants.EnvironmentV1Config)
	return envImpl.GetEnvironment(params)
}

func (c *CommandImplementation) GetDataPlane(params api.GetDataPlaneParams) error {
	dpImpl := getdataplane.NewGetDataPlaneImpl(constants.DataPlaneV1Config)
	return dpImpl.GetDataPlane(params)
}

func (c *CommandImplementation) GetDeploymentTrack(params api.GetDeploymentTrackParams) error {
	trackImpl := getdeploymenttrack.NewGetDeploymentTrackImpl(constants.DeploymentTrackV1Config)
	return trackImpl.GetDeploymentTrack(params)
}

func (c *CommandImplementation) GetEndpoint(params api.GetEndpointParams) error {
	endpointImpl := getendpoint.NewGetEndpointImpl(constants.EndpointV1Config)
	return endpointImpl.GetEndpoint(params)
}

// Create Operations

func (c *CommandImplementation) CreateOrganization(params api.CreateOrganizationParams) error {
	orgImpl := organization.NewCreateOrgImpl(constants.OrganizationV1Config)
	return orgImpl.CreateOrganization(params)
}

func (c *CommandImplementation) CreateProject(params api.CreateProjectParams) error {
	projImpl := project.NewCreateProjImpl(constants.ProjectV1Config)
	return projImpl.CreateProject(params)
}

func (c *CommandImplementation) CreateComponent(params api.CreateComponentParams) error {
	compImpl := component.NewCreateCompImpl(constants.ComponentV1Config)
	return compImpl.CreateComponent(params)
}

func (c *CommandImplementation) CreateBuild(params api.CreateBuildParams) error {
	buildImpl := build.NewCreateBuildImpl(constants.ComponentV1Config)
	return buildImpl.CreateBuild(params)
}

func (c *CommandImplementation) CreateDeployment(params api.CreateDeploymentParams) error {
	deployImpl := deployment.NewCreateDeploymentImpl(constants.DeploymentV1Config)
	return deployImpl.CreateDeployment(params)
}

func (c *CommandImplementation) CreateEnvironment(params api.CreateEnvironmentParams) error {
	envImpl := environment.NewCreateEnvironmentImpl(constants.EnvironmentV1Config)
	return envImpl.CreateEnvironment(params)
}

func (c *CommandImplementation) CreateDataPlane(params api.CreateDataPlaneParams) error {
	dpImpl := dataplane.NewCreateDataPlaneImpl(constants.DataPlaneV1Config)
	return dpImpl.CreateDataPlane(params)
}

func (c *CommandImplementation) CreateDeploymentTrack(params api.CreateDeploymentTrackParams) error {
	trackImpl := deploymenttrack.NewCreateDeploymentTrackImpl(constants.DeploymentTrackV1Config)
	return trackImpl.CreateDeploymentTrack(params)
}

func (c *CommandImplementation) CreateDeployableArtifact(params api.CreateDeployableArtifactParams) error {
	daImpl := deployableartifact.NewCreateDeployableArtifactImpl(constants.DeployableArtifactV1Config)
	return daImpl.CreateDeployableArtifact(params)
}

// Authentication Operations

func (c *CommandImplementation) Login(params api.LoginParams) error {
	loginImpl := login.NewAuthImpl()
	return loginImpl.Login(params)
}

func (c *CommandImplementation) IsLoggedIn() bool {
	loginImpl := login.NewAuthImpl()
	return loginImpl.IsLoggedIn()
}

func (c *CommandImplementation) GetLoginPrompt() string {
	loginImpl := login.NewAuthImpl()
	return loginImpl.GetLoginPrompt()
}

func (c *CommandImplementation) Logout() error {
	logoutImpl := logout.NewLogoutImpl()
	return logoutImpl.Logout()
}

// Configuration Operations

func (c *CommandImplementation) Apply(params api.ApplyParams) error {
	applyImpl := apply.NewApplyImpl()
	return applyImpl.Apply(params)
}

// Logs Operations

func (c *CommandImplementation) GetLogs(params api.LogParams) error {
	logsImpl := logs.NewLogsImpl()
	return logsImpl.GetLogs(params)
}

// Config Context Operations

func (c *CommandImplementation) GetContexts() error {
	configContextImpl := config.NewConfigContextImpl()
	return configContextImpl.GetContexts()
}

func (c *CommandImplementation) GetCurrentContext() error {
	configContextImpl := config.NewConfigContextImpl()
	return configContextImpl.GetCurrentContext()
}

func (c *CommandImplementation) SetContext(params api.SetContextParams) error {
	configContextImpl := config.NewConfigContextImpl()
	return configContextImpl.SetContext(params)
}

func (c *CommandImplementation) UseContext(params api.UseContextParams) error {
	configContextImpl := config.NewConfigContextImpl()
	return configContextImpl.UseContext(params)
}
