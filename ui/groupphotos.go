package ui

import (
	"fmt"
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// Photo represents a photo from a Flickr group
type Photo struct {
	Title  string
	Author string
	URL    string
	Image  []byte // Raw image data after download
}

// GroupPhotos represents a collection of photos from a Flickr group
type GroupPhotos struct {
	GroupName string
	Photos    []Photo
}

type groupPhotos struct {
	ma                     *myApp
	gridWrap               *fyne.Container
	photoCards             []fyne.CanvasObject
	mutex                  sync.Mutex // For thread-safe operations
	batchSize              int        // Number of photos to load in each batch
	maxConcurrentDownloads int        // Maximum number of concurrent image downloads
}

func (p *groupPhotos) makeUI() *fyne.Container {
	p.gridWrap = container.NewGridWrap(fyne.NewSize(200, 200)) // Using hardcoded values for now
	scrollingGrid := container.NewScroll(p.gridWrap)

	// Set default values if not already set
	if p.batchSize == 0 {
		p.batchSize = 10 // Default batch size
	}
	if p.maxConcurrentDownloads == 0 {
		p.maxConcurrentDownloads = 5 // Default concurrent downloads
	}

	go func() {
		<-p.ma.groupPhotosChan // Wait for signal

		// Create a slice to hold all the new objects
		var newObjects []fyne.CanvasObject
		for _, group := range p.ma.usersGroups {
			groupCard := container.NewVBox(
				widget.NewLabel(group.Name),
				widget.NewButton("More...", func() {}),
			)

			// If we have photos for this group, add them to the card
			// Find the group photos for this group
			for _, groupPhotos := range p.ma.usersGroupPhotos {
				if groupPhotos.GroupName == group.Name {
					if len(groupPhotos.Photos) > 0 {
						// Get the first batch of photos
						batchSize := p.batchSize
						if batchSize > len(groupPhotos.Photos) {
							batchSize = len(groupPhotos.Photos)
						}

						// Convert api.Photo to our Photo type
						var firstBatch []Photo
						for i := 0; i < batchSize; i++ {
							apiPhoto := groupPhotos.Photos[i]
							photo := Photo{
								Title:  apiPhoto.Title,
								Author: apiPhoto.Username,
								URL: fmt.Sprintf("https://live.staticflickr.com/%s/%s_%s.jpg",
									apiPhoto.Server, apiPhoto.ID, apiPhoto.Secret),
							}
							firstBatch = append(firstBatch, photo)
						}

						// Add photos to the group card
						p.addPhotosToGroupCard(groupCard, firstBatch)
					}
					break
				}
			}

			newObjects = append(newObjects, groupCard)
		}

		// Update the UI in a single operation
		p.gridWrap.Objects = newObjects
		p.gridWrap.Refresh()
	}()

	return container.NewStack(scrollingGrid)
}

// createPhotoCard creates a card for a single photo
func (p *groupPhotos) createPhotoCard(photo Photo) fyne.CanvasObject {
	// Create labels for title and author
	titleLabel := widget.NewLabel(photo.Title)
	authorLabel := widget.NewLabel(fmt.Sprintf("By: %s", photo.Author))

	// Create a placeholder for the image
	imgPlaceholder := widget.NewIcon(nil)

	// Create the photo card container
	photoCard := container.NewVBox(
		titleLabel,
		authorLabel,
		imgPlaceholder,
	)

	// Start downloading the image in the background
	go p.downloadAndSetImage(photo.URL, imgPlaceholder)

	return photoCard
}

// addPhotosToGroupCard adds photo cards to a group card
func (p *groupPhotos) addPhotosToGroupCard(groupCard *fyne.Container, photos []Photo) {
	// We need to insert the photo cards before the "More..." button
	// which is the last element in the group card

	// Create photo cards
	var photoCards []fyne.CanvasObject
	for _, photo := range photos {
		photoCard := p.createPhotoCard(photo)
		photoCards = append(photoCards, photoCard)
	}

	// Insert photo cards before the "More..." button
	newObjects := make([]fyne.CanvasObject, 0, len(groupCard.Objects)+len(photoCards))
	newObjects = append(newObjects, groupCard.Objects[0]) // Group name label
	newObjects = append(newObjects, photoCards...)
	newObjects = append(newObjects, groupCard.Objects[len(groupCard.Objects)-1]) // "More..." button

	// Update the group card
	groupCard.Objects = newObjects
	groupCard.Refresh()
}

// downloadAndSetImage downloads an image from a URL and sets it to the provided image widget
func (p *groupPhotos) downloadAndSetImage(url string, imgWidget *widget.Icon) {
	// TODO: Implement actual image downloading logic
	// This would involve:
	// 1. Making an HTTP request to download the image
	// 2. Processing the image data
	// 3. Creating a resource from the image data
	// 4. Setting the resource to the image widget

	// For now, we'll just use a placeholder
	// In a real implementation, you would use a semaphore to limit concurrent downloads
	// and handle errors appropriately
}
