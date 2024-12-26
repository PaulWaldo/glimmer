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

type FlickrClientInterface interface {
	DoGet(response interface{}) error
	Init()
	OAuthSign()
	Args() map[string]string
	SetArg(key, value string)
	EndpointUrl() string
	SetEndpointUrl(url string)
}

func GetGroups(client FlickrClientInterface, userId string, extras string) (*GetGroupsResponse, error) {
	client.Init()
	client.SetEndpointUrl(flickr.API_ENDPOINT)

	client.SetArg("method", "flickr.people.getGroups")
	client.SetArg("user_id", userId)
	if extras != "" {
		client.SetArg("extras", extras)
	}

	client.OAuthSign()
	response := &GetGroupsResponse{}
	err := client.DoGet(response)

	if err != nil {
		return nil, err
	}

	return response, nil
}
