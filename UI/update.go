package UI

import tea "github.com/charmbracelet/bubbletea"

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m.Step {
	case StepSourceCred, StepDestCred:
		m = UpdateDBCred(m, msg)
	case StepSelectSourceTable, StepSelectSourceColumns, StepSelectDestTable, StepSelectDestColumns:
		m = UpdateSelection(m, msg)
	case StepMapping:
		m = UpdateMapping(m, msg)
	case StepDumpOption:
		m = UpdateDumpOption(m, msg)
	}

	// Global key handling
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" || msg.String() == "q" {
			return m, tea.Quit
		}
	}
	

	return m, nil
}
