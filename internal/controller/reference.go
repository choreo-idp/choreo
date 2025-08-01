// Copyright 2025 The OpenChoreo Authors
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	"context"
	"fmt"

	"sigs.k8s.io/controller-runtime/pkg/client"

	openchoreov1alpha1 "github.com/openchoreo/openchoreo/api/v1alpha1"
	"github.com/openchoreo/openchoreo/internal/labels"
)

func GetDataplaneOfEnv(ctx context.Context, c client.Client, env *openchoreov1alpha1.Environment) (*openchoreov1alpha1.DataPlane, error) {
	dataplaneList := &openchoreov1alpha1.DataPlaneList{}
	listOpts := []client.ListOption{
		client.InNamespace(env.GetNamespace()),
		client.MatchingLabels{
			labels.LabelKeyOrganizationName: GetOrganizationName(env),
			labels.LabelKeyName:             env.Spec.DataPlaneRef,
		},
	}

	if err := c.List(ctx, dataplaneList, listOpts...); err != nil {
		return nil, fmt.Errorf("failed to list dataplanes: %w", err)
	}

	if len(dataplaneList.Items) > 0 {
		return &dataplaneList.Items[0], nil
	}

	return nil, fmt.Errorf("failed to find dataplane for environment: %s", env.Spec.DataPlaneRef)
}
