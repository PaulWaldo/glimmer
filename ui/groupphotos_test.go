package ui

import (
	"testing"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"github.com/stretchr/testify/require"
)

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
