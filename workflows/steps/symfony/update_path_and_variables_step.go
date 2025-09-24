package symfony

import (
	"fmt"

	"github.com/racastellanosm/tools-php-skeleton-creator/utilities"
	"github.com/racastellanosm/tools-php-skeleton-creator/workflows/steps"
)

type UpdatePathsAndVariablesStep struct{}

func (s *UpdatePathsAndVariablesStep) Execute(parameters steps.StepParameters) error {
	fmt.Println("* Update Paths And Variables Step")

	// doctrine migration configuration
	doctrineMigrationsFile := parameters.ProjectName + "/config/packages/doctrine_migrations.php"
	oldString := "'%kernel.project_dir%/migrations'"
	newString := "'%kernel.project_dir%/src/Infrastructure/Resources/Migrations'"

	if err := utilities.ReplaceAndCreate(doctrineMigrationsFile, oldString, newString); err != nil {
		return fmt.Errorf("failed to replace and create file: %w", err)
	}

	// doctrine file
	doctrineFile := parameters.ProjectName + "/config/packages/doctrine.php"
	oldString = "'%env(resolve:DATABASE_URL)%'"
	newString = "'%env(resolve:APP_DATABASE_URL)%'"

	if err := utilities.ReplaceAndCreate(doctrineFile, oldString, newString); err != nil {
		return fmt.Errorf("failed to replace and create file: %w", err)
	}

	// router file
	routerFile := parameters.ProjectName + "/config/packages/routing.php"
	oldString = "'%env(DEFAULT_URI)%'"
	newString = "'%env(APP_DEFAULT_URI)%'"

	if err := utilities.ReplaceAndCreate(routerFile, oldString, newString); err != nil {
		return fmt.Errorf("failed to replace and create file: %w", err)
	}

	// routing file
	routingFile := parameters.ProjectName + "/config/routes.php"
	oldString = "''../src/Controller/'"
	newString = "'../src/UI/Http/Controller/'"

	if err := utilities.ReplaceAndCreate(routingFile, oldString, newString); err != nil {
		return fmt.Errorf("failed to replace and create file: %w", err)
	}

	return nil
}
