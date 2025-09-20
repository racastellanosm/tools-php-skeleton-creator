package create

import (
	"fmt"
	"os/exec"

	"github.com/racastellanosm/tools-php-skeleton-creator/workflows/steps"
)

type RewriteYamlConfigToPhpStep struct{}

func (s *RewriteYamlConfigToPhpStep) Execute(parameters steps.StepParameters) error {
	fmt.Println("* Rewrite YAML Config to PHP Step for " + parameters.ProjectName)

	command := exec.Command("vendor/bin/config-transformer")

	if err := command.Run(); err != nil {
		return fmt.Errorf("failed to execute command: %w", err)
	}

	return nil
}
