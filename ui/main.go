package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"gopkg.in/masci/flickr.v3"
)

const AppID = "com.github.PaulWaldo.glimmer"

type myApp struct {
	prefs  appPrefs
	app    fyne.App
	window fyne.Window
	client *flickr.FlickrClient
	loginMenu, logoutMenu *fyne.MenuItem
}

func (ma *myApp) isLoggedIn() bool {
	s, err := ma.prefs.secrets.apiKey.Get()
	if err != nil || len(s) == 0 {
		return false
	}
	s, err = ma.prefs.secrets.apiSecret.Get()
	if err != nil || len(s) == 0 {
		return false
	}
	s, err = ma.prefs.secrets.oAuthToken.Get()
	if err != nil || len(s) == 0 {
		return false
	}
	s, err = ma.prefs.secrets.oAuthSecret.Get()
	if err != nil || len(s) == 0 {
		return false
	}

	return true
}


func Run() {
	ma := myApp{}
	ma.app = app.NewWithID(AppID)
	ma.prefs = NewPreferences(ma.app)
	ma.window = ma.app.NewWindow("Glimmer")
	ma.loginMenu = fyne.NewMenuItem("Log In", ma.authenticate)
	ma.logoutMenu = fyne.NewMenuItem("Log Out", ma.forgetCredentials)
	ma.window.SetMainMenu(fyne.NewMainMenu(
		fyne.NewMenu("Server", ma.loginMenu, ma.logoutMenu)),
	)
	ma.setAuthMenuStatus()
	// e := apiInfoEntry{}
	// ma.window.SetContent(e.makeUI())
	ma.window.Resize(fyne.Size{Width: 1000, Height: 500})
	if ma.isLoggedIn() {
	// 	ma.refreshFollowedTags()
	} else {
		ma.authenticate()
	}
	// ma.authenticate()
	ma.window.ShowAndRun()
}

func (ma *myApp) setAuthMenuStatus() {
	ma.logoutMenu.Disabled = !ma.isLoggedIn()
	ma.loginMenu.Disabled = ma.isLoggedIn()
}
