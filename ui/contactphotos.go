package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"gopkg.in/masci/flickr.v3"
)

type contactPhotos struct {
	ma        *myApp
	container *fyne.Container
	title     *widget.Label
	photoList *widget.List
	ptotos flickr.
}

func (p *contactPhotos) makeUI() fyne.CanvasObject {
	p.title = widget.NewLabel("Contact Photos")
	p.photoList=widget.NewList()
	p.container = container.NewVBox(p.title)
	return p.container
}
