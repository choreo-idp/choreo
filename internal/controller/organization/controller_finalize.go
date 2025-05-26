// Copyright OpenChoreo Authors 2025
// SPDX-License-Identifier: Apache-2.0

package organization

import (
	"context"
	"errors"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	choreov1 "github.com/openchoreo/openchoreo/api/v1"
	"github.com/openchoreo/openchoreo/internal/controller"
)

const (
	// OrgCleanUpFinalizer is the finalizer for cleaning up organization resources
	OrgCleanUpFinalizer = "core.choreo.dev/delete-namespace"
)

var ErrNamespaceDeletionWait = errors.New("waiting for namespace to be deleted")

func (r *Reconciler) ensureFinalizer(ctx context.Context, organization *choreov1.Organization) (bool, error) {
	// If the organization is being deleted, no need to add the finalizer
	if !organization.DeletionTimestamp.IsZero() {
		return false, nil
	}

	if controllerutil.AddFinalizer(organization, OrgCleanUpFinalizer) {
		return true, r.Update(ctx, organization)
	}

	return false, nil
}

// finalize cleans up the data plane resources associated with the organization.
func (r *Reconciler) finalize(ctx context.Context, old, organization *choreov1.Organization) (ctrl.Result, error) {
	// If the finalizer is not there, no need to do anything
	if !controllerutil.ContainsFinalizer(organization, OrgCleanUpFinalizer) {
		return ctrl.Result{}, nil
	}

	if meta.SetStatusCondition(&organization.Status.Conditions, NewOrganizationFinalizingCondition(organization.Generation)) {
		if err := controller.UpdateStatusConditions(ctx, r.Client, old, organization); err != nil {
			return ctrl.Result{}, err
		}
		return ctrl.Result{}, nil
	}

	// Ensure the namespace is deleted
	if err := r.handleNamespaceDeletion(ctx, organization); err != nil {
		if errors.Is(err, ErrNamespaceDeletionWait) {
			// Returns non error result here.
			// Next reconcile will trigger once the namespace is deleted, as the org reconciler is watching the namespace.
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, err
	}

	// Remove finalizer once cleanup is done
	base := client.MergeFrom(organization.DeepCopy())
	if controllerutil.RemoveFinalizer(organization, OrgCleanUpFinalizer) {
		if err := r.Patch(ctx, organization, base); err != nil {
			return ctrl.Result{}, err
		}
	}

	return ctrl.Result{}, nil
}

// handleNamespaceDeletion Ensures the namespace is deleted when the organization is deleted
func (r *Reconciler) handleNamespaceDeletion(ctx context.Context, org *choreov1.Organization) error {
	namespace := &corev1.Namespace{}
	err := r.Get(ctx, types.NamespacedName{Name: org.Name}, namespace)

	if apierrors.IsNotFound(err) {
		return nil
	} else if err != nil {
		return fmt.Errorf("failed to get namespace: %w", err)
	}

	// If namespace still exists, attempt deletion
	if namespace.DeletionTimestamp.IsZero() {
		if err := r.Delete(ctx, namespace); err != nil {
			return fmt.Errorf("failed to delete namespace: %w", err)
		}
	}

	return fmt.Errorf("%w: %s", ErrNamespaceDeletionWait, org.Name)
}
