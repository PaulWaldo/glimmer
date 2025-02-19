package ui

import (
	"testing"

	"fyne.io/fyne/v2/container"
	"github.com/stretchr/testify/require"
)

func Test_groupPhotos_makeUI_containsEmptyGrid(t *testing.T) {
	gp := groupPhotos{}
	cont := gp.makeUI()
	scroll := cont.Objects[0].(*container.Scroll)
	gw := scroll.Content
	require.Equal(t, gw, gp.gridWrap)
}
