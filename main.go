package main

import (
	"fmt"
	"log"

	"github.com/qlfzn/tydi/cmd"
	"github.com/qlfzn/tydi/internal/organiser"
)

func main() {
	cliConf := cmd.ParseCLIArgs()

	f := organiser.File{
		DirPath: cliConf.InputPath,
		GroupBy: cliConf.GroupBy,
	}

	dirEntries, err := f.GetAllFilesInDir(f.DirPath)
	if err != nil {
		log.Fatal(err)
	}

	var strategy organiser.GroupingStrategy

	switch cliConf.GroupBy {
	case "extension":
		strategy = f.GroupByExtension
	case "prefix":
		strategy = f.GroupByPrefix
	default:
		log.Fatalf("unknown grouping strategy: %s", cliConf.GroupBy)
	}

	groupResult := strategy(dirEntries)

	fmt.Println("Welcome, user! Let's organise with tydi")
	f.ShowUniqueCount(groupResult)
}
