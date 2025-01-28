package main

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/pelletier/go-toml/v2"
	"github.com/yiffyi/autope"
)

func tryRichOutput() {
	var style1 = lipgloss.NewStyle().
		Bold(true).
		Italic(true).
		Faint(true).
		Blink(true).
		Strikethrough(true).
		Underline(true).
		Reverse(true)

	fmt.Println(style1.Render("Styled Text"))

	var style2 = lipgloss.NewStyle().
		Foreground(lipgloss.ANSIColor(5))
	fmt.Println(style2.Render("ANSI 5"))

	var style3 = lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("63"))
	fmt.Println(style3.Render("I have border"))
}

type errMsg error

type model struct {
	spinner  spinner.Model
	progress progress.Model
	quitting bool
	err      error
}

type tickMsg time.Time

func initialModel() model {
	s := spinner.New()
	s.Spinner = spinner.Line
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	return model{spinner: s, progress: progress.New(progress.WithDefaultGradient())}
}

func tickCmd() tea.Cmd {
	return tea.Tick(time.Second*1, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func (m model) Init() tea.Cmd {
	return tickCmd()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			m.quitting = true
			return m, tea.Quit
		default:
			return m, nil
		}

	case tickMsg:
		if m.progress.Percent() == 1.0 {
			return m, tea.Quit
		}

		// Note that you can also use progress.Model.SetPercent to set the
		// percentage value explicitly, too.
		return m, tea.Batch(tickCmd(), m.progress.IncrPercent(0.25), m.spinner.Tick)

	// FrameMsg is sent when the progress bar wants to animate itself
	case progress.FrameMsg:
		progressModel, cmd := m.progress.Update(msg)
		m.progress = progressModel.(progress.Model)
		return m, cmd

	case errMsg:
		m.err = msg
		return m, nil

	default:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}
}

func (m model) View() string {
	if m.err != nil {
		return m.err.Error()
	}
	str := fmt.Sprintf("\n\n  %s Loading forever\n  %s\npress q to quit\n\n", m.spinner.View(), m.progress.View())
	if m.quitting {
		return str + "\n"
	}
	return str
}

func trySpinner() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Println(err)
	}
}

func main() {

	t, err := toml.Marshal(autope.Playbook{})
	fmt.Println(string(t), err)

	tryRichOutput()
	trySpinner()
}
