package ui

import tea "github.com/charmbracelet/bubbletea"

type inputModel struct {
	prompt string
	input  string
}

func NewInputPrompt(prompt string) (string, error) {
	p := tea.NewProgram(inputModel{
		prompt: prompt,
	})
	m, err := p.StartReturningModel()
	if err != nil {
		return "", err
	}
	return m.(inputModel).input, nil
}

func (m inputModel) Init() tea.Cmd {
	return nil
}

func (m inputModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyRunes:
			m.input += string(msg.Runes)
		case tea.KeyBackspace:
			if len(m.input) > 0 {
				m.input = m.input[:len(m.input)-1]
			}
		case tea.KeyEnter, tea.KeyBreak:
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m inputModel) View() string {
	return m.prompt + ": " + m.input
}
