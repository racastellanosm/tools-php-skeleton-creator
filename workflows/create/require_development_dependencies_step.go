package create

import (
	"fmt"

	"github.com/Equation-Labs-I-O/eqlabs-tools-php-skeleton-creator/utilities"
)

type RequireDevelopmentDependenciesStep struct{}

func (s *RequireDevelopmentDependenciesStep) Execute(projectName string) error {
	fmt.Println("* Require Development Dependencies Step")

	dependencies := []string{
		"symfony/test-pack",
		"symplify/config-transformer",
	}

	for _, dep := range dependencies {
		if output, err := utilities.RunComposer(utilities.ComposerRunner{}, []string{"require", "--dev", dep}); err != nil {
			return fmt.Errorf("failed to add dependency %s: %w\n%s", dep, err, output)
		}
	}

	return nil
}
