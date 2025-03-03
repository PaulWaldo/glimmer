package ui

import (
	"testing"

	"github.com/PaulWaldo/glimmer/api"
	"github.com/stretchr/testify/assert"
	"gopkg.in/masci/flickr.v3"
)

func TestNewPhotoCard(t *testing.T) {
	photo := api.Photo{
		ID:       "123",
		Owner:    "testuser",
		Secret:   "secret",
		Server:   "server",
		Farm:     "1",
		Title:    "Test Photo",
		Username: "testuser",
	}

	// Use a test canvas for the image loading
	// test.NewApp()
	// defer test.NewApp()

	client := flickr.GetTestClient() // Use the test client
	card := NewPhotoCard(photo, client, func() {})

	assert.NotNil(t, card.Card.Content, "PhotoCard.Card.Content should not be nil")
	assert.Equal(t, photo.Title, card.Title, "Card title should match photo title")
	assert.Equal(t, photo.Username, card.Subtitle, "Card subtitle should match photo username")
}
