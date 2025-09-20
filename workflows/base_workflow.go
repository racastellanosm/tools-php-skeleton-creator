package workflows

import (
	"fmt"
	"os"

	"github.com/racastellanosm/tools-php-skeleton-creator/utilities"
	"github.com/racastellanosm/tools-php-skeleton-creator/workflows/steps"
)

// BaseWorkflow provides the common structure and execution logic for all workflows.
type BaseWorkflow struct {
	steps []steps.Step
}

// AddStep adds a step to the workflow's execution list.
func (w *BaseWorkflow) AddStep(step steps.Step) {
	w.steps = append(w.steps, step)
}

// Handle executes the workflow's steps sequentially.
// It takes a spinnerColor to allow customization for different workflows.
func (w *BaseWorkflow) Handle(parameters steps.StepParameters, spinnerColor string) error {
	spinner := utilities.NewSpinner(spinnerColor)

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

func rollbackInCaseOfFailure(projectName string) error {
	fmt.Println("Rollback & Cleaning Everything Step")
	// This assumes the current working directory is the parent of the project directory.
	// This can be fragile if steps change the directory unexpectedly.
	return os.RemoveAll(projectName)
}
