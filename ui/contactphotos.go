package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type contactPhotos struct {
	ma        *myApp
	container *fyne.Container
	title     *widget.Label
}

func (p *contactPhotos) makeUI() fyne.CanvasObject {
	p.title = widget.NewLabel("Contact Photos")
	p.container = container.NewVBox(p.title)
	return p.container
}
