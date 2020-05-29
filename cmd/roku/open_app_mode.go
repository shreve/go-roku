package main

import (
	"strings"
	"github.com/shreve/go-roku/roku"
	"github.com/shreve/tui"
	"github.com/shreve/tui/ansi"
)

type OpenAppMode struct {
	query string
}

func (o *OpenAppMode) InputHandler(in string) {
	switch in {
	case "\x1B":
		o.query = ""
		app.SetMode(Main)
	case "\u007F":
		if len(o.query) > 0 {
			o.query = o.query[0:len(o.query)-1]
		}
	case "\r":
		res := o.findApp()
		o.query = ""
		if res.Name != "" {
			go client.Launch(res.Id)
			app.SetMode(Main)
		}
	default:
		o.query += in
	}
	app.Redraw()
}

func (o *OpenAppMode) Render(height, width int) tui.View {
	view := make(tui.View, height)
	comp := o.findCompletion()
	view[0] = "App to open: " + o.query + suggestText + comp + ansi.DisplayResetCode
	return view
}

func (o *OpenAppMode) findApp() roku.App {
	if o.query == "" {
		return roku.App{}
	}
	for a := range apps {
		app := apps[a]
		i := strings.Index(
			strings.ToLower(app.Name),
			strings.ToLower(o.query))
		if i == 0 {
			return app
		}
	}
	return roku.App{}
}

func (o *OpenAppMode) findCompletion() string {
	app := o.findApp()
	if app.Name != "" {
		return app.Name[len(o.query):len(app.Name)]
	} else {
		return ""
	}
}
