package components

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	// There is also this package, you may check it out
	// https://github.com/common-nighthawk/go-figure
	"github.com/zs5460/art"

	"github.com/aristidebm/pomodoro/internal/events"
)

type Timer struct {
	duration  int64
	elapsed   int64
	isRunning bool
	value     time.Duration
}

func (s *Timer) Init() tea.Cmd {
	return nil
}

func (s *Timer) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case events.TickMsg:
		// the first thick at applicaton start so
		// compute and set the value
		if s.isZero() {
			s.reset()
		}

		if s.isRunning {
			s.update(time.Time(msg))
		}
	}
	return s, nil
}

func (s *Timer) reset() {
	s.value = time.Duration(s.duration) * time.Second
}

func (s *Timer) isZero() bool {
	return s.value.Microseconds() == 0
}

func (s *Timer) update(value time.Time) {
    if v := time.Duration(s.duration-s.elapsed) * time.Second; v >= 0 {
        s.value = v
        s.elapsed += 1
    } else {
        // reset the timer
        s.elapsed = 0
        s.isRunning = false
    }
}

func (s *Timer) play() {
	// toggle the running state
	s.isRunning = !s.isRunning
}

var timeStyle = lipgloss.NewStyle().
	// BorderStyle(lipgloss.NormalBorder()).
	// BorderForeground(lipgloss.Color("#FFFFFF")).
	Align(lipgloss.Center).
	Height(2).
	Bold(true).
	Padding(0, 4)

func (s *Timer) View() string {
	v := s.value.Seconds()
	minute := int(v / 60)
	second := int(v) % 60
	return timeStyle.Render(art.String(fmt.Sprintf("%s:%s", normalizeNumber(minute), normalizeNumber(second))))
}

func normalizeNumber(number int) string {
	if number < 10 {
		return "0" + fmt.Sprintf("%d", number)
	}
	return fmt.Sprintf("%d", number)
}
