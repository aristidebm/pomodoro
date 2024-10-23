package main

import (
	// "fmt"
	// "os"
	// "time"

	"flag"
	// "github.com/aristidebm/pomodoro/components"
	// tea "github.com/charmbracelet/bubbletea"
)

var filename = flag.String("filename", "", "mp3 filename")

func main() {
	flag.Parse()
	if *filename != "" {
		Play(*filename)
	}

	// initialState := &components.App{
	// 	IsRunning: false,
	// 	Timer:     &components.Timer{ETA: time.Now().Add(1200 * time.Second)},
	// }
	// app := tea.NewProgram(initialState, tea.WithAltScreen())
	// if _, err := app.Run(); err != nil {
	// 	fmt.Printf("Alas, there's been an error: %v", err)
	// 	os.Exit(1)
	// }
}
