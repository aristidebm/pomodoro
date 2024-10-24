package components

import (
	"errors"
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/aristidebm/pomodoro/events"
)

func tickEvery(duration time.Duration) tea.Cmd {
	return tea.Every(duration, func(t time.Time) tea.Msg {
		return events.TickMsg(t)
	})
}

var PlayListLoadingError = errors.New("cannot load the playlist")

type app struct {
	isRunning bool
	isHelp    bool
	timer     *timer
	player    *player
	help      *Help
	width     int
	height    int
}

func (s *app) Init() tea.Cmd {
	// returns an initial command for the application to run
	return tickEvery(time.Microsecond)
}

func (s *app) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// that handles incoming events and updates the state accordingly.
	// and returns the updated state and the cmd to run on it
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:
		switch msg.String() {
		// These keys should exit the program.
		case "q":
			return s, tea.Quit

		case "p":
			s.play()

		case "h":
			s.isHelp = !s.isHelp
			s.help.Update(msg)
		}
	case tea.WindowSizeMsg:
		s.width = msg.Width
		s.height = msg.Height

	case events.TickMsg:
		// Return your Every command again to loop.
		s.timer.Update(msg)
		return s, tickEvery(time.Second)
	}

	return s, nil
}

var appStyle = lipgloss.NewStyle().
	// BorderStyle(lipgloss.NormalBorder()).
	// BorderForeground(lipgloss.Color("#FF0000")).
	Width(100).
	Height(10).
	Padding(2, 4).
	Align(lipgloss.Center)

func (s *app) View() string {
	// renders the UI based on the data in the model.
	if s.width == 0 {
		return ""
	}

	page := appStyle.Render(lipgloss.JoinVertical(lipgloss.Left,
		s.timer.View(),
		s.player.View(),
	))

	if s.isHelp {
		page = appStyle.Render(s.help.View())
	}

	return lipgloss.Place(s.width,
		s.height,
		lipgloss.Center,
		lipgloss.Center,
		page,
	)
}

func (s *app) play() {
	s.isRunning = !s.isRunning
	s.timer.play()
	s.player.play()
}

type option struct {
	songPath string
}

type optionSetter func(o *option)

func WithPlayList(path string) optionSetter {
	return func(o *option) {
		o.songPath = path
	}
}

func NewApp(duration int64, opts ...optionSetter) (*app, error) {
	// set options
	opt := &option{}
	for _, f := range opts {
		f(opt)
	}

	var songs []song
	if opt.songPath != "" {
		if _, err := os.Stat(opt.songPath); err != nil {
			return nil, fmt.Errorf("%w: %w", PlayListLoadingError, err)
		}

		result, err := generatePlayList(opt.songPath)
		if err != nil {
			return nil, fmt.Errorf("%w: %w", PlayListLoadingError, err)
		}
		songs = result
	}

	return &app{
		isRunning: false,
		isHelp:    false,
		timer:     &timer{duration: duration},
		player:    newPlayer(songs),
		help:      NewHelp(),
	}, nil
}
