package workflows

import (
	"fmt"

	"github.com/racastellanosm/tools-php-skeleton-creator/utilities"
	create "github.com/racastellanosm/tools-php-skeleton-creator/workflows/create_slim_project"
	"github.com/racastellanosm/tools-php-skeleton-creator/workflows/shared"
	"github.com/racastellanosm/tools-php-skeleton-creator/workflows/steps"
)

var _ Workflow = (*CreateSlimProjectWorkflow)(nil)

// implement the Workflow interface in this class

type CreateSlimProjectWorkflow struct {
	steps []steps.Step
}

func (w *CreateSlimProjectWorkflow) AddStep(step steps.Step) {
	w.steps = append(w.steps, step)
}

func NewCreateSlimProjectWorkflow() *CreateSlimProjectWorkflow {
	newCreateSlimProjectWorkflow := &CreateSlimProjectWorkflow{}

	newCreateSlimProjectWorkflow.AddStep(&create.RequireSlimSkeletonStep{})
	newCreateSlimProjectWorkflow.AddStep(&create.RequireDevelopmentDependenciesStep{})
	newCreateSlimProjectWorkflow.AddStep(&create.ReorganizeNeededFoldersStep{})
	newCreateSlimProjectWorkflow.AddStep(&shared.AddDockerComposeFileStep{})
	newCreateSlimProjectWorkflow.AddStep(&shared.AddMakefileStep{})

	return newCreateSlimProjectWorkflow
}

func (w *CreateSlimProjectWorkflow) Handle(parameters steps.StepParameters) error {
	spinner := utilities.NewSpinner("yellow")

	for _, step := range w.steps {
		spinner.Start()

		stepError := step.Execute(parameters)
		if stepError != nil {
			// use rollback step to clean everything up
			rollbackError := rollbackInCaseOfFailure(parameters.ProjectName)
			if rollbackError != nil {
				return fmt.Errorf("failed to cleanup workspace for %s : %w", parameters.ProjectName, rollbackError)
			}
			return stepError
		}

		spinner.Stop()
	}

	return nil
}
