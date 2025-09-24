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
	createSlimProjectWorkflow := &CreateSlimProjectWorkflow{}

	createSlimProjectWorkflow.AddStep(&slim.RequireSlimSkeletonStep{})
	createSlimProjectWorkflow.AddStep(&slim.RequireDevelopmentDependenciesStep{})
	createSlimProjectWorkflow.AddStep(&slim.ReorganizeNeededFoldersStep{})
	createSlimProjectWorkflow.AddStep(&shared.MergeEnvironmentFilesStep{})
	createSlimProjectWorkflow.AddStep(&shared.AddDockerComposeFileStep{})
	createSlimProjectWorkflow.AddStep(&shared.AddRoadrunnerFilesStep{})
	createSlimProjectWorkflow.AddStep(&shared.AddMakefileStep{})
	createSlimProjectWorkflow.AddStep(&shared.CleanUpStep{})

	return createSlimProjectWorkflow
}

func (w *CreateSlimProjectWorkflow) Handle(parameters steps.StepParameters) error {
	return w.BaseWorkflow.Handle(parameters, "yellow")
}
