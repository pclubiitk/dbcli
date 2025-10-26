package UI

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"gorm.io/gorm"
)

const (
	StepSourceCred = iota
	StepDestCred
	StepSelectSourceTable
	StepSelectSourceColumns
	StepSelectDestTable
	StepSelectDestColumns
	StepMapping
	StepDumpOption
	StepMigrationConfirm
)

type Model struct {
	Step int

	// ---- DB CREDENTIAL INPUTS ----
	SourceCred  map[string]string
	DestCred    map[string]string
	Source	    *gorm.DB  //these are the most imp fields
	Dest        *gorm.DB  //they are direct connections to databases
	CredInput   textinput.Model
	CredKeys    []string
	CredIndex   int
	IsSource    bool

	// ---- SOURCE DATABASE ----
	SourceTables      []string
	SelectedSourceTbl string
	SourceTableList   list.Model
	SourceColumns     []string
	SelectedSourceCols []string
	SourceColumnList  list.Model

	// ---- DESTINATION DATABASE ----
	DestTables      []string
	SelectedDestTbl string
	DestTableList   list.Model
	DestColumns     []string
	SelectedDestCols []string
	DestColumnList  list.Model

	// ---- MAPPING ----
	ColumnMapping map[string]string
	CurrentMapIdx int
	MapInput      textinput.Model

	// ---- DUMP OPTION ----
	WantDump    bool
	DumpPath    string
	DumpPathInp textinput.Model

	// ---- MISC ----
	ErrMsg string
}

func (m Model) Init() tea.Cmd {
	return nil
}
