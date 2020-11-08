package main

import (
	"time"

	"github.com/shreve/go-roku/roku"
	"github.com/shreve/tui"
)

type MainMode struct {
}

func (m *MainMode) InputHandler(in string) {
	switch in {
	case "q", tui.CtrlC:
		app.Done()
	case "o":
		app.SetMode(OpenApp)

	// Vim keys to navigate
	case "h":
		Press(roku.Left)
	case "j":
		Press(roku.Down)
	case "k":
		Press(roku.Up)
	case "l":
		Press(roku.Right)

	// Arrow keys also navigate
	case "\x1B[A":
		Press(roku.Up)
	case "\x1B[B":
		Press(roku.Down)
	case "\x1B[C":
		Press(roku.Right)
	case "\x1B[D":
		Press(roku.Left)

	// Control + arrow keys to change volume, sideways to mute
	case "\x1B[1;5A":
		Press(roku.VolumeUp)
	case "\x1B[1;5B":
		Press(roku.VolumeDown)
	case "\x1B[1;5C":
		Press(roku.VolumeMute)
	case "\x1B[1;5D":
		Press(roku.VolumeMute)

	// Space bar to play / pause
	case " ":
		Press(roku.Play)

	// Match asterisk button on remote
	case "*":
		Press(roku.Info)

	// Escape and backspace to back / exit
	case "\x1B":
		Press(roku.Back)
	case "\x7F":
		Press(roku.Back)

	// Enter to select
	case "\r":
		Press(roku.Select)

	// Ctrl + q to power off
	case "\x11":
		Press(roku.PowerOff)
	}
}

func (m *MainMode) Render(height, width int) tui.View {
	view := make(tui.View, height)

	if !client.Ready() {
		view[0] = "Searching for Roku device..."
	} else {
		app := client.ActiveApp()
		name := info.FriendlyDeviceName
		view[0] = "Current Roku: " + name + "          Uptime: " + humanTime(info.Uptime)
		view[1] = "Current App: " + app.Name
		view[2] = "Last Button: " + lastPressed
		view[len(view)-1] = client.Address
	}

	return view
}

func humanTime(seconds uint) string {
	return (time.Second * time.Duration(seconds)).String()
}
