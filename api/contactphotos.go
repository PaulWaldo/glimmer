package api

import (
	"gopkg.in/masci/flickr.v3"
)

type ContactPhotos struct {
	Page    int     `xml:"page,attr"`
	Pages   int     `xml:"pages,attr"`
	PerPage int     `xml:"perpage,attr"`
	Total   int     `xml:"total,attr"`
	Photos  []Photo `xml:"photo"`
}

type GetContactPhotosResponse struct {
	flickr.BasicResponse
	Photos ContactPhotos `xml:"photos"`
}

func GetContactPhotos(client *flickr.FlickrClient) (*GetContactPhotosResponse, error) {
	client.Init()
	client.EndpointUrl = flickr.API_ENDPOINT

	client.Args.Set("method", "flickr.photos.getContactsPhotos")
	client.Args.Set("per_page", "25")

	client.OAuthSign()
	response := &GetContactPhotosResponse{}
	err := flickr.DoGet(client, response)

	if err != nil {
		return nil, err
	}

	return response, nil
}
