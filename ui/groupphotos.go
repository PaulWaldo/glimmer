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

	fmt.Println("User's groups:")
	for _, group := range groups {
		fmt.Println(group.Name)

		// Create a container for the group's photos
		photoContainer := container.NewGridWrap(fyne.NewSize(150, 150))

		// Get photos for this group
		if groupPhotos, ok := p.ma.usersGroupPhotos[group.Nsid]; ok {
			// Add up to 4 photos initially (or all if less than 4)
			photosToShow := len(groupPhotos.Photos)
			if photosToShow > 4 {
				photosToShow = 4
			}

			for i := 0; i < photosToShow; i++ {
				if i < len(groupPhotos.Photos) {
					photoCard := NewGroupPhotoCard(groupPhotos.Photos[i], group.Nsid, p.ma.client)
					photoContainer.Add(photoCard)
				}
			}

			// Add "More..." button if there are more photos
			if len(groupPhotos.Photos) > 4 {
				moreButton := widget.NewButton("More...", func() {
					// This will be implemented in Story 5
					fmt.Println("More button clicked for group:", group.Name)
				})
				photoContainer.Add(moreButton)
			}
		}

		// Create the group card with the photo container
		card := &GroupCard{
			Card: widget.Card{
				Title: group.Name,
				Content: container.NewVBox(
					photoContainer,
					widget.NewButton("Collapse", func() {
						// This will be implemented in Story 6
						fmt.Println("Collapse button clicked for group:", group.Name)
					}),
				),
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
	photo  api.Photo
	client *flickr.FlickrClient
	info   photos.PhotoInfo
}

// loadImage loads the image for a group photo card
func (c *GroupPhotoCard) loadImage() {
	// Get photo info
	resp, err := photos.GetInfo(c.client, c.photo.ID, c.photo.Secret)
	if err != nil {
		fyne.LogError("Failed to get photo info", err)
		c.SetContent(widget.NewLabel("Failed to load image"))
		return
	}

	c.info = resp.Photo

	// Construct the photo URL
	photoUrl := fmt.Sprintf("https://live.staticflickr.com/%s/%s_%s_%s.jpg",
		c.info.Server, c.info.Id, c.info.Secret, "z")

	// Parse the URL
	uri, err := storage.ParseURI(photoUrl)
	if err != nil {
		fyne.LogError("parsing url", err)
		c.SetContent(widget.NewLabel("Failed to parse URL"))
		return
	}

	// Create and set the image
	image := canvas.NewImageFromURI(uri)
	image.FillMode = canvas.ImageFillContain
	c.SetContent(image)
}

// NewGroupPhotoCard creates a new photo card for group photos
func NewGroupPhotoCard(photo api.Photo, groupID string, client *flickr.FlickrClient) *GroupPhotoCard {
	clone := api.CloneClient(client)
	card := &GroupPhotoCard{
		Card: widget.Card{
			// Title:    photo.Title,
			// Subtitle: photo.Username,
			Content: widget.NewProgressBarInfinite(),
		},
		photo:  photo,
		client: clone,
	}
	card.ExtendBaseWidget(card)

	// Start loading the image asynchronously
	go card.loadImage()

	return card
}
