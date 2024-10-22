package main

import (
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type TickMsg time.Time

// Send a message every second.
func tickEvery(duration time.Duration) tea.Cmd {
    return tea.Every(duration, func(t time.Time) tea.Msg {
        return TickMsg(t)
    })
}

type TimerState struct {
	Timer time.Time
}

func (s *TimerState) Init() tea.Cmd {
    return nil
}

func (s *TimerState) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case TickMsg:
        // Return your Every command again to loop.
        s.Timer = time.Time(msg)
    }
	return s, nil
}

func (s *TimerState) View() string {
	return s.Timer.Format(time.TimeOnly)
}

type Song struct {
	Artist string
	Title  string
}

type PlayListState struct {
	PlayList []Song
}

func (s *PlayListState) Init() tea.Cmd {
	return nil
}

func (s *PlayListState) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return s, nil
}

func (s *PlayListState) View() string {
	return ""
}

type AppState struct {
	IsRunning bool
	Timer     *TimerState
	PlayList  *PlayListState
}

func (s *AppState) Init() tea.Cmd {
	// returns an initial command for the application to run
	return tickEvery(time.Second)
}

func (s *AppState) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
    case TickMsg:
        // Return your Every command again to loop.
        s.Timer.Update(msg)
        return s, tickEvery(time.Second)
	}

	return s, nil
}

func (s *AppState) View() string {
	// renders the UI based on the data in the model.
	return s.Timer.View()
}

func main() {
	initialState := &AppState{
		IsRunning: false,
		PlayList:  &PlayListState{PlayList: []Song{}},
		Timer:     &TimerState{Timer: time.Now()},
	}

	app := tea.NewProgram(initialState, tea.WithAltScreen())
	if _, err := app.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
