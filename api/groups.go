package api

import (
	"gopkg.in/masci/flickr.v3"
	"gopkg.in/masci/flickr.v3/groups"
)

type GetGroupPhotosResponse struct {
	flickr.BasicResponse
	Photos []Photo `xml:"photos>photo"`
}

func GetGroupPhotos(client *flickr.FlickrClient, groupID string, params map[string]string) (*GetGroupPhotosResponse, error) {
	client.Init()
	client.Args.Set("method", "flickr.groups.pools.getPhotos")
	client.Args.Set("group_id", groupID)
	for k, v := range params {
		client.Args.Set(k, v)
	}
	client.ApiSign()

	var response GetGroupPhotosResponse
	err := flickr.DoPost(client, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

type GetUserGroupsResponse struct {
	flickr.BasicResponse
	Groups []groups.Group `xml:"groups>group"`
}

func GetUserGroups(client *flickr.FlickrClient, userID string, params map[string]string) (*GetUserGroupsResponse, error) {
	client.Init()
	client.Args.Set("method", "flickr.people.getGroups")
	client.Args.Set("user_id", userID)
	for k, v := range params {
		client.Args.Set(k, v)
	}
	client.ApiSign()

	var response GetUserGroupsResponse
	err := flickr.DoPost(client, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
