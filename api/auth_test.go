package api_test

import (
	"errors"
	"testing"

	"github.com/PaulWaldo/glimmer/api"
	"github.com/PaulWaldo/glimmer/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
		authorize   api.Authorize
		expectedURL string
		expectError bool
	}{
		{
			desc: "URL Request success",
			authorize: api.Authorize{
				Secrets: api.Secrets{
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
			client := flickr.NewFlickrClient(tC.authorize.Secrets.ApiKey, tC.authorize.Secrets.ApiSecret)

			mockAuthorizer := mocks.NewAuthorizer(t)
			expectedRequestToken := &flickr.RequestToken{OauthToken: "token", OauthTokenSecret: "secret"}
			mockAuthorizer.EXPECT().GetRequestToken(client).Return(expectedRequestToken, nil)
			mockAuthorizer.EXPECT().GetAuthorizeUrl(client, expectedRequestToken).Return(tC.expectedURL, nil)

			authorize := api.Authorize{Secrets: tC.authorize.Secrets, Authorizer: mockAuthorizer, Client: client}

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

func TestGetAccessToken(t *testing.T) {
	testCases := []struct {
		desc        string
		expectError bool
		err         error
		// authorize   Authorize
	}{
		{
			desc:        "Successful Token",
			expectError: false,
			err:         nil,
		},
		{
			desc:        "Unsuccessful Token",
			expectError: true,
			err:         errors.New("test error"),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			client := flickr.NewFlickrClient("", "")
			confirmationCode := "abc123"

			authorize := api.Authorize{}
			mockAuthorizer := mocks.NewAuthorizer(t)
			expectedRequestToken := &flickr.RequestToken{OauthToken: "token", OauthTokenSecret: "secret"}
			authorize.RequestToken = expectedRequestToken
			authorize.Client = client
			mockAuthorizer.EXPECT().GetAccessToken(client, expectedRequestToken, confirmationCode).Return(&flickr.OAuthToken{OAuthToken: "token", OAuthTokenSecret: "secret"}, tC.err)
			authorize.Authorizer = mockAuthorizer

			err := authorize.GetAccessToken(confirmationCode)

			if tC.expectError == true {
				require.Error(t, errors.New("test error"))
				return
			}
			require.NoError(t, err)
			assert.Equal(t, "token", authorize.Client.OAuthToken, "expecting OAuth Token to be %q but got %q", "token", authorize.Client.OAuthToken)
			assert.Equal(t, "secret", authorize.Client.OAuthTokenSecret, "expecting OAuth Secret to be %q but got %q", "secret", authorize.Client.OAuthTokenSecret)
		})
	}
}
