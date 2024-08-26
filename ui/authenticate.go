package ui

import (
	"fmt"
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/PaulWaldo/glimmer/api"
	"gopkg.in/masci/flickr.v3"
	"gopkg.in/masci/flickr.v3/test"
)

// const PREF_KEY_API_KEY = "ApiKey"
// const PREF_KEY_API_SECRET = "ApiSecret"
// const PREF_KEY_ACCESS_TOKEN = "AccessToken"
// const PREF_KEY_OAUTH_TOKEN = "OAuthToken"
// const PREF_KEY_OAUTH_SECRET = "OAuthSecret"

// func (ma *myApp) LoadSecrets() api.Secrets {
// 	sec := api.Secrets{}
// 	pref := ma.app.Preferences()
// 	sec.ApiKey = pref.String(PREF_KEY_API_KEY)
// 	sec.ApiSecret = pref.String(PREF_KEY_API_SECRET)
// 	sec.AccessToken = pref.String(PREF_KEY_ACCESS_TOKEN)
// 	sec.OAuthToken = pref.String(PREF_KEY_OAUTH_TOKEN)
// 	sec.OAuthSecret = pref.String(PREF_KEY_OAUTH_SECRET)
// 	return sec
// }

// func (ma *myApp) CreateSecrets() (api.Secrets, error) {
// 	sec := api.Secrets{}
// 	w := ma.app.NewWindow("Authorize to Flickr")
// 	e := apiInfoEntry{}
// 	w.SetContent(e.makeUI(ma))

// 	return sec, nil
// }

type apiInfoEntry struct {
	apiKeyEntry    *widget.Entry
	apiSecretEntry *widget.Entry
	form           *widget.Form
	myApp          *myApp
}

func (e *apiInfoEntry) makeUI(ma *myApp) fyne.CanvasObject {
	e.apiKeyEntry = widget.NewPasswordEntry()
	e.apiSecretEntry = widget.NewPasswordEntry()
	e.form = widget.NewForm(
		widget.NewFormItem("API Key", e.apiKeyEntry),
		widget.NewFormItem("API Secret", e.apiSecretEntry),
	)
	e.form.OnSubmit = func() {}
	return e.form
}

type accessTokenEntry struct {
	accessTokenEntry *widget.Entry
	form             *widget.Form
}

func (e *accessTokenEntry) makeUI() fyne.CanvasObject {
	e.accessTokenEntry = widget.NewPasswordEntry()
	e.form = widget.NewForm(
		widget.NewFormItem("Access Token", e.accessTokenEntry),
	)
	e.form.OnSubmit = func() {}
	return e.form
}

func (ma *myApp) isAuthenticated() bool {
	_, err := api.GetContactList(ma.client)
	return err == nil
}

func (ma *myApp) authenticate() {
	ma.client = NewClientFromPrefs(ma.prefs.secrets)
	if ma.isAuthenticated() {
		return
	}
	apiKeyEntry := widget.NewEntryWithData(ma.prefs.secrets.apiKey)
	apiKeyEntry.Validator = nil
	apiKeyEntry.Password = true
	apiSecretEntry := widget.NewEntryWithData(ma.prefs.secrets.apiSecret)
	apiSecretEntry.Validator = nil
	apiSecretEntry.Password = true
	formContents := container.NewVBox(apiKeyEntry, apiSecretEntry)
	var auth *api.Authorization
	form := dialog.NewCustomConfirm("Your Flickr API Credentials", "Authenticate", "Abort", formContents, func(confirmed bool) {
		if confirmed {
			ma.prefs.secrets.apiKey.Set(apiKeyEntry.Text)
			ma.prefs.secrets.apiSecret.Set(apiSecretEntry.Text)
			ma.client = flickr.NewFlickrClient(apiKeyEntry.Text, apiSecretEntry.Text)

			auth = api.NewAuth(
				api.Secrets{
					ApiKey:    apiKeyEntry.Text,
					ApiSecret: apiSecretEntry.Text,
				})
			uri, err := auth.GetUrl()
			if err != nil {
				fyne.LogError("Getting Auth URL: ", err)
				dialog.NewError(err, ma.window).Show()
				return
			}

			authURI, err := url.Parse(uri)
			if err != nil {
				fyne.LogError("Parsing authentication URI", err)
				dialog.NewError(err, ma.window).Show()
				return
			}

			err = ma.app.OpenURL(authURI)
			if err != nil {
				dialog.NewError(err, ma.window).Show()
				fyne.LogError(fmt.Sprintf("Calling URL.open on '%s'", authURI), err)
				return
			}
			confirmationEntry := widget.NewEntry()
			formContents = container.NewVBox(confirmationEntry)
			form := dialog.NewCustomConfirm("Your Flickr Authorization Code", "OK", "Abort", formContents, func(confirmed bool) {
				if confirmed {
					ma.prefs.secrets.apiKey.Set(apiKeyEntry.Text)
					ma.prefs.secrets.apiSecret.Set(apiSecretEntry.Text)
					ma.client = flickr.NewFlickrClient(apiKeyEntry.Text, apiSecretEntry.Text)
					err := auth.GetAccessToken(confirmationEntry.Text)
					if err != nil {
						dialog.NewError(err, ma.window).Show()
						fyne.LogError("calling GetAccessToken", err)
						return
					}

					ma.prefs.secrets.accessToken.Set(auth.Secrets.AccessToken)
					ma.prefs.secrets.oAuthToken.Set(auth.Secrets.OAuthToken)
					ma.prefs.secrets.oAuthSecret.Set(auth.Secrets.OAuthSecret)
					ma.client.OAuthToken = auth.Client.OAuthToken
					ma.client.OAuthTokenSecret = auth.Client.OAuthTokenSecret
					ma.client.Id = auth.Client.Id

					r, err := test.Login(auth.Client) //ma.client)
					fmt.Println(r)
					if err != nil {
						dialog.NewError(err, ma.window).Show()
						fyne.LogError("testing login check", err)
						return
					}

				}
			}, ma.window)
			form.Show()
		}
	}, ma.window)

	form.Resize(fyne.Size{Width: 300, Height: 300})
	form.Show()
}

// func (ma *myApp) forgetCredentials() {
// 	dialog.NewConfirm("Log out", "Logging out will remove your authentication data", func(b bool) {
// 		if b {
// 			ClearCredentialsPrefs()
// 			ma.setAuthMenuStatus()
// 			ma.SetFollowedTags([]*mastodon.FollowedTag{})
// 		}
// 	}, ma.window).Show()
// }

// // getAuthCode allows the user to input the Authentication Token provided by Mastodon into the preferences
// func (ma *myApp) getAuthCode() {
// 	accessTokenEntry := widget.NewEntry()
// 	accessTokenEntry.Validator = nil
// 	dialog.NewForm("Authorization Code", "Save", "Cancel", []*widget.FormItem{
// 		{
// 			Text:     "Authorization Code",
// 			Widget:   accessTokenEntry,
// 			HintText: "XXX-XXX-XXX",
// 		}},
// 		func(confirmed bool) {
// 			if confirmed {
// 				c := NewClientFromPrefs(ma.prefs)
// 				// fmt.Printf("After authorizing, client is \n%+v\n", c.Config)
// 				err := c.AuthenticateToken(context.Background(), accessTokenEntry.Text, "urn:ietf:wg:oauth:2.0:oob")
// 				if err != nil {
// 					dialog.NewError(err, ma.window).Show()
// 					fyne.LogError("Authenticating token", err)
// 					return
// 				}
// 				_ = ma.prefs.AccessToken.Set(c.Config.AccessToken)
// 				ma.setAuthMenuStatus()
// 				ma.refreshFollowedTags()
// 			}
// 		},
// 		ma.window).Show()
// }
