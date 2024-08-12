package glimmer

type authInfo struct {
	apiKey      string
	apiSecret   string
	accessToken string
	oauthToken  string
	oauthSecret string
}

func (a authInfo) NeedsAuthentication() bool {
	return len(a.apiKey) == 0 ||
		len(a.apiSecret) == 0 ||
		len(a.oauthToken) == 0 ||
		len(a.oauthSecret) == 0
}
func (a authInfo) GetAuthorizeUrl() (string, error) {
	return "", nil
}

// // first, get a request token
// requestTok, err := flickr.GetRequestToken(client)
// if err != nil {
// 	fmt.Println("GetRequestToken: ", err)
// 	return
// }

// url, err := flickr.GetAuthorizeUrl(client, requestTok)
// if err != nil {
// 	fmt.Println("GetAuthorizeUrl: ", err)
// 	return
// }

// fmt.Println("Authorize at ", url)
