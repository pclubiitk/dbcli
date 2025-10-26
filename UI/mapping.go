package UI

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

//TODO::Put logic to map selected coloumns of both source
//and destination database

func UpdateMapping(m Model, msg tea.Msg) Model {
	if m.ColumnMapping == nil {
		m.ColumnMapping = make(map[string]string)
	}
	if len(m.SelectedSourceCols) > 0 && len(m.SelectedDestCols) > 0 {
		for i, col := range m.SelectedSourceCols {
			if i < len(m.SelectedDestCols) {
				m.ColumnMapping[col] = m.SelectedDestCols[i]
			}
		}
	}
	m.Step++
	return m
}

func ViewMapping(m Model) string {
	out := "Column Mapping:\n"
	for src, dest := range m.ColumnMapping {
		out += fmt.Sprintf("%s -> %s\n", src, dest)
	}
	out += "Press Enter to continue."
	return out
}
