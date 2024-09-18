package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
	"gopkg.in/masci/flickr.v3/photos"
)

type photoView struct {
	info photos.PhotoInfoResponse
	im   *canvas.Image
	url  string
}

func (p photoView) makeUI() (*fyne.Container, error) {
	uri, err := storage.ParseURI(p.url)
	if err != nil {
		fyne.LogError("parsing url", err)
		return nil, err
	}
	fmt.Println("Downloading ", uri)
	p.im = canvas.NewImageFromURI(uri)
	p.im.FillMode = canvas.ImageFillContain
	cont := container.NewStack(p.im)
	return cont, nil
}
