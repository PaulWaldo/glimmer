package ui

import (
	"fmt"
	"strconv"

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
	// Find the smallest photo that is larger than the current window
	photoIndex := -1
	winSize := p.ma.window.Canvas().Size()
	for i := range pai.Sizes {
		picWidth, err := strconv.ParseFloat(pai.Sizes[i].Width, 32)
		if err != nil {
			return nil, fmt.Errorf("converting width '%s' to float: %w", pai.Sizes[i].Width, err)
		}
		picHeight, err := strconv.ParseFloat(pai.Sizes[i].Height, 32)
		if err != nil {
			return nil, fmt.Errorf("converting height '%s' to float: %w", pai.Sizes[i].Height, err)
		}
		if picWidth >= float64(winSize.Width) && picHeight >= float64(winSize.Height) {
			photoIndex = i
			break
		}
	}
	if photoIndex == -1 {
		return nil, fmt.Errorf("no suitable photo found")
	}
	info := pai.Sizes[photoIndex]
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
