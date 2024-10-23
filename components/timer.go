package components

import (
	"fmt"
	"time"

	"github.com/aristidebm/pomodoro/events"

	tea "github.com/charmbracelet/bubbletea"
)

type Timer struct {
	ETA               time.Time
	CountDownInSecond int64
	value             time.Duration
}

func (s *Timer) getETA() time.Time {
	if !s.ETA.IsZero() {
		return s.ETA
	}
	s.ETA = time.Now().Add(time.Duration(s.CountDownInSecond) * time.Second)
	return s.ETA
}

func (s *Timer) Init() tea.Cmd {
	return nil
}

func (s *Timer) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case events.TickMsg:
		v := s.getETA().Sub(time.Time(msg))
		if v >= 0 {
			s.value = v
		}
	}
	return s, nil
}

func (s *Timer) View() string {
	v := s.value.Seconds()
	minute := int(v / 60)
	second := int(v) % 60
	return fmt.Sprintf("%d:%d", minute, second)
}
