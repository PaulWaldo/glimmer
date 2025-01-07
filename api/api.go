package api

import (
	"gopkg.in/masci/flickr.v3"
)

// ...

func GetContactList(client *flickr.FlickrClient) (*GetContactListResponse, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	// ...
}
