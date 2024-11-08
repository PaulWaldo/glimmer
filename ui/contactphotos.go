package ui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/PaulWaldo/glimmer/api"
)

// import (
// 	"fmt"
// 	"image/color"

// 	"fyne.io/fyne/v2"
// 	"fyne.io/fyne/v2/canvas"
// 	"fyne.io/fyne/v2/container"
// 	"fyne.io/fyne/v2/storage"
// 	"fyne.io/fyne/v2/widget"
// 	"github.com/PaulWaldo/glimmer/api"
// 	"gopkg.in/masci/flickr.v3"
// 	"gopkg.in/masci/flickr.v3/photos"
// )

// const (
//
//	GridSizeWidth  = 500
//	GridSizeHeight = 500
//
// )
const (
	ImageWidth  = 200
	ImageHeight = 200
)

type contactPhotos struct {
	ma        *myApp
	container *fyne.Container
	title     *widget.Label
	photoList *widget.List
	photos    []api.Photo
	// photoCards []fyne.CanvasObject
}

func (p *contactPhotos) makeUI() *fyne.Container {
	p.title = widget.NewLabel("Contact Photos")

	// 	// Create cards for each photo
	// 	p.photoCards = make([]fyne.CanvasObject, len(p.photos))
	// 	for i, photo := range p.photos {
	// 		card := NewPhotoCard(photo, p.ma.client, func() {
	// 			pv := &photoView{ma: p.ma, photo: photo}
	// 			cont, err := pv.makeUI()
	// 			if err != nil {
	// 				fyne.LogError("parsing url", err)
	// 				return
	// 			}
	// 			p.ma.vs.Push(cont)
	// 		})
	// 		p.photoCards[i] = card
	// 	}

	p.photoList = widget.NewList(
		func() int {
			return len(p.photos)
		},
		func() fyne.CanvasObject {
			img := canvas.NewRectangle(color.Black)
			img.SetMinSize(fyne.Size{Width: ImageWidth, Height: ImageHeight})
			return NewPhotoCard("Title", "Author", img).makeUI()
			// return NewPhotoCard(api.Photo{}, p.ma.client, func() {
			// 	pv := &photoView{ma: p.ma}
			// 	cont, err := pv.makeUI()
			// 	if err != nil {
			// 		fyne.LogError("parsing url", err)
			// 		return
			// 	}
			// 	p.ma.vs.Push(cont)
			// })
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			cont := o.(*fyne.Container)
			infoCont := cont.Objects[1].(*fyne.Container)
			title := infoCont.Objects[0].(*widget.Label)
			author := infoCont.Objects[1].(*widget.Label)
			title.SetText(p.photos[i].Title)
			author.SetText(p.photos[i].Username)
		},
	)
	p.container = container.NewStack(p.photoList)
	return p.container

	// 	// gw := container.NewGridWrap(fyne.NewSize(GridSizeWidth, GridSizeHeight), p.photoCards...)
	// 	// scrollingGrid := container.NewScroll(gw)

	// 	// p.container = container.NewStack(
	// 	// 	container.NewStack(),
	// 	// 	container.NewBorder(p.title, nil, nil, nil, scrollingGrid),
	// 	// )
	// 	p.container = container.NewBorder(p.title, nil, nil, nil, p.photoList)
	// 	return p.container
	// }

	// type PhotoCard struct {
	// 	*widget.Card
	// 	info   photos.PhotoInfo
	// 	photo  api.Photo
	// 	client *flickr.FlickrClient
	// 	tap    func()
	// }

	// func NewPhotoCard(photo api.Photo /*content fyne.CanvasObject,*/, client *flickr.FlickrClient, onTapped func()) *PhotoCard {
	// 	clone := CloneClient(client)
	// 	i := &PhotoCard{tap: onTapped, photo: photo, client: clone}
	// 	i.Card = widget.NewCard(photo.Title, photo.Username, canvas.NewRectangle(color.Black))
	// 	i.ExtendBaseWidget(i)
	// 	go i.loadImage()
	// 	return i
	// }

	// func CloneClient(orig *flickr.FlickrClient) *flickr.FlickrClient {
	// 	clone := flickr.NewFlickrClient(orig.ApiKey, orig.ApiSecret)
	// 	clone.OAuthToken = orig.OAuthToken
	// 	clone.OAuthTokenSecret = orig.OAuthTokenSecret
	// 	return clone
	// }

	// func (c *PhotoCard) loadImage() {
	// 	resp, err := photos.GetInfo(c.client, c.photo.Id, c.photo.Secret)
	// 	if err != nil {
	// 		fyne.LogError("Failed to get photo info", err)
	// 		return
	// 	}
	// 	c.info = resp.Photo
	// 	photoUrl := fmt.Sprintf("https://live.staticflickr.com/%s/%s_%s_%s.jpg", c.info.Server, c.info.Id, c.info.Secret, "z")
	// 	uri, err := storage.ParseURI(photoUrl)
	// 	if err != nil {
	// 		fyne.LogError("parsing url", err)
	// 		c.Content = widget.NewLabel("Failed to load image")
	// 		return
	// 	}
	// 	fmt.Println("Downloading ", uri)
	// 	image := canvas.NewImageFromURI(uri)
	// 	fmt.Println("Got ", uri)
	// 	c.Content = image
	// 	image.FillMode = canvas.ImageFillContain
	// 	c.Refresh()
	// }

	// func (c *PhotoCard) Tapped(e *fyne.PointEvent) {
	// 	if c.tap == nil {
	// 		return
	// 	}
	// 	c.tap()
	// }

	// type photoList struct {
	// 	photos *[]api.Photo
	// }

	// const (
	// 	ImageWidth  = 200
	// 	ImageHeight = 200
	// )

	//	func (pl photoList) makeUI() fyne.Container {
	//		l := widget.NewList(
	//			func() int {
	//				return len(*pl.photos)
	//			},
	//			func() fyne.CanvasObject {
	//				img := canvas.NewRectangle(color.Black)
	//				img.SetMinSize(fyne.Size{Width: ImageWidth, Height: ImageHeight})
	//				return NewPhotoCard("Title", "Author", img).makeUI()
	//			},
	//			func(i widget.ListItemID, o fyne.CanvasObject) {
	//				card, ok := o.(*photoCard)
	//				if !ok {
	//					return
	//				}
	//				o.(*widget.Label).SetText((*pl.photos)[i].Title)
	//			},
	//		)
	//	}
}
