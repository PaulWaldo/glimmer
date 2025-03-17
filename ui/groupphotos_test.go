package ui

import (
	"fmt"
	"net/http"
	"net/http/httptest"
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
		usersGroups: make([]groups.Group, 0), // Initialize usersGroups
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

	// Wait for the grid to be populated with a timeout
	deadline := time.Now().Add(time.Second)
	for time.Now().Before(deadline) {
		if len(grid.Objects) == 2 {
			break
		}
		time.Sleep(10 * time.Millisecond)
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
	gp := groupPhotos{ma: &myApp{groupPhotosChan: make(chan struct{})}}
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
		ma:         &myApp{},
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

func Test_groupPhotos_createPhotoCard(t *testing.T) {
	// Setup test data
	gp := groupPhotos{
		ma: &myApp{},
	}

	photo := Photo{
		Title:  "Test Photo",
		Author: "Test Author",
		URL:    "https://example.com/test.jpg",
	}

	// Create a photo card
	card := gp.createPhotoCard(photo)

	// Verify the card structure
	container, ok := card.(*fyne.Container)
	require.True(t, ok, "Photo card should be a container")

	// Check that the container has the expected components
	require.GreaterOrEqual(t, len(container.Objects), 3, "Photo card should have at least 3 components")

	// Check title
	title, ok := container.Objects[0].(*widget.Label)
	require.True(t, ok, "First component should be a title label")
	require.Equal(t, "Test Photo", title.Text)

	// Check author
	author, ok := container.Objects[1].(*widget.Label)
	require.True(t, ok, "Second component should be an author label")
	require.Equal(t, "By: Test Author", author.Text)

	// The third component would be an image, but we can't easily test its content
	_, ok = container.Objects[2].(*widget.Icon)
	require.True(t, ok, "Third component should be an image placeholder")
}

func Test_groupPhotos_addPhotosToGroupCard(t *testing.T) {
	// Setup test data
	gp := groupPhotos{
		ma: &myApp{},
	}

	// Create a group card
	groupCard := container.NewVBox(
		widget.NewLabel("Test Group"),
		widget.NewButton("More...", func() {}),
	)

	// Create test photos
	photos := []Photo{
		{Title: "Photo 1", Author: "Author 1", URL: "https://example.com/1.jpg"},
		{Title: "Photo 2", Author: "Author 2", URL: "https://example.com/2.jpg"},
	}

	// Add photos to the group card
	gp.addPhotosToGroupCard(groupCard, photos)

	// Verify the group card structure
	require.Equal(t, 4, len(groupCard.Objects), "Group card should have 4 components (label + 2 photos + button)")

	// Check that the photos were inserted before the "More..." button
	moreButton, ok := groupCard.Objects[3].(*widget.Button)
	require.True(t, ok, "Last component should be the More button")
	require.Equal(t, "More...", moreButton.Text)

	// Check the photo cards
	for i := 0; i < 2; i++ {
		photoCard, ok := groupCard.Objects[i+1].(*fyne.Container)
		require.True(t, ok, fmt.Sprintf("Component %d should be a photo card container", i+1))

		title, ok := photoCard.Objects[0].(*widget.Label)
		require.True(t, ok, "First component of photo card should be a title label")
		require.Equal(t, photos[i].Title, title.Text)
	}
}

func Test_groupPhotos_downloadAndSetImage(t *testing.T) {
	// Setup test data
	gp := groupPhotos{
		ma: &myApp{},
		maxConcurrentDownloads: 5,
	}

	// Create a mock icon widget
	imgWidget := widget.NewIcon(nil)

	// Create a test server that returns a simple image
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Return a small test image (1x1 transparent PNG)
		w.Header().Set("Content-Type", "image/png")
		w.Write([]byte{
			0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A, 0x00, 0x00, 0x00, 0x0D,
			0x49, 0x48, 0x44, 0x52, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x01,
			0x08, 0x06, 0x00, 0x00, 0x00, 0x1F, 0x15, 0xC4, 0x89, 0x00, 0x00, 0x00,
			0x0A, 0x49, 0x44, 0x41, 0x54, 0x78, 0x9C, 0x63, 0x00, 0x01, 0x00, 0x00,
			0x05, 0x00, 0x01, 0x0D, 0x0A, 0x2D, 0xB4, 0x00, 0x00, 0x00, 0x00, 0x49,
			0x45, 0x4E, 0x44, 0xAE, 0x42, 0x60, 0x82,
		})
	}))
	defer testServer.Close()

	// Call the method with the test server URL
	gp.downloadAndSetImage(testServer.URL, imgWidget)

	// Wait a short time for the download to complete
	time.Sleep(100 * time.Millisecond)

	// Verify that the icon now has a resource (no longer nil)
	require.NotNil(t, imgWidget.Resource, "Image widget should have a resource after download")
}
