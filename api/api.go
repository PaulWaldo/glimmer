package api

import (
	"gopkg.in/masci/flickr.v3"
)

type Contact struct {
	// <contact nsid="12037949629@N01" username="Eric" iconserver="1" realname="Eric Costello" friend="1" family="0" ignored="1" />
	Nsid       string `xml:"nsid,attr"`
	UserName   string `xml:"username,attr"`
	IconServer string `xml:"iconserver,attr"`
	RealName   string `xml:"realname,attr"`
	Friend     string `xml:"friend,attr"`
	Family     string `xml:"family,attr"`
	Ignored    string `xml:"ignored,attr"`
}

type Contacts struct {
	Page    int       `xml:"page,attr"`
	Pages   int       `xml:"pages,attr"`
	PerPage int       `xml:"perpage,attr"`
	Total   int       `xml:"total,attr"`
	Contact []Contact `xml:"contact"`
}

type GetContactListResponse struct {
	flickr.BasicResponse
	Contacts Contacts `xml:"contacts"`
}

func GetContactList(client *flickr.FlickrClient) (*GetContactListResponse, error) {
	client.Init()
	client.EndpointUrl = flickr.API_ENDPOINT

	client.Args.Set("method", "flickr.contacts.getList")
	// client.Args.Set("brand", "nikon")

	client.OAuthSign()
	response := &GetContactListResponse{}
	err := flickr.DoGet(client, response)

	if err != nil {
		return nil, err
	}

	return response, nil
}
