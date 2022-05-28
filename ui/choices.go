package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type choicesModel struct {
	prompt  string
	choices []string
	cursor  int
}

func NewChoicesPrompt(prompt string, choices []string) (int, error) {
	p := tea.NewProgram(choicesModel{
		prompt:  prompt,
		choices: choices,
	})
	m, err := p.StartReturningModel()
	if err != nil {
		return 0, err
	}
	cursor := m.(choicesModel).cursor
	if cursor < 0 {
		return 0, fmt.Errorf("user cancelled")
	}
	return cursor, nil
}

func (m choicesModel) Init() tea.Cmd {
	return nil
}

func (m choicesModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			m.cursor = -1
			return m, tea.Quit
		case tea.KeyDown:
			if m.cursor < (len(m.choices) - 1) {
				m.cursor++
			}
		case tea.KeyUp:
			if m.cursor > 0 {
				m.cursor--
			}
		case tea.KeyEnter:
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m choicesModel) View() string {
	s := m.prompt + "\n\n"

	for i, c := range m.choices {
		cur := " "
		if m.cursor == i {
			cur = lipgloss.NewStyle().Foreground(lipgloss.Color("#7D56F4")).Render("»")
			c = lipgloss.NewStyle().Bold(true).Render(c)
		}
		s += fmt.Sprintf(" %s %s\n", cur, c)
	}

	return s + "\n"
}
