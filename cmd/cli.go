package cmd

import (
	"flag"
	"log"
	"slices"
)

type CLIConfig struct {
	InputPath string
	GroupBy   string
}

func ParseCLIArgs() *CLIConfig {
	var groupPatterns = []string{"extension", "prefix"}

	path := flag.String("path", ".", "Directory path to organise")
	groupBy := flag.String("groupby", "extension", "Grouping patterns: extension, prefix")

	flag.Parse()

	if !slices.Contains(groupPatterns, *groupBy) {
		log.Fatalf("pattern not recognised: %s", *groupBy)
	}

	return &CLIConfig{
		InputPath: *path,
		GroupBy:   *groupBy,
	}
}
