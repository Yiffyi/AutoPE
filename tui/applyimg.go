package tui

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/rs/zerolog/log"
	"github.com/yiffyi/autope/native"
)

type TUIApplyImage struct {
	progress progress.Model
	spinner  spinner.Model

	cs              *native.WIMMessageChannelList
	curFileBaseName string
	curETA          time.Duration
	curProgress     float64

	err error
}

type wimUpdateMsg struct {
	progress bool
	quit     bool
}

func CreateTUIAppltImage(channels *native.WIMMessageChannelList) *TUIApplyImage {
	s := spinner.New()
	s.Spinner = spinner.Line
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	return &TUIApplyImage{cs: channels, spinner: s, progress: progress.New(progress.WithDefaultGradient())}
}

func (m *TUIApplyImage) receiveUpdateCmd() tea.Cmd {
	return func() tea.Msg {
		for {
			select {
			case p1 := <-m.cs.Process:
				// fmt.Println("Process", p1)
				m.curFileBaseName = filepath.Base(p1)
				return wimUpdateMsg{progress: false, quit: false}
			case p2 := <-m.cs.Progress:
				// fmt.Println("Progress", p2)
				m.curProgress = float64(p2) / 100
				return wimUpdateMsg{progress: true, quit: false}
			case p3 := <-m.cs.ETA:
				// tea.Println("WIM ETA", p3)
				m.curETA = time.Duration(p3) * time.Millisecond
				return wimUpdateMsg{progress: false, quit: false}
			case <-m.cs.Others:
				// there are many undocumented WIM Messages, so we only pick what we use
				// log.Info().Str("msgId", p4.String()).Msg("received other WIM Message")
			case p5 := <-m.cs.Quit:
				log.Info().Err(p5).Msg("WIM Quit")
				m.err = p5
				return wimUpdateMsg{progress: false, quit: true}
			}
		}
	}
}

func (m *TUIApplyImage) Init() tea.Cmd {
	return tea.Batch(m.receiveUpdateCmd(), m.spinner.Tick)
}

func (m *TUIApplyImage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return m, tea.Quit
		default:
			return m, nil
		}

	// from receiveUpdateCmd
	case wimUpdateMsg:
		if msg.progress {
			return m, tea.Batch(m.receiveUpdateCmd(), m.progress.SetPercent(m.curProgress))
		} else if msg.quit {
			return m, tea.Quit
		} else {
			return m, m.receiveUpdateCmd()
		}

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

func (m *TUIApplyImage) View() string {
	if m.err != nil {
		return m.err.Error()
	}
	str := fmt.Sprintf("\n\n  %s %s\n  %s ETA: %s\npress q to quit\n\n", m.spinner.View(), m.curFileBaseName, m.progress.View(), m.curETA)

	return str
}
