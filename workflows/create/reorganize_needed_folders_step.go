package create

import (
	"fmt"
	"time"
)

type ReorganizeNeededFoldersStep struct{}

func (s *ReorganizeNeededFoldersStep) Execute(projectName string) error {
	fmt.Println("Reorganize Needed Folders Step" + projectName)
	time.Sleep(1 * time.Second)

	return nil
}
