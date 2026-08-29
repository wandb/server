package pkgm

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	doneStyle = lipgloss.NewStyle().Margin(1, 2)
	checkMark = lipgloss.NewStyle().Foreground(lipgloss.Color("42")).SetString("✓")
)

type CallbackFunc func(string)

func New(packages []string, callback CallbackFunc) *tea.Program {
	p := progress.New(
		progress.WithDefaultGradient(),
		progress.WithWidth(40),
		progress.WithoutPercentage(),
	)
	s := spinner.New()
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("63"))
	return tea.NewProgram(&model{
		packages: packages,
		callback: callback,

		spinner:  s,
		progress: p,
	})
}

type model struct {
	packages []string
	callback func(string)

	index    int
	width    int
	height   int
	spinner  spinner.Model
	progress progress.Model
	done     bool
}

type taskComplete struct {
	name string
}

func (m model) Init() tea.Cmd {
	execute := func(pkg string) tea.Cmd {
		return func() tea.Msg {
			if m.callback != nil {
				m.callback(pkg)
			}
			return taskComplete{name: pkg}
		}
	}

	var cmds []tea.Cmd
	for _, pkg := range m.packages {
		cmds = append(cmds, execute(pkg))
	}

	return tea.Batch(cmds...)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" || msg.String() == "esc" || msg.String() == "q" {
			return m, tea.Quit
		}
	case taskComplete:
		m.index++
		progressCmd := m.progress.SetPercent(float64(m.index) / float64(len(m.packages)))
		if m.index == len(m.packages) {
			m.done = true
			return m, tea.Quit
		}
		return m, tea.Batch(
			progressCmd,
			tea.Printf("%s %s", checkMark, msg.name),
		)
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	case progress.FrameMsg:
		newModel, cmd := m.progress.Update(msg)
		if newModel, ok := newModel.(progress.Model); ok {
			m.progress = newModel
		}
		return m, cmd
	}
	return m, nil
}

func (m model) View() string {
	n := len(m.packages)
	w := lipgloss.Width(fmt.Sprintf("%d", n))

	if m.done {
		return doneStyle.Render(fmt.Sprintf("Done! Installed %d packages.\n", n))
	}

	pkgCount := fmt.Sprintf(" %*d/%*d", w, m.index, w, n)
	prog := m.progress.View()
	gap := strings.Repeat(" ", 2)

	return pkgCount + gap + prog
}
