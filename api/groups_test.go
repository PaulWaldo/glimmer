package api_test

import (
	"fmt"
	"io"
	"mime"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/PaulWaldo/glimmer/api"
	"github.com/stretchr/testify/assert"
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
		groupsResponse += fmt.Sprintf(`<group nsid="%d" name="Test Group %d" member_count="10" privacy="1" iconserver="1" iconfarm="1" admin="1" />`, i, i)
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
		var testUsersGroups []groups.Group
		var testUsersGroupPhotos []api.UsersGroupPhotos

		err := api.GetUsersGroupPhotos(fclient, userID, nil, &testUsersGroups, &testUsersGroupPhotos)
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
			name:       "success returns photos from a group",
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
			name:       "server error returns error",
			statusCode: 500,
			response:   "",
			want:       nil,
			wantErr:    true,
		},
		{
			name:       "invalid xml returns error",
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
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
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
			name:       "success returns user's groups",
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
			name:       "server error returns error",
			statusCode: 500,
			response:   "",
			want:       nil,
			wantErr:    true,
		},
		{
			name:       "invalid server xml returns error",
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
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
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
		wantGroups     []groups.Group
		wantPhotos     []api.UsersGroupPhotos
		wantErr        bool
	}{
		{
			name: "success returns user's group photos",
			groupsResponse: `<?xml version="1.0" encoding="utf-8" ?>
                <rsp stat="ok">
                    <groups>
                        <group nsid="12345" name="Test Group" member_count="10" privacy="1" iconserver="1" iconfarm="1" admin="1" />
                        <group nsid="67890" name="Another Group" member_count="20" privacy="2" iconserver="2" iconfarm="2" admin="0" />
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
			wantGroups: []groups.Group{
				{
					Nsid:        "12345",
					Name:        "Test Group",
					Privacy:     "1",
					Admin:       "1",
					Iconserver:  "1",
					Iconfarm:    "1",
					MemberCount: "10",
				},
				{
					Nsid:        "67890",
					Name:        "Another Group",
					Privacy:     "2",
					Admin:       "0",
					Iconserver:  "2",
					Iconfarm:    "2",
					MemberCount: "20",
				},
			},
			wantPhotos: []api.UsersGroupPhotos{
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
			name: "server error returns error",
			groupsResponse: `<?xml version="1.0" encoding="utf-8" ?>
                <rsp stat="fail">
                    <err code="1" msg="Group not found" />
                </rsp>`,
			photosResponse: nil,
			wantGroups:     []groups.Group{},
			wantPhotos:     nil,
			wantErr:        true,
		},
		{
			name:           "invalid server xml returns error",
			groupsResponse: "invalid xml",
			photosResponse: nil,
			wantGroups:     []groups.Group{},
			wantPhotos:     nil,
			wantErr:        true,
		},
		{
			name: "no groups returns empty list",
			groupsResponse: `<?xml version="1.0" encoding="utf-8" ?>
                <rsp stat="ok">
                    <groups />
                </rsp>`,
			photosResponse: nil,
			wantGroups:     []groups.Group{},
			wantPhotos:     []api.UsersGroupPhotos{}, // Expect an empty list, not nil
			wantErr:        false,                    // Expect no error
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

			testUsersGroups := []groups.Group{}
			testUsersGroupPhotos := []api.UsersGroupPhotos{}

			err := api.GetUsersGroupPhotos(fclient, userID, nil, &testUsersGroups, &testUsersGroupPhotos)
			if tt.wantErr {
				assert.Error(t, err)
				return
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantPhotos, testUsersGroupPhotos)

				if tt.wantGroups != nil { // Only check groups if expected
					assert.Equal(t, len(tt.wantGroups), len(testUsersGroups))
					if len(tt.wantGroups) > 0 {
						assert.Equal(t, tt.wantGroups[0], testUsersGroups[0])
						assert.Equal(t, tt.wantGroups[1], testUsersGroups[1])
					}
				}
			}
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

	if req.Method == "POST" && method == "" && groupID == "" {
		mediaType, params, err := mime.ParseMediaType(req.Header.Get("Content-Type"))
		if err != nil {
			return nil, fmt.Errorf("parsing media type: %w", err)
		}
		if strings.HasPrefix(mediaType, "multipart/") {
			mr := multipart.NewReader(req.Body, params["boundary"])
			values := make(map[string][]string)
			for {
				p, err := mr.NextPart()
				if err == io.EOF {
					break
				}
				if err != nil {
					return nil, fmt.Errorf("reading next part: %w", err)
				}
				sl, err := io.ReadAll(p)
				if err != nil {
					return nil, fmt.Errorf("reading all from part: %w", err)
				}
				values[p.FormName()] = append(values[p.FormName()], string(sl))
			}
			req.Form = url.Values(values)
			method = req.FormValue("method")
			groupID = req.FormValue("group_id")
		}
	}

	if method == "" && groupID == "" {
		err := req.ParseMultipartForm(1024 * 1024) // Adjust limit as needed
		if err != nil {
			return nil, err
		}

		method = req.FormValue("method")
		groupID = req.FormValue("group_id")

		if method == "" || groupID == "" {
			return nil, fmt.Errorf("method or group_id not found in multipart form")
		}
	}

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
