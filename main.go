package main

import (
	"fmt"
	"os"
	"time"

	"github.com/aristidebm/pomodoro/components"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	initialState := &components.App{
		IsRunning: false,
		Timer:     &components.Timer{ETA: time.Now().Add(60 * time.Second)},
	}
	app := tea.NewProgram(initialState, tea.WithAltScreen())
	if _, err := app.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
