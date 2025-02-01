package main

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/pelletier/go-toml/v2"
	"github.com/yiffyi/autope"
	"github.com/yiffyi/autope/native"
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

type fileNameMsg string
type progressMsg int
type etaMsg time.Duration
type errMsg error

type model struct {
	progress progress.Model
	spinner  spinner.Model

	cs          *native.WIMMessageChannelList
	curFileName string
	curETA      time.Duration

	quitting bool
	err      error
}

func initialModel() model {
	s := spinner.New()
	s.Spinner = spinner.Line
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	cs := &native.WIMMessageChannelList{
		Process:  make(chan string, 16),
		Progress: make(chan int, 16),
		ETA:      make(chan uint64, 16),
		Others:   make(chan native.WimMessageId, 16),
		Quit:     make(chan error, 16),
	}

	go native.WIMApplyImageByPath(`D:\sources\install.wim`, 1, `X:\`, cs)
	return model{cs: cs, spinner: s, progress: progress.New(progress.WithDefaultGradient())}
}

func (m *model) receiveUpdateCmd() tea.Cmd {
	return func() tea.Msg {
		for {
			select {
			case p1 := <-m.cs.Process:
				// fmt.Println("Process", p1)
				return fileNameMsg(p1)
			case p2 := <-m.cs.Progress:
				// fmt.Println("Progress", p2)
				return progressMsg(p2)
			case p3 := <-m.cs.ETA:
				// tea.Println("WIM ETA", p3)
				return etaMsg(time.Duration(p3) * time.Millisecond)
			case p4 := <-m.cs.Others:
				tea.Println("WIM Message:", p4)
			case p5 := <-m.cs.Quit:
				// tea.Println("WIM Quit", p5)
				return errMsg(p5)
			}
		}
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(m.receiveUpdateCmd(), m.spinner.Tick)
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

	// from receiveUpdateCmd
	case fileNameMsg:
		m.curFileName = string(msg)
		return m, m.receiveUpdateCmd()
	case progressMsg:
		cmd := m.progress.SetPercent(float64(msg))
		return m, tea.Batch(cmd, m.receiveUpdateCmd())
	case etaMsg:
		m.curETA = time.Duration(msg)
		return m, m.receiveUpdateCmd()
	case errMsg:
		m.err = msg
		return m, tea.Quit

	// from spinner.Tick
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd

	// from progress.Update
	case progress.FrameMsg:
		progressModel, cmd := m.progress.Update(msg)
		m.progress = progressModel.(progress.Model)
		return m, cmd

	default:
		return m, nil
	}
}

func (m model) View() string {
	if m.err != nil {
		return m.err.Error()
	}
	baseName := filepath.Base(m.curFileName)
	str := fmt.Sprintf("\n\n  %s %s\n  %s %f\npress q to quit\n\n", m.spinner.View(), baseName, m.progress.View(), m.curETA.Seconds())
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
