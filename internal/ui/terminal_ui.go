package ui

import (
	"fmt"
	"os"
	"sort"

	"github.com/olekukonko/tablewriter"
)

var header = `

 ██████████                   ██ ░██
    ░██                      ░██    
    ░██    ░██    ░██  ░████████ ░██
    ░██    ░██    ░██ ░██    ░██ ░██
    ░██    ░██    ░██ ░██    ░██ ░██
    ░██    ░██   ░███ ░██   ░███ ░██
    ░██     ░█████░██  ░█████░██ ░██
                  ░██               
            ░███████                
                                    

`

type TerminalUI struct {
	Dir     string
	Group   string
	DryRun  bool
	Folders []string
}

func (t *TerminalUI) PrintHeader() {
	fmt.Println(header)
}

func (t *TerminalUI) PrintBanner() {
	fmt.Printf("  ► Organising directory: %s\n", t.Dir)
	fmt.Printf("  ◨ Grouping by: %s\n", t.Group)

	if t.DryRun == true {
		fmt.Println(" (!) Dry run enabled — no files will be moved")
	}
}

func (t *TerminalUI) PrintGroupTable(unique map[string]int) {
	headers := []string{"Group", "Count"}

	table := tablewriter.NewWriter(os.Stdout)
	table.Header(headers)

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

func (t *TerminalUI) PrintDestinationPath(folderGroups []string) {
	for _, folder := range t.Folders {
		fmt.Printf("\n # Moving files into: %s", folder)
	}
}
