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

// func loadImage(card *PhotoCard) {
// 	// fmt.Println("Sleep start")
// 	time.Sleep(time.Second * time.Duration(rand.Int64N(10))) // Simulate a really long download
// 	// fmt.Println("Waking up")
// 	resp, err := photos.GetInfo(card.client, card.photo.Id, card.photo.Secret)
// 	if err != nil {
// 		fyne.LogError("Failed to get photo info", err)
// 		return
// 	}
// 	card.info = resp.Photo
// 	photoUrl := fmt.Sprintf("https://live.staticflickr.com/%s/%s_%s_%s.jpg", card.info.Server, card.info.Id, card.info.Secret, "z")
// 	uri, err := storage.ParseURI(photoUrl)
// 	if err != nil {
// 		fyne.LogError("parsing url", err)
// 		card.Content = widget.NewLabel("Failed to load image")
// 		return
// 	}
// 	fmt.Println("Downloading ", uri)
// 	// fmt.Printf("\"%s\",\n", uri)
// 	image := canvas.NewImageFromURI(uri)
// 	if image == nil || image.Resource == nil {
// 		panic("Image is nil")
// 	}
// 	// fmt.Printf("Image size is %d\n", len(image.Resource.Content()))
// 	image.FillMode = canvas.ImageFillContain
// 	// fmt.Println("Got ", uri)
// 	card.SetContent(image)
// }

var runloopStarted = false

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

	// // val, _ := ma.prefs.userName.Get()
	// // groups, err := api.GetGroups(ma.client, val)
	// // if err != nil {
	// // 	panic(err)
	// // }
	// // fmt.Printf("\n\n\nGroups:\n%#v\n", groups)

	// x, err := api.Feed(ma.client)
	// // if err != nil {
	// // 	panic(err)
	// // }
	// fmt.Printf("\n\n\nFeed:\n%#v\n", x)

	cp.photos = photos.Photos.Photos
	ma.window.SetContent(cp.makeUI())
	// ma.vs.Push(cp.makeUI())
	ma.window.Resize(fyne.Size{
		Width:  GridSizeWidth*2 + theme.Padding()*3,
		Height: GridSizeHeight*2 + theme.Padding()*3,
	})
	fmt.Println("Contact Photos container created")
	fmt.Println("All photos scheduled")
	ma.window.ShowAndRun()
}

func (ma *myApp) setAuthMenuStatus() {
	ma.logoutMenu.Disabled = !ma.isLoggedIn()
	ma.loginMenu.Disabled = ma.isLoggedIn()
}
