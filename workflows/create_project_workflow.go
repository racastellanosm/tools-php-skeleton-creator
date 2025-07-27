package workflows

import (
	"fmt"
	"os"

	"github.com/Equation-Labs-I-O/eqlabs-tools-php-skeleton-creator/utilities"
)

type CreteProjectWorkflow struct {
	steps []Step
}

func (w *CreteProjectWorkflow) AddStep(step Step) {
	w.steps = append(w.steps, step)
}

func (w *CreteProjectWorkflow) Handle(dependencies WorkflowDependencies) error {
	spinner := utilities.NewSpinner("yellow")

	for _, step := range w.steps {
		spinner.Start()

		err := step.Execute(dependencies.ProjectName)
		if err != nil {
			// use rollback step to clean everithing up
			rollbackInCaseOfFailure(dependencies.ProjectName)
			return err
		}

		spinner.Stop()
	}

	return nil
}

func rollbackInCaseOfFailure(projectName string) error {
	fmt.Println("Rollback & Cleaning Everithing Step")

	if err := os.Chdir(".."); err != nil {
		return fmt.Errorf("failed to change directory to %s: %w", projectName, err)
	}

	if err := os.RemoveAll(projectName); err != nil {
		return fmt.Errorf("failed to remove directory %s: %w", projectName, err)
	}

	return nil
}

func NewCreteProjectWorkflow() *CreteProjectWorkflow {
	return &CreteProjectWorkflow{}
}
