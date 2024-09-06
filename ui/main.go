package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/PaulWaldo/glimmer/api"
	"gopkg.in/masci/flickr.v3"
)

const AppID = "com.github.PaulWaldo.glimmer"

type myApp struct {
	prefs                 appPrefs
	app                   fyne.App
	window                fyne.Window
	client                *flickr.FlickrClient
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
	s, err = ma.prefs.secrets.oAuthTokenSecret.Get()
	if err != nil || len(s) == 0 {
		return false
	}

	return true
}

func Run() {
	ma := &myApp{}
	ma.app = app.NewWithID(AppID)
	ma.prefs = NewPreferences(ma.app)
	ma.client = NewClientFromPrefs(ma.prefs)
	ma.window = ma.app.NewWindow("Glimmer")
	ma.loginMenu = fyne.NewMenuItem("Log In", ma.authenticate)
	ma.logoutMenu = fyne.NewMenuItem("Log Out", ma.forgetCredentials)
	ma.window.SetMainMenu(fyne.NewMainMenu(
		fyne.NewMenu("Server", ma.loginMenu, ma.logoutMenu)),
	)
	ma.setAuthMenuStatus()
	ma.window.Resize(fyne.Size{Width: 1000, Height: 500})
	if ma.isLoggedIn() {
	} else {
		ma.authenticate()
	}
	cp := contactPhotos{ma: ma}
	photos, err := api.GetContactPhotos(ma.client)
	// x, err := api.Feed(ma.client)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("\n\n\nFeed:\n%#v\n", x)
	if err != nil {
		panic(err)
	}
	cp.photos = photos.Photos.Photos
	ma.window.SetContent(cp.makeUI())
	ma.window.ShowAndRun()
}

func (ma *myApp) setAuthMenuStatus() {
	ma.logoutMenu.Disabled = !ma.isLoggedIn()
	ma.loginMenu.Disabled = ma.isLoggedIn()
}
