package utilities

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

// copyFileContents copies the contents of the file named src to the file named
// by dst. The file will be created if it does not already exist. If the
// destination file exists, all it's contents will be replaced by the contents
// of the source file.
func CopyFile(src, dst string) error {

	// Absolute path to the folder
	filepath, err := filepath.Abs(src)
	if err != nil {
		log.Fatalln(err)
	}

	// Open the source file for reading
	srcFile, err := os.Open(filepath)
	if err != nil {
		return fmt.Errorf("failed to open the %s: %w", dst, err)
	}
	defer srcFile.Close()

	// Open the destination file for writing
	dstFile, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("failed to create the %s buffer: %w", dst, err)
	}
	// Return any errors that result from closing the destination file
	// Will return nil if no errors occurred
	defer func() {
		cerr := dstFile.Close()
		if err == nil {
			err = cerr
		}
	}()

	// Copy the contents of the source file into the destination files
	if _, err = io.Copy(dstFile, srcFile); err != nil {
		return fmt.Errorf("failed to copy %s into directory: %w", dst, err)
	}

	err = dstFile.Sync()
	return nil
}
