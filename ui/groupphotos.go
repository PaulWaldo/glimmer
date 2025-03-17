package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/PaulWaldo/glimmer/api"
	"gopkg.in/masci/flickr.v3/groups"
)

// GroupPhotos represents a collection of photos from a Flickr group
type GroupPhotos struct {
	GroupName string
	Photos    []api.Photo
}

type GroupCard struct {
	widget.Card
	ma *myApp
}

type groupPhotosUI struct {
	ma         *myApp
	gridWrap   *fyne.Container
	groupCards []fyne.CanvasObject
	// cardByID allows access to the group card for a given group ID
	cardByID map[string]*fyne.CanvasObject
}

func (p *groupPhotosUI) makeUI() *fyne.Container {
	p.gridWrap = container.NewGridWrap(fyne.NewSize(200, 200)) // Using hardcoded values for now
	scrollingGrid := container.NewScroll(p.gridWrap)
	return container.NewStack(scrollingGrid)
}

func (p *groupPhotosUI) setGroups(groups []groups.Group) {
	// TODO: for each group,
	// * create a group card, assigning the group name to the title
	// * append it to p.groupCards
	// * add the mapping of the group's ID to p.cardById

}

// createPhotoCard creates a card for a single photo
func (p *groupPhotosUI) createPhotoCard(photo api.Photo) *PhotoCard {
	return nil
}

