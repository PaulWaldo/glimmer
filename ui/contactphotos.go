package ui

import (
	"fmt"
	"image"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"github.com/PaulWaldo/glimmer/api"
	"gopkg.in/masci/flickr.v3/photos"
)

type contactPhotos struct {
	ma        *myApp
	container *fyne.Container
	title     *widget.Label
	photoList *widget.List
	photos    []api.Photo
}

func (p *contactPhotos) makeUI() fyne.CanvasObject {
	p.title = widget.NewLabel("Contact Photos")
	m := image.NewRGBA(image.Rect(0, 0, 640, 640))
	m.Set(5, 5, color.RGBA{255, 0, 0, 255})
	placeholderImage := canvas.NewImageFromImage(m)
	p.photoList = widget.NewList(
		func() int { return len(p.photos) },
		func() fyne.CanvasObject {
			card := widget.NewCard("Some", "Some Subtitle", nil)
			card.Image = placeholderImage
			return container.NewStack(card)
		},
		func(index widget.ListItemID, template fyne.CanvasObject) {
			if index > 10 {
				return
			}
			cont := template.(*fyne.Container)
			card := cont.Objects[0].(*widget.Card)
			photo := p.photos[index]
			card.SetTitle(photo.Title)
			card.SetSubTitle(photo.Username)

			info, err := photos.GetInfo(p.ma.client, photo.Id, photo.Secret)
			if err != nil {
				fyne.LogError("Failed to get photo info", err)
				return
			}
			fmt.Printf("%#v", info)
			photoUrl := fmt.Sprintf("https://live.staticflickr.com/%s/%s_%s_%s.jpg", info.Photo.Server, info.Photo.Id, info.Photo.Secret, "z")
			uri, err := storage.ParseURI(photoUrl)
			if err != nil {
				fyne.LogError("parsing url", err)
			}
			c := canvas.NewImageFromURI(uri)
			card.Image = c
			card.Image.FillMode = canvas.ImageFillOriginal
			// card.SetContent(c)

		},
	)
	p.container = container.NewBorder(p.title, nil, nil, nil, p.photoList)
	return p.container
}
