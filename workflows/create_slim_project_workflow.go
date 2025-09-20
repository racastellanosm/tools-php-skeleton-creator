package workflows

import (
	"github.com/racastellanosm/tools-php-skeleton-creator/workflows/steps"
	"github.com/racastellanosm/tools-php-skeleton-creator/workflows/steps/shared"
	"github.com/racastellanosm/tools-php-skeleton-creator/workflows/steps/slim"
)

var _ Workflow = (*CreateSlimProjectWorkflow)(nil)

type CreateSlimProjectWorkflow struct {
	BaseWorkflow
}

func NewCreateSlimProjectWorkflow() *CreateSlimProjectWorkflow {
	newCreateSlimProjectWorkflow := &CreateSlimProjectWorkflow{}

	newCreateSlimProjectWorkflow.AddStep(&slim.RequireSlimSkeletonStep{})
	newCreateSlimProjectWorkflow.AddStep(&slim.RequireDevelopmentDependenciesStep{})
	newCreateSlimProjectWorkflow.AddStep(&slim.ReorganizeNeededFoldersStep{})
	newCreateSlimProjectWorkflow.AddStep(&shared.AddDockerComposeFileStep{})
	newCreateSlimProjectWorkflow.AddStep(&shared.AddRoadrunnerFiles{})
	newCreateSlimProjectWorkflow.AddStep(&shared.AddMakefileStep{})

	return newCreateSlimProjectWorkflow
}

func (w *CreateSlimProjectWorkflow) Handle(parameters steps.StepParameters) error {
	return w.BaseWorkflow.Handle(parameters, "yellow")
}
