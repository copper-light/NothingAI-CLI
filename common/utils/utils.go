package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func IsExsit(list []string, keyword string) bool {
	for _, v := range list {
		if v == keyword {
			return true
		}
	}
	return false
}

func IsHangul(r rune) bool {
	return r >= 0xAC00 && r <= 0xD7A3
}

func CountingHangul(s string) int {
	var cnt int = 0
	for _, c := range s {
		if IsHangul(c) {
			cnt++
		}
	}
	return cnt
}

func GetFileList(path string) []string {
	result := make([]string, 0)
	fullPath := ""
	file, err := os.Stat(path)
	if err != nil {
		return nil
	}
	if file.IsDir() {
		files, err := os.ReadDir(path)
		if path == "." {
			path = ""
		}
		if err != nil {
			fmt.Printf("Failed to read directory : %s\n", path)
			return nil
		}
		for _, file := range files {
			fullPath = filepath.Join(path, file.Name())
			if file.IsDir() {
				files := GetFileList(fullPath)
				result = append(result, files...)
			} else {
				result = append(result, fullPath)
			}
		}
	} else {
		result = append(result, path)
	}

	return result
}
