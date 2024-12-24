package api

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"io/ioutil"
	"strings"
	"gopkg.in/masci/flickr.v3"
)

// MockFlickrClient is a mock implementation of the FlickrClient for testing.
type MockFlickrClient struct {
	*flickr.FlickrClient
	ResponseBody string
}

// DoGet overrides the original DoGet method to return a mock response.
func (m *MockFlickrClient) DoGet(response interface{}) error {
	resp := httptest.NewRecorder()
	resp.Body = ioutil.NopCloser(strings.NewReader(m.ResponseBody))
	return flickr.ParseResponse(resp.Result(), response)
}

// TestGetGroups tests the GetGroups function.
func TestGetGroups(t *testing.T) {
	mockResponse := `<?xml version="1.0" encoding="utf-8" ?>
	<rsp stat="ok">
		<groups>
			<group nsid="12345678901@N01" name="Test Group" iconfarm="1" iconserver="1" admin="1" eighteenplus="0" invitation_only="0" members="10" pool_count="5"/>
		</groups>
	</rsp>`

	mockClient := &MockFlickrClient{
		FlickrClient: &flickr.FlickrClient{},
		ResponseBody: mockResponse,
	}

	response, err := GetGroups(mockClient, "12345678901@N01", "")
	if err != nil {
		t.Errorf("GetGroups returned an error: %v", err)
	}

	if response.Groups.Group[0].Nsid != "12345678901@N01" {
		t.Errorf("GetGroups returned wrong Nsid: got %v, want %v", response.Groups.Group[0].Nsid, "12345678901@N01")
	}

	if response.Groups.Group[0].Name != "Test Group" {
		t.Errorf("GetGroups returned wrong Name: got %v, want %v", response.Groups.Group[0].Name, "Test Group")
	}

	if response.Groups.Group[0].Iconfarm != "1" {
		t.Errorf("GetGroups returned wrong Iconfarm: got %v, want %v", response.Groups.Group[0].Iconfarm, "1")
	}

	if response.Groups.Group[0].Iconserver != "1" {
		t.Errorf("GetGroups returned wrong Iconserver: got %v, want %v", response.Groups.Group[0].Iconserver, "1")
	}

	if response.Groups.Group[0].Admin != "1" {
		t.Errorf("GetGroups returned wrong Admin: got %v, want %v", response.Groups.Group[0].Admin, "1")
	}

	if response.Groups.Group[0].Eighteenplus != "0" {
		t.Errorf("GetGroups returned wrong Eighteenplus: got %v, want %v", response.Groups.Group[0].Eighteenplus, "0")
	}

	if response.Groups.Group[0].InvitationOnly != "0" {
		t.Errorf("GetGroups returned wrong InvitationOnly: got %v, want %v", response.Groups.Group[0].InvitationOnly, "0")
	}

	if response.Groups.Group[0].Members != "10" {
		t.Errorf("GetGroups returned wrong Members: got %v, want %v", response.Groups.Group[0].Members, "10")
	}

	if response.Groups.Group[0].PoolCount != "5" {
		t.Errorf("GetGroups returned wrong PoolCount: got %v, want %v", response.Groups.Group[0].PoolCount, "5")
	}
}
