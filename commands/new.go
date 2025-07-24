package commands

import (
	"fmt"
	"os"
	"time"

	"github.com/Equation-Labs-I-O/eqlabs-tools-php-skeleton-creator/utilities"
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
		spinner := utilities.NewSpinner("green")

		if err := createProjectWithComposer(name, spinner); err != nil {
			utilities.PrintErrorBox(os.Stdout, fmt.Sprintf("[ERROR] %v", err))
			return
		}

		if err := addAdditonalComposerDependencies(name, spinner); err != nil {
			utilities.PrintErrorBox(os.Stdout, fmt.Sprintf("[ERROR] %v", err))
			return
		}

		if err := createFolderStructure(spinner); err != nil {
			utilities.PrintErrorBox(os.Stdout, fmt.Sprintf("[ERROR] %v", err))
			return
		}

		if err := overrideYamlConfigForPhpf(spinner); err != nil {
			utilities.PrintErrorBox(os.Stdout, fmt.Sprintf("[ERROR] %v", err))
			return
		}

		if err := addDockerComposeFiles(spinner); err != nil {
			utilities.PrintErrorBox(os.Stdout, fmt.Sprintf("[ERROR] %v", err))
			return
		}

		if err := addMakefile(spinner); err != nil {
			utilities.PrintErrorBox(os.Stdout, fmt.Sprintf("[ERROR] %v", err))
			return
		}

		dir, _ := os.Getwd()
		var newVar = "[OK] Project created successfully in " + dir + "/" + name
		utilities.PrintSuccessBox(os.Stdout, newVar)
	},
}

func createProjectWithComposer(projectName string, s utilities.Spinner) error {
	s.Start()
	fmt.Println("* Creating a new Symfony project with Composer")

	runner := utilities.ComposerRunner{}
	if output, err := utilities.RunComposer(runner, []string{"create-project", "symfony/skeleton", projectName}); err != nil {
		return fmt.Errorf("failed to create Symfony project: %w\n%s", err, output)
	}
	s.Stop()
	return nil
}

func addAdditonalComposerDependencies(projectName string, s utilities.Spinner) error {
	s.Start()
	fmt.Println("* Adding additional Composer dependencies to comply with CQRS and DDD patterns")

	if err := os.Chdir(projectName); err != nil {
		return fmt.Errorf("failed to change directory to %s: %w", projectName, err)
	}

	runner := utilities.ComposerRunner{}
	if output, err := utilities.RunComposer(runner, []string{"require", "symfony/messenger", "symfony/http-client", "symfony/uid"}); err != nil {
		return fmt.Errorf("failed to add additional Composer dependencies: %w\n%s", err, output)
	}

	if output, err := utilities.RunComposer(runner, []string{"require", "--dev", "symfony/test-pack"}); err != nil {
		return fmt.Errorf("failed to add development Composer dependencies: %w\n%s", err, output)
	}

	s.Stop()
	return nil
}

func createFolderStructure(s utilities.Spinner) error {
	s.Start()
	fmt.Println("* Creating DDD folder structure for project")
	time.Sleep(1 * time.Second)
	s.Stop()
	return nil
}

func overrideYamlConfigForPhpf(s utilities.Spinner) error {
	s.Start()
	fmt.Println("* Overriding YAML configuration for PHP")
	time.Sleep(1 * time.Second)
	s.Stop()
	return nil
}

func addDockerComposeFiles(s utilities.Spinner) error {
	s.Start()
	fmt.Println("* Adding Docker Compose files for project")
	time.Sleep(1 * time.Second)
	s.Stop()
	return nil
}

func addMakefile(s utilities.Spinner) error {
	s.Start()
	fmt.Println("* Adding Makefile for project management")
	time.Sleep(1 * time.Second)
	s.Stop()
	return nil
}
