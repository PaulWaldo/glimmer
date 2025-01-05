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
					Photo: []api.Photo {
						{
							ID:       "12345",
							Owner:    "testuser",
							Secret:   "abcdef",
							Server:   "123",
							Farm:     "1",
							Title:    "Test Photo",
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
				assert.Equal(t, tt.want, resp)
			}
		})
	}
}
