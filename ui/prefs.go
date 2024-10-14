package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"github.com/PaulWaldo/glimmer/api"
	"gopkg.in/masci/flickr.v3"
)

const (
	PrefKeyAPIKey      = "ApiKey"
	PrefKeyAPISecret   = "ApiSecret"
	PrefKeyAccessToken = "AccessToken"
	PrefKeyOauthToken  = "OAuthToken"
	PrefKeyOauthSecret = "OAuthSecret"
	PrefKeyUserNsID    = "UserNsID"
	PrefKeyUserName    = "UserName"
	PrefKeyFullName    = "FullName"
)

// appPrefs stores user data locally between application runs
type appPrefs struct {
	secrets  secrets
	userNsID binding.String
	userName binding.String
	fullName binding.String
}

type secrets struct {
	apiKey    binding.String
	apiSecret binding.String
	// accessToken      binding.String
	oAuthToken       binding.String
	oAuthTokenSecret binding.String
}

// NewClientFromPrefs creates a new Flickr client from user preferences
func NewClientFromPrefs(prefs appPrefs) *flickr.FlickrClient {
	apiKey, _ := prefs.secrets.apiKey.Get()
	apiSecret, _ := prefs.secrets.apiSecret.Get()
	oAuthToken, _ := prefs.secrets.oAuthToken.Get()
	oAuthSecret, _ := prefs.secrets.oAuthTokenSecret.Get()
	c := flickr.NewFlickrClient(apiKey, apiSecret)
	c.OAuthToken = oAuthToken
	c.OAuthTokenSecret = oAuthSecret
	return c
}

func NewPreferences(a fyne.App) appPrefs {
	return appPrefs{
		secrets: secrets{
			apiKey:    binding.BindPreferenceString(PrefKeyAPIKey, a.Preferences()),
			apiSecret: binding.BindPreferenceString(PrefKeyAPISecret, a.Preferences()),
			// accessToken:      binding.BindPreferenceString(PrefKeyAccessToken, a.Preferences()),
			oAuthToken:       binding.BindPreferenceString(PrefKeyOauthToken, a.Preferences()),
			oAuthTokenSecret: binding.BindPreferenceString(PrefKeyOauthSecret, a.Preferences()),
		},
		userNsID: binding.BindPreferenceString(PrefKeyUserNsID, a.Preferences()),
		userName: binding.BindPreferenceString(PrefKeyUserName, a.Preferences()),
		fullName: binding.BindPreferenceString(PrefKeyFullName, a.Preferences()),
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

func (p *appPrefs) StoreAuthPrefs(a api.Authorization) {
	// _ = p.secrets.oAuthToken.Set(a.RequestToken.OauthToken)
	// _ = p.secrets.oAuthTokenSecret.Set(a.RequestToken.OauthTokenSecret)
	_ = p.userNsID.Set(a.OAuthToken.UserNsid)
	_ = p.userName.Set(a.OAuthToken.Username)
	_ = p.fullName.Set(a.OAuthToken.Fullname)
}

func (ma *myApp) UpdateSecrefPrefs() {
	_ = ma.prefs.secrets.apiKey.Set(ma.client.ApiKey)
	_ = ma.prefs.secrets.apiSecret.Set(ma.client.ApiSecret)
	_ = ma.prefs.secrets.oAuthToken.Set(ma.client.OAuthToken)
	_ = ma.prefs.secrets.oAuthTokenSecret.Set(ma.client.OAuthTokenSecret)
	ma.logAuth("UpdateSecrefPrefs")
}

func (ma *myApp) SaveAuth(a api.Authorization) {
	ma.UpdateSecrefPrefs()
	ma.prefs.StoreAuthPrefs(a)
}