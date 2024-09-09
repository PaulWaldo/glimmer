package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"github.com/PaulWaldo/glimmer/api"
	"gopkg.in/masci/flickr.v3/photos"
)

type contactPhotos struct {
	ma         *myApp
	container  *fyne.Container
	title      *widget.Label
	photoList  *fyne.Container
	photos     []api.Photo
	photoCards []fyne.CanvasObject
}

func (p *contactPhotos) makeUI() fyne.CanvasObject {
	p.title = widget.NewLabel("Contact Photos")
	// m := image.NewRGBA(image.Rect(0, 0, 640, 640))
	// for x := range 640 {
	// 	for y := range 640 {
	// 		m.Set(x, y, color.RGBA{255, 0, 0, 255})
	// 	}
	// }
	// placeholderImage := canvas.NewImageFromImage(m)

	// Create cards for each photo
	p.photoCards = make([]fyne.CanvasObject, len(p.photos))
	for i, photo := range p.photos {
		info, err := photos.GetInfo(p.ma.client, photo.Id, photo.Secret)
		if err != nil {
			fyne.LogError("Failed to get photo info", err)
			continue
		}
		photoUrl := fmt.Sprintf("https://live.staticflickr.com/%s/%s_%s_%s.jpg", info.Photo.Server, info.Photo.Id, info.Photo.Secret, "z")
		uri, err := storage.ParseURI(photoUrl)
		if err != nil {
			fyne.LogError("parsing url", err)
			continue
		}
		fmt.Println("Downloading ", uri)
		c := canvas.NewImageFromURI(uri)
		// card.Image = placeholderImage
		// card := widget.NewCard(photo.Title, photo.Username, nil)
		card := newTapCard(photo.Title, photo.Username, nil, func() {
			photoUrl := fmt.Sprintf("https://live.staticflickr.com/%s/%s_%s_%s.jpg", info.Photo.Server, info.Photo.Id, info.Photo.Secret, "k")
			uri, err := storage.ParseURI(photoUrl)
			if err != nil {
				fyne.LogError("parsing url", err)
				return
			}
			fmt.Println("Downloading ", uri)
			c := canvas.NewImageFromURI(uri)
			cont := container.NewStack(c)
			p.ma.window.SetContent(cont)
		})
		card.Content = c
		c.FillMode = canvas.ImageFillContain
		// card.Image.FillMode = canvas.ImageFillOriginal

		p.photoCards[i] = card
	}

	gw := container.NewGridWrap(fyne.NewSize(500, 500), p.photoCards...)
	scrollingGrid := container.NewScroll(gw)
	p.container = container.NewBorder(p.title, nil, nil, nil, scrollingGrid)
	return p.container
}

type tapCard struct {
	*widget.Card
	tap func()
}

func newTapCard(title, subtitle string, content fyne.CanvasObject, fn func()) *tapCard {
	i := &tapCard{tap: fn}
	i.Card = widget.NewCard(title, subtitle, content)
	i.ExtendBaseWidget(i)
	return i
}
