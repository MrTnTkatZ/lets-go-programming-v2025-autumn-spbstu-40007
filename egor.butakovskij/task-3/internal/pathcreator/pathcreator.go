package pathcreator

import (
	"fmt"
	"os"
	"path/filepath"
)

func PathCreator(outputFile string) error {
	outputDir := filepath.Dir(outputFile)

	if outputDir != "" {
		err := os.MkdirAll(outputDir, 0755)

		if err != nil {
			return fmt.Errorf("%w", err)
		}
	}

	return nil
}
