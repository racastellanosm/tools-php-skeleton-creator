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
    createProjetWorkflow := &CreteSymfonyProjectWorkflow{}

    createProjetWorkflow.AddStep(&create.RequireSymfonySkeletonStep{})
    createProjetWorkflow.AddStep(&create.RequireAdditionalDependenciesStep{})
    createProjetWorkflow.AddStep(&create.RequireDevelopmentDependenciesStep{})
    createProjetWorkflow.AddStep(&create.RewriteYamlConfigToPhpStep{})
    createProjetWorkflow.AddStep(&create.ReorganizeNeededFoldersStep{})
    createProjetWorkflow.AddStep(&create.AddDockerComposeFileStep{})
    createProjetWorkflow.AddStep(&create.AddMakefileStep{})

    return createProjetWorkflow
}

func (w *CreteSymfonyProjectWorkflow) Handle(dependencies WorkflowDependencies) error {
    spinner := utilities.NewSpinner("green")

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
