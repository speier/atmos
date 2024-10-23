package utils

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	message  string
	choices  []string
	cursor   int
	selected bool
}

func initialModel(message string) model {
	return model{
		message:  message,
		choices:  []string{"Yes", "No"},
		cursor:   0,
		selected: false,
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
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter", " ":
			m.selected = true
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	if m.selected {
		return fmt.Sprintf("You chose %s\n", m.choices[m.cursor])
	}

	s := fmt.Sprintf("%s\n\n", m.message)
	for i, choice := range m.choices {
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor
		}
		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}
	s += "\nPress q to quit.\n"
	return s
}

// Confirm displays a confirmation prompt with a custom message and returns the user's choice
func Confirm(message string) (bool, error) {
	p := tea.NewProgram(initialModel(message))
	m, err := p.Run()
	if err != nil {
		return false, fmt.Errorf("error starting confirm program: %w", err)
	}

	return m.(model).choices[m.(model).cursor] == "Yes", nil
}