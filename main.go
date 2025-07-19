package main

import (
	"log"

	"github.com/qlfzn/tydi/cmd"
	"github.com/qlfzn/tydi/internal/organiser"
	"github.com/qlfzn/tydi/internal/ui"
)

func main() {
	// initialise entrypoint
	cliConf := cmd.ParseCLIArgs()

	// create file instance
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

	folderGroup := f.GetFolderPath(f.DirPath, groupResult)

	// initialise UI props
	tui := ui.TerminalUI{
		Dir:     cliConf.InputPath,
		Group:   cliConf.GroupBy,
		DryRun:  true,
		Folders: folderGroup,
	}

	tui.PrintHeader()
	tui.PrintBanner()
	tui.PrintGroupTable(groupResult)
	tui.PrintDestinationPath(folderGroup)
}
