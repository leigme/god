package utils

import (
	"os"
	"strings"
)

func CheckFileExist(fileName string) bool {
	if dir, err := os.Getwd(); err == nil {
		if files, err := os.ReadDir(dir); err == nil && len(files) > 0 {
			for _, f := range files {
				if strings.EqualFold(f.Name(), fileName) {
					return true
				}
			}
		}
	}
	return false
}
