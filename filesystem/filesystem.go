package filesystem

import (
	"fmt"
	"os"
)

func HomeFolder() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return homeDir
}

func CreateFolder(folderPath string) error {
	err := os.MkdirAll(folderPath, 0755)
	if err != nil {
		return fmt.Errorf("error creating directory: %s: %v", folderPath, err)
	}
	return nil
}
