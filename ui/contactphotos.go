package ui

import (
	"fmt"

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

var (
	imageLoadSemaphore = make(chan struct{}, 5) // Allow up to 5 concurrent image loads
	runloopStarted     = make(chan struct{})
)

type contactPhotos struct {
	ma         *myApp
	container  *fyne.Container
	title      *widget.Label
	photoList  *fyne.Container
	gridWrap   *fyne.Container
	photos     []api.Photo
	photoCards []fyne.CanvasObject
	page       int
	totalPages int
}

func (p *contactPhotos) makeUI() *fyne.Container {
	p.title = widget.NewLabel("Contact Photos")

	p.gridWrap = container.NewGridWrap(fyne.NewSize(GridSizeWidth, GridSizeHeight), p.photoCards...)
	scrollingGrid := container.NewScroll(p.gridWrap)
	scrollingGrid.OnScrolled = func(pos fyne.Position) {
		if pos.Y+scrollingGrid.Size().Height >= p.gridWrap.Size().Height {
			p.loadNextPage()
		}
	}

	p.container = container.NewStack(
		container.NewStack(),
		container.NewBorder(p.title, nil, nil, nil, scrollingGrid),
	)
	p.loadNextPage()
	return p.container
}

func (p *contactPhotos) loadNextPage() {
	p.page++
	photos, err := api.GetContactPhotos(p.ma.client, p.page)
	if err != nil {
		fmt.Println(err)
		return
	}
	p.photos = append(p.photos, photos.Photos.Photos...)
	p.gridWrap.Objects = append(p.gridWrap.Objects, p.makePhotoCards(photos.Photos.Photos)...)
	p.gridWrap.Refresh()
}

func (p *contactPhotos) makePhotoCards(photos []api.Photo) []fyne.CanvasObject {
	cards := make([]fyne.CanvasObject, len(photos))
	for i, photo := range photos {
		card := NewPhotoCard(photo, p.ma.client, func() {
			pv := &photoView{ma: p.ma, photo: photo}
			cont, err := pv.makeUI()
			if err != nil {
				fyne.LogError("parsing url", err)
				return
			}
			p.ma.vs.Push(cont)
		})
		cards[i] = card
	}
	return cards
}

type PhotoCard struct {
	widget.Card
	info   photos.PhotoInfo
	photo  api.Photo
	client *flickr.FlickrClient
	tap    func()
}

func NewPhotoCard(photo api.Photo, client *flickr.FlickrClient, onTapped func()) *PhotoCard {
	clone := CloneClient(client)
	i := &PhotoCard{
		Card: widget.Card{
			Title:    photo.Title,
			Subtitle: photo.Username,
			Content:  widget.NewProgressBarInfinite(),
		},
		tap:    onTapped,
		photo:  photo,
		client: clone,
	}
	i.ExtendBaseWidget(i)
	go func() {
		<-runloopStarted
		imageLoadSemaphore <- struct{}{} // Acquire a semaphore slot
		i.loadImage(func() {
			<-imageLoadSemaphore // Release the semaphore slot
		})
	}()
	return i
}

func CloneClient(orig *flickr.FlickrClient) *flickr.FlickrClient {
	clone := flickr.NewFlickrClient(orig.ApiKey, orig.ApiSecret)
	clone.OAuthToken = orig.OAuthToken
	clone.OAuthTokenSecret = orig.OAuthTokenSecret
	return clone
}

func (c *PhotoCard) loadImage(callback func()) {
	// Load the image...
	resp, err := photos.GetInfo(c.client, c.photo.Id, c.photo.Secret)
	if err != nil {
		fyne.LogError("Failed to get photo info", err)
		callback() // Release the semaphore slot
		return
	}
	c.info = resp.Photo
	photoUrl := fmt.Sprintf("https://live.staticflickr.com/%s/%s_%s_%s.jpg", c.info.Server, c.info.Id, c.info.Secret, "z")
	uri, err := storage.ParseURI(photoUrl)
	if err != nil {
		fyne.LogError("parsing url", err)
		c.Content = widget.NewLabel("Failed to load image")
		callback() // Release the semaphore slot
		return
	}

	image := canvas.NewImageFromURI(uri)
	if image == nil || image.Resource == nil {
		panic("Image is nil")
	}
	image.FillMode = canvas.ImageFillContain
	c.SetContent(image)
	callback() // Release the semaphore slot
}

func (c *PhotoCard) Tapped(e *fyne.PointEvent) {
	if c.tap == nil {
		return
	}
	c.tap()
}
