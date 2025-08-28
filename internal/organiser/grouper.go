package organiser

import (
	"os"
	"path/filepath"
	"strings"
)

// Register functions
type GroupFunc func(file []os.DirEntry) map[string]int

var groups = map[string]GroupFunc{
	"extension": func(files []os.DirEntry) map[string]int {
		return new(File).groupByExtension(files)
	},
	"prefix": func(files []os.DirEntry) map[string]int {
		return new(File).groupByPrefix(files)
	},
}

// Main functino for group file
func (f *File) GroupFiles(fileEntries []os.DirEntry, groupMethod string) map[string]int {
	fn := groups[groupMethod]
	return fn(fileEntries)
}

// Group files by unique extensions
// Returns all extensions and their counts
func (f *File) groupByExtension(files []os.DirEntry) map[string]int {
	extResult := make(map[string]int)

	for _, f := range files {
		if f.IsDir() {
			continue
		}

		ext := strings.ToUpper(strings.TrimPrefix(filepath.Ext(f.Name()), "."))
		extResult[ext]++
	}

	return extResult
}

// Group files by unique prefix (string before underscore in filename is considered prefix)
// Returns all prefixes and their counts
func (f *File) groupByPrefix(files []os.DirEntry) map[string]int {
	prefixResult := make(map[string]int)

	for _, f := range files {
		if f.IsDir() {
			continue
		}

		index := strings.IndexRune(f.Name(), '_')
		if index == -1 {
			continue
		}

		prefix := f.Name()[:index]
		prefixResult[prefix]++
	}

	return prefixResult
}

// Set destination path for grouped files
// Returns all defined destination paths
func (f *File) GetFolderPath(pathPrefix string, fileGroups map[string]int) []string {
	var keyGroups []string

	for key := range fileGroups {
		keyGroups = append(keyGroups, pathPrefix+key+"/")
	}

	return keyGroups
}
