package main

import (
	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
)

func main() {

	js := mewn.String("./frontend/build/static/js/main.js")
	css := mewn.String("./frontend/build/static/css/main.css")

	app := wails.CreateApp(&wails.AppConfig{
		Width:  1280,
		Height: 1024,
		Title:  "Hello, World",
		JS:     js,
		CSS:    css,
		Colour: "#131313",
	})
	app.Bind(GetMyBridge())
	app.Run()

}
