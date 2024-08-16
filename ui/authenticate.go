package ui

import "github.com/PaulWaldo/glimmer"

const PREF_KEY_API_KEY = "ApiKey"
const PREF_KEY_API_SECRET = "ApiSecret"
const PREF_KEY_ACCESS_TOKEN = "AccessToken"
const PREF_KEY_OAUTH_TOKEN = "OAuthToken"
const PREF_KEY_OAUTH_SECRET = "OAuthSecret"

func (ma *myApp) LoadSecrets() glimmer.Secrets {
	sec := glimmer.Secrets{}
	pref := ma.app.Preferences()
	sec.ApiKey = pref.String(PREF_KEY_API_KEY)
	sec.ApiSecret = pref.String(PREF_KEY_API_SECRET)
	sec.AccessToken = pref.String(PREF_KEY_ACCESS_TOKEN)
	sec.OAuthToken = pref.String(PREF_KEY_OAUTH_TOKEN)
	sec.OAuthSecret = pref.String(PREF_KEY_OAUTH_SECRET)
	return sec
}

// func (ma *myApp) authenticate() {
// 	serverUrlEntry := widget.NewEntryWithData(ma.prefs.MastodonServer)
// 	serverUrlEntry.Validator = nil
// 	formContents := container.NewVBox(serverUrlEntry)
// 	serverUrlEntry.SetPlaceHolder("https://MyMastodonServer.com")
// 	form := dialog.NewCustomConfirm("URL of your Mastodon server", "Authenticate", "Abort", formContents, func(confirmed bool) {
// 		if confirmed {
// 			val, _ := ma.prefs.MastodonServer.Get()
// 			app, err := mastodon.RegisterApp(context.Background(), app.NewAuthenticationConfig(val))
// 			if err != nil {
// 				fyne.LogError("Calling (Mastodon) App registration", err)
// 				ma.serverText.Text = fmt.Sprintf("Error contacting Mastodon server %s", val)
// 				dialog.NewError(err, ma.window).Show()
// 				return
// 			}
// 			_ = ma.prefs.ClientID.Set(app.ClientID)
// 			_ = ma.prefs.ClientSecret.Set(app.ClientSecret)
// 			authURI, err := url.Parse(app.AuthURI)
// 			if err != nil {
// 				fyne.LogError("Parsing authentication URI", err)
// 				dialog.NewError(err, ma.window).Show()
// 				return
// 			}
// 			err = ma.app.OpenURL(authURI)
// 			if err != nil {
// 				dialog.NewError(err, ma.window).Show()
// 				fyne.LogError(fmt.Sprintf("Calling URL.open on '%s'", authURI), err)
// 				return
// 			}
// 			ma.getAuthCode()
// 		}
// 	}, ma.window)
// 	form.Resize(fyne.Size{Width: 300, Height: 300})
// 	form.Show()
// }

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
