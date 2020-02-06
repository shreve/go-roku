package main

import (
	"github.com/shreve/go-roku/roku"
	"github.com/shreve/tui"
	"os"
	"time"
)

const (
	Main = iota
	OpenApp
)

var currentMode int = Main
func SetMode(app *tui.App, mode int) {
	if mode == currentMode {
		return
	}
	currentMode = mode
	switch mode {
	case Main:
		app.InputHandler = mainModeInput
		app.CurrentView = indexView
	case OpenApp:
		app.InputHandler = openAppModeInput
		app.CurrentView = openAppView
	}
}

var lastPressed string
func Press(key string) {
	if client.Ready {
		go client.Keypress(key)
		lastPressed = key
	}
}

var query = ""
func openAppModeInput(input string, app *tui.App) {
	switch input {
	case "\x1B":
		query = ""
		SetMode(app, Main)
	case "\u007F":
		query = query[0:len(query)-1]
	case "\r":
		// Run search
	default:
		query += input
	}
	app.Redraw()
}

func mainModeInput(input string, app *tui.App) {
	switch input {
	case "q", tui.CtrlC: app.Done()
	case "o": SetMode(app, OpenApp)

	// Vim keys to navigate
	case "h": Press(roku.Left)
	case "j": Press(roku.Up)
	case "k": Press(roku.Down)
	case "l": Press(roku.Right)

	// Arrow keys also navigate
	case "\x1B[A": Press(roku.Up)
	case "\x1B[B": Press(roku.Down)
	case "\x1B[C": Press(roku.Right)
	case "\x1B[D": Press(roku.Left)

	// Control + arrow keys to change volume, sideways to mute
	case "\x1B[1;5A": Press(roku.VolumeUp)
	case "\x1B[1;5B": Press(roku.VolumeDown)
	case "\x1B[1;5C": Press(roku.VolumeMute)
	case "\x1B[1;5D": Press(roku.VolumeMute)

	// Space bar to play / pause
	case " ": Press(roku.Play)

	// Match asterisk button on remote
	case "*": Press(roku.Info)

	// Escape to back / exit
	case "\x1B": Press(roku.Back)

	// Enter to select
	case "\r": Press(roku.Select)

	// Ctrl + q to power off
	case "\x11": Press(roku.PowerOff)
	}

	app.Redraw()
}

func humanTime(seconds uint) string {
	return (time.Second * time.Duration(seconds)).String()
}

func indexView(height, width int) tui.View {
	view := make(tui.View, height)

	if ! client.Ready {
		view[0] = "Searching for Roku device..."
	} else {
		app := client.ActiveApp()
		name := info.FriendlyDeviceName
		view[0] = "Current Roku: " + name + "          Uptime: " + humanTime(info.Uptime)
		view[1] = "Current App: " + app.Name
		view[2] = "Last Button: " +  lastPressed
	}

	return view
}

func openAppView(height, width int) tui.View {
	view := make(tui.View, height)
	view[0] = "App to open: " + query
	return view
}

var client roku.Client
var info roku.DeviceInfo
var apps []roku.App

func main() {
	tui := tui.NewApp()

	go (func() {
		var err error
		host := os.Getenv("ROKU_HOST")
		if host != "" {
			client, err = roku.Connect(host)
		} else {
			client, err = roku.Discover()
		}
		if err != nil {
			tui.Panic("Unable to find Roku device")
		}
		info = client.DeviceInfo()
		tui.Redraw()
		apps = client.Apps()

		for {
			time.Sleep(time.Second)
			info = client.DeviceInfo()
			tui.Redraw()
		}
	})()

	tui.InputHandler = mainModeInput
	tui.CurrentView = indexView
	tui.Run()
}
