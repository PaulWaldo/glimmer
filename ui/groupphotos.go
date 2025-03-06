package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
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
	p.gridWrap = container.NewGridWrap(fyne.NewSize(GridSizeWidth, GridSizeHeight)) // Initialize empty gridWrap
	scrollingGrid := container.NewScroll(container.NewVBox())                       // Start with empty VBox in Scroll

	go func() {
		<-p.ma.groupPhotosChan // Wait for signal
		for _, group := range p.ma.usersGroups {
			groupCard := container.NewVBox(widget.NewLabel(group.Name), widget.NewButton("More...", func() {}))
			p.gridWrap.Objects = append(p.gridWrap.Objects, groupCard) // Add group card to gridWrap
			// p.ma.window.Canvas().Refresh(p.ma.window)                  // Refresh UI to show changes
		}
		scrollingGrid.Content = p.gridWrap // Set gridWrap as content after it's populated
	}()

	return container.NewStack(scrollingGrid)
}
