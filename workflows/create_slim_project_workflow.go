package workflows

import (
	"fmt"
	"github.com/Equation-Labs-I-O/eqlabs-tools-php-skeleton-creator/utilities"
	create "github.com/Equation-Labs-I-O/eqlabs-tools-php-skeleton-creator/workflows/create_slim_project"
)

type CreatelimProjectWorkflow struct {
	steps []Step
}

func (w *CreatelimProjectWorkflow) AddStep(step Step) {
	w.steps = append(w.steps, step)
}

func NewCreateSlimProjectWorkflow() *CreatelimProjectWorkflow {
	newSlimProjectWorkflow := &CreatelimProjectWorkflow{}

	newSlimProjectWorkflow.AddStep(&create.RequireSlimSkeletonStep{})

	return newSlimProjectWorkflow
}

func (w *CreatelimProjectWorkflow) Handle(dependencies WorkflowDependencies) error {
	spinner := utilities.NewSpinner("yellow")

	for _, step := range w.steps {
		spinner.Start()

		err := step.Execute(dependencies.ProjectName)
		if err != nil {
			// use rollback step to clean everithing up
			err := rollbackInCaseOfFailure(dependencies.ProjectName)
			if err != nil {
				return fmt.Errorf("failed to cleanup workspace for %s : %w", dependencies.ProjectName, err)
			}
			return err
		}

		spinner.Stop()
	}

	return nil
}
