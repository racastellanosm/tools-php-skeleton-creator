package commands

import (
	"os"

	"github.com/racastellanosm/tools-php-skeleton-creator/utilities"
	"github.com/racastellanosm/tools-php-skeleton-creator/workflows"
	"github.com/spf13/cobra"
)

var (
	slimProjectName string
	slimDatabase    string
)

var CreateProjectWithSlimCommand = &cobra.Command{
	Use:   "create-with-slim --name your-project-name [--db postgresql|mysql ]",
	Short: "Create a new PHP skeleton with Slim Framework",
	Long: `Create a new PHP project wiht Slim framework following recommended development guidelines (DDD, CQRS, Testing, Automation).
This command initializes a new project directory with the necessary files and structure for a PHP application.`,
	Run: func(cmd *cobra.Command, args []string) {

		workflowDependencies := workflows.WorkflowDependencies{
			ProjectName: slimProjectName,
			Database:    slimDatabase,
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

func init() {
	CreateProjectWithSlimCommand.Flags().StringVar(&slimProjectName, "name", "", "The desired project name")
	CreateProjectWithSlimCommand.Flags().StringVar(&slimDatabase, "db", "postgresql", "The desired database name")

	if err := CreateProjectWithSlimCommand.MarkFlagRequired("name"); err != nil {
		utilities.PrintErrorBox(os.Stdout, "Project name is required: ´--name your-project-name´")
	}
}
