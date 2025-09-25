package ui

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/olekukonko/tablewriter"
	"github.com/qlfzn/tydi/internal/organiser"
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

type TerminalUI struct {}

func (t *TerminalUI) CreateUI(result *organiser.OrganiseResult) {
	t.PrintHeader()
	t.PrintBanner(result.Dir, result.GroupBy)
	t.PrintGroupTable(result.Groups)
	t.PrintDestinationPath(result.DestPaths)
	t.Prompt()
}

func (t *TerminalUI) PrintHeader() {
	fmt.Println(header)
}

func (t *TerminalUI) PrintBanner(dir string, groupBy string) {
	fmt.Printf("  ► Organising directory: %s\n", dir)
	fmt.Printf("  ◨ Grouping by: %s\n", groupBy)
}

func (t *TerminalUI) PrintGroupTable(groups map[string][]os.DirEntry) {
	table := tablewriter.NewWriter(os.Stdout)
	headers := []string{"Group", "Count"}
	table.Header(headers)

	for name, g := range groups {
		table.Append([]string{name, fmt.Sprintf("%d", len(g))})
	}
	table.Render()
}

func (t *TerminalUI) PrintDestinationPath(folderGroups []string) {
	for _, folder := range folderGroups {
		fmt.Printf("\n # Destination path: %s", folder)
	}
}

func (t *TerminalUI) Prompt() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\n\nProceed with moving files? (y/N): ")
	confirm, _ := reader.ReadString('\n')
	confirm = strings.TrimSpace(strings.ToLower(confirm))

	if confirm != "y" && confirm != "yes" {
		fmt.Println("\nNo files were moved.")
		os.Exit(1)
	}
}