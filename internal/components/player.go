package components

import (
	"context"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/dhowden/tag"

	// Check this tutorial for introduction
	// https://reintech.io/blog/writing-a-go-based-music-player
	"github.com/gopxl/beep/v2"
	"github.com/gopxl/beep/v2/mp3"
	"github.com/gopxl/beep/v2/speaker"
)

// check here for a better implementation
// https://github.com/Malwarize/retro/blob/dd50e3722eef83406d1322031a4ca9f3d8707410/server/player/musics.go#L13
type Song struct {
	filename string
	artist   string
	title    string
}

// To know how to handle pauses and resumes check here
// https://github.com/Malwarize/retro/blob/dd50e3722eef83406d1322031a4ca9f3d8707410/server/player/player.go#L275
type Player struct {
	icon           string
	songs          []Song
	cursor         int
	isRunning      bool
	trackIsPlaying bool
	isInitialized  bool
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
	return playerStyle.Render(fmt.Sprintf("%s %s - %s", p.icon, p.songs[p.cursor].artist, p.songs[p.cursor].title))
}

func (p *Player) play() {
	p.isRunning = !p.isRunning
	if p.isRunning && !p.trackIsPlaying && p.cursor < len(p.songs) {
		ctx := context.Background()
		ctx, _ = context.WithCancel(ctx)
		go p.playTrack(ctx)
	}
}

func (p *Player) next() {
	p.cursor = (p.cursor + 1) % len(p.songs)
}

func newPlayer(songs []Song) *Player {
	return &Player{
		icon:  "",
		songs: songs,
	}
}

func generatePlayList(path string) ([]Song, error) {
	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	if !info.IsDir() {
		return getSong(path)
	}

	songs := []Song{}
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

func getSong(path string) ([]Song, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	meta, err := tag.ReadFrom(f)
	if err != nil {
		return nil, err
	}

	return []Song{{
		filename: path,
		artist:   meta.Artist(),
		title:    meta.Title(),
	}}, nil
}

func (p *Player) playTrack(ctx context.Context) error {
	for {
		song := p.songs[p.cursor]
		f, err := os.Open(song.filename)
		if err != nil {
			// try the next item in the list
			p.next()
			continue
		}
		defer f.Close()

		streamer, format, err := mp3.Decode(f)
		if err != nil {
			log.Fatal(err)
		}
		defer streamer.Close()

		if !p.isInitialized && p.isRunning {
			err = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
			if err != nil {
				return err
			}
			p.isInitialized = true
		}

		done := make(chan bool)
		// Play the audio
		speaker.Play(beep.Seq(streamer, beep.Callback(func() {
			done <- true
		})))

		<-done

		p.next()

		streamer.Close()
		f.Close()
	}
}
