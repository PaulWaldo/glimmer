package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
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
	vs                    *ViewStack
	userNsID              string
	userName              string
	fullName              string
}

func (ma *myApp) logAuth(marker string) {
	fmt.Printf("**********************************************\n%s\n", marker)
	fmt.Printf("API Key: %s\n", ma.client.ApiKey)
	fmt.Printf("API Secret: %s\n", ma.client.ApiSecret)
	fmt.Printf("OAuthToken: %s\n", ma.client.OAuthToken)
	fmt.Printf("OAuthTokenSecret: %s\n", ma.client.OAuthTokenSecret)

	prefs := ma.prefs
	apiKey, _ := prefs.secrets.apiKey.Get()
	apiSecret, _ := prefs.secrets.apiSecret.Get()
	oAuthToken, _ := prefs.secrets.oAuthToken.Get()
	oAuthSecret, _ := prefs.secrets.oAuthTokenSecret.Get()
	fmt.Printf("Prefs apiKey: %s\n", apiKey)
	fmt.Printf("Prefs apiSecret: %s\n", apiSecret)
	fmt.Printf("Prefs oAuthToken: %s\n", oAuthToken)
	fmt.Printf("Prefs oAuthSecret: %s\n", oAuthSecret)
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
	ma.vs = NewViewStack(ma.window)
	ma.window.SetMainMenu(fyne.NewMainMenu(
		fyne.NewMenu("Server", ma.loginMenu, ma.logoutMenu)),
	)
	ma.setAuthMenuStatus()
	// ma.logAuth("Before auth check")
	if ma.isLoggedIn() {
	} else {
		ma.authenticate()
	}

	cp := contactPhotos{ma: ma}
	photos, err := api.GetContactPhotos(ma.client)
	// ma.logAuth("main GetContactPhotos")

	if err != nil {
		fmt.Println(err)
		photos = &api.GetContactPhotosResponse{Photos: api.ContactPhotos{Photos: []api.Photo{}}}
	}
	// fmt.Printf("\n\n\nPhotos:\n%#v\n", photos)

	// Create placeholder images for each photo
	for _, photo := range photos.Photos.Photos {
		img := fyne.NewImage(nil)
		img.SetMinSize(fyne.Size{Width: 100, Height: 100}) // Set a minimum size for the placeholder
		cp.imageContainer.Add(img)
	}

	// Start a goroutine for each photo to download the image data
	for _, photo := range photos.Photos.Photos {
		go func(photo api.Photo) {
			// Download the image data
			data, err := api.DownloadImage(ma.client, photo.Url)
			if err != nil {
				// Handle error
				return
			}
			// Request a refresh of the placeholder image
			img := cp.imageContainer.Objects[cp.photos.IndexOf(photo)]
			img.Refresh()
			// Update the image data
			img.SetData(data)
		}(photo)
	}

	cp.photos = photos.Photos.Photos
	ma.vs.Push(cp.makeUI())
	ma.window.Resize(fyne.Size{
		Width:  GridSizeWidth*2 + theme.Padding()*3,
		Height: GridSizeHeight*2 + theme.Padding()*3,
	})
	ma.window.ShowAndRun()
}

func (ma *myApp) setAuthMenuStatus() {
	ma.logoutMenu.Disabled = !ma.isLoggedIn()
	ma.loginMenu.Disabled = ma.isLoggedIn()
}
