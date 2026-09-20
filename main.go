package main

import (
	"fmt"
	"os"
	"strings"

	tea "charm.land/bubbletea/v2"
)

type model struct {
	choices  []string
	cursor   int
	selected map[int]struct{}
}

func initialModel() model {
	return model{
		choices: []string{
			"I have to do house chores.",
			"I have to get groceries",
			"I have to do laundry",
		},
		selected: make(map[int]struct{}),
	}
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
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter", "space":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}
	return m, nil
}

func (m model) View() tea.View {

	var s strings.Builder
	s.WriteString("Choose from the list:\n\n")
	for i, c := range m.choices {
		cur := " "
		if m.cursor == i {
			cur = ">"
		}

		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "x"
		}

		fmt.Fprintf(&s, "%s [%s] %s\n", cur, checked, c)
	}

	s.WriteString("\nPress q to quit.\n")

	return tea.NewView(s.String())
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		os.Exit(1)
	}
}
