package main

import (
	"fmt"
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var style_select = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FF3333"))

type Model struct {
	board      *Board
	next_turn  Team
	turn_count int

	selx, sely int
	errmsg     string
	outcome    string
}

func initialModel() Model {
	b := NewBoard()
	return Model{
		board:      &b,
		next_turn:  Team_X,
		turn_count: 0,
		selx:       1,
		sely:       1,
	}
}

func (m Model) Init() tea.Cmd {
	return tea.SetWindowTitle("Tic Tac Toe")
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "up":
			if m.sely > 0 {
				m.sely--
			}
		case "down":
			if m.sely < 2 {
				m.sely++
			}
		case "left":
			if m.selx > 0 {
				m.selx--
			}
		case "right":
			if m.selx < 2 {
				m.selx++
			}
		case " ", "enter":
			if m.outcome == "" {
				m.play()
			}
		}
	}

	return m, nil
}

func (m *Model) play() {
	if m.board.get(m.selx, m.sely) != Team_None {
		m.errmsg = fmt.Sprintf(
			"Position %d, %d is already taken by %s\n",
			m.selx+1,
			m.sely+1,
			m.board.get(m.selx, m.sely).Name(),
		)
		return
	}

	m.errmsg = ""
	m.board.set(m.selx, m.sely, m.next_turn)
	m.turn_count++
	switch m.next_turn {
	case Team_O:
		m.next_turn = Team_X
	case Team_X:
		m.next_turn = Team_O
	}

	w := m.board.getWinner()
	if w != Team_None {
		m.outcome = fmt.Sprintf("%s won!", w.Name())
	} else if m.turn_count >= 9 {
		m.outcome = "Draw!"
	}
}

func (m Model) View() string {
	s := fmt.Sprintf("Turn %d: %s\n", m.turn_count, m.next_turn.Name())
	for i := 0; i < 9; i++ {
		if i == m.selx+3*m.sely && m.outcome == "" {
			s += style_select.Render(string(m.board.data[i]))
		} else {
			s += string(m.board.data[i])
		}
		if i%3 == 2 {
			s += "\n"
		}
	}
	s += m.errmsg + "\n"
	s += m.outcome + "\n"
	return s
}

func main() {
	p := tea.NewProgram(initialModel())
	_, err := p.Run()
	if err != nil {
		log.Fatalf("Alas, there's been an error: %v", err)
	}
}
