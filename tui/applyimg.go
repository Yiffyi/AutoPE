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

	ctx             *native.WIMMessageContext
	curFileBaseName string
	curETA          time.Duration
	curProgress     float64

	Error error
}

type wimUpdateMsg struct {
	progress bool
	quit     bool
}

func CreateTUIAppltImage(ctx *native.WIMMessageContext) *TUIApplyImage {
	s := spinner.New()
	s.Spinner = spinner.Line
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	return &TUIApplyImage{ctx: ctx, spinner: s, progress: progress.New(progress.WithDefaultGradient())}
}

func (m *TUIApplyImage) receiveUpdateCmd() tea.Cmd {
	return func() tea.Msg {
		for {
			select {
			case p1 := <-m.ctx.Process:
				// fmt.Println("Process", p1)
				m.curFileBaseName = filepath.Base(p1)
				return wimUpdateMsg{progress: false, quit: false}
			case p2 := <-m.ctx.Progress:
				// fmt.Println("Progress", p2)
				m.curProgress = float64(p2) / 100
				return wimUpdateMsg{progress: true, quit: false}
			case p3 := <-m.ctx.ETA:
				// tea.Println("WIM ETA", p3)
				m.curETA = time.Duration(p3) * time.Millisecond
				return wimUpdateMsg{progress: false, quit: false}
			case <-m.ctx.Others:
				// there are many undocumented WIM Messages, so we only pick what we use
				// log.Info().Str("msgId", p4.String()).Msg("received other WIM Message")
			case p5 := <-m.ctx.Quit:
				log.Info().Err(p5).Msg("WIM Quit")
				m.Error = p5
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
			m.ctx.Cancel = true
			return m, nil
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
	if m.Error != nil {
		return m.Error.Error()
	}
	str := fmt.Sprintf("\n  %s %s\n  %s [%d / %d] ETA: %s\npress q to quit\n", m.spinner.View(), m.curFileBaseName, m.progress.View(), m.ctx.FileIndex, m.ctx.FileCount, m.curETA)

	return str
}
