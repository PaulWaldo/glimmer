package api

import (
	"gopkg.in/masci/flickr.v3"
)

type Group struct {
	Nsid         string `xml:"nsid,attr"`
	Name         string `xml:"name,attr"`
	Iconfarm     string `xml:"iconfarm,attr"`
	Iconserver   string `xml:"iconserver,attr"`
	Admin        string `xml:"admin,attr"`
	Eighteenplus string `xml:"eighteenplus,attr"`
	InvitationOnly string `xml:"invitation_only,attr"`
	Members      string `xml:"members,attr"`
	PoolCount    string `xml:"pool_count,attr"`
}

type Groups struct {
	Group []Group `xml:"group"`
}

type GetGroupsResponse struct {
	flickr.BasicResponse
	Groups Groups `xml:"groups"`
}

func GetGroups(client *flickr.FlickrClient, userId string, extras string) (*GetGroupsResponse, error) {
	client.Init()
	client.EndpointUrl = flickr.API_ENDPOINT

	client.Args.Set("method", "flickr.people.getGroups")
	client.Args.Set("user_id", userId)
	if extras != "" {
		client.Args.Set("extras", extras)
	}

	client.OAuthSign()
	response := &GetGroupsResponse{}
	err := flickr.DoGet(client, response)

	if err != nil {
		return nil, err
	}

	return response, nil
}
