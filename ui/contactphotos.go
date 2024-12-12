package ui

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"github.com/PaulWaldo/glimmer/api"
	"gopkg.in/masci/flickr.v3"
	"gopkg.in/masci/flickr.v3/photos"
)

const (
	GridSizeWidth  = 500
	GridSizeHeight = 500
)

type contactPhotos struct {
	ma         *myApp
	container  *fyne.Container
	title      *widget.Label
	photoList  *fyne.Container
	photos     []api.Photo
	photoCards []fyne.CanvasObject
}

func (p *contactPhotos) makeUI() *fyne.Container {
	p.title = widget.NewLabel("Contact Photos")
	
	// Create cards for each photo
	p.photoCards = make([]fyne.CanvasObject, len(p.photos))
	for i, photo := range p.photos {
		card := NewPhotoCard(photo, p.ma.client, func() {
			pv := &photoView{ma: p.ma, photo: photo}
			cont, err := pv.makeUI()
			if err != nil {
				fyne.LogError("parsing url", err)
				return
			}
			p.ma.vs.Push(cont)
		})
		p.photoCards[i] = card
	}

	gw := container.NewGridWrap(fyne.NewSize(GridSizeWidth, GridSizeHeight), p.photoCards...)
	scrollingGrid := container.NewScroll(gw)

	p.container = container.NewStack(
		container.NewStack(),
		container.NewBorder(p.title, nil, nil, nil, scrollingGrid),
	)
	return p.container
}

type PhotoCard struct {
	*widget.Card
	info   photos.PhotoInfo
	photo  api.Photo
	client *flickr.FlickrClient
	tap    func()
}

func NewPhotoCard(photo api.Photo /*content fyne.CanvasObject,*/, client *flickr.FlickrClient, onTapped func()) *PhotoCard {
	clone := CloneClient(client)
	i := &PhotoCard{tap: onTapped, photo: photo, client: clone}
	i.Card = widget.NewCard(photo.Title, photo.Username, canvas.NewRectangle(color.Black))
	i.ExtendBaseWidget(i)
	go i.loadImage()
	return i
}

func CloneClient(orig *flickr.FlickrClient) *flickr.FlickrClient {
	clone := flickr.NewFlickrClient(orig.ApiKey, orig.ApiSecret)
	clone.OAuthToken = orig.OAuthToken
	clone.OAuthTokenSecret = orig.OAuthTokenSecret
	return clone
}

func (c *PhotoCard) loadImage() {
	// fmt.Printf("Sleep start on card %p\n", c)
	// time.Sleep(time.Second * time.Duration(rand.Int64N(4))) // Simulate a really long download
	// fmt.Println("Waking up")

	resp, err := photos.GetInfo(c.client, c.photo.Id, c.photo.Secret)
	if err != nil {
		fyne.LogError("Failed to get photo info", err)
		return
	}
	c.info = resp.Photo
	photoUrl := fmt.Sprintf("https://live.staticflickr.com/%s/%s_%s_%s.jpg", c.info.Server, c.info.Id, c.info.Secret, "z")
	uri, err := storage.ParseURI(photoUrl)
	if err != nil {
		fyne.LogError("parsing url", err)
		c.Content = widget.NewLabel("Failed to load image")
		return
	}
	fmt.Println("Downloading ", uri)
	image := canvas.NewImageFromURI(uri)
	if image == nil || image.Resource == nil {
		panic("Image is nil")
	}
	fmt.Printf("Image size is %d\n", len(image.Resource.Content()))
	image.FillMode = canvas.ImageFillContain
	fmt.Println("Got ", uri)
	c.SetContent(image)
	// c.SetContent(canvas.NewRectangle(color.RGBA{R: 250, G: 10, B: 10, A: 255}))
}

func (c *PhotoCard) Tapped(e *fyne.PointEvent) {
	if c.tap == nil {
		return
	}
	c.tap()
}
