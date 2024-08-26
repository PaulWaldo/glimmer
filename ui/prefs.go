package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"gopkg.in/masci/flickr.v3"
)

const (
	PrefKeyAPIKey      = "ApiKey"
	PrefKeyAPISecret   = "ApiSecret"
	PrefKeyAccessToken = "AccessToken"
	PrefKeyOauthToken  = "OAuthToken"
	PrefKeyOauthSecret = "OAuthSecret"
)

// appPrefs stores user data locally between application runs
type appPrefs struct {
	secrets secrets
}

type secrets struct {
	apiKey      binding.String
	apiSecret   binding.String
	accessToken binding.String
	oAuthToken  binding.String
	oAuthSecret binding.String
}

// NewClientFromPrefs creates a new Flickr client from user preferences
func NewClientFromPrefs(secrets secrets) *flickr.FlickrClient {
	apiKey, _ := secrets.apiKey.Get()
	apiSecret, _ := secrets.apiSecret.Get()
	// accessToken, _ := p.secrets.accessToken.Get()
	oAuthToken, _ := secrets.oAuthToken.Get()
	oAuthSecret, _ := secrets.oAuthSecret.Get()
	c := flickr.NewFlickrClient(apiKey, apiSecret)
	c.OAuthToken = oAuthToken
	c.OAuthTokenSecret = oAuthSecret
	return c
}

func NewPreferences(a fyne.App) appPrefs {
	return appPrefs{
		secrets: secrets{
			apiKey:      binding.BindPreferenceString(PrefKeyAPIKey, a.Preferences()),
			apiSecret:   binding.BindPreferenceString(PrefKeyAPISecret, a.Preferences()),
			accessToken: binding.BindPreferenceString(PrefKeyAccessToken, a.Preferences()),
			oAuthToken:  binding.BindPreferenceString(PrefKeyOauthToken, a.Preferences()),
			oAuthSecret: binding.BindPreferenceString(PrefKeyOauthSecret, a.Preferences()),
		},
	}
}

func ClearCredentialsPrefs() {
	p := fyne.CurrentApp().Preferences()
	p.RemoveValue(PrefKeyAPIKey)
	p.RemoveValue(PrefKeyAPISecret)
	p.RemoveValue(PrefKeyAccessToken)
	p.RemoveValue(PrefKeyOauthToken)
	p.RemoveValue(PrefKeyOauthSecret)
}
