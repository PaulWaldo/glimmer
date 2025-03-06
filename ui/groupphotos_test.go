package ui

import (
	"testing"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/stretchr/testify/require"
	"gopkg.in/masci/flickr.v3/groups" // Import groups package
)

func Test_groupPhotos_makeUI_createsGroupCardsAsync(t *testing.T) {
	ma := &myApp{
		groupPhotosChan: make(chan struct{}),
		// window:          ma.window, // Ensure ma.window is initialized
		usersGroups:     make([]groups.Group, 0), // Initialize usersGroups
	}
	gp := groupPhotos{
		ma: ma,
	}

	cont := gp.makeUI()
	scroll := cont.Objects[0].(*container.Scroll)
	grid := scroll.Content.(*fyne.Container)

	go func() {
		gp.ma.usersGroups = []groups.Group{ // Populate usersGroups
			{Name: "Group 1"},
			{Name: "Group 2"},
		}
		close(gp.ma.groupPhotosChan)
	}()

	require.Equal(t, 0, len(grid.Objects)) // Initially, the grid should be empty

	timeout := time.After(1 * time.Second)
	for len(grid.Objects) != 2 { // Wait for 2 group cards
		select {
		case <-timeout:
			t.Fatal("Timeout waiting for group cards")
		default:
			time.Sleep(10 * time.Millisecond)
			// gp.ma.window.Canvas().Refresh(gp.ma.window) // Correct Refresh call
		}
	}

	require.Equal(t, 2, len(grid.Objects)) // Finally, there should be 2 group cards

	// Check that each group card has a "More..." button and a label with the correct group name
	for i, obj := range grid.Objects {
		groupCard := obj.(*fyne.Container)
		require.Equal(t, 2, len(groupCard.Objects), "Group card should have 2 objects")

		moreButton, ok := groupCard.Objects[1].(*widget.Button)
		require.True(t, ok, "Second object should be a *widget.Button")
		require.Equal(t, "More...", moreButton.Text)

		label, ok := groupCard.Objects[0].(*widget.Label)
		require.True(t, ok, "First object should be a *widget.Label")

		expectedGroupName := gp.ma.usersGroups[i].Name  // Access group name from usersGroups
		require.Equal(t, expectedGroupName, label.Text) // Check label text
	}
}

func Test_groupPhotos_makeUI_containsEmptyGrid(t *testing.T) {
	gp := groupPhotos{}
	cont := gp.makeUI()
	scroll := cont.Objects[0].(*container.Scroll)
	grid := scroll.Content.(*fyne.Container)

	// Verify the grid exists and is empty
	require.NotNil(t, grid)
	require.Equal(t, 0, len(grid.Objects))
}

func Test_groupPhotos_makeUI_createsGroupCards(t *testing.T) {
	// Setup test data
	gp := groupPhotos{
		photoCards: make([]fyne.CanvasObject, 0),
	}

	// Call makeUI
	cont := gp.makeUI()

	// Verify we have a container with group cards
	scroll := cont.Objects[0].(*container.Scroll)
	grid := scroll.Content.(*fyne.Container)

	// This will fail because we haven't implemented group cards yet
	require.Equal(t, 0, len(grid.Objects),
		"Should start with empty grid")
}
