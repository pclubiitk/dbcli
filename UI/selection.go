package UI

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)
//TODO::using m.Source and m.Dest to 
// show tables and add selection rule


func UpdateSelection(m Model, msg tea.Msg) Model {
	
	switch m.Step {
	case StepSelectSourceTable:
		if len(m.SourceTables) > 0 {
			m.SelectedSourceTbl = m.SourceTables[0]
			m.Step++
		}
	case StepSelectSourceColumns:
		m.SelectedSourceCols = m.SourceColumns
		m.Step++
	case StepSelectDestTable:
		if len(m.DestTables) > 0 {
			m.SelectedDestTbl = m.DestTables[0]
			m.Step++
		}
	case StepSelectDestColumns:
		m.SelectedDestCols = m.DestColumns
		m.Step++
	}
	return m
}

func ViewSelection(m Model) string {
	switch m.Step {
	case StepSelectSourceTable:
		return fmt.Sprintf("Select source table:\n%s\nPress Enter to continue.", m.SourceTables)
	case StepSelectSourceColumns:
		return fmt.Sprintf("Select source columns:\n%s\nPress Enter to continue.", m.SourceColumns)
	case StepSelectDestTable:
		return fmt.Sprintf("Select destination table:\n%s\nPress Enter to continue.", m.DestTables)
	case StepSelectDestColumns:
		return fmt.Sprintf("Select destination columns:\n%s\nPress Enter to continue.", m.DestColumns)
	}
	return "Selection step"
}
