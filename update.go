package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/someprofessional/smallfilessearch/src/reader"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case reader.FileLoadedMsg:
		m.State = LoadedState
		m.Loading = false
		m.Files = msg.Files
		return m, nil

	case tea.KeyMsg:
		keyMsg := msg
		switch keyMsg.String() {
		case "ctrl+c": // Example: quit when 'q' is pressed
			return m, tea.Quit
		case "enter":
			if m.State == InteractiveState {
				m.State = seeResultsState
				return m, nil
			}
			m.State = InteractiveState
			return m, nil
		case "backspace":
			if len(m.SearchQuery) > 0 {
				m.SearchQuery = m.SearchQuery[:len(m.SearchQuery)-1] // Remove last character
			}
			return m, nil
		default:
			if m.State == InteractiveState {
				m.SearchQuery += keyMsg.String()
				return m, nil
			}
		}
	}
	return m, nil
}
