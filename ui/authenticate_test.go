package ui

import (
	"testing"

	"github.com/PaulWaldo/glimmer/api"
	"github.com/stretchr/testify/assert"

	"fyne.io/fyne/v2/test"
)

func Test_myApp_LoadSecrets(t *testing.T) {
	app := test.NewApp()
	ma := &myApp{
		app: app,
	}
	expectedSecrets := api.Secrets{
		ApiKey:      "apikey",
		ApiSecret:   "apisecret",
		AccessToken: "accesstoken",
		OAuthToken:  "oauthtoken",
		OAuthSecret: "oauthsecret",
	}

	// Store secrets into preferences
	p := app.Preferences()
	p.SetString(PREF_KEY_API_KEY, expectedSecrets.ApiKey)
	p.SetString(PREF_KEY_API_SECRET, expectedSecrets.ApiSecret)
	p.SetString(PREF_KEY_ACCESS_TOKEN, expectedSecrets.AccessToken)
	p.SetString(PREF_KEY_OAUTH_TOKEN, expectedSecrets.OAuthToken)
	p.SetString(PREF_KEY_OAUTH_SECRET, expectedSecrets.OAuthSecret)

	secrets := ma.LoadSecrets()

	assert.Equal(t, expectedSecrets.ApiKey, secrets.ApiKey)
	assert.Equal(t, expectedSecrets.ApiSecret, secrets.ApiSecret)
	assert.Equal(t, expectedSecrets.AccessToken, secrets.AccessToken)
	assert.Equal(t, expectedSecrets.OAuthSecret, secrets.OAuthSecret)
	assert.Equal(t, expectedSecrets.OAuthToken, secrets.OAuthToken)
}

func Test_apiInfoEntry_InfoHandling(t *testing.T) {
	e := apiInfoEntry{}
	e.makeUI()
	test.Type(e.apiKeyEntry, "abc123")
	test.Type(e.apiSecretEntry, "xyz789")

	e.form.OnSubmit()
	
	assert.Equal(t, "abc123", e.apiKeyEntry.Text)
}
