package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/pclubiitk/dbcli/UI"
)

func main() {
	model := UI.Model{
		Step:     UI.StepSourceCred,
		CredKeys: []string{"dbVendor", "host", "port", "user", "password", "dbname"},
		IsSource: true,
	}

	p := tea.NewProgram(model)
	if err := p.Start(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
