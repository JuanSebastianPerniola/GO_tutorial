// Here we will implement a simple Snake game in Go using the Buble Tea library.

package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type position struct {
	x, y int
}
type model struct {
	snake  []position
	food   position
	width  int
	height int
	score  int
}

func initialModel() model {
	return model{
		snake:  []position{{x: 10, y: 5}, {x: 10, y: 6}, {x: 10, y: 7}},
		food:   position{x: 5, y: 5},
		width:  20,
		height: 10,
		score:  0,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}

	}
	return m, nil
}
func (m model) View() string {
	return fmt.Sprintf("Puntaje: %d\n", m.score)
}
func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alerta, hubo un error: %v", err)
		os.Exit(1)
	}
}
