package autope

import (
	"os"
	"path/filepath"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func SetupDefaultLogger() {
	zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
		return filepath.Base(file) + ":" + strconv.Itoa(line)
	}

	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Caller().Logger()
	log.Logger = logger
}

type logToBubbleTea struct {
	p *tea.Program
}

func (l *logToBubbleTea) Write(p []byte) (n int, err error) {
	n = len(p)
	if n > 0 && p[n-1] == '\n' {
		// Trim CR.
		p = p[0 : n-1]
	}
	l.p.Println(string(p))
	return n, nil
}

func SetupBubbleTeaLogger(p *tea.Program) {
	zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
		return filepath.Base(file) + ":" + strconv.Itoa(line)
	}

	o := logToBubbleTea{p}
	logger := zerolog.New(zerolog.ConsoleWriter{Out: &o}).With().Timestamp().Caller().Logger()
	log.Logger = logger
}
