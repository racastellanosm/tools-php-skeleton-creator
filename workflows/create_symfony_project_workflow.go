package workflows

import (
	"github.com/racastellanosm/tools-php-skeleton-creator/workflows/steps"
	"github.com/racastellanosm/tools-php-skeleton-creator/workflows/steps/shared"
	"github.com/racastellanosm/tools-php-skeleton-creator/workflows/steps/symfony"
)

var _ Workflow = (*CreateSymfonyProjectWorkflow)(nil)

type CreateSymfonyProjectWorkflow struct {
	BaseWorkflow
}

func NewCreateSymfonyProjectWorkflow() *CreateSymfonyProjectWorkflow {
	createSymfonyProjectWorkflow := &CreateSymfonyProjectWorkflow{}

	createSymfonyProjectWorkflow.AddStep(&symfony.RequireSymfonySkeletonStep{})
	createSymfonyProjectWorkflow.AddStep(&symfony.RequireAdditionalDependenciesStep{})
	createSymfonyProjectWorkflow.AddStep(&symfony.RequireDevelopmentDependenciesStep{})
	createSymfonyProjectWorkflow.AddStep(&symfony.RequireDoctrineDependencies{})
	createSymfonyProjectWorkflow.AddStep(&symfony.RewriteYamlConfigToPhpStep{})
	createSymfonyProjectWorkflow.AddStep(&symfony.ReorganizeNeededFoldersStep{})
	createSymfonyProjectWorkflow.AddStep(&symfony.MergeEnvironmentFiles{})
	createSymfonyProjectWorkflow.AddStep(&shared.AddDockerComposeFileStep{})
	createSymfonyProjectWorkflow.AddStep(&shared.AddRoadrunnerFiles{})
	createSymfonyProjectWorkflow.AddStep(&shared.AddMakefileStep{})
	createSymfonyProjectWorkflow.AddStep(&shared.RemoveUnneededFiles{})

	return createSymfonyProjectWorkflow
}

func (w *CreateSymfonyProjectWorkflow) Handle(parameters steps.StepParameters) error {
	return w.BaseWorkflow.Handle(parameters, "green")
}
