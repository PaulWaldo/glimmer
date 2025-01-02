package api

import (
	"github.com/your/flickr"
)

type GetPhotosResponse struct {
	// Add fields for the response here
}

func GetPhotos(client *flickr.FlickrClient, groupID string, params map[string]string) (*GetPhotosResponse, error) {
    client.Init()
    client.Args.Set("method", "flickr.groups.pools.getPhotos")
    client.Args.Set("group_id", groupID)
    for k, v := range params {
        client.Args.Set(k, v)
    }
    client.ApiSign()

    var response GetPhotosResponse
    err := flickr.DoPost(client, &response)
    if err != nil {
        return nil, err
    }

    return &response, nil
}
