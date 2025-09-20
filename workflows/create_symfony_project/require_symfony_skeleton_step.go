package create

import (
	"fmt"

	"github.com/racastellanosm/tools-php-skeleton-creator/utilities"
	"github.com/racastellanosm/tools-php-skeleton-creator/workflows/steps"
)

type RequireSymfonySkeletonStep struct{}

func (s *RequireSymfonySkeletonStep) Execute(parameters steps.StepParameters) error {
	fmt.Println("* Require Symfony Skeleton Step")

	composerParameters := []string{"create-project", "symfony/skeleton", parameters.ProjectName}

	if _, err := utilities.RunComposer(utilities.ComposerRunner{}, composerParameters); err != nil {
		return fmt.Errorf("failed to create project with Composer: %w", err)
	}

	return nil
}
