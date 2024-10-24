package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"

	"flag"

	"github.com/aristidebm/pomodoro/components"
	tea "github.com/charmbracelet/bubbletea"
)

var durationParser = regexp.MustCompile(`(\d+)(d|h|m|s)?`)

var path = flag.String("path", "", "path to playlist (.mp3)")
var duration = flag.String("duration", "25m", "session duration, example 1d|24h|1444m|86400")

func main() {
	flag.Parse()

	match := durationParser.FindStringSubmatch(*duration)
	if len(match) < 2 {
		fmt.Printf("Alas, there's been an error")
		os.Exit(1)
	}

	d, err := strconv.Atoi(match[1])
	if err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}

	if len(match) > 2 {
		switch match[2] {
		case "m":
			d = d * 60
		case "h":
			d = d * 60 * 60
		case "d":
			d = d * 24 * 60 * 60
		default:
			//
		}
	}

	app, err := components.NewApp(int64(d), components.WithPlayList(*path))
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
