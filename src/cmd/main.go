package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/someprofessional/smallfilesearch/src/reader"
)

type Model struct {
	State       AppState
	Files       []string
	Interactive bool
	Loading     bool
	SearchQuery string
	Results     string
}

type Msg string

type AppState int

const (
	LoadingState AppState = iota
	LoadedState
	InteractiveState
	seeResultsState

	FileLoadedMsg Msg = "file_loaded"
)

func main() {

	styledTitle := titleStyle.Render("Hello smallfilesearch !")

	fmt.Println(styledTitle)
	model := Model{
		State:   LoadingState,
		Loading: true,
	}
	p := tea.NewProgram(model)
	if err := p.Start(); err != nil {
		fmt.Println("Error:", err)
	}
}

func (m Model) Init() tea.Cmd {
	if m.State == LoadingState {
		//TODO: add a control to see if there are aldready files in the output
		// if there are continue to next part
		return reader.LoadFilesCmd()
	}

	return nil
}
