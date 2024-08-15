package glimmer

import (
	"testing"

	"github.com/PaulWaldo/glimmer/mocks"
	"github.com/stretchr/testify/assert"
	"gopkg.in/masci/flickr.v3"
)

// func TestAuth_NeedsAuthentication(t *testing.T) {
// 	testCases := []struct {
// 		desc          string
// 		authInfo      AuthInfo
// 		expectedNeeds bool
// 	}{
// 		{
// 			desc:          "Empty auth info needs auth",
// 			authInfo:      AuthInfo{},
// 			expectedNeeds: true,
// 		},
// 		{
// 			desc: "Full auth info does not need auth",
// 			authInfo: AuthInfo{
// 				ApiKey:      "abc",
// 				ApiSecret:   "def",
// 				AccessToken: "123",
// 				OAuthToken:  "456",
// 				OAuthSecret: "789",
// 			},
// 			expectedNeeds: false,
// 		},
// 	}
// 	for _, tC := range testCases {
// 		t.Run(tC.desc, func(t *testing.T) {
// 			needsAuth := NeedsAuthentication(tC.authInfo)
// 			assert.Equal(t, tC.expectedNeeds, needsAuth)
// 		})
// 	}
// }

// type tokenSuccess struct{}

// func (t tokenSuccess) GetRequestToken(client *flickr.FlickrClient) (*flickr.RequestToken, error) {
// 	return &flickr.RequestToken{OauthToken: ""}, nil
// }

// type urlSuccess struct{}

// func (u urlSuccess) GetAuthorizeUrl(client *flickr.FlickrClient, reqToken *flickr.RequestToken) (string, error) {
// 	return "https://example.com/auth_here", nil
// }

func TestAuth_GetAuthorizeUrl(t *testing.T) {
	testCases := []struct {
		desc        string
		authorize   Authorize
		expectedURL string
		expectError bool
	}{
		{
			desc: "URL Request success",
			authorize: Authorize{
				secrets: Secrets{
					ApiKey:      "abc",
					ApiSecret:   "def",
					AccessToken: "123",
					OAuthToken:  "456",
					OAuthSecret: "789",
				},
			},
			expectedURL: "https://example.com/auth_here",
			expectError: false,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			client := flickr.NewFlickrClient(tC.authorize.secrets.ApiKey, tC.authorize.secrets.ApiSecret)

			mockAuthorizer := mocks.NewAuthorizer(t)
			expectedRequestToken := &flickr.RequestToken{OauthToken: "token", OauthTokenSecret: "secret"}
			mockAuthorizer.EXPECT().GetRequestToken(client).Return(expectedRequestToken, nil)
			mockAuthorizer.EXPECT().GetAuthorizeUrl(client, expectedRequestToken).Return(tC.expectedURL, nil)

			authorize := Authorize{secrets: tC.authorize.secrets, authorizer: mockAuthorizer, client: client}

			url, err := authorize.GetUrl()
			if tC.expectError == true {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tC.expectedURL, url)
			}
		})
	}
}
