package organiser

import (
	"os"
	"path/filepath"
	"strings"
)

// function type for one behaviour, different implementations
type GroupingStrategy func([]os.DirEntry) map[string]int

func (f *File) GroupByExtension(files []os.DirEntry) map[string]int {
	extResult := make(map[string]int)

	for _, f := range files {
		if f.IsDir() {
			continue
		}

		ext := filepath.Ext(f.Name())
		extResult[ext]++
	}

	return extResult
}

func (f *File) GroupByPrefix(files []os.DirEntry) map[string]int {
	prefixResult := make(map[string]int)

	// string before underscore is prefix
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

// TODO: implement group by date added
