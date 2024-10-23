package components

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/aristidebm/pomodoro/events"
)

func tickEvery(duration time.Duration) tea.Cmd {
	return tea.Every(duration, func(t time.Time) tea.Msg {
		return events.TickMsg(t)
	})
}

type App struct {
	IsRunning bool
	Timer     *Timer
}

func (s *App) Init() tea.Cmd {
	// returns an initial command for the application to run
	return tickEvery(time.Second)
}

func (s *App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// that handles incoming events and updates the state accordingly.
	// and returns the updated state and the cmd to run on it
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:
		switch msg.String() {
		// These keys should exit the program.
		case "q":
			return s, tea.Quit
		}
	case events.TickMsg:
		// Return your Every command again to loop.
		s.Timer.Update(msg)
		return s, tickEvery(time.Second)
	}

	return s, nil
}

func (s *App) View() string {
	// renders the UI based on the data in the model.
	return s.Timer.View()
}
