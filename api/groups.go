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

type UsersGroupPhotos struct {
	GroupID   string
	GroupName string
	Photos    []Photo
}

func GetUsersGroupPhotos(client *flickr.FlickrClient, userID string) ([]UsersGroupPhotos, error) {
	userGroups, err := GetUserGroups(client, userID, nil)
	if err != nil {
		return nil, err
	}

	var usersGroupPhotos []UsersGroupPhotos
	for _, group := range userGroups.Groups {
		groupPhotos, err := GetGroupPhotos(client, group.NSID, nil)
		if err != nil {
			return nil, err
		}
		usersGroupPhotos = append(usersGroupPhotos, UsersGroupPhotos{
			GroupID:   group.NSID,
			GroupName: group.Name,
			Photos:    groupPhotos.Photos,
		})
	}

	return usersGroupPhotos, nil
}
