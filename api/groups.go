package api

import (
	"fmt"

	"gopkg.in/masci/flickr.v3"
	"gopkg.in/masci/flickr.v3/groups"
)

type GetGroupPhotosResponse struct {
	flickr.BasicResponse
	Photos []Photo `xml:"photos>photo"`
}

func GetGroupPhotos(client *flickr.FlickrClient, groupID string, params map[string]string) (*GetGroupPhotosResponse, error) {
	client.Init()
	client.EndpointUrl = flickr.API_ENDPOINT
	client.Args.Set("method", "flickr.groups.pools.getPhotos")
	client.Args.Set("group_id", groupID)
	for k, v := range params {
		client.Args.Set(k, v)
	}
	client.OAuthSign()

	response := &GetGroupPhotosResponse{}
	err := flickr.DoPost(client, response)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Got %d group photos\n", len(response.Photos))
	return response, nil
}

type GetUserGroupsResponse struct {
	flickr.BasicResponse
	Groups []groups.Group `xml:"groups>group"`
}

func GetUserGroups(client *flickr.FlickrClient, userID string, params map[string]string) (*GetUserGroupsResponse, error) {
	client.Init()
	client.EndpointUrl = flickr.API_ENDPOINT
	client.Args.Set("method", "flickr.people.getGroups")
	client.Args.Set("user_id", userID)
	for k, v := range params {
		client.Args.Set(k, v)
	}
	client.OAuthSign()

	var response GetUserGroupsResponse
	err := flickr.DoGet(client, &response)
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

func GetUsersGroupPhotos(client *flickr.FlickrClient, userID string, params map[string]string, groups *[]groups.Group, photos *[]UsersGroupPhotos) error {
	// Find all the groups the user belongs too
	fmt.Println("Getting User Groups")
	clonedClient := CloneClient(client)
	userGroups, err := GetUserGroups(clonedClient, userID, params)
	if err != nil {
		return err
	}
	fmt.Println("Done with User Groups")

	// For each group, get photos for the group
	var usersGroupPhotos = make([]UsersGroupPhotos, len(userGroups.Groups))
	fmt.Printf("Getting %d user groups\n", len(userGroups.Groups))
	*groups = userGroups.Groups
	for i, group := range userGroups.Groups {
		fmt.Println(i)
		clonedClient = CloneClient(client)
		groupPhotos, err := GetGroupPhotos(clonedClient, group.Nsid, params)
		if err != nil {
			return err
		}
		usersGroupPhotos[i] = UsersGroupPhotos{
			GroupID:   group.Nsid,
			GroupName: group.Name,
			Photos:    groupPhotos.Photos,
		}
		*photos = usersGroupPhotos
	}
	fmt.Println("Done with group photos")

	return nil
}
