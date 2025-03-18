package _folder

import (
	"os"

	_path "github.com/crlspe/frame-go/util/path"
)

type location struct {
	OSName string
}

func Locations() location {

	return location{
		OSName: "linux",
	}
}

func (l location) GetUserHomeFolder() string {
	return HomeFolder()
}

func HomeFolder() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return homeDir
}

func Exists(filePath string) bool {
	var exists, info = _path.Exist(filePath)
	return exists && info.IsDir()
}

func CreateFolder(folderPath string) error {
	err := os.MkdirAll(folderPath, 0755)
	if err != nil {
		return err
	}
	return nil
}
