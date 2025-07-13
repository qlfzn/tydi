package organiser

// function type - take a file, return a string
type GrouperFunc func(file File) string

func CheckGroups(files []File, groupFn GrouperFunc) map[string][]File {
	groups := make(map[string][]File)

	for _, file := range files {
		key := groupFn(file)
		groups[key] = append(groups[key], file)
	}

	return groups
}
