package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var newCommand = &cobra.Command{
	Use:   "new [project-name]",
	Short: "Create a new PHP project skeleton",
	Long: `Create a new PHP project skeleton with the specified name.
This command initializes a new project directory with the necessary files and structure for a PHP application.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		// if not arg is provided, print usage
		if name == "" {
			cmd.Help()
			return
		}

		// Logic to create a new PHP project skeleton
		fmt.Printf("Creating new PHP project skeleton: %s\n", name)
		// Here you would add the logic to actually create the project
	},
}
