package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

type groupPhotos struct {
	ma *myApp
	// container *fyne.Container
	// photoList  *fyne.Container
	gridWrap *fyne.Container
	// photos     []api.Photo
	photoCards []fyne.CanvasObject
	// page       int
	// totalPages int
}

func (p *groupPhotos) makeUI() *fyne.Container {
	p.gridWrap = container.NewGridWrap(fyne.NewSize(GridSizeWidth, GridSizeHeight), p.photoCards...)
	scrollingGrid := container.NewScroll(p.gridWrap)
	// scrollingGrid.OnScrolled = func(pos fyne.Position) {
	// 	if pos.Y+scrollingGrid.Size().Height >= p.gridWrap.Size().Height {
	// 		p.loadNextPage()
	// 	}
	// }

	// p.loadNextPage()
	return container.NewStack(scrollingGrid)
}
