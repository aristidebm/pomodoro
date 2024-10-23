package components

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/aristidebm/pomodoro/events"
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

var timeStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("#FFFFFF")).
	Align(lipgloss.Center).
	Padding(2, 4)

func (s *Timer) View() string {
	v := s.value.Seconds()
	minute := int(v / 60)
	second := int(v) % 60
	return timeStyle.Render(fmt.Sprintf("%s:%s", normalizeNumber(minute), normalizeNumber(second)))
}

func normalizeNumber(number int) string {
	if number < 10 {
		return "0" + fmt.Sprintf("%d", number)
	}
	return fmt.Sprintf("%d", number)
}
