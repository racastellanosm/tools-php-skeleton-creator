package create

import (
	"fmt"

	"github.com/Equation-Labs-I-O/eqlabs-tools-php-skeleton-creator/utilities"
)

type RequireSymfonySkeletonStep struct{}

func (s *RequireSymfonySkeletonStep) Execute(projectName string) error {
	fmt.Println("* Require Symfony Skeleton Step")

	parameters := []string{"create-project", "symfony/skeleton", projectName}

	if _, err := utilities.RunComposer(utilities.ComposerRunner{}, parameters); err != nil {
		return fmt.Errorf("failed to create project with Composer: %w", err)
	}

	return nil
}
