package _path

import (
	"os"
)

func Exist(path string) (bool, os.FileInfo) {
	var info, err = os.Stat(path)
	if os.IsExist(err) {
		return true, info
	}
	return false, nil
}
