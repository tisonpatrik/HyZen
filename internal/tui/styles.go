package tui

import (
	"github.com/charmbracelet/lipgloss"

	"hyprzen/internal/utils"
)

const (
	progressBarWidth  = 71
	progressFullChar  = "█"
	progressEmptyChar = "░"
	dotChar           = " • "
)

var (
	keywordStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("211"))
	subtleStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
	ticksStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("79"))
	checkboxStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("212"))
	progressEmpty = subtleStyle.Render(progressEmptyChar)
	dotStyle      = lipgloss.NewStyle().
			Foreground(lipgloss.Color("236")).
			Render(dotChar)
	mainStyle = lipgloss.NewStyle().MarginLeft(2)

	// Gradient colors we'll use for the progress bar
	ramp = utils.MakeRampStyles("#B14FFF", "#00FFA3", progressBarWidth)
)
