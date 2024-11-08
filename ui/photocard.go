package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type photoCard struct {
	title  widget.Label
	author widget.Label
	img    fyne.CanvasObject
}

func NewPhotoCard(title, author string, img fyne.CanvasObject) *photoCard {
	return &photoCard{
		title:  widget.Label{Text: title},
		author: widget.Label{Text: author},
		img:    img,
	}
}

func (pc *photoCard) makeUI() fyne.CanvasObject {
	imageInfo := container.NewVBox(&pc.title, &pc.author)
	return container.NewBorder(nil, imageInfo, nil, nil, pc.img)
}
