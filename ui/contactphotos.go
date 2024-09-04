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
	ma        *myApp
	container *fyne.Container
	title     *widget.Label
	photoList *widget.List
	photos    []api.Photo
}

func (p *contactPhotos) makeUI() fyne.CanvasObject {
	p.title = widget.NewLabel("Contact Photos")
	p.photoList = widget.NewList(
		func() int { return len(p.photos) },
		func() fyne.CanvasObject {
			card := widget.NewCard("Some", "Some Subtitle", nil)
			return container.NewStack(card)
		},
		func(index widget.ListItemID, template fyne.CanvasObject) {
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
			// Download the image at url and convert to a canvas.Image
			// uri,err := url.Parse(photoUrl)
			uri, err := storage.ParseURI(photoUrl)
			if err != nil {
				fyne.LogError("parsing url", err)
			}
			c:=canvas.NewImageFromURI(uri)
			card.SetContent(c)

		},
	)
	p.container = container.NewBorder(p.title, nil, nil, nil, p.photoList)
	return p.container
}
