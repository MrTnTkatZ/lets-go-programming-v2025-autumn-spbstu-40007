package pathcreator

import (
	"fmt"
	"os"
	"path/filepath"
)

const DefaultDirPermissions = 0o755

func PathCreator(outputFile string) error {
	outputDir := filepath.Dir(outputFile)

	if outputDir != "" {
		err := os.MkdirAll(outputDir, DefaultDirPermissions)
		if err != nil {
			return fmt.Errorf("%w", err)
		}
	}

	return nil
}
