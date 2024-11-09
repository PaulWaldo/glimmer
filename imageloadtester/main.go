package main

import (
	"fmt"
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

func loadImage(photoUrl string, card *widget.Card) {
	fmt.Println("Sleep start")
	time.Sleep(time.Second * 5) // Simulate a really long download
	fmt.Println("Waking up")
	uri, err := storage.ParseURI(photoUrl)
	if err != nil {
		fyne.LogError("parsing url", err)
		return
	}
	image := canvas.NewImageFromURI(uri)
	image.FillMode = canvas.ImageFillContain
	card.SetContent(image)
}

func main() {
	app := app.New()
	win := app.NewWindow("Image Loading Tester")
	win.Resize(fyne.NewSize(400, 400))

	card := widget.NewCard("Image Loading Tester", "This is a test to see if the image loading is working.", nil)
	card.SetContent(canvas.NewRectangle(color.Black))
	win.SetContent(container.NewStack(card))
	go loadImage("https://live.staticflickr.com/65535/54128319919_2ee12b9544_z.jpg", card)

	win.ShowAndRun()
}
