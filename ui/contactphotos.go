package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/PaulWaldo/glimmer/api"
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
		},
	)
	p.container = container.NewBorder(p.title, nil, nil, nil, p.photoList)
	return p.container
}
