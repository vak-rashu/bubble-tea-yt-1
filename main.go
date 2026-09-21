package main

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
)

type model struct {
	choiceList []string
	cursor     int
	selected   map[int]struct{}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "j":
			if m.cursor > 0 {
				m.cursor--
			}
		case "k":
			if m.cursor < len(m.choiceList)-1 {
				m.cursor++
			}
		case "space":
			if _, ok := m.selected[m.cursor]; ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}

	return m, nil
}

func (m model) View() tea.View {
	s := "Choose from the list:\n\n"

	for i, v := range m.choiceList {
		cursor := ""
		check := ""

		if m.cursor == i {
			cursor = ">"
		}

		if _, ok := m.selected[i]; ok {
			check = "+"
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, check, v)
	}
	return tea.NewView(s)
}

func main() {
	p := tea.NewProgram(model{
		choiceList: []string{"I have to do laundry",
			"I have to do house chores at 9pm.",
			"I have to go to the garden",
		},

		selected: make(map[int]struct{}),
	})

	if _, err := p.Run(); err != nil {
		panic(err)
	}
}
