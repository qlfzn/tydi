package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// Create struct for File
type File struct {
	Path            string
	Name            string
	Extension       string
	Category        string
	Prefix          string
	Date            string
	Error           string
	DestinationPath string
}

func (f *File) getAllFilesInDir(path string) ([]os.DirEntry, error) {
	dir, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	return dir, err
}

func (f *File) getCountUniqueFileTypes(files []os.DirEntry) map[string]int {
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

func (f *File) printExtensionTable(uniqueExt map[string]int) {
	fmt.Printf("+------------+-------+\n")
	fmt.Printf("| %-10s | %-5s |\n", "Extension", "Count")
	fmt.Printf("+------------+-------+\n")

	for ext, count := range uniqueExt {
		fmt.Printf("| %-10s | %5d |\n", ext, count)
	}

	fmt.Printf("+------------+-------+\n")
}

func main() {
	f := File{}

	files, err := f.getAllFilesInDir("/Users/qlfzn/Downloads/")
	if err != nil {
		log.Fatal(err)
	}

	unique := f.getCountUniqueFileTypes(files)

	f.printExtensionTable(unique)

}
