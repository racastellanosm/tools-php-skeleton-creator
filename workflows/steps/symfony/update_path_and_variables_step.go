package symfony

import (
	"fmt"
	"os"
	"strings"

	"github.com/racastellanosm/tools-php-skeleton-creator/workflows/steps"
)

type UpdatePathsAndVariablesStep struct{}

func (s *UpdatePathsAndVariablesStep) Execute(parameters steps.StepParameters) error {
	fmt.Println("* Update Paths And Variables Step")

	// doctrine migration configuration
	doctrineMigrationsFile := parameters.ProjectName + "/config/packages/doctrine_migrations.php"
	content, err := os.ReadFile(doctrineMigrationsFile)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	oldString := "'%kernel.project_dir%/migrations'"
	newString := "'%kernel.project_dir%/src/Infrastructure/Resources/Migrations'"
	modified := strings.ReplaceAll(string(content), oldString, newString)

	err = os.WriteFile(doctrineMigrationsFile, []byte(modified), 0644)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	// doctrine file
	doctrineFile := parameters.ProjectName + "/config/packages/doctrine.php"
	content, err = os.ReadFile(doctrineFile)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	oldString = "'%env(resolve:DATABASE_URL)%'"
	newString = "'%env(resolve:APP_DATABASE_URL)%'"
	modified = strings.ReplaceAll(string(content), oldString, newString)

	err = os.WriteFile(doctrineFile, []byte(modified), 0644)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}
