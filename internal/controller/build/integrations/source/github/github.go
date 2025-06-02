// Copyright 2025 The OpenChoreo Authors
// SPDX-License-Identifier: Apache-2.0

package github

import (
	"context"
	"fmt"

	"github.com/google/go-github/v69/github"
	"gopkg.in/yaml.v3"

	"github.com/openchoreo/openchoreo/internal/controller/build/integrations"
	"github.com/openchoreo/openchoreo/internal/controller/build/integrations/source"
)

type githubHandler struct {
	githubClient *github.Client
}

var _ source.SourceHandler[integrations.BuildContext] = (*githubHandler)(nil)

func NewGithubHandler(githubClient *github.Client) source.SourceHandler[integrations.BuildContext] {
	return &githubHandler{
		githubClient: githubClient,
	}
}

func (h *githubHandler) Name(ctx context.Context, builtCtx *integrations.BuildContext) string {
	return "SourceGithub"
}

func (h *githubHandler) FetchComponentDescriptor(ctx context.Context, buildCtx *integrations.BuildContext) (*source.Config, error) {
	owner, repositoryName, err := source.ExtractRepositoryInfo(buildCtx.Component.Spec.Source.GitRepository.URL)
	if err != nil {
		return nil, fmt.Errorf("bad git repository url: %w", err)
	}
	// If the build has a specific git revision, use it. Otherwise, use the default branch.
	ref := buildCtx.Build.Spec.Branch
	if buildCtx.Build.Spec.GitRevision != "" {
		ref = buildCtx.Build.Spec.GitRevision
	}

	componentYaml, _, _, err := h.githubClient.Repositories.GetContents(ctx, owner, repositoryName,
		source.MakeComponentDescriptorPath(buildCtx), &github.RepositoryContentGetOptions{Ref: ref})
	if err != nil {
		return nil, fmt.Errorf("failed to get component.yaml from the repository buildName:%s;owner:%s;repo:%s;%w", buildCtx.Build.Name, owner, repositoryName, err)
	}
	componentYamlContent, err := componentYaml.GetContent()
	if err != nil {
		return nil, fmt.Errorf("failed to get content of component.yaml from the repository  buildName:%s;owner:%s;repo:%s;%w", buildCtx.Build.Name, owner, repositoryName, err)
	}
	config := source.Config{}
	err = yaml.Unmarshal([]byte(componentYamlContent), &config)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal component.yaml from the repository buildName:%s;owner:%s;repo:%s;%w", buildCtx.Build.Name, owner, repositoryName, err)
	}

	return &config, nil
}
