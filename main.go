package glimmer

import (
	"fmt"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"gopkg.in/masci/flickr.v3"
	"gopkg.in/masci/flickr.v3/photosets"
	"gopkg.in/masci/flickr.v3/test"
)

// type authInfo struct {
// 	apiKey      binding.String
// 	apiSecret   binding.String
// 	accessToken binding.String
// 	oauthToken  binding.String
// 	oauthSecret binding.String
// }

// func NewAuthInfo(app fyne.App) *authInfo {
// 	auth := &authInfo{
// 		apiKey:      binding.BindPreferenceString("apiKey", app.Preferences()),
// 		apiSecret:   binding.BindPreferenceString("apiSecret", app.Preferences()),
// 		accessToken: binding.BindPreferenceString("accessToken", app.Preferences()),
// 		oauthToken:  binding.BindPreferenceString("oauthToken", app.Preferences()),
// 		oauthSecret: binding.BindPreferenceString("oauthSecret", app.Preferences()),
// 	}
// 	if _, err := auth.apiKey.Get(); err != nil {
// 		auth.apiKey.Set("96ba2e57a2c76d40b70b6a3c491015da")
// 		auth.apiSecret.Set("c9620849b48fee91")
// 	}
// 	return auth
// }

// func authenticate(ai authInfo, client *flickr.FlickrClient) error {
// 	// first, get a request token
// 	requestTok, err := flickr.GetRequestToken(client)
// 	if err != nil {
// 		fmt.Println("GetRequestToken: ", err)
// 		return err
// 	}

// 	url, err := flickr.GetAuthorizeUrl(client, requestTok)
// 	if err != nil {
// 		fmt.Println("GetAuthorizeUrl: ", err)
// 		return err
// 	}

// 	fmt.Println("Authorize at ", url)
// 	fmt.Println("Enter confirmation code")
// 	var code string
// 	_, err = fmt.Scanln(&code)
// 	if err != nil {
// 		fmt.Println("Reading confirmation code: ", err)
// 		return err
// 	}

// 	// finally, get the access token, setup the client and start making requests
// 	accessTok, err := flickr.GetAccessToken(client, requestTok, code)
// 	if err != nil {
// 		fmt.Println("Getting access token: ", err)
// 		return err
// 	}
// 	ai.accessToken.Set(accessTok)
// 	client.OAuthToken = accessTok.OAuthToken
// 	client.OAuthTokenSecret = accessTok.OAuthTokenSecret
// 	return nil
// }

// type prefs struct{}

// func (p prefs) makeUI() fyne.CanvasObject {
// 	return container.NewVBox(widg)
// }

// When using the masci/flickr Go package for Flickr API authentication, you typically go through an OAuth authentication process that involves obtaining access tokens. To persist authentication across application runs, you can store the following data:
//  Access Token: This is the token that allows your application to make authenticated requests on behalf of the user. It is usually obtained after the user has authenticated and authorized your application.
// Access Token Secret: This is a secret associated with the access token. It is used in conjunction with the access token to sign requests.
// User ID: This is the unique identifier for the user on Flickr. It can be useful for making user-specific requests.
//  Refresh Token (if applicable): Some OAuth implementations provide a refresh token that can be used to obtain a new access token when the current one expires. However, Flickr's OAuth implementation does not typically use refresh tokens, so this may not be relevant.

// func NewFlickrClient(app fyne.App) (flickr.FlickrClient, error) {
// 	authInfo := NewAuthInfo(app)
// 	apiKey, err := authInfo.apiKey.Get()
// 	if err != nil {
// 		return flickr.FlickrClient{}, err
// 	}
// 	apiSecret, err := authInfo.apiSecret.Get()
// 	if err != nil {
// 		return flickr.FlickrClient{}, err
// 	}
// 	client := flickr.NewFlickrClient(apiKey, apiSecret)
// 	err := authenticate(*authInfo, client)
// }

func main() {
	a := app.NewWithID("com.github.PaulWaldo.flickr")
	// auth := NewAuthInfo(a)
	// ma := myApp{}
	// ma.apiKey, _ = os.LookupEnv("API_KEY")
	// ma.apiSecret, _ = os.LookupEnv("API_SECRET")
	w := a.NewWindow("My Fyne App")
	w.SetContent(widget.NewLabel("Basic Fyne app created."))

	// client := flickr.NewFlickrClient("96ba2e57a2c76d40b70b6a3c491015da", "c9620849b48fee91")
	// Works!
	client := flickr.NewFlickrClient("aba7e80c0aae6a896446b2046f1fbddb", "ec104b69e6afcfe2")

	// a.Preferences().SetString("apikey", "96ba2e57a2c76d40b70b6a3c491015da")
	// a.Preferences().SetString("apisecret", "c9620849b48fee91")
	// first, get a request token
	requestTok, err := flickr.GetRequestToken(client)
	if err != nil {
		fmt.Println("GetRequestToken: ", err)
		return
	}

	url, err := flickr.GetAuthorizeUrl(client, requestTok)
	if err != nil {
		fmt.Println("GetAuthorizeUrl: ", err)
		return
	}

	fmt.Println("Authorize at ", url)
	fmt.Println("Enter confirmation code")
	var code string
	_, err = fmt.Scanln(&code)
	if err != nil {
		fmt.Println("Reading confirmation code: ", err)
		return
	}

	// finally, get the access token, setup the client and start making requests
	accessTok, err := flickr.GetAccessToken(client, requestTok, code)
	if err != nil {
		fmt.Println("Getting access token: ", err)
		return
	}
	client.OAuthToken = accessTok.OAuthToken
	client.OAuthTokenSecret = accessTok.OAuthTokenSecret

	presponse, err := photosets.Create(client, "My Set", "Description", "primary_photo_id")
	if err != nil {
		fmt.Println("Error creating photoset ", err)
		// return
	}
	fmt.Println("New photoset created:", presponse.Set.Id)

	eresponse, err := test.Echo(client)
	if err != nil {
		fmt.Println("Error testing echo: ", err)
		// return
	}
	fmt.Printf("Echo test response = %v\n", eresponse)

	response, err := test.Login(client)
	if err != nil {
		fmt.Println("Error testing login: ", err)
		// return
	}
	fmt.Printf("Login test response = %v\n", response)

	client2 := &flickr.FlickrClient{
		ApiKey:           client.ApiKey,
		ApiSecret:        client.ApiSecret,
		OAuthToken:       client.OAuthToken,
		OAuthTokenSecret: client.OAuthTokenSecret,
	}
	response, err = test.Login(client2)
	if err != nil {
		fmt.Println("Error testing login2: ", err)
		return
	}
	fmt.Printf("Login test response2 = %v\n", response)

	w.ShowAndRun()
}
