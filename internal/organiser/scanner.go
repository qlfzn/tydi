package organiser

import (
	"fmt"
	"os"
	"sort"

	"github.com/olekukonko/tablewriter"
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
	headers := []string{"Group", "Count"}

	table := tablewriter.NewWriter(os.Stdout)
	table.Header(headers)

	// Optional: sort keys for consistent output
	keys := make([]string, 0, len(unique))
	for k := range unique {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, key := range keys {
		row := []string{key, fmt.Sprintf("%d", unique[key])}
		table.Append(row)
	}

	table.Render()
}
