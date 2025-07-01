// Copyright 2025 The OpenChoreo Authors
// SPDX-License-Identifier: Apache-2.0

package scheduledtaskrelease

import (
	"context"
	"fmt"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	choreov1 "github.com/openchoreo/openchoreo/api/v1"
	dpKubernetes "github.com/openchoreo/openchoreo/internal/dataplane/kubernetes"
	"github.com/openchoreo/openchoreo/internal/labels"
)

// Reconciler reconciles a ScheduledTaskRelease object
type Reconciler struct {
	client.Client
	DpClientMgr *dpKubernetes.KubeClientManager
	Scheme      *runtime.Scheme
}

// +kubebuilder:rbac:groups=core.choreo.dev,resources=scheduledtaskreleases,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core.choreo.dev,resources=scheduledtaskreleases/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=core.choreo.dev,resources=scheduledtaskreleases/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
func (r *Reconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	// Fetch the ScheduledTaskRelease instance
	scheduledTaskRelease := &choreov1.ScheduledTaskRelease{}
	if err := r.Get(ctx, req.NamespacedName, scheduledTaskRelease); err != nil {
		if apierrors.IsNotFound(err) {
			logger.Info("ScheduledTaskRelease resource not found. Ignoring since it must be deleted.")
			return ctrl.Result{}, nil
		}
		logger.Error(err, "Failed to get ScheduledTaskRelease")
		return ctrl.Result{}, err
	}

	// Get dataplane client for the environment
	dpClient, err := r.getDPClient(ctx, scheduledTaskRelease.Spec.EnvironmentName)
	if err != nil {
		logger.Error(err, "Failed to get dataplane client")
		return ctrl.Result{}, err
	}

	// Apply all resources to the dataplane
	if err := r.applyResources(ctx, dpClient, scheduledTaskRelease); err != nil {
		logger.Error(err, "Failed to apply resources to dataplane")
		return ctrl.Result{}, err
	}

	logger.Info("Successfully applied ScheduledTaskRelease resources to dataplane")
	return ctrl.Result{}, nil
}

// getDPClient gets the dataplane client for the specified environment
func (r *Reconciler) getDPClient(ctx context.Context, environmentName string) (client.Client, error) {
	// Fetch the environment from default namespace
	env := &choreov1.Environment{}
	if err := r.Get(ctx, client.ObjectKey{Name: environmentName, Namespace: "default"}, env); err != nil {
		return nil, fmt.Errorf("failed to get environment %s: %w", environmentName, err)
	}

	// Get the dataplane using the direct reference from default namespace
	dataplane := &choreov1.DataPlane{}
	if err := r.Get(ctx, client.ObjectKey{Name: env.Spec.DataPlaneRef, Namespace: "default"}, dataplane); err != nil {
		return nil, fmt.Errorf("failed to get dataplane %s for environment %s: %w", env.Spec.DataPlaneRef, environmentName, err)
	}

	// Get the dataplane client
	dpClient, err := r.DpClientMgr.GetClient(dataplane.Name, dataplane.Spec.KubernetesCluster.Credentials)
	if err != nil {
		return nil, fmt.Errorf("failed to create dataplane client for %s: %w", dataplane.Name, err)
	}

	return dpClient, nil
}

// applyResources applies all resources from ScheduledTaskRelease to the dataplane
func (r *Reconciler) applyResources(ctx context.Context, dpClient client.Client, scheduledTaskRelease *choreov1.ScheduledTaskRelease) error {
	logger := log.FromContext(ctx)

	for _, resource := range scheduledTaskRelease.Spec.Resources {
		logger.Info("Applying resource", "resourceID", resource.ID)

		// Convert RawExtension to Unstructured
		obj := &unstructured.Unstructured{}
		if err := obj.UnmarshalJSON(resource.Object.Raw); err != nil {
			return fmt.Errorf("failed to unmarshal resource %s: %w", resource.ID, err)
		}

		// Add ownership labels for tracking
		r.addOwnershipLabels(obj, scheduledTaskRelease)

		// Apply the resource using server-side apply
		if err := r.applyResource(ctx, dpClient, obj); err != nil {
			return fmt.Errorf("failed to apply resource %s: %w", resource.ID, err)
		}

		logger.Info("Successfully applied resource", "resourceID", resource.ID, "kind", obj.GetKind(), "name", obj.GetName())
	}

	return nil
}

// addOwnershipLabels adds OpenChoreo ownership labels to the resource
func (r *Reconciler) addOwnershipLabels(obj *unstructured.Unstructured, scheduledTaskRelease *choreov1.ScheduledTaskRelease) {
	resourceLabels := obj.GetLabels()
	if resourceLabels == nil {
		resourceLabels = make(map[string]string)
	}

	// Add OpenChoreo tracking labels
	resourceLabels[labels.LabelKeyProjectName] = scheduledTaskRelease.Spec.Owner.ProjectName
	resourceLabels[labels.LabelKeyComponentName] = scheduledTaskRelease.Spec.Owner.ComponentName
	resourceLabels[labels.LabelKeyEnvironmentName] = scheduledTaskRelease.Spec.EnvironmentName

	obj.SetLabels(resourceLabels)
}

// applyResource applies a single resource to the dataplane using server-side apply
func (r *Reconciler) applyResource(ctx context.Context, dpClient client.Client, obj *unstructured.Unstructured) error {
	// Use server-side apply for better conflict resolution
	return dpClient.Patch(ctx, obj, client.Apply, client.ForceOwnership, client.FieldOwner("scheduledtaskrelease-controller"))
}

// SetupWithManager sets up the controller with the Manager.
func (r *Reconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&choreov1.ScheduledTaskRelease{}).
		Named("scheduledtaskrelease").
		Complete(r)
}
