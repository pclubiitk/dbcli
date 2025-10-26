package UI

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

func UpdateDumpOption(m Model, msg tea.Msg) Model {
	switch msg := msg.(type) {
	case tea.KeyMsg:
 		if msg.String() == "enter" {
			m.WantDump = true
			m.DumpPath = "./dump.sql" // default placeholder
			m.Step++
		}
	}
	return m
}

func ViewDumpOption(m Model) string {
	return fmt.Sprintf("Dump option:\n%s\nPress Enter to continue.", "./dump.sql")
}
