package create

import (
	"fmt"

	"github.com/racastellanosm/tools-php-skeleton-creator/utilities"
	"github.com/racastellanosm/tools-php-skeleton-creator/workflows/steps"
)

type RequireDevelopmentDependenciesStep struct{}

func (s *RequireDevelopmentDependenciesStep) Execute(parameters steps.StepParameters) error {
	fmt.Println("* Require Development Dependencies Step")

	prefixCommand := []string{
		"require",
		"--dev",
	}

	dependencies := []string{
		"symfony/test-pack",
		"symplify/config-transformer",
	}

	if output, err := utilities.RunComposer(utilities.ComposerRunner{}, append(prefixCommand, dependencies...)); err != nil {
		return fmt.Errorf("failed to add dependency: %w\n%s", err, output)
	}

	return nil
}
