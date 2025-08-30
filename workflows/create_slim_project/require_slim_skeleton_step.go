package create

import (
	"fmt"

	"github.com/racastellanosm/tools-php-skeleton-creator/utilities"
)

type RequireSlimSkeletonStep struct{}

func (s *RequireSlimSkeletonStep) Execute(projectName string) error {
	fmt.Println("* Require Slim Skeleton Step")

	parameters := []string{"create-project", "slim/slim-skeleton", projectName}

	if _, err := utilities.RunComposer(utilities.ComposerRunner{}, parameters); err != nil {
		return fmt.Errorf("failed to create project with Composer: %w", err)
	}

	return nil
}
