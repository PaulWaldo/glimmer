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
	if args.Get(0) != nil {
		*response.(*GetGroupsResponse) = *args.Get(0).(*GetGroupsResponse)
	}
	return args.Error(1)
}

func (m *MockFlickrClient) Init() {}
func (m *MockFlickrClient) OAuthSign() {}
func (m *MockFlickrClient) Args() map[string]string { return map[string]string{} }
func (m *MockFlickrClient) SetArg(key, value string) {}
func (m *MockFlickrClient) EndpointUrl() string { return "" }
func (m *MockFlickrClient) SetEndpointUrl(url string) {}

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

	mockClient.On("DoGet", mock.Anything).Return(fakeGroups, nil)

	groups, err := GetGroups(mockClient, "some_user_id", "")

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if len(groups.Groups.Group) != 2 {
		t.Errorf("Expected 2 groups, but got %d", len(groups.Groups.Group))
	}
}
