package UI

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/pclubiitk/dbcli/DB"
)

func UpdateDBCred(m Model, msg tea.Msg) Model {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			// Save input
			if m.IsSource {
				if m.SourceCred == nil {
					m.SourceCred = make(map[string]string)
				}
				m.SourceCred[m.CredKeys[m.CredIndex]] = m.CredInput.Value()
			} else {
				if m.DestCred == nil {
					m.DestCred = make(map[string]string)
				}
				m.DestCred[m.CredKeys[m.CredIndex]] = m.CredInput.Value()
			}

			// Reset input
			m.CredInput.SetValue("")
			m.CredIndex++
			if m.CredIndex >= len(m.CredKeys) {
				m.CredIndex = 0
				m.Step++
				if(m.IsSource){
					loweredVendor := strings.ToLower(m.SourceCred["dbVendor"])
					switch loweredVendor{
					case "oracle":
						DB.ConnectOracle( m.SourceCred["host"],  m.SourceCred["port"],  m.SourceCred["password"], m.SourceCred["dbname"], m.SourceCred["user"] )
						m.Source=&DB.SQLWrapper{DB:DB.OracleDB}
						break
					case "mysql":
						DB.ConnectMySQL( m.SourceCred["host"],  m.SourceCred["port"],  m.SourceCred["password"], m.SourceCred["dbname"], m.SourceCred["user"] )
						m.Source=&DB.GormWrapper{DB:DB.MySQLDB}
						break
					}
				}else{
					loweredVendor := strings.ToLower(m.DestCred["dbVendor"])
					switch loweredVendor{
					case "oracle":
						DB.ConnectOracle( m.DestCred["host"],  m.DestCred["port"],  m.DestCred["password"], m.DestCred["dbname"], m.DestCred["user"] )
						m.Dest=&DB.SQLWrapper{DB:DB.OracleDB}
						break
					case "mysql":
						DB.ConnectMySQL( m.DestCred["host"],  m.DestCred["port"],  m.DestCred["password"], m.DestCred["dbname"], m.DestCred["user"] )
						m.Dest=&DB.GormWrapper{DB:DB.MySQLDB}
						break
					}
				}
				m.IsSource = !m.IsSource
			}
		case "backspace":
			if len(m.CredInput.Value()) > 0 {
				m.CredInput.SetValue(m.CredInput.Value()[:len(m.CredInput.Value())-1])
			}
		default:
			m.CredInput.SetValue(m.CredInput.Value() + msg.String())
		}
	}
	return m
}

func ViewDBCred(m Model) string {
	var sb strings.Builder
	dbType := "Source"
	if !m.IsSource {
		dbType = "Destination"
	}

	sb.WriteString(fmt.Sprintf("Step: Enter %s DB Credentials\n", dbType))
	if m.CredIndex < len(m.CredKeys) {
		key := m.CredKeys[m.CredIndex]
		sb.WriteString(fmt.Sprintf("%s: %s\n", key, m.CredInput.Value()))
		sb.WriteString("Type your input and press Enter to continue.\n")
	} else {
		sb.WriteString("All credentials entered. Press Enter to continue.\n")
	}

	if m.ErrMsg != "" {
		sb.WriteString(fmt.Sprintf("Error: %s\n", m.ErrMsg))
	}


	sb.WriteString("\nPress q to quit.\n")
	return sb.String()
}
