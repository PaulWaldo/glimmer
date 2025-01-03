package api_test

import (
	"testing"

	"github.com/PaulWaldo/glimmer/api"
	"gopkg.in/masci/flickr.v3"
)

func TestGetGroupPhotos(t *testing.T) {
	tests := []struct {
		name       string
		statusCode int
		response   string
		wantErr    bool
	}{
		{
			name:       "returns valid group photos for valid input",
			statusCode: 200,
			response: `
				<rsp stat="ok">
					<photos page="1" pages="1" perpage="100" total="3">
						<photo id="12345" owner="testuser" secret="abcdef" server="123" farm="1" title="Test Photo" ispublic="1" isfriend="0" isfamily="0" />
						<photo id="67890" owner="anotheruser" secret="ghijkl" server="456" farm="2" title="Another Photo" ispublic="1" isfriend="0" isfamily="0" />
						<photo id="34567" owner="yetanotheruser" secret="mnopqr" server="789" farm="3" title="Yet Another Photo" ispublic="1" isfriend="0" isfamily="0" />
					</photos>
				</rsp>
			`,
			wantErr: false,
		},
		{
			name:       "returns error for server error",
			statusCode: 500,
			response:   "",
			wantErr:    true,
		},
		{
			name:       "returns error for invalid xml from server",
			statusCode: 200,
			response:   " invalid xml ",
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fclient := flickr.GetTestClient()
			mockServer, client := flickr.FlickrMock(tt.statusCode, tt.response, "text/xml")
			defer mockServer.Close()
			fclient.HTTPClient = client

			groupID := "12345"
			params := map[string]string{}

			_, err := api.GetGroupPhotos(fclient, groupID, params)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetGroupPhotos() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
