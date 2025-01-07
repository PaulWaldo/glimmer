package api

import (
	"fmt"
	"gopkg.in/masci/flickr.v3"
)

// ...

func GetContactPhotos(client *flickr.FlickrClient, page int) (*GetContactPhotosResponse, error) {
	if page < 1 {
		return nil, errors.New("page must be greater than 0")
	}
	// ...
}
