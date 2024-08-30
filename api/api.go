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

// <photos>
//   <photo id="2801" owner="12037949629@N01" secret="123456" server="1" username="Eric is the best" title="grease" />
//   <photo id="2499" owner="33853651809@N01" secret="123456" server="1" username="cal18" title="36679_o" />
//   <photo id="2437" owner="12037951898@N01" secret="123456" server="1" username="georgie parker" title="phoenix9_stewart" />
// </photos>

type Photo struct {
	Id       string `xml:"id,attr"`
	Owner    string `xml:"owner,attr"`
	Secret   string `xml:"secret,attr"`
	Server   string `xml:"server,attr"`
	Username string `xml:"username,attr"`
	Title    string `xml:"title,attr"`
}

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

	client.OAuthSign()
	response := &GetContactPhotosResponse{}
	err := flickr.DoGet(client, response)

	if err != nil {
		return nil, err
	}

	return response, nil
}
