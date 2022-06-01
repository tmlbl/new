package ui

import tea "github.com/charmbracelet/bubbletea"

type yesNoModel struct {
	prompt string
	answer bool
}

func NewYesNoPrompt(prompt string, defaultAnswer bool) (bool, error) {
	p := tea.NewProgram(yesNoModel{
		prompt: prompt,
		answer: defaultAnswer,
	})
	m, err := p.StartReturningModel()
	if err != nil {
		return defaultAnswer, err
	}
	return m.(yesNoModel).answer, nil
}

func (m yesNoModel) Init() tea.Cmd {
	return nil
}

func (m yesNoModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyRunes:
			r := msg.Runes[0]
			if r == 'y' || r == 'Y' {
				m.answer = true
			}
			return m, tea.Quit
		case tea.KeyEnter:
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m yesNoModel) View() string {
	s := m.prompt
	if m.answer {
		s += " [Y/n]"
	} else {
		s += " [y/N]"
	}
	return s
}
