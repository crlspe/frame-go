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

func PathExist(path string) (bool, os.FileInfo) {
	var info , err = os.Stat(path)
	if os.IsExist(err) {
		return true, info
	}
	return false, nil
}

func FileExists(filePath string) bool {
	var exists, info = PathExist(filePath)
	return exists && !info.IsDir()
}

func FolderExists(filePath string) bool {
	var exists, info = PathExist(filePath)
	return exists && info.IsDir()
}
