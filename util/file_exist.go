package util

import (
	"os"
)

func FileExists(f string) bool {
	_, err := os.Stat(f)
	return os.IsExist(err)
}
