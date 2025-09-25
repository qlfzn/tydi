package main

import (
	"fmt"
	"log"
	"time"

	"github.com/qlfzn/tydi/cmd"
	"github.com/qlfzn/tydi/internal/organiser"
	"github.com/qlfzn/tydi/internal/ui"
)

func main() {
	cliConf := cmd.ParseCLIArgs()

	f, err := organiser.NewFileGroup(cliConf.InputPath, cliConf.GroupBy)
	if err != nil {
		log.Fatal(err)
	}

	folders, err := f.Organise()
	if err != nil {
		log.Fatal(err)
	}

	tui := ui.TerminalUI{}
	tui.CreateUI(folders)

	fmt.Println("\n  Starting moving files")
	startTime := time.Now()

	err = f.MoveFiles(folders.Dir, folders.Groups)
	if err != nil {
		log.Fatal(err)
	}

	elapsed := time.Since(startTime)
	fmt.Printf("\n  File organiser took %s\n", elapsed)
}