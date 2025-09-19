package workflows

import (
	"fmt"
	"os"

	"github.com/racastellanosm/tools-php-skeleton-creator/utilities"
	create "github.com/racastellanosm/tools-php-skeleton-creator/workflows/create_symfony_project"
	shared "github.com/racastellanosm/tools-php-skeleton-creator/workflows/shared"
)

type CreateSymfonyProjectWorkflow struct {
	steps []Step
}

func (w *CreateSymfonyProjectWorkflow) AddStep(step Step) {
	w.steps = append(w.steps, step)
}

func NewCreateSymfonyProjectWorkflow() *CreateSymfonyProjectWorkflow {
	createSymfonyProjectWorkflow := &CreateSymfonyProjectWorkflow{}

	createSymfonyProjectWorkflow.AddStep(&create.RequireSymfonySkeletonStep{})
	createSymfonyProjectWorkflow.AddStep(&create.RequireAdditionalDependenciesStep{})
	createSymfonyProjectWorkflow.AddStep(&create.RequireDevelopmentDependenciesStep{})
	createSymfonyProjectWorkflow.AddStep(&create.RewriteYamlConfigToPhpStep{})
	createSymfonyProjectWorkflow.AddStep(&create.ReorganizeNeededFoldersStep{})
	createSymfonyProjectWorkflow.AddStep(&shared.AddDockerComposeFileStep{})
	createSymfonyProjectWorkflow.AddStep(&shared.AddMakefileStep{})

	return createSymfonyProjectWorkflow
}

func (w *CreateSymfonyProjectWorkflow) Handle(dependencies WorkflowDependencies) error {
	spinner := utilities.NewSpinner("green")

	for _, step := range w.steps {
		spinner.Start()

		stepError := step.Execute(dependencies.ProjectName)
		if stepError != nil {
			// use rollback step to clean everything up
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

func rollbackInCaseOfFailure(projectName string) error {
	fmt.Println("Rollback & Cleaning Everything Step")

	if err := os.Chdir(".."); err != nil {
		return fmt.Errorf("failed to change directory to %s: %w", projectName, err)
	}

	if err := os.RemoveAll(projectName); err != nil {
		return fmt.Errorf("failed to remove directory %s: %w", projectName, err)
	}

	return nil
}
