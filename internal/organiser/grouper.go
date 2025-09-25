package organiser

import (
	"os"
	"path/filepath"
	"strings"
)

// Register functions
type GroupFunc func(file []os.DirEntry) map[string][]os.DirEntry

var groups = map[string]GroupFunc{
	"extension": func(files []os.DirEntry) map[string][]os.DirEntry {
		return new(FileGroup).groupByExtension(files)
	},
	"prefix": func(files []os.DirEntry) map[string][]os.DirEntry {
		return new(FileGroup).groupByPrefix(files)
	},
}

// Main functino for group file
func (f *FileGroup) groupFiles(fileEntries []os.DirEntry, groupMethod string) map[string][]os.DirEntry {
	fn := groups[groupMethod]
	return fn(fileEntries)
}

// Group files by unique extensions
// Returns all extensions and their file names
func (f *FileGroup) groupByExtension(files []os.DirEntry) map[string][]os.DirEntry {
	extResult := make(map[string][]os.DirEntry)

	for _, f := range files {
		if f.IsDir() {
			continue
		}

		ext := strings.ToUpper(strings.TrimPrefix(filepath.Ext(f.Name()), "."))
		extResult[ext] = append(extResult[ext], f)

	}

	return extResult
}

// Group files by unique prefix (string before underscore in filename is considered prefix)
// Returns all prefixes and their file names
func (f *FileGroup) groupByPrefix(files []os.DirEntry) map[string][]os.DirEntry {
	prefixResult := make(map[string][]os.DirEntry)

	for _, f := range files {
		if f.IsDir() {
			continue
		}

		index := strings.IndexAny(f.Name(), "_- ")
		if index == -1 {
			continue
		}

		prefix := f.Name()[:index]
		prefixResult[prefix] = append(prefixResult[prefix], f)
	}

	return prefixResult
}
