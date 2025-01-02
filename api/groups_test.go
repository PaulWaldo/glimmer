package api

import (
	"testing"
	"github.com/PaulWaldo/glimmer/mocks"
)

func TestGetGroupPhotos(t *testing.T) {
	mockServer, client := mocks.NewFlickrMock(200, `
		<rsp stat="ok">
			<photos page="1" pages="1" perpage="100" total="1">
				<photo id="12345" owner="testuser" secret="abcdef" server="123" farm="1" title="Test Photo" ispublic="1" isfriend="0" isfamily="0" />
			</photos>
		</rsp>
	`)
	defer mockServer.Close()

	groupID := "12345"
	params := map[string]string{}

	response, err := GetGroupPhotos(client, groupID, params)
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
	mockServer, client := mocks.NewFlickrMock(500, "")
	defer mockServer.Close()

	groupID := "12345"
	params := map[string]string{}

	_, err := GetGroupPhotos(client, groupID, params)
	if err == nil {
		t.Fatal("Expected error, got nil")
	}
}

func TestGetGroupPhotosInvalidXML(t *testing.T) {
	mockServer, client := mocks.NewFlickrMock(200, " invalid xml ")
	defer mockServer.Close()

	groupID := "12345"
	params := map[string]string{}

	_, err := GetGroupPhotos(client, groupID, params)
	if err == nil {
		t.Fatal("Expected error, got nil")
	}
}
