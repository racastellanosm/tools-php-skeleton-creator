package workflows

import (
	"fmt"
	"github.com/racastellanosm/tools-php-skeleton-creator/utilities"
	create "github.com/racastellanosm/tools-php-skeleton-creator/workflows/create_slim_project"
)

type CreatelimProjectWorkflow struct {
	steps []Step
}

func (w *CreatelimProjectWorkflow) AddStep(step Step) {
	w.steps = append(w.steps, step)
}

func NewCreateSlimProjectWorkflow() *CreatelimProjectWorkflow {
	newCreateSlimProjectWorkflow := &CreatelimProjectWorkflow{}

	newCreateSlimProjectWorkflow.AddStep(&create.RequireSlimSkeletonStep{})

	return newCreateSlimProjectWorkflow
}

func (w *CreatelimProjectWorkflow) Handle(dependencies WorkflowDependencies) error {
	spinner := utilities.NewSpinner("yellow")

	for _, step := range w.steps {
		spinner.Start()

		stepError := step.Execute(dependencies.ProjectName)
		if stepError != nil {
			// use rollback step to clean everithing up
			rollbackError := rollbackInCaseOfFailure(dependencies.ProjectName)
			if rollbackError != nil {
				return fmt.Errorf("failed to cleanup workspace for %s : %w", dependencies.ProjectName, rollbackError)
			}
			return stepError
		}

		spinner.Stop()
	}

	return nil
}
