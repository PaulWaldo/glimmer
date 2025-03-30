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
	"gopkg.in/masci/flickr.v3/groups"
	"gopkg.in/masci/flickr.v3/photos"
)

// GroupPhotos represents a collection of photos from a Flickr group
type GroupPhotos struct {
	GroupName string
	Photos    []api.Photo
}

type GroupCard struct {
	widget.Card
	ma        *myApp
	GroupName string
}

type groupPhotosUI struct {
	ma         *myApp
	gridWrap   *fyne.Container
	groupCards []fyne.CanvasObject
	// cardByID allows access to the group card for a given group ID
	cardByID map[string]*fyne.CanvasObject
}

func (p *groupPhotosUI) makeUI() *fyne.Container {
	p.gridWrap = container.NewGridWrap(fyne.NewSize(GridSizeWidth, GridSizeHeight)) // Using same constants as contact photos
	scrollingGrid := container.NewScroll(p.gridWrap)
	return container.NewStack(scrollingGrid)
}

func (p *groupPhotosUI) setGroups(groups []groups.Group) {
	p.groupCards = make([]fyne.CanvasObject, 0, len(groups))
	p.cardByID = make(map[string]*fyne.CanvasObject)

	fmt.Println("User's group:")
	for _, group := range groups {
		fmt.Println(group.Name)
		card := &GroupCard{
			Card: widget.Card{
				Title: group.Name,
			},
			ma:        p.ma,
			GroupName: group.Name,
		}
		p.groupCards = append(p.groupCards, card)
		cardObj := fyne.CanvasObject(card)
		p.cardByID[group.Nsid] = &cardObj
	}

	// Update the grid with the new cards
	if p.gridWrap != nil {
		p.gridWrap.Objects = p.groupCards
		p.gridWrap.Refresh()
	}
}

// GroupPhotoCard represents a card displaying a photo from a group
type GroupPhotoCard struct {
	widget.Card
	info   photos.PhotoInfo
	photo  api.Photo
	client *flickr.FlickrClient
	tap    func()
}


// loadImage loads the image for a group photo card
func (c *GroupPhotoCard) loadImage() {
	// Load the image...
	resp, err := photos.GetInfo(c.client, c.photo.ID, c.photo.Secret)
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

	image := canvas.NewImageFromURI(uri)
	if image == nil || image.Resource == nil {
		panic("Image is nil")
	}
	image.FillMode = canvas.ImageFillContain
	c.SetContent(image)
}

// NewGroupPhotoCard creates a new photo card for group photos
func NewGroupPhotoCard(photo api.Photo, client *flickr.FlickrClient) *GroupPhotoCard {
	clone := api.CloneClient(client)
	i := &GroupPhotoCard{
		Card: widget.Card{
			Title:    photo.Title,
			Subtitle: photo.Username,
			Content:  widget.NewProgressBarInfinite(),
		},
		photo:  photo,
		client: clone,
	}
	i.ExtendBaseWidget(i)
	go func() {
		i.loadImage()  // Start loading the image in a goroutine
	}()
	return i
}
