package _file

import (
	"os"
	"strings"

	_path "github.com/crlspe/frame-go/util/path"
)

func GetLines(filename string) ([]string, error) {
	var content, err = os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return strings.Split(string(content), "\n"), nil
}

func AppendContent(filename string, content string) error {
	if len(strings.TrimSpace(content)) == 0 {
		return nil
	}
	var file, err = os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write([]byte(content))
	if err != nil {
		return err
	}

	return nil
}

func Exists(filePath string) bool {
	var exists, info = _path.Exist(filePath)
	return exists && !info.IsDir()
}
