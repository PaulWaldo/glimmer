package api_test

import (
	"testing"

	"github.com/PaulWaldo/glimmer/api"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/masci/flickr.v3"
	"gopkg.in/masci/flickr.v3/groups"
)

func TestGetGroupPhotos(t *testing.T) {
	fclient := flickr.GetTestClient()
	groupID := "12345"
	params := map[string]string{}

	tests := []struct {
		name       string
		statusCode int
		response   string
		want       *api.GetGroupPhotosResponse
		wantErr    bool
	}{
		{
			name:       "success",
			statusCode: 200,
			response: `
                 <rsp stat="ok">
                     <photos page="1" pages="1" perpage="100" total="3">
                         <photo id="12345" owner="testuser" secret="abcdef" server="123" farm="1" title="Te
 Photo" ispublic="1" isfriend="1" isfamily="0" />
                         <photo id="67890" owner="anotheruser" secret="ghijkl" server="456" farm="2"
 title="Another Photo" ispublic="1" isfriend="0" isfamily="1" />
                         <photo id="34567" owner="yetanotheruser" secret="mnopqr" server="789" farm="3"
 title="Yet Another Photo" ispublic="1" isfriend="1" isfamily="1" />
                     </photos>
                 </rsp>
             `,
			want: &api.GetGroupPhotosResponse{
				BasicResponse: flickr.BasicResponse{
					Status: "ok",
				},
				Photos: []api.Photo{
					{
						ID:       "12345",
						Owner:    "testuser",
						Secret:   "abcdef",
						Server:   "123",
						Farm:     "1",
						Title:    "Te\n Photo",
						IsPublic: 1,
						IsFriend: 1,
						IsFamily: 0,
					},
					{
						ID:       "67890",
						Owner:    "anotheruser",
						Secret:   "ghijkl",
						Server:   "456",
						Farm:     "2",
						Title:    "Another Photo",
						IsPublic: 1,
						IsFriend: 0,
						IsFamily: 1,
					},
					{
						ID:       "34567",
						Owner:    "yetanotheruser",
						Secret:   "mnopqr",
						Server:   "789",
						Farm:     "3",
						Title:    "Yet Another Photo",
						IsPublic: 1,
						IsFriend: 1,
						IsFamily: 1,
					},
				},
			},
			wantErr: false,
		},
		{
			name:       "error",
			statusCode: 500,
			response:   "",
			want:       nil,
			wantErr:    true,
		},
		{
			name:       "invalid xml",
			statusCode: 200,
			response:   " invalid xml ",
			want:       nil,
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockServer, client := flickr.FlickrMock(tt.statusCode, tt.response, "text/xml")
			defer mockServer.Close()
			fclient.HTTPClient = client

			resp, err := api.GetGroupPhotos(fclient, groupID, params)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				// Don't worry about the Extra field, it is just the raw XML
				tt.want.Extra = resp.Extra
				assert.Equal(t, tt.want, resp)
			}
		})
	}
}

func TestGetUserGroups(t *testing.T) {
	fclient := flickr.GetTestClient()
	userID := "12345"
	params := map[string]string{}

	tests := []struct {
		name       string
		statusCode int
		response   string
		want       *api.GetUserGroupsResponse
		wantErr    bool
	}{
		{
			name:       "success",
			statusCode: 200,
			response: `
				<rsp stat="ok">
					<groups>
						<group id="12345" name="Test Group" member_count="10" privacy="1" admin="1" />
						<group id="67890" name="Another Group" member_count="20" privacy="2" admin="0" />
					</groups>
				</rsp>
			`,
			want: &api.GetUserGroupsResponse{
				BasicResponse: flickr.BasicResponse{
					Status: "ok",
				},
				Groups: []groups.Group{
					{
						ID:          "12345",
						Name:        "Test Group",
						MemberCount: "10",
						Privacy:     "1",
						Admin:       "1",
					},
					{
						ID:          "67890",
						Name:        "Another Group",
						MemberCount: "20",
						Privacy:     "2",
						Admin:       "0",
					},
				},
			},
			wantErr: false,
		},
		{
			name:       "error",
			statusCode: 500,
			response:   "",
			want:       nil,
			wantErr:    true,
		},
		{
			name:       "invalid xml",
			statusCode: 200,
			response:   " invalid xml ",
			want:       nil,
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockServer, client := flickr.FlickrMock(tt.statusCode, tt.response, "text/xml")
			defer mockServer.Close()
			fclient.HTTPClient = client

			resp, err := api.GetUserGroups(fclient, userID, params)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				tt.want.Extra = resp.Extra
				assert.Equal(t, tt.want, resp)
			}
		})
	}
}

func TestGetUsersGroupPhotos(t *testing.T) {
	fclient := flickr.GetTestClient()
	userID := "12345"

	tests := []struct {
		name           string
		groupsResponse string
		photosResponse map[string]string
		want           []api.UsersGroupPhotos
		wantErr        bool
	}{
		{
			name: "success",
			groupsResponse: `
				<rsp stat="ok">
					<groups>
						<group id="12345" name="Test Group" member_count="10" privacy="1" admin="1" />
						<group id="67890" name="Another Group" member_count="20" privacy="2" admin="0" />
					</groups>
				</rsp>
			`,
			photosResponse: map[string]string{
				"12345": `
					<rsp stat="ok">
						<photos page="1" pages="1" perpage="100" total="3">
							<photo id="12345" owner="testuser" secret="abcdef" server="123" farm="1" title="Te\n Photo" ispublic="1" isfriend="1" isfamily="0" />
						</photos>
					</rsp>
				`,
				"67890": `
					<rsp stat="ok">
						<photos page="1" pages="1" perpage="100" total="3">
							<photo id="67890" owner="anotheruser" secret="ghijkl" server="456" farm="2" title="Another Photo" ispublic="1" isfriend="0" isfamily="1" />
						</photos>
					</rsp>
				`,
			},
			want: []api.UsersGroupPhotos{
				{
					GroupID:   "12345",
					GroupName: "Test Group",
					Photos: []api.Photo{
						{
							ID:       "12345",
							Owner:    "testuser",
							Secret:   "abcdef",
							Server:   "123",
							Farm:     "1",
							Title:    "Te\n Photo",
							IsPublic: 1,
							IsFriend: 1,
							IsFamily: 0,
						},
					},
				},
				{
					GroupID:   "67890",
					GroupName: "Another Group",
					Photos: []api.Photo{
						{
							ID:       "67890",
							Owner:    "anotheruser",
							Secret:   "ghijkl",
							Server:   "456",
							Farm:     "2",
							Title:    "Another Photo",
							IsPublic: 1,
							IsFriend: 0,
							IsFamily: 1,
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name:           "error",
			groupsResponse: "",
			photosResponse: nil,
			want:           nil,
			wantErr:        true,
		},
		{
			name:           "invalid xml",
			groupsResponse: " invalid xml ",
			photosResponse: nil,
			want:           nil,
			wantErr:        true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			groupsMockServer, groupsClient := flickr.FlickrMock(200, tt.groupsResponse, "text/xml")
			defer groupsMockServer.Close()
			fclient.HTTPClient = groupsClient.HTTPClient

			photosMockServers := make([]*flickr.MockServer, 0)
			photosClients := make(map[string]*flickr.FlickrClient, 0)
			for groupID, response := range tt.photosResponse {
				mockServer, client := flickr.FlickrMock(200, response, "text/xml")
				photosMockServers = append(photosMockServers, mockServer)
				photosClients[groupID] = client
			}
			defer func() {
				for _, server := range photosMockServers {
					server.Close()
				}
			}()

			var originalDoPost = flickr.DoPost
			defer func() { flickr.DoPost = originalDoPost }()
			flickr.DoPost = func(client *flickr.FlickrClient, response interface{}) error {
				if client.Args.Get("method") == "flickr.people.getGroups" {
					return originalDoPost(groupsClient, response)
				}
				groupID := client.Args.Get("group_id")
				if client.Args.Get("method") == "flickr.groups.pools.getPhotos" {
					return originalDoPost(photosClients[groupID], response)
				}
				return nil
			}

			resp, err := api.GetUsersGroupPhotos(fclient, userID)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.want, resp)
			}
		})
	}
}
