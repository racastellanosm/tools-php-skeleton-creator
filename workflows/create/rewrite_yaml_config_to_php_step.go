package create

import (
	"fmt"
	"os/exec"
)

type RewriteYamlConfigToPhpStep struct{}

func (s *RewriteYamlConfigToPhpStep) Execute(projectName string) error {
	fmt.Println("* Rewrite YAML Config to PHP Step for " + projectName)

	command := exec.Command("vendor/bin/config-transformer")

	if err := command.Run(); err != nil {
		return fmt.Errorf("failed to execute command: %w", err)
	}

	return nil
}
