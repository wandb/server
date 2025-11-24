package task

import (
	"fmt"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	doneStyle = lipgloss.NewStyle().Margin(1, 2)
)


func New(text string, callback func() error) *tea.Program {
	s := spinner.New()
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("63"))
	s.Spinner = spinner.Dot
	return tea.NewProgram(&model{
		text: text,
		callback: callback,
		spinner: s,
	})
}

type taskComplete struct {}

type model struct {
	text string
	callback func() error
	spinner  spinner.Model
	quitting bool
	done bool
	err error
}

func (m model) Init() tea.Cmd {
	return tea.Batch(
		func() tea.Msg {
			if m.callback != nil {
				m.err = m.callback()
			}
			return taskComplete{}
		}, 
		m.spinner.Tick,
	)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case taskComplete:
		m.done = true
		return m, tea.Quit

	case tea.KeyMsg:
		switch msg.String() {
		case  "ctrl+c":
			m.quitting = true
			return m, tea.Quit
		default:
			return m, nil
		}

	default:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}
}

func (m model) View() string {
	if m.err != nil {
		return doneStyle.Render(fmt.Sprintf("%s exited: %s\n",m.text, m.err))
	}
	if m.quitting {
		return doneStyle.Render(fmt.Sprintf("%s exited.\n", m.text))
	}

	if m.done {
		return doneStyle.Render(fmt.Sprintf("%s completed.\n", m.text))
	}

	return fmt.Sprintf("\n\n  %s %s\n\n", m.spinner.View(), m.text)
}