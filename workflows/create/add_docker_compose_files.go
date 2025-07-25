package create

import (
	"fmt"
	"time"
)

type AddDockerComposeFilesStep struct{}

func (s *AddDockerComposeFilesStep) Execute(projectName string) error {
	fmt.Println("Add Docker Compose Files Step" + projectName)
	time.Sleep(1 * time.Second)

	return nil
}
