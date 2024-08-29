package ui

import (
	"testing"

	"fyne.io/fyne/v2/test"
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
