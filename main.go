package main

import (
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	state       state
	choices     []string
	subChoices  map[string][]string
	cursor      int
	selected    int
	subSelected int
	blink       bool
	message     string
}

type state int

const (
	mainChoices state = iota
	subChoices
)

func initialModel() model {
	return model{
		state:   mainChoices,
		choices: []string{"Linked List", "Choice 2", "Choice 3", "Choice 4", "Choice 5"},
		subChoices: map[string][]string{
			"Linked List": {"Insert at End", "Insert at Beginning", "Insert in Between"},
			"Choice 2":    {"Sub-choice 2.1", "Sub-choice 2.2"},
			"Choice 3":    {"Sub-choice 3.1", "Sub-choice 3.2"},
			"Choice 4":    {"Sub-choice 4.1", "Sub-choice 4.2"},
			"Choice 5":    {"Sub-choice 5.1", "Sub-choice 5.2"},
		},
		cursor:      0,
		selected:    -1,
		subSelected: -1,
		blink:       true,
		message:     "",
	}
}

func (m model) Init() tea.Cmd {
	return tea.Tick(time.Millisecond*500, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

type tickMsg time.Time

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.state == mainChoices {
				if m.cursor > 0 {
					m.cursor--
				}
			} else if m.state == subChoices {
				if m.cursor > 0 {
					m.cursor--
				}
			}
		case "down", "j":
			if m.state == mainChoices {
				if m.cursor < len(m.choices)-1 {
					m.cursor++
				}
			} else if m.state == subChoices {
				subChoiceList := m.subChoices[m.choices[m.selected]]
				if m.cursor < len(subChoiceList)-1 {
					m.cursor++
				}
			}
		case "enter":
			if m.state == mainChoices {
				m.selected = m.cursor
				m.cursor = 0
				m.state = subChoices
			} else if m.state == subChoices {
				subChoiceList := m.subChoices[m.choices[m.selected]]
				m.subSelected = m.cursor
				m.message = fmt.Sprintf("You selected: %s > %s", m.choices[m.selected], subChoiceList[m.subSelected])
				m.state = mainChoices
				m.cursor = 0
				m.selected = -1
				m.subSelected = -1
			}
		}
	case tickMsg:
		m.blink = !m.blink
		return m, tea.Tick(time.Millisecond*500, func(t time.Time) tea.Msg {
			return tickMsg(t)
		})
	}

	return m, nil
}

func (m model) View() string {
	var s string
	if m.state == mainChoices {
		s = "Choose a Data structure:\n\n"
		for i, choice := range m.choices {
			cursor := " "
			if m.cursor == i && m.blink {
				cursor = ">"
			}
			if m.selected == i {
				cursor = "[>]"
			}
			line := fmt.Sprintf("%s %s", cursor, choice)
			// if m.selected == i {
			// 	line = lipgloss.NewStyle().Foreground(lipgloss.Color("205")).Render(line)
			// }
			s += line + "\n"
		}
	} else if m.state == subChoices {
		s = fmt.Sprintf("Operations for %s:\n\n", m.choices[m.selected])
		subChoiceList := m.subChoices[m.choices[m.selected]]
		for i, subChoice := range subChoiceList {
			cursor := " "
			if m.cursor == i && m.blink {
				cursor = ">"
			}
			if m.subSelected == i {
				cursor = "[>]"
			}
			line := fmt.Sprintf("%s %s", cursor, subChoice)
			// if m.subSelected == i {
			// 	line = lipgloss.NewStyle().Foreground(lipgloss.Color("205")).Render(line)
			// }
			s += line + "\n"
		}
	}

	if m.message != "" {
		s += "\n" + m.message + "\n"
	}

	s += "\nPress q to quit.\n"
	return s
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
