package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/someprofessional/smallfilessearch/src/reader"
)

type Model struct {
	State       string
	Files       []string
	Interactive bool
	Loading     bool
	SearchQuery string
}

type Msg string

const (
	LoadingState     = "loading"
	LoadedState      = "loaded"
	InteractiveState = "interactive"
	seeResultsState  = "results"

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
		return reader.LoadFilesCmd()
	}

	return nil
}
