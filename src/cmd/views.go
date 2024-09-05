package main

import (
	"fmt"
	"strings"
)

func (m Model) View() string {
	switch m.State {
	case LoadingState:
		return "Loading files...\n"
	case LoadedState:
		return m.renderFilesView()
	case InteractiveState:
		return m.renderInteractiveView()
	case seeResultsState:
		return m.RenderSeeResultsView()
	default:
		return "Unkown state"
	}

}

func (m Model) renderFilesView() string {
	var sb strings.Builder
	sb.WriteString("Current files in ./input/: ")
	for _, file := range m.Files {
		sb.WriteString(fmt.Sprintf(" %s/", file))
	}
	sb.WriteString("\n\nPress Enter to pursue\n")

	content := sb.String()

	styledRenderFile := filesStyle.Render(content)

	return styledRenderFile
}

func (m Model) renderInteractiveView() string {

	// TODO : find a way to make this interactive
	// maybe by using the view thing -- before that make it searchable
	//searchParams := [2]string{"word", "sentences"}
	//activeParam := 0
	//wordArround := 0

	//fmt.Printf("Choose your search option : %s", searchParams)

	//if activeParam == 1 {
	//	fmt.Println("how many words to you want arround the searched term ?")
	//	fmt.Scan(&wordArround)
	//}

	return fmt.Sprintf("Searching for: %s\n", m.SearchQuery)
}

func (m Model) RenderSeeResultsView() string {
	return fmt.Sprintln("Here are your results")
}
