package utilities

import (
	"embed"
	"fmt"
	"os"
)

//go:embed templates/*/*
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

func CopyFileFromDirectory(sourceDir string, destinationDir string, files []string) error {
	for _, file := range files {
		sourcePath := fmt.Sprintf("%s/%s", sourceDir, file)
		destinationPath := fmt.Sprintf("%s/%s", destinationDir, file)

		if err := CopyFile(sourcePath, destinationPath); err != nil {
			return fmt.Errorf("failed to copy file: %w", err)
		}
	}

	return nil
}
