package create

import (
	"fmt"

	"github.com/Equation-Labs-I-O/eqlabs-tools-php-skeleton-creator/utilities"
)

type RequireDevelopmentDependenciesStep struct{}

func (s *RequireDevelopmentDependenciesStep) Execute(projectName string) error {
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
