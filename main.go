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
	cliConf := cmd.ParseCLIArgs()

	f, err := organiser.NewFileGroup(cliConf.InputPath, cliConf.GroupBy)
	if err != nil {
		log.Fatal(err)
	}

	folders, err := f.Organise()
	if err != nil {
		log.Fatal(err)
	}

	// tui := ui.TerminalUI{
	// 	Dir:     cliConf.InputPath,
	// 	Group:   cliConf.GroupBy,
	// 	Folders: folders,
	// }

	// tui.PrintHeader()
	// tui.PrintBanner()
	// tui.PrintGroupTable(groupResult)
	// tui.PrintDestinationPath(folders)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\n\n Proceed with moving files? (y/N): ")
	confirm, _ := reader.ReadString('\n')
	confirm = strings.TrimSpace(strings.ToLower(confirm))

	if confirm != "y" && confirm != "yes" {
		fmt.Println("\n No files were moved.")
		return
	}

	fmt.Println("\n  Starting moving files")
	startTime := time.Now()

	err = f.MoveFiles(f.DirPath, groupResult)
	if err != nil {
		log.Fatal(err)
	}

	elapsed := time.Since(startTime)
	fmt.Printf("\n  File organiser took %s\n", elapsed)
}