// Copyright 2025 The OpenChoreo Authors
// SPDX-License-Identifier: Apache-2.0

package services

import "errors"

// Common service errors
var (
	ErrProjectAlreadyExists       = errors.New("project already exists")
	ErrProjectNotFound            = errors.New("project not found")
	ErrComponentAlreadyExists     = errors.New("component already exists")
	ErrComponentNotFound          = errors.New("component not found")
	ErrOrganizationNotFound       = errors.New("organization not found")
	ErrEnvironmentNotFound        = errors.New("environment not found")
	ErrEnvironmentAlreadyExists   = errors.New("environment already exists")
	ErrDataPlaneNotFound          = errors.New("dataplane not found")
	ErrDataPlaneAlreadyExists     = errors.New("dataplane already exists")
	ErrBindingNotFound            = errors.New("binding not found")
	ErrDeploymentPipelineNotFound = errors.New("deployment pipeline not found")
	ErrInvalidPromotionPath       = errors.New("invalid promotion path")
)

// Error codes for API responses
const (
	CodeProjectExists              = "PROJECT_EXISTS"
	CodeProjectNotFound            = "PROJECT_NOT_FOUND"
	CodeComponentExists            = "COMPONENT_EXISTS"
	CodeComponentNotFound          = "COMPONENT_NOT_FOUND"
	CodeOrganizationNotFound       = "ORGANIZATION_NOT_FOUND"
	CodeEnvironmentNotFound        = "ENVIRONMENT_NOT_FOUND"
	CodeEnvironmentExists          = "ENVIRONMENT_EXISTS"
	CodeDataPlaneNotFound          = "DATAPLANE_NOT_FOUND"
	CodeDataPlaneExists            = "DATAPLANE_EXISTS"
	CodeBindingNotFound            = "BINDING_NOT_FOUND"
	CodeDeploymentPipelineNotFound = "DEPLOYMENT_PIPELINE_NOT_FOUND"
	CodeInvalidPromotionPath       = "INVALID_PROMOTION_PATH"
	CodeInvalidInput               = "INVALID_INPUT"
	CodeInternalError              = "INTERNAL_ERROR"
)
