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

package deploymenttrack

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/choreo-idp/choreo/internal/controller"
)

const (
	// ConditionAvailable represents whether the deploymentTrack is created
	ConditionAvailable controller.ConditionType = "Available"
	// ConditionReady represents whether the deploymentTrack is ready
	ConditionReady controller.ConditionType = "Ready"
)

const (
	// ReasonDeploymentTrackAvailable is the reason used when a deploymentTrack is available
	ReasonDeploymentTrackAvailable controller.ConditionReason = "DeploymentTrackAvailable"

	// ReasonDeploymentTrackFinalizing is the reason used when a component is being finalized
	ReasonDeploymentTrackFinalizing controller.ConditionReason = "DeploymentTrackFinalizing"
)

// NewDeploymentTrackAvailableCondition creates a condition to indicate the deploymentTrack is available
func NewDeploymentTrackAvailableCondition(generation int64) metav1.Condition {
	return controller.NewCondition(
		ConditionAvailable,
		metav1.ConditionTrue,
		ReasonDeploymentTrackAvailable,
		"DeploymentTrack is available",
		generation,
	)
}

// NewDeploymentTrackFinalizingCondition creates a condition to indicate the component is being finalized
func NewDeploymentTrackFinalizingCondition(generation int64) metav1.Condition {
	return controller.NewCondition(
		ConditionReady,
		metav1.ConditionFalse,
		ReasonDeploymentTrackFinalizing,
		"DeploymentTrack is being finalized",
		generation,
	)
}
