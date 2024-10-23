package components

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type song struct {
	filename string
	artist   string
	title    string
}

type Player struct {
	Icon  string
	songs []song
	curor int
}

func (p *Player) Init() tea.Cmd {
	return nil
}

func (p *Player) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return p, nil
}

var playerStyle = lipgloss.NewStyle().
	// BorderStyle(lipgloss.NormalBorder()).
	// BorderForeground(lipgloss.Color("#0000FF")).
	Align(lipgloss.Center)
	// Height(2).
	// Padding(0, 4)

func (p *Player) View() string {
	if len(p.songs) <= 0 {
		return ""
	}
	return playerStyle.Render(fmt.Sprintf("%s %s - %s", p.Icon, p.songs[p.curor].artist, p.songs[p.curor].title))
}

func NewPlayer() *Player {
	return &Player{
		Icon: "",
		songs: []song{
			{
				filename: "/tmp/song.mp3",
				artist:   "Adele",
				title:    "Hello (25)",
			},
		},
	}
}
