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
	"github.com/wso2-enterprise/choreo-cp-declarative-api/internal/controller"
	"github.com/wso2-enterprise/choreo-cp-declarative-api/internal/dataplane"
	dpkubernetes "github.com/wso2-enterprise/choreo-cp-declarative-api/internal/dataplane/kubernetes"
)

func makeLabels(epCtx *dataplane.EndpointContext) map[string]string {
	return map[string]string{
		dpkubernetes.LabelKeyOrganizationName:    controller.GetOrganizationName(epCtx.Project),
		dpkubernetes.LabelKeyProjectName:         controller.GetName(epCtx.Project),
		dpkubernetes.LabelKeyProjectID:           string(epCtx.Project.UID),
		dpkubernetes.LabelKeyComponentName:       controller.GetName(epCtx.Component),
		dpkubernetes.LabelKeyComponentID:         string(epCtx.Component.UID),
		dpkubernetes.LabelKeyDeploymentTrackName: controller.GetName(epCtx.DeploymentTrack),
		dpkubernetes.LabelKeyDeploymentTrackID:   string(epCtx.DeploymentTrack.UID),
		dpkubernetes.LabelKeyEnvironmentName:     controller.GetName(epCtx.Environment),
		dpkubernetes.LabelKeyEnvironmentID:       string(epCtx.Environment.UID),
		dpkubernetes.LabelKeyDeploymentName:      controller.GetName(epCtx.Deployment),
		dpkubernetes.LabelKeyDeploymentID:        string(epCtx.Deployment.UID),
		dpkubernetes.LabelKeyManagedBy:           dpkubernetes.LabelValueManagedBy,
		dpkubernetes.LabelKeyBelongTo:            dpkubernetes.LabelValueBelongTo,
	}
}

func makeWorkloadLabels(epCtx *dataplane.EndpointContext) map[string]string {
	labels := makeLabels(epCtx)
	labels[dpkubernetes.LabelKeyComponentType] = string(epCtx.Component.Spec.Type)
	return labels
}

func extractManagedLabels(labels map[string]string) map[string]string {
	return map[string]string{
		dpkubernetes.LabelKeyOrganizationName:    labels[dpkubernetes.LabelKeyOrganizationName],
		dpkubernetes.LabelKeyProjectName:         labels[dpkubernetes.LabelKeyProjectName],
		dpkubernetes.LabelKeyProjectID:           labels[dpkubernetes.LabelKeyProjectID],
		dpkubernetes.LabelKeyComponentName:       labels[dpkubernetes.LabelKeyComponentName],
		dpkubernetes.LabelKeyComponentID:         labels[dpkubernetes.LabelKeyComponentID],
		dpkubernetes.LabelKeyDeploymentTrackName: labels[dpkubernetes.LabelKeyDeploymentTrackName],
		dpkubernetes.LabelKeyDeploymentTrackID:   labels[dpkubernetes.LabelKeyDeploymentTrackID],
		dpkubernetes.LabelKeyEnvironmentName:     labels[dpkubernetes.LabelKeyEnvironmentName],
		dpkubernetes.LabelKeyEnvironmentID:       labels[dpkubernetes.LabelKeyEnvironmentID],
		dpkubernetes.LabelKeyDeploymentName:      labels[dpkubernetes.LabelKeyDeploymentName],
		dpkubernetes.LabelKeyDeploymentID:        labels[dpkubernetes.LabelKeyDeploymentID],
		dpkubernetes.LabelKeyManagedBy:           labels[dpkubernetes.LabelKeyManagedBy],
		dpkubernetes.LabelKeyBelongTo:            labels[dpkubernetes.LabelKeyBelongTo],
		dpkubernetes.LabelKeyComponentType:       labels[dpkubernetes.LabelKeyComponentType],
	}
}
