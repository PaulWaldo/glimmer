package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	app := app.New()
	win := app.NewWindow("Card Test")
	win.Resize(fyne.Size{Width: 500, Height: 500})

	cards := []fyne.CanvasObject{
		widget.NewCard("ReaLLLLLLLLLLLLy Long Title", "ReaLLLLLLLLLLLLLLLLLLLLLLLLLLLy Long Subtitle", nil),
		widget.NewCard("Really Long Title", "Really Long Subtitle", nil),
		widget.NewCard("Really Long Title", "Really Long Subtitle", nil),
		widget.NewCard("Title", "Subtitle", nil),
	}

	gw := container.NewGridWrap(fyne.NewSize(200, 200), cards...)

	win.SetContent(gw)
	win.ShowAndRun()
}
