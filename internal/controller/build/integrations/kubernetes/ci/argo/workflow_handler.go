// Copyright OpenChoreo Authors 2025
// SPDX-License-Identifier: Apache-2.0

package argo

import (
	"context"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/openchoreo/openchoreo/internal/controller/build/integrations"
	"github.com/openchoreo/openchoreo/internal/controller/build/integrations/kubernetes"
	"github.com/openchoreo/openchoreo/internal/dataplane"
	dpkubernetes "github.com/openchoreo/openchoreo/internal/dataplane/kubernetes"
	argoproj "github.com/openchoreo/openchoreo/internal/dataplane/kubernetes/types/argoproj.io/workflow/v1alpha1"
)

type workflowHandler struct {
	kubernetesClient client.Client
}

var _ dataplane.ResourceHandler[integrations.BuildContext] = (*workflowHandler)(nil)

func NewWorkflowHandler(kubernetesClient client.Client) dataplane.ResourceHandler[integrations.BuildContext] {
	return &workflowHandler{
		kubernetesClient: kubernetesClient,
	}
}

func (h *workflowHandler) Name() string {
	return "ArgoWorkflow"
}

func (h *workflowHandler) GetCurrentState(ctx context.Context, builtCtx *integrations.BuildContext) (interface{}, error) {
	name := makeWorkflowName(builtCtx)
	workflow := argoproj.Workflow{}
	err := h.kubernetesClient.Get(ctx, client.ObjectKey{Name: name, Namespace: kubernetes.MakeNamespaceName(builtCtx)}, &workflow)
	if apierrors.IsNotFound(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return workflow, nil
}

func (h *workflowHandler) Create(ctx context.Context, builtCtx *integrations.BuildContext) error {
	workflow := makeArgoWorkflow(builtCtx)
	return h.kubernetesClient.Create(ctx, workflow)
}

func (h *workflowHandler) Update(ctx context.Context, builtCtx *integrations.BuildContext, currentState interface{}) error {
	return nil
}

func (h *workflowHandler) Delete(ctx context.Context, buildCtx *integrations.BuildContext) error {
	workflow := &argoproj.Workflow{
		ObjectMeta: metav1.ObjectMeta{
			Name:      makeWorkflowName(buildCtx),
			Namespace: kubernetes.MakeNamespaceName(buildCtx),
			Labels: map[string]string{
				dpkubernetes.LabelKeyManagedBy: dpkubernetes.LabelBuildControllerCreated,
			},
		},
	}
	err := h.kubernetesClient.Delete(ctx, workflow)
	if apierrors.IsNotFound(err) {
		return nil
	}
	return err
}

func (h *workflowHandler) IsRequired(builtCtx *integrations.BuildContext) bool {
	return true
}

// makeWorkflowName generates the workflow name using the build name.
// WorkflowName is limited to 63 characters.
func makeWorkflowName(buildCtx *integrations.BuildContext) string {
	return dpkubernetes.GenerateK8sNameWithLengthLimit(63, buildCtx.Build.ObjectMeta.Name)
}

func GetStepPhase(phase argoproj.NodePhase) integrations.StepPhase {
	switch phase {
	case argoproj.NodeRunning, argoproj.NodePending:
		return integrations.Running
	case argoproj.NodeFailed, argoproj.NodeError, argoproj.NodeSkipped:
		return integrations.Failed
	default:
		return integrations.Succeeded
	}
}

func GetStepByTemplateName(nodes argoproj.Nodes, step integrations.BuildWorkflowStep) (*argoproj.NodeStatus, bool) {
	for _, node := range nodes {
		if node.TemplateName == string(step) {
			return &node, true
		}
	}
	return nil, false
}

func GetImageNameFromWorkflow(output argoproj.Outputs) string {
	for _, param := range output.Parameters {
		if param.Name == "image" {
			return *param.Value
		}
	}
	return ""
}
