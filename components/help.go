package components

import (
	// "fmt"
	//    "slices"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/zs5460/art"
)

type HelpEntry struct {
	cmd string
	msg string
}

type Help struct {
	entries []HelpEntry
}

func (h *Help) Init() tea.Cmd {
	return nil
}

func (h *Help) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return h, nil
}

var leftBlockStyle = lipgloss.NewStyle().
	// BorderStyle(lipgloss.NormalBorder()).
	Bold(true).
	// BorderForeground(lipgloss.Color("#FF0000")).
	// Height(10).
	Align(lipgloss.Center)

var rightBlockStyle = lipgloss.NewStyle().
	// BorderStyle(lipgloss.NormalBorder()).
	Bold(true).
	// BorderForeground(lipgloss.Color("#00FF00")).
	MarginLeft(2).
	Align(lipgloss.Center)

var titleStyle = lipgloss.NewStyle().
	Align(lipgloss.Center)

func (h *Help) View() string {

	cmds := []string{}
	msgs := []string{}

	for _, e := range h.entries {
		cmds = append(cmds, e.cmd)
		msgs = append(msgs, e.msg)
	}

	leftBlock := leftBlockStyle.Render(lipgloss.JoinVertical(lipgloss.Left,
		cmds...,
	))

	// clear the slice
	rightBlock := rightBlockStyle.Render(lipgloss.JoinVertical(lipgloss.Left,
		msgs...,
	))

	body := lipgloss.JoinHorizontal(lipgloss.Center, leftBlock, rightBlock)
	title := titleStyle.Render(art.String("Help!"))

	return lipgloss.JoinVertical(lipgloss.Center, title, body)
}

// func (h *Help) format(e HelpEntry) string {
// 	return helpStyle.Render(fmt.Sprintf("%s \t\t %s", e.cmd, e.msg))
// }

func NewHelp() *Help {
	return &Help{
		entries: []HelpEntry{
			{cmd: "h", msg: "show or hide this help page"},
			{cmd: "p", msg: "toggle play and pause"},
			{cmd: "r", msg: "reset the timer"},
		},
	}
}
