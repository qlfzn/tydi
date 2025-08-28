package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

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

	groupResult := f.GroupFiles(dirEntries, f.GroupBy)

	folderGroup := f.GetFolderPath(f.DirPath, groupResult)

	// initialise UI props
	tui := ui.TerminalUI{
		Dir:     cliConf.InputPath,
		Group:   cliConf.GroupBy,
		Folders: folderGroup,
	}

	tui.PrintHeader()
	tui.PrintBanner()
	tui.PrintGroupTable(groupResult)
	tui.PrintDestinationPath(folderGroup)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\n\n Proceed with moving files? (y/N): ")
	confirm, _ := reader.ReadString('\n')
	confirm = strings.TrimSpace(strings.ToLower(confirm))

	if confirm != "y" && confirm != "yes" {
		fmt.Println("\n No files were moved.")
		return
	}

	// move files
	fmt.Println("\n  Starting moving files")
	startTime := time.Now()

	err = f.MoveFiles(f.DirPath, groupResult)
	if err != nil {
		log.Fatal(err)
	}

	elapsed := time.Since(startTime)
	fmt.Printf("\n  File organiser took %s\n", elapsed)
}
