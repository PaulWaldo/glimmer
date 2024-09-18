package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
	"github.com/PaulWaldo/glimmer/api"
	"gopkg.in/masci/flickr.v3/photos"
)

type photoView struct {
	ma    *myApp
	photo api.Photo
}

func (p photoView) makeUI() (*fyne.Container, error) {
	pai, err := photos.GetSizes(p.ma.client, p.photo.Id)
	if err != nil {
		fyne.LogError("getting sizes", err)
		return nil, fmt.Errorf("getting sizes: %w", err)
	}
	fmt.Printf("access Info: %+v\n", pai)
	// Find the largest photo.  Assumes they are listed is ascending order
	info := pai.Sizes[len(pai.Sizes)-1]
	url := info.Source
	uri, err := storage.ParseURI(url)
	if err != nil {
		fyne.LogError("parsing url", err)
		return nil, err
	}
	fmt.Println("Downloading ", uri)
	im := canvas.NewImageFromURI(uri)
	im.FillMode = canvas.ImageFillContain
	cont := container.NewStack(im)
	return cont, nil
}
