package main

import (
	"fmt"
	"os"

	// "flag"
	"github.com/aristidebm/pomodoro/components"
	tea "github.com/charmbracelet/bubbletea"
)

// var filename = flag.String("filename", "", "mp3 filename")

func main() {
	// flag.Parse()
	// if *filename != "" {
	// 	Play(*filename)
	// }

	app, err := components.NewApp(25 * 60)
	if err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}

	p := tea.NewProgram(app, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
