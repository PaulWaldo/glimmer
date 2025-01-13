package api

import (
	"gopkg.in/masci/flickr.v3"
)

type GetGroupPhotosResponse struct {
	flickr.BasicResponse
	Photos struct {
		Page    int     `xml:"page,attr"`
		Pages   int     `xml:"pages,attr"`
		PerPage int     `xml:"perpage,attr"`
		Total   int     `xml:"total,attr"`
		Photo   []Photo `xml:"photo"`
	} `xml:"photos"`
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
	Groups struct {
		Group []flickr.Group `xml:"group"`
	} `xml:"groups"`
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
