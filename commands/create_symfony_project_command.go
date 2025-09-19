package commands

import (
	"os"

	"github.com/racastellanosm/tools-php-skeleton-creator/utilities"
	"github.com/racastellanosm/tools-php-skeleton-creator/workflows"
	"github.com/spf13/cobra"
)

var (
	projectName string
	database    string
)

var CreateProjectWithSymfonyCommand = &cobra.Command{
	Use:   "create-with-symfony --name your-project-name [--db postgresql|mysql ]",
	Short: "Create a new PHP project with Symfony Framework",
	Long: `Create a new PHP project with Symfony framework following recommended development guidelines (DDD, CQRS, Testing, Automation).
This command initializes a new project directory with the necessary files and structure for a PHP application.`,
	Run: func(cmd *cobra.Command, args []string) {
		workflowDependencies := workflows.WorkflowDependencies{
			ProjectName: projectName,
			Database:    database,
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

func init() {
	CreateProjectWithSymfonyCommand.Flags().StringVar(&projectName, "name", "", "The desired project name")
	CreateProjectWithSymfonyCommand.Flags().StringVar(&database, "db", "postgresql", "The desired database name (mysql|postgresql)")

	// Required Flags
	if err := CreateProjectWithSymfonyCommand.MarkFlagRequired("name"); err != nil {
		utilities.PrintErrorBox(os.Stderr, "Project name is required: ´--name your-project-name´")
	}
}
