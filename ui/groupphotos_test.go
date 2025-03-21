package ui

import (
	"testing"

	"fyne.io/fyne/v2"
	"github.com/PaulWaldo/glimmer/api"
	"github.com/stretchr/testify/assert" // Added for assertions
	"gopkg.in/masci/flickr.v3"
	"gopkg.in/masci/flickr.v3/groups"
)

func TestSetGroups(t *testing.T) {
	p := &groupPhotosUI{
		ma:       &myApp{}, // Initialize myApp
		cardByID: make(map[string]*fyne.CanvasObject),
	}
	testGroups := []groups.Group{
		{Nsid: "1", Name: "Group 1"},
		{Nsid: "2", Name: "Group 2"},
	}

	p.setGroups(testGroups)

	assert.Equal(t, len(testGroups), len(p.groupCards))
	for _, group := range testGroups {
		card := (*p.cardByID[group.Nsid]).(*GroupCard) // Type assertion
		assert.NotNil(t, card)
		assert.Equal(t, group.Name, card.Title)
	}
}

// TestNewGroupPhotoCard tests the creation of a photo card for group photos
func TestNewGroupPhotoCard(t *testing.T) {
	// Use the Flickr test client
	client := flickr.GetTestClient()
	
	photo := api.Photo{
		ID:       "12345",
		Owner:    "owner123",
		Secret:   "secret123",
		Server:   "server123",
		Username: "testuser",
		Title:    "Test Photo",
	}
	
	photoCard := NewGroupPhotoCard(photo, client)
	
	assert.NotNil(t, photoCard)
	assert.Equal(t, photo.Title, photoCard.Title)
	assert.Equal(t, photo.Username, photoCard.Subtitle)
	assert.Equal(t, photo, photoCard.photo)
}
