package organiser

import (
	"fmt"
	"os"
	"path/filepath"
)

// Create struct for File
type File struct {
	CurrDir string
}

// this is just experimentally created by me
type UniqueFile interface {
	GetUniqueFile([]os.DirEntry, string) map[string]int
}

func (f *File) GetAllFilesInDir(path string) ([]os.DirEntry, error) {
	dir, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	return dir, err
}

func (f *File) GetCountUniqueFileTypes(files []os.DirEntry) map[string]int {
	extCount := make(map[string]int)
	fileExt := []string{}

	for _, f := range files {
		fileExt = append(fileExt, filepath.Ext(f.Name()))
	}

	for _, item := range fileExt {
		if _, found := extCount[item]; found {
			extCount[item]++
		} else {
			extCount[item] = 1
		}
	}

	return extCount
}

func (f *File) ShowExtCount(uniqueExt map[string]int) {
	fmt.Printf("+------------+-------+\n")
	fmt.Printf("| %-10s | %-5s |\n", "Extension", "Count")
	fmt.Printf("+------------+-------+\n")

	for ext, count := range uniqueExt {
		fmt.Printf("| %-10s | %5d |\n", ext, count)
	}

	fmt.Printf("+------------+-------+\n")
}
