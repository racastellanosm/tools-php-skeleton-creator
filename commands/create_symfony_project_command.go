package commands

import (
    "os"

    "github.com/racastellanosm/tools-php-skeleton-creator/utilities"
    "github.com/racastellanosm/tools-php-skeleton-creator/workflows"
    "github.com/spf13/cobra"
)

var CreateProjectWithSymfonyCommand = &cobra.Command{
    Use:   "create-with-symfony [project-name]",
    Short: "Create a new PHP project with Symfony Framework",
    Long: `Create a new PHP project with Symfony framework following recommended development guidelines (DDD, CQRS, Testing, Automation).
This command initializes a new project directory with the necessary files and structure for a PHP application.`,
    Args: cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        projectName := args[0]
        workflowDependencies := workflows.WorkflowDependencies{
            ProjectName: projectName,
        }

        commandWorkflow := workflows.NewCreateSymfonyProjectWorkflow()

        if err := commandWorkflow.Handle(workflowDependencies); err != nil {
            utilities.PrintErrorBox(os.Stdout, `Error: `+err.Error())
            os.Exit(1)
        }

        dir, _ := os.Getwd()
        var newVar = "[OK] Project created successfully in " + dir + "/" + projectName
        utilities.PrintSuccessBox(os.Stdout, newVar)
    },
}
