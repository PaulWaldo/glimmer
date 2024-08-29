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
)

func (ma *myApp) isAuthenticated() bool {
	_, err := api.GetContactList(ma.client)
	return err == nil
}

func (ma *myApp) authenticate() {
	ma.client = NewClientFromPrefs(ma.prefs)
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

			auth = api.NewAuthorizer()
			uri, err := auth.GetUrl(ma.client)
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
					err := auth.GetAccessToken(ma.client, confirmationEntry.Text)
					if err != nil {
						dialog.NewError(err, ma.window).Show()
						fyne.LogError("calling GetAccessToken", err)
						return
					}

					ma.UpdateSecrefPrefs()

					r, err := api.GetContactList(ma.client)
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

func (ma *myApp) forgetCredentials() {
	dialog.NewConfirm("Log out", "Logging out will remove your authentication data", func(b bool) {
		if b {
			ClearCredentialsPrefs()
			ma.setAuthMenuStatus()
		}
	}, ma.window).Show()
}
