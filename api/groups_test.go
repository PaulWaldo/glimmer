package api_test

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/PaulWaldo/glimmer/api"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/masci/flickr.v3"
	"gopkg.in/masci/flickr.v3/groups"
)

func BenchmarkGetUsersGroupPhotos(b *testing.B) {
	fclient := flickr.GetTestClient()
	userID := "12345"

	// Set up mock responses for a large number of groups and photos
	numGroups := 100 // Example: Simulate 100 groups

	groupsResponse := `<?xml version="1.0" encoding="utf-8" ?>
        <rsp stat="ok">
            <groups>`
	for i := 0; i < numGroups; i++ {
		groupsResponse += fmt.Sprintf(`<group nsid="%d" name="Test Group %d" members="10" privacy="1" iconserver="1" iconfarm="1" admin="1" />`, i, i)
	}
	groupsResponse += `</groups>
        </rsp>`

	photosResponse := make(map[string]string)
	for i := 0; i < numGroups; i++ {
		photosResponse[fmt.Sprintf("%d", i)] = `<?xml version="1.0" encoding="utf-8" ?>
            <rsp stat="ok">
                <photos page="1" pages="1" perpage="100" total="1">
                    <photo id="1" owner="testuser" secret="abcdef" server="123" farm="1" title="Test Photo" ispublic="1" isfriend="1" isfamily="0" />
                </photos>
            </rsp>`
	}

	transport := &mockTransport{
		responses: make(map[string]mockResponse),
	}
	transport.responses["flickr.people.getGroups"] = mockResponse{
		statusCode: 200,
		body:       groupsResponse,
	}
	for groupID, response := range photosResponse {
		transport.responses[fmt.Sprintf("flickr.groups.pools.getPhotos-%s", groupID)] = mockResponse{
			statusCode: 200,
			body:       response,
		}
	}
	fclient.HTTPClient = &http.Client{
		Transport: transport,
	}

	b.ResetTimer() // Reset timer to exclude setup time

	for i := 0; i < b.N; i++ { // Loop for accurate benchmarking
		_, err := api.GetUsersGroupPhotos(fclient, userID)
		if err != nil {
			b.Fatalf("GetUsersGroupPhotos failed: %v", err) // Use b.Fatal to signal errors
		}
	}
}

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
			groupsResponse: `<?xml version="1.0" encoding="utf-8" ?>
                <rsp stat="ok">
                    <groups>
                        <group nsid="12345" name="Test Group" members="10" privacy="1" iconserver="1" iconfarm="1" admin="1" />
                        <group nsid="67890" name="Another Group" members="20" privacy="2" iconserver="2" iconfarm="2" admin="0" />
                    </groups>
                </rsp>`,
			photosResponse: map[string]string{
				"12345": `<?xml version="1.0" encoding="utf-8" ?>
                    <rsp stat="ok">
                        <photos page="1" pages="1" perpage="100" total="1">
                            <photo id="12345" owner="testuser" secret="abcdef" server="123" farm="1" title="Test Photo" ispublic="1" isfriend="1" isfamily="0" />
                        </photos>
                    </rsp>`,
				"67890": `<?xml version="1.0" encoding="utf-8" ?>
                    <rsp stat="ok">
                        <photos page="1" pages="1" perpage="100" total="1">
                            <photo id="67890" owner="anotheruser" secret="ghijkl" server="456" farm="2" title="Another Photo" ispublic="1" isfriend="0" isfamily="1" />
                        </photos>
                    </rsp>`,
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
							Title:    "Test Photo",
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
			name: "error",
			groupsResponse: `<?xml version="1.0" encoding="utf-8" ?>
                <rsp stat="fail">
                    <err code="1" msg="Group not found" />
                </rsp>`,
			photosResponse: nil,
			want:           nil,
			wantErr:        true,
		},
		{
			name:           "invalid xml",
			groupsResponse: "invalid xml",
			photosResponse: nil,
			want:           nil,
			wantErr:        true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transport := &mockTransport{
				responses: make(map[string]mockResponse),
			}

			// Add response for groups endpoint
			transport.responses["flickr.people.getGroups"] = mockResponse{
				statusCode: 200,
				body:       tt.groupsResponse,
			}

			// Add responses for photos endpoints
			for groupID, response := range tt.photosResponse {
				transport.responses[fmt.Sprintf("flickr.groups.pools.getPhotos-%s", groupID)] = mockResponse{
					statusCode: 200,
					body:       response,
				}
			}
			fclient.HTTPClient = &http.Client{
				Transport: transport,
			}

			resp, err := api.GetUsersGroupPhotos(fclient, userID)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.want, resp)
		})
	}
}

type mockTransport struct {
	responses map[string]mockResponse
}

type mockResponse struct {
	statusCode int
	body       string
}

func (t *mockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if err := req.ParseForm(); err != nil {
		return nil, err
	}

	method := req.FormValue("method")
	groupID := req.FormValue("group_id")


	key := method
	if method == "flickr.groups.pools.getPhotos" {
		key = fmt.Sprintf("%s-%s", method, groupID)
	}

	response, ok := t.responses[key]
	if !ok {
		return nil, fmt.Errorf("no mock response found for key %q", key)
	}

	return &http.Response{
		StatusCode: response.statusCode,
		Body:       io.NopCloser(strings.NewReader(response.body)),
		Header:     make(http.Header),
	}, nil
}
