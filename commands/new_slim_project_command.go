package commands

import (
    "github.com/Equation-Labs-I-O/eqlabs-tools-php-skeleton-creator/utilities"
    "github.com/Equation-Labs-I-O/eqlabs-tools-php-skeleton-creator/workflows"
    "github.com/spf13/cobra"
    "os"
)

var NewSlimProjectCommand = &cobra.Command{
    Use:   "slim [project-name]",
    Short: "Create a new slim project boilerplate",
    Long:  "Creates a new Slim framework skeleton project boilerplate including EquaotionLabs development guidelines",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        projectName := args[0]

        workflowDependencies := workflows.WorkflowDependencies{
            ProjectName: projectName,
        }

        commandWorkflow := workflows.NewCreateSlimProjectWorkflow()

        if err := commandWorkflow.Handle(workflowDependencies); err != nil {
            utilities.PrintErrorBox(os.Stdout, `Error: `+err.Error())
            os.Exit(1)
        }

        dir, _ := os.Getwd()
        var newVar = "[OK] Project created successfully in " + dir + "/" + projectName
        utilities.PrintSuccessBox(os.Stdout, newVar)
    },
}
