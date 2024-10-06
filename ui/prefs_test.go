package ui

import (
	"testing"

	"fyne.io/fyne/v2/test"
	"github.com/PaulWaldo/glimmer/api"
	"github.com/stretchr/testify/assert"
	"gopkg.in/masci/flickr.v3"
)

func TestNewClientFromPrefs(t *testing.T) {
	apiKey := "api_key"
	apiSecret := "api_secret"
	oAuthToken := "oauthtoken"
	oAuthTokenSecret := "oauthsecret"

	app := test.NewTempApp(t)
	prefs := NewPreferences(app)
	var err error
	err = prefs.secrets.apiKey.Set(apiKey)
	assert.NoError(t, err)
	err = prefs.secrets.apiSecret.Set(apiSecret)
	assert.NoError(t, err)
	err = prefs.secrets.oAuthToken.Set(oAuthToken)
	assert.NoError(t, err)
	err = prefs.secrets.oAuthTokenSecret.Set(oAuthTokenSecret)
	assert.NoError(t, err)

	client := NewClientFromPrefs(prefs)

	assert.Equal(t, apiKey, client.ApiKey)
	assert.Equal(t, apiSecret, client.ApiSecret)
	assert.Equal(t, oAuthToken, client.OAuthToken)
	assert.Equal(t, oAuthTokenSecret, client.OAuthTokenSecret)
}

func TestClearCredentialsPrefs(t *testing.T) {
	app := test.NewTempApp(t)
	prefs := NewPreferences(app)
	var err error
	err = prefs.secrets.apiKey.Set("apiKey")
	assert.NoError(t, err)
	err = prefs.secrets.apiSecret.Set("apiSecret")
	assert.NoError(t, err)
	err = prefs.secrets.oAuthToken.Set("oAuthToken")
	assert.NoError(t, err)
	err = prefs.secrets.oAuthTokenSecret.Set("oAuthTokenSecret")
	assert.NoError(t, err)

	ClearCredentialsPrefs()

	var val string
	val, _ = prefs.secrets.apiKey.Get()
	assert.Equal(t, "", val)
	val, _ = prefs.secrets.apiSecret.Get()
	assert.Equal(t, "", val)
	val, _ = prefs.secrets.oAuthToken.Get()
	assert.Equal(t, "", val)
	val, _ = prefs.secrets.oAuthTokenSecret.Get()
	assert.Equal(t, "", val)
}

func Test_myApp_UpdateSecretsPrefs(t *testing.T) {
	ma := myApp{prefs: NewPreferences(test.NewTempApp(t))}
	ma.client = flickr.NewFlickrClient("apiKey", "apiSecret")
	ma.client.OAuthToken = "oAuthToken"
	ma.client.OAuthTokenSecret = "oAuthTokenSecret"

	ma.UpdateSecrefPrefs()

	var val string
	val, _ = ma.prefs.secrets.apiKey.Get()
	assert.Equal(t, "apiKey", val)
	val, _ = ma.prefs.secrets.apiSecret.Get()
	assert.Equal(t, "apiSecret", val)
	val, _ = ma.prefs.secrets.oAuthToken.Get()
	assert.Equal(t, "oAuthToken", val)
	val, _ = ma.prefs.secrets.oAuthTokenSecret.Get()
	assert.Equal(t, "oAuthTokenSecret", val)
}

func Test_appPrefs_StoreAuthPrefs(t *testing.T) {
	app := test.NewApp()
	p := NewPreferences(app)
	auth := api.NewAuthorizer()

	auth.RequestToken = &flickr.RequestToken{}
	auth.RequestToken.OauthToken = "oauthtoken"
	auth.RequestToken.OauthTokenSecret = "oauthsecret"

	auth.OAuthToken = &flickr.OAuthToken{}
	auth.OAuthToken.Fullname = "fullname"
	auth.OAuthToken.Username = "username"
	auth.OAuthToken.UserNsid = "usernsid"

	p.StoreAuthPrefs(*auth)

	var val string
	// val, _ := p.secrets.oAuthToken.Get()
	// assert.Equal(t, "oauthtoken", val)
	// val, _ = p.secrets.oAuthTokenSecret.Get()
	// assert.Equal(t, "oauthsecret", val)
	val, _ = p.userNsID.Get()
	assert.Equal(t, "usernsid", val)
	val, _ = p.userName.Get()
	assert.Equal(t, "username", val)
	val, _ = p.fullName.Get()
	assert.Equal(t, "fullname", val)
}

func Test_myApp_SaveAuth_StoresAllAuthDataInPrefs(t *testing.T) {
	expectedApiKey := "apikey"
	expectedApiSecret := "apisecret"
	expectedOAuthToken := "oauthtoken"
	expectedOAuthTokenSecret := "oauthsecret"
	client := &flickr.FlickrClient{}
	client.ApiKey = expectedApiKey
	client.ApiSecret = expectedApiSecret
	// client.OAuthToken = expectedOAuthToken
	// client.OAuthTokenSecret = expectedOAuthTokenSecret
	expectedUserNsID := "usernsid"
	expectedUserName := "username"
	expectedFullName := "fullname"
	auth := api.Authorization{
		RequestToken: &flickr.RequestToken{
			OauthToken:       expectedOAuthToken,
			OauthTokenSecret: expectedOAuthTokenSecret,
		},
		OAuthToken: &flickr.OAuthToken{
			OAuthToken:       expectedOAuthToken,
			OAuthTokenSecret: expectedOAuthTokenSecret,
			UserNsid:         expectedUserNsID,
			Username:         expectedUserName,
			Fullname:         expectedFullName,
		},
	}
	app := test.NewApp()

	ma := myApp{prefs: NewPreferences(app), client: client}
	ma.SaveAuth(auth)

	client2 := NewClientFromPrefs(ma.prefs)
	assert.Equal(t, expectedApiKey, client2.ApiKey)
	assert.Equal(t, expectedApiSecret, client2.ApiSecret)
	// assert.Equal(t, expectedOAuthToken, client2.OAuthToken)
	// assert.Equal(t, expectedOAuthTokenSecret, client2.OAuthTokenSecret)
}
