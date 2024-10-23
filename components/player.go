package components

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/dhowden/tag"
)

type song struct {
	filename string
	artist   string
	title    string
}

type player struct {
	icon      string
	songs     []song
	cursor    int
	isRunning bool
}

func (p *player) Init() tea.Cmd {
	return nil
}

func (p *player) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return p, nil
}

var playerStyle = lipgloss.NewStyle().
	// BorderStyle(lipgloss.NormalBorder()).
	// BorderForeground(lipgloss.Color("#0000FF")).
	Align(lipgloss.Center)
	// Height(2).
	// Padding(0, 4)

func (p *player) View() string {
	if len(p.songs) <= 0 {
		return ""
	}
	return playerStyle.Render(fmt.Sprintf("%s %s - %s", p.icon, p.songs[p.cursor].artist, p.songs[p.cursor].title))
}

func (p *player) play() {
	p.isRunning = !p.isRunning
	if p.isRunning {

	} else {

	}
}

func newPlayer(songs []song) *player {
	return &player{
		icon:  "",
		songs: songs,
	}
}

func generatePlayList(path string) ([]song, error) {
	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	if !info.IsDir() {
		return getSong(path)
	}

	songs := []song{}
	err = filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
		// FIXME: check for errors we want to ignore
		if err != nil {
			return nil
		}

		if d.IsDir() {
			// ingnore dirs
			return nil
		}
		result, err := getSong(path)

		// FIXME: check for errors we want to ignore
		if err != nil {
			return nil
		}
		songs = append(songs, result...)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return songs, nil
}

func getSong(path string) ([]song, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	meta, err := tag.ReadFrom(f)
	if err != nil {
		return nil, err
	}
	return []song{{
		filename: path,
		artist:   meta.Artist(),
		title:    meta.Title(),
	}}, nil
}
