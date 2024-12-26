package api

import (
	"testing"

	"github.com/stretchr/testify/mock"

	flickr "gopkg.in/masci/flickr.v3"
)

type MockFlickrClient struct {
	mock.Mock
}

func (m *MockFlickrClient) DoGet(response interface{}) error {
	args := m.Called(response)
	return args.Error(0)
}

func TestGetGroups(t *testing.T) {
	mockClient := &MockFlickrClient{}
	fakeGroups := &GetGroupsResponse{
		Groups: Groups{
			Group: []Group{
				{Nsid: "12345", Name: "Fake Group 1"},
				{Nsid: "67890", Name: "Fake Group 2"},
			},
		},
	}

	mockClient.On("DoGet", mock.Any()).Return(fakeGroups, nil)

	groups, err := GetGroups(mockClient, "some_user_id", "")

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if len(groups.Groups.Group) != 2 {
		t.Errorf("Expected 2 groups, but got %d", len(groups.Groups.Group))
	}
}
