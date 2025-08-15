package utilities

import (
	"embed"
	"fmt"
	"os"
)

//go:embed templates/*
var staticTemplateSources embed.FS

func CopyFile(sourceFileName string, targetPath string) error {

	// Get the File by its name from staticTemplateSources
	originalFileBytes, err := staticTemplateSources.ReadFile(sourceFileName)
	if err != nil {
		return fmt.Errorf("failed to open the static file from sources: %w", err)
	}

	err = os.WriteFile(targetPath, originalFileBytes, 0644)
	if err != nil {
		return fmt.Errorf("failed to write the file on destination path: %w", err)
	}

	return nil
}
