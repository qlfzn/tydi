package organiser

import (
	"fmt"
	"os"
)

// Create struct for File
type File struct {
	DirPath string
	GroupBy string
}

func (f *File) GetAllFilesInDir(path string) ([]os.DirEntry, error) {
	dir, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	return dir, err
}

func (f *File) ShowUniqueCount(unique map[string]int) {
	fmt.Printf("+------------+-------+\n")
	fmt.Printf("| %-10s | %-5s |\n", "Pattern", "Count")
	fmt.Printf("+------------+-------+\n")

	for ext, count := range unique {
		fmt.Printf("| %-10s | %5d |\n", ext, count)
	}

	fmt.Printf("+------------+-------+\n")
}
