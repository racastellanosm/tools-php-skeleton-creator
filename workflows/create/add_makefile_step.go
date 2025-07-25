package create

import (
	"fmt"
	"time"
)

type AddMakefileStep struct{}

func (s *AddMakefileStep) Execute(projectName string) error {
	fmt.Println("Add Makefile Step" + projectName)
	time.Sleep(1 * time.Second)

	return nil
}
