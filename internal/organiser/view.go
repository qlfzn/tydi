package organiser

import "os"

type OrganiseResult struct {
	Dir       string
	GroupBy   string
	Groups    map[string][]os.DirEntry
	DestPaths []string
}
