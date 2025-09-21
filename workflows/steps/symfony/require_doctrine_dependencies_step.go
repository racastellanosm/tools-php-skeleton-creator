package symfony

import (
	"fmt"

	"github.com/racastellanosm/tools-php-skeleton-creator/utilities"
	"github.com/racastellanosm/tools-php-skeleton-creator/workflows/steps"
)

type RequireDoctrineDependencies struct{}

func (s *RequireDoctrineDependencies) Execute(parameters steps.StepParameters) error {
	fmt.Println("* Require Additional Dependencies Step")

	dependenciesNoScript := []string{
		"doctrine/doctrine-migrations-bundle",
	}

	for _, dep := range dependenciesNoScript {
		if output, err := utilities.RunComposer(utilities.ComposerRunner{}, []string{"require", dep, "--no-scripts"}); err != nil {
			return fmt.Errorf("failed to add dependency %s: %w\n%s", dep, err, output)
		}
	}

	return nil
}
