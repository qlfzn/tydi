package organiser

import (
	"fmt"
	"os"
	"path/filepath"
)

// Create struct for FileGroup
type FileGroup struct {
	dirPath string
	groupBy string
}

// Create new file group instance
func NewFileGroup(dirPath string, groupBy string) (*FileGroup, error) {
	cleanedPath := filepath.Clean(dirPath)

	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("path does not exist: %s", err)
	}

	return &FileGroup{
		dirPath: cleanedPath,
		groupBy: groupBy,
	}, nil
}

// Organise directory entries
// Returns defined folder groups and their destination path
func (f *FileGroup) Organise() (*OrganiseResult, error) {
	dirEntries, err := f.getAllFilesInDir(f.dirPath)
	if err != nil {
		return nil, err
	}

	groupResult := f.groupFiles(dirEntries, f.groupBy)
	folderGroup := f.getFolderPath(f.dirPath, groupResult)

	return &OrganiseResult{
		Dir:       f.dirPath,
		GroupBy:   f.groupBy,
		Groups:    groupResult,
		DestPaths: folderGroup,
	}, nil
}

func (f *FileGroup) getAllFilesInDir(path string) ([]os.DirEntry, error) {
	dir, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	return dir, err
}

// Set destination path for grouped files
// Returns all defined destination paths
func (f *FileGroup) getFolderPath(pathPrefix string, fileGroups map[string][]os.DirEntry) []string {
	var keyGroups []string

	for key := range fileGroups {
		keyGroups = append(keyGroups, pathPrefix+key+"/")
	}

	return keyGroups
}
