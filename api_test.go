package glimmer

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/masci/flickr.v3"
)

var contactListBody = `<?xml version="1.0" encoding="utf-8"?>
	<rsp stat="ok">
	<contacts page="1" pages="1" perpage="1000" total="3">
		<contact nsid="12037949629@N01" username="Eric" iconserver="1" realname="Eric Costello"
		friend="1" family="0" ignored="1" />
		<contact nsid="12037949631@N01" username="neb" iconserver="1" realname="Ben Cerveny" friend="0"
		family="0" ignored="0" />
		<contact nsid="41578656547@N01" username="cal_abc" iconserver="1" realname="Cal Henderson"
		friend="1" family="1" ignored="0" />
	</contacts>
	</rsp>`

func TestGetContactList(t *testing.T) {
	fclient := flickr.GetTestClient()
	server, client := flickr.FlickrMock(200, contactListBody, "text/xml")
	defer server.Close()
	fclient.HTTPClient = client

	resp, err := GetContactList(fclient)

	assert.NoError(t, err)
	assert.Equal(t, 3, len(resp.Contacts.Contact), "Expected 3 contacts, but got %d", len(resp.Contacts.Contact))

	// flickr.Expect(t, err, nil)
	// flickr.Expect(t, resp.Photos.Page, 2)
	// flickr.Expect(t, resp.Photos.Pages, 89)
	// flickr.Expect(t, resp.Photos.PerPage, 10)
	// flickr.Expect(t, resp.Photos.Total, 881)
	// flickr.Expect(t, len(resp.Photos.Photos), 4)
	// flickr.Expect(t, resp.Photos.Photos[0], Photo{
	// 	Id:       "2636",
	// 	Owner:    "47058503995@N01",
	// 	Secret:   "a123456",
	// 	Server:   "2",
	// 	Title:    "test_04",
	// 	IsPublic: true,
	// 	IsFriend: false,
	// 	IsFamily: false,
	// })
	// flickr.Expect(t, resp.Photos.Photos[3], Photo{
	// 	Id:       "2610",
	// 	Owner:    "12037949754@N01",
	// 	Secret:   "d123456",
	// 	Server:   "2",
	// 	Title:    "00_tall",
	// 	IsPublic: true,
	// 	IsFriend: false,
	// 	IsFamily: false,
	// })
}
