package ci

import (
	choreov1 "github.com/choreo-idp/choreo/api/v1"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"testing"
)

func TestDeploymentIntegrationKubernetes(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Build Integration Kubernetes CI Suite")
}

func newBuildpackBasedBuild() *choreov1.Build {
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
