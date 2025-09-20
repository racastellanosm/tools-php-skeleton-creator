package symfony

import (
	"fmt"
	"os"

	"github.com/racastellanosm/tools-php-skeleton-creator/utilities"
	"github.com/racastellanosm/tools-php-skeleton-creator/workflows/steps"
)

type RequireAdditionalDependenciesStep struct{}

func (s *RequireAdditionalDependenciesStep) Execute(parameters steps.StepParameters) error {
	fmt.Println("* Require Additional Dependencies Step")

	if err := os.Chdir(parameters.ProjectName); err != nil {
		return fmt.Errorf("failed to change directory to %s: %w", parameters.ProjectName, err)
	}

	dependencies := []string{
		"symfony/messenger",
		"symfony/console",
		"symfony/http-client",
		"runtime/roadrunner-symfony-nyholm",
	}

	for _, dep := range dependencies {
		if output, err := utilities.RunComposer(utilities.ComposerRunner{}, []string{"require", dep}); err != nil {
			return fmt.Errorf("failed to add dependency %s: %w\n%s", dep, err, output)
		}
	}

	return nil
}
