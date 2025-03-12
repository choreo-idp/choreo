package source

import (
	choreov1 "github.com/choreo-idp/choreo/api/v1"
	"github.com/choreo-idp/choreo/internal/controller/build/integrations"
	"github.com/choreo-idp/choreo/internal/labels"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"testing"
)

func TestDeploymentIntegrationKubernetes(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Build Integration Kubernetes CI Suite")
}

// Create a new BuildContext for testing. Each test should create a new context
// and set the required fields for the test.
func newTestBuildContext() *integrations.BuildContext {
	buildCtx := &integrations.BuildContext{}

	buildCtx.Component = &choreov1.Component{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-component",
			Namespace: "test-organization",
			Labels: map[string]string{
				labels.LabelKeyOrganizationName: "test-organization",
				labels.LabelKeyProjectName:      "test-project",
				labels.LabelKeyName:             "test-component",
			},
		},
		Spec: choreov1.ComponentSpec{
			Type: choreov1.ComponentTypeService,
			Source: choreov1.ComponentSource{
				GitRepository: &choreov1.GitRepository{
					URL: "https://github.com/choreo-idp/test",
				},
			},
		},
	}
	buildCtx.DeploymentTrack = &choreov1.DeploymentTrack{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-main-track",
			Namespace: "test-organization",
			Labels: map[string]string{
				labels.LabelKeyOrganizationName: "test-organization",
				labels.LabelKeyProjectName:      "test-project",
				labels.LabelKeyComponentName:    "test-component",
				labels.LabelKeyName:             "test-main",
			},
		},
	}
	buildCtx.Build = &choreov1.Build{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-build",
			Namespace: "test-organization",
			Labels: map[string]string{
				labels.LabelKeyOrganizationName:    "test-organization",
				labels.LabelKeyProjectName:         "test-project",
				labels.LabelKeyComponentName:       "test-component",
				labels.LabelKeyDeploymentTrackName: "test-main",
				labels.LabelKeyName:                "test-build",
			},
		},
	}

	return buildCtx
}

func newTestBuildpackBasedBuild() *choreov1.Build {
	return &choreov1.Build{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test-build",
			Labels: map[string]string{
				"core.choreo.dev/organization":     "test-organization",
				"core.choreo.dev/project":          "test-project",
				"core.choreo.dev/component":        "test-component",
				"core.choreo.dev/deployment-track": "test-main",
				"core.choreo.dev/name":             "test-build",
			},
			Namespace: "test-organization",
		},
		Spec: choreov1.BuildSpec{
			Branch: "main",
			Path:   "/test-service",
			BuildConfiguration: choreov1.BuildConfiguration{
				Buildpack: &choreov1.BuildpackConfiguration{
					Name:    "Go",
					Version: "1.x",
				},
			},
		},
	}
}

func newTestDeployableArtifact() *choreov1.DeployableArtifact {
	return &choreov1.DeployableArtifact{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test-build",
			Labels: map[string]string{
				"core.choreo.dev/organization":     "test-organization",
				"core.choreo.dev/project":          "test-project",
				"core.choreo.dev/component":        "test-component",
				"core.choreo.dev/deployment-track": "test-main",
				"core.choreo.dev/name":             "test-build",
			},
			Namespace: "test-organization",
		},
		Spec: choreov1.DeployableArtifactSpec{
			TargetArtifact: choreov1.TargetArtifact{
				FromBuildRef: &choreov1.FromBuildRef{
					Name: "test-build",
				},
			},
		},
	}
}

func newTestEndpoints() *[]choreov1.EndpointTemplate {
	return &[]choreov1.EndpointTemplate{
		{
			ObjectMeta: metav1.ObjectMeta{
				Name: "endpoint-1",
			},
			Spec: choreov1.EndpointSpec{
				Type: "HTTP",
				Service: choreov1.EndpointServiceSpec{
					BasePath: "/api/v1",
					Port:     80,
				},
			},
		},
	}
}
