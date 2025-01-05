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
	// <contacts page="1" pages="1" perpage="1000" total="3">
	//   <contact nsid="12037949629@N01" username="Eric" iconserver="1" realname="Eric Costello" friend="1" family="0" ignored="1" />
	//   <contact nsid="12037949631@N01" username="neb" iconserver="1" realname="Ben Cerveny" friend="0" family="0" ignored="0" />
	//   <contact nsid="41578656547@N01" username="cal_abc" iconserver="1" realname="Cal Henderson" friend="1" family="1" ignored="0" />
	// </contacts>
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

	client.OAuthSign()
	response := &GetContactListResponse{}
	err := flickr.DoGet(client, response)

	if err != nil {
		return nil, err
	}

	return response, nil
}

type Photo struct {
	ID          string `xml:"id,attr"`
	Owner       string `xml:"owner,attr"`
	Secret      string `xml:"secret,attr"`
	Server      string `xml:"server,attr"`
	Username    string `xml:"username,attr"`
	Title       string `xml:"title,attr"`
	Farm        string `xml:"farm,attr"`
	IsPublic    int    `xml:"ispublic,attr"`
	IsFriend    int    `xml:"isfriend,attr"`
	IsFamily    int    `xml:"isfamily,attr"`
	Description string `xml:"description"`
	Dates       struct {
		Posted           int    `xml:"posted,attr"`
		Taken            string `xml:"taken,attr"`
		TakenGranularity int    `xml:"takengranularity,attr"`
	} `xml:"dates"`
	Views    int `xml:"views"`
	Comments int `xml:"comments"`
}

func Feed(client *flickr.FlickrClient) (*GetContactPhotosResponse, error) {
	client.Init()
	client.EndpointUrl = flickr.API_ENDPOINT

	client.Args.Set("method", "flickr.feed.getFeed")

	client.OAuthSign()
	response := &GetContactPhotosResponse{}
	err := flickr.DoGet(client, response)

	if err != nil {
		return nil, err
	}

	return response, nil
}
