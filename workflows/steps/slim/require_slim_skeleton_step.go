package slim

import (
	"fmt"
	"os"

	"github.com/racastellanosm/tools-php-skeleton-creator/utilities"
	"github.com/racastellanosm/tools-php-skeleton-creator/workflows/steps"
)

type RequireSlimSkeletonStep struct{}

func (s *RequireSlimSkeletonStep) Execute(parameters steps.StepParameters) error {
	fmt.Println("* Require Slim Skeleton Step")

	// Create project dir
	if err := os.MkdirAll(parameters.ProjectName, 0755); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", parameters.ProjectName, err)
	}

	if err := os.Chdir(parameters.ProjectName); err != nil {
		return fmt.Errorf("failed to change directory to %s: %w", parameters.ProjectName, err)
	}

	prefixCommand := []string{
		"require",
	}

	dependencies := []string{
		"slim/slim",
		"slim/psr7",
		"doctrine/dbal",
		"runtime/roadrunner-nyholm",
	}

	if output, err := utilities.RunComposer(utilities.ComposerRunner{}, append(prefixCommand, dependencies...)); err != nil {
		return fmt.Errorf("failed to add dependency: %w\n%s", err, output)
	}

	return nil
}
