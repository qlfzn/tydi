package organiser

import (
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
