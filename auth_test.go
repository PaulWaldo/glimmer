package glimmer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuth_NeedsAuthentication(t *testing.T) {
	testCases := []struct {
		desc          string
		authInfo      authInfo
		expectedNeeds bool
	}{
		{
			desc:          "Empty auth info needs auth",
			authInfo:      authInfo{},
			expectedNeeds: true,
		},
		{
			desc: "Full auth info does not need auth",
			authInfo: authInfo{
				apiKey:      "abc",
				apiSecret:   "def",
				accessToken: "123",
				oauthToken:  "456",
				oauthSecret: "789",
			},
			expectedNeeds: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			needsAuth := tC.authInfo.NeedsAuthentication()
			assert.Equal(t, tC.expectedNeeds, needsAuth)
		})
	}
}

func TestAuth_GetAuthorizeUrl(t *testing.T) {
	testCases := []struct {
		desc        string
		authInfo    authInfo
		expectedURL string
		expectError bool
	}{
		{
			desc: "URL Request success",
			authInfo: authInfo{
				apiKey:      "abc",
				apiSecret:   "def",
				accessToken: "123",
				oauthToken:  "456",
				oauthSecret: "789",
			},
			expectedURL: "https://nhjghjg",
			expectError: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			url, err := tC.authInfo.GetAuthorizeUrl()
			if tC.expectError == true {
				assert.Error(t, err)
			} else {
				assert.Equal(t, tC.expectedURL, url)
			}
		})
	}
}
