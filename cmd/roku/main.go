package main

import (
	"os"
	"time"

	"github.com/shreve/go-roku/roku"
	"github.com/shreve/tui"
	"github.com/shreve/tui/ansi"
)

const (
	Main = iota
	OpenApp
)

var suggestText = ansi.NewDisplay(90, ansi.Black).Code()

var lastPressed string
func Press(key string) {
	if client.Ready() {
		go client.Keypress(key)
		lastPressed = key
	}
}

var client roku.Client
var info roku.DeviceInfo
var apps []roku.App
var app *tui.App

func main() {
	app = tui.NewApp()
	app.AddMode(Main, &MainMode{})
	app.AddMode(OpenApp, &OpenAppMode{})

	go (func() {
		host := os.Getenv("ROKU_HOST")
		if host != "" {
			client = roku.Connect(host)
		} else {
			var err error
			client, err = roku.Discover()
			if err != nil {
				app.Panic("Unable to find Roku device")
			}
		}
		info = client.DeviceInfo()
		app.Redraw()
		apps = client.Apps()

		for {
			time.Sleep(time.Second)
			info = client.DeviceInfo()
			app.Redraw()
		}
	})()

	app.Run()
}
