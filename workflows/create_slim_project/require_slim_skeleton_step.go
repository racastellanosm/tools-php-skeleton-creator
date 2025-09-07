package create

import (
	"fmt"
	"os"

	"github.com/racastellanosm/tools-php-skeleton-creator/utilities"
)

type RequireSlimSkeletonStep struct{}

func (s *RequireSlimSkeletonStep) Execute(projectName string) error {
	fmt.Println("* Require Slim Skeleton Step")
	
	// Create project dir
	if err := os.MkdirAll(projectName, 0755); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", projectName, err)
	}
	
	if err := os.Chdir(projectName); err != nil {
		return fmt.Errorf("failed to change directory to %s: %w", projectName, err)
	}

	prefixCommand := []string{
		"require",
	}

	dependencies := []string{
		"slim/slim",
		"slim/psr7",
	}

	if output, err := utilities.RunComposer(utilities.ComposerRunner{}, append(prefixCommand, dependencies...)); err != nil {
		return fmt.Errorf("failed to add dependency: %w\n%s", err, output)
	}

	return nil
}
