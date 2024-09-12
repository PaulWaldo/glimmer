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
}

func (p photoView) makeUI() (*fyne.Container, error) {
	photoUrl := fmt.Sprintf("https://live.staticflickr.com/%s/%s_%s_%s.jpg", p.info.Photo.Server, p.info.Photo.Id, p.info.Photo.OriginalSecret, "k")
	uri, err := storage.ParseURI(photoUrl)
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
