package api_test

import (
	"testing"

	"github.com/PaulWaldo/glimmer/api"
	"github.com/PaulWaldo/glimmer/mocks"
	"gopkg.in/masci/flickr.v3"
)

func TestGetGroupPhotos(t *testing.T) {
	mockServer, client := flickr.FlickrMock(200, `
		<rsp stat="ok">
			<photos page="1" pages="1" perpage="100" total="1">
				<photo id="12345" owner="testuser" secret="abcdef" server="123" farm="1" title="Test Photo" ispublic="1" isfriend="0" isfamily="0" />
			</photos>
		</rsp>
	`, "text/xml")
	defer mockServer.Close()

	groupID := "12345"
	params := map[string]string{}

	response, err := api.GetGroupPhotos(client, groupID, params)
	if err != nil {
		t.Fatal(err)
	}

	if response.Photos.Page != 1 {
		t.Errorf("Expected page 1, got %d", response.Photos.Page)
	}

	if len(response.Photos.Photo) != 1 {
		t.Errorf("Expected 1 photo, got %d", len(response.Photos.Photo))
	}
}

func TestGetGroupPhotosError(t *testing.T) {
	mockServer, client := flickr.FlickrMock(500, "", "text/xml")
	defer mockServer.Close()

	groupID := "12345"
	params := map[string]string{}

	_, err := api.GetGroupPhotos(client, groupID, params)
	if err == nil {
		t.Fatal("Expected error, got nil")
	}
}

func TestGetGroupPhotosInvalidXML(t *testing.T) {
	mockServer, client := flickr.FlickrMock(200, " invalid xml ", "text/xml")
	defer mockServer.Close()

	groupID := "12345"
	params := map[string]string{}

	_, err := api.GetGroupPhotos(client, groupID, params)
	if err == nil {
		t.Fatal("Expected error, got nil")
	}
}
