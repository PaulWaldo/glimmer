package api_test

import (
	"testing"

	"github.com/PaulWaldo/glimmer/api"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/masci/flickr.v3"
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
				Photos: struct {
					Page    int `xml:"page,attr"`
					Pages   int `xml:"pages,attr"`
					PerPage int `xml:"perpage,attr"`
					Total   int `xml:"total,attr"`
					Photo   []api.Photo `xml:"photo"`
				}{
					Page:    1,
					Pages:   1,
					PerPage: 100,
					Total:   3,
					Photo: []api.Photo{
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
						<group id="12345" name="Test Group" members="10" privacy="1" admin="1" invitation="0" needs_invite="0" />
						<group id="67890" name="Another Group" members="20" privacy="2" admin="0" invitation="1" needs_invite="1" />
					</groups>
				</rsp>
			`,
			want: &api.GetUserGroupsResponse{
				BasicResponse: flickr.BasicResponse{
					Status: "ok",
				},
				Groups: struct {
					Group []flickr.Group `xml:"group"`
				}{
					Group: []flickr.Group{
						{
							ID:          "12345",
							Name:        "Test Group",
							Members:     10,
							Privacy:     1,
							Admin:       1,
							Invitation:  0,
							NeedsInvite: 0,
						},
						{
							ID:          "67890",
							Name:        "Another Group",
							Members:     20,
							Privacy:     2,
							Admin:       0,
							Invitation:  1,
							NeedsInvite: 1,
						},
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
