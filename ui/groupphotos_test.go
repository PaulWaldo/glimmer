package ui

import (
	"testing"

	"fyne.io/fyne/v2"
	"github.com/stretchr/testify/assert" // Added for assertions
	"gopkg.in/masci/flickr.v3/groups"
)

func TestSetGroups(t *testing.T) {
	p := &groupPhotosUI{
		ma:       &myApp{}, // Initialize myApp
		cardByID: make(map[string]*fyne.CanvasObject),
	}
	testGroups := []groups.Group{
		{ID: "1", Name: "Group 1"},
		{ID: "2", Name: "Group 2"},
	}

	p.setGroups(testGroups)

	assert.Equal(t, len(testGroups), len(p.groupCards))
	for _, group := range testGroups {
		card := (*p.cardByID[group.ID]).(*GroupCard) // Type assertion
		assert.NotNil(t, card)
		assert.Equal(t, group.Name, card.Title)
	}
}
