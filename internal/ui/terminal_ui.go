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
	Folders []string
}

func (t *TerminalUI) PrintHeader() {
	fmt.Println(header)
}

func (t *TerminalUI) PrintBanner() {
	fmt.Printf("  ► Organising directory: %s\n", t.Dir)
	fmt.Printf("  ◨ Grouping by: %s\n", t.Group)
}

func (t *TerminalUI) PrintGroupTable(unique map[string][]os.DirEntry) {
	headers := []string{"Group", "Count"}

	table := tablewriter.NewWriter(os.Stdout)
	table.Header(headers)

	keys := make([]string, 0, len(unique))
	for k := range unique {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, key := range keys {
		row := []string{key, fmt.Sprintf("%d", len(unique[key]))}
		table.Append(row)
	}

	table.Render()
}

func (t *TerminalUI) PrintDestinationPath(folderGroups []string) {
	for _, folder := range t.Folders {
		fmt.Printf("\n # Destination path: %s", folder)
	}
}
