package workflows

import (
    "fmt"
    "os"

    "github.com/Equation-Labs-I-O/eqlabs-tools-php-skeleton-creator/utilities"
    "github.com/Equation-Labs-I-O/eqlabs-tools-php-skeleton-creator/workflows/create_symfony_project"
)

type CreteSymfonyProjectWorkflow struct {
    steps []Step
}

func (w *CreteSymfonyProjectWorkflow) AddStep(step Step) {
    w.steps = append(w.steps, step)
}

func NewCreateSymfonyProjectWorkflow() *CreteSymfonyProjectWorkflow {
    createSymfonyProjetWorkflow := &CreteSymfonyProjectWorkflow{}

    createSymfonyProjetWorkflow.AddStep(&create.RequireSymfonySkeletonStep{})
    createSymfonyProjetWorkflow.AddStep(&create.RequireAdditionalDependenciesStep{})
    createSymfonyProjetWorkflow.AddStep(&create.RequireDevelopmentDependenciesStep{})
    createSymfonyProjetWorkflow.AddStep(&create.RewriteYamlConfigToPhpStep{})
    createSymfonyProjetWorkflow.AddStep(&create.ReorganizeNeededFoldersStep{})
    createSymfonyProjetWorkflow.AddStep(&create.AddDockerComposeFileStep{})
    createSymfonyProjetWorkflow.AddStep(&create.AddMakefileStep{})

    return createSymfonyProjetWorkflow
}

func (w *CreteSymfonyProjectWorkflow) Handle(dependencies WorkflowDependencies) error {
    spinner := utilities.NewSpinner("green")

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
