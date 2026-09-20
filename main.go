package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
)

type isEven bool

type model struct {
	num int
	is  isEven
}

func initialModel() model {
	return model{num: 5}
}

func isEvenOrNot(num int) tea.Cmd {
	return func() tea.Msg {
		if num%2 == 0 {
			return isEven(true)
		}
		return isEven(false)
	}
}

func (m model) Init() tea.Cmd {
	return isEvenOrNot(m.num)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case isEven:
		m.is = msg

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "k":
			isEvenOrNot(m.num)
		}
	}
	return m, nil
}

func (m model) View() tea.View {

	// _, err := fmt.Scan(&m.num)
	// if err != nil {
	// 	panic(err)
	// }

	s := fmt.Sprintf("Num: %d\nStatus: %v", m.num, m.is)
	return tea.NewView(s)
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		os.Exit(1)
	}
}
