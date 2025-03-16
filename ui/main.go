package ui

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"github.com/PaulWaldo/glimmer/api"
	"gopkg.in/masci/flickr.v3"
	"gopkg.in/masci/flickr.v3/groups"
)

const AppID = "com.github.PaulWaldo.glimmer"

type myApp struct {
	prefs                 appPrefs
	app                   fyne.App
	window                fyne.Window
	client                *flickr.FlickrClient
	loginMenu, logoutMenu *fyne.MenuItem
	vs                    *ViewStack
	tabsUI                *apptabs
	userNsID              string
	userName              string
	fullName              string
	usersGroups           []groups.Group
	groupPhotosChan       chan struct{}
	usersGroupPhotos      []api.UsersGroupPhotos
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
	ma.groupPhotosChan = make(chan struct{})

	ma.prefs = NewPreferences(ma.app)
	ma.userNsID, _ = ma.prefs.userNsID.Get()
	ma.userName, _ = ma.prefs.userName.Get()

	ma.client = NewClientFromPrefs(ma.prefs)
	ma.window = ma.app.NewWindow("Glimmer")
	ma.loginMenu = fyne.NewMenuItem("Log In", ma.authenticate)
	ma.logoutMenu = fyne.NewMenuItem("Log Out", ma.forgetCredentials)
	ma.tabsUI = &apptabs{ma: ma}
	ma.tabsUI.makeUI()

	ma.window.SetMainMenu(fyne.NewMainMenu(
		fyne.NewMenu("Server", ma.loginMenu, ma.logoutMenu)),
	)
	ma.setAuthMenuStatus()

	at := &apptabs{ma: ma}
	at.makeUI()

	if ma.isLoggedIn() {
		// Fetch group photos synchronously.  Consider making this asynchronous later.
		go func() {
			var err error
			fmt.Println("Starting fetching group photos")
			client := api.CloneClient(ma.client)
			client.Args.Set("per_page", strconv.Itoa(10))
			params := map[string]string{"per_page": strconv.Itoa(10)}
			err = api.GetUsersGroupPhotos(client, ma.userNsID, params, &ma.usersGroups, &ma.usersGroupPhotos)
			if err != nil {
				fyne.LogError("getting users group photos", err)
				// Handle the error appropriately, e.g., display an error message.
				// For now, we'll just log the error and continue.
			}
			fmt.Println("Group photos fetched:", len(ma.usersGroupPhotos))
			close(ma.groupPhotosChan)
		}()
	} else {
		ma.authenticate()
	}

	cp := contactPhotos{ma: ma}
	ma.tabsUI.appTabs.SetItems([]*container.TabItem{
		container.NewTabItem("Contacts", cp.makeUI()),
		container.NewTabItem("Groups", container.NewVBox(at.appTabs)),
	})
	ma.window.SetContent(ma.tabsUI.appTabs)
	ma.window.Resize(fyne.Size{
		Width:  GridSizeWidth*2 + theme.Padding()*3,
		Height: GridSizeHeight*2 + theme.Padding()*3,
	})
	fmt.Println("Contact Photos container created")
	fmt.Println("All photos scheduled")
	go func() {
		close(runloopStarted)
	}()
	ma.window.ShowAndRun()
}

func (ma *myApp) setAuthMenuStatus() {
	ma.logoutMenu.Disabled = !ma.isLoggedIn()
	ma.loginMenu.Disabled = ma.isLoggedIn()
}
