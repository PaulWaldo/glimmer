package glimmer

import (
	"fmt"

	"gopkg.in/masci/flickr.v3"
)

type Secrets struct {
	ApiKey      string
	ApiSecret   string
	AccessToken string
	OAuthToken  string
	OAuthSecret string
}

type Authorizer interface {
	GetRequestToken(client *flickr.FlickrClient) (*flickr.RequestToken, error)
	GetAuthorizeUrl(client *flickr.FlickrClient, reqToken *flickr.RequestToken) (string, error)
	GetAccessToken(client *flickr.FlickrClient, reqToken *flickr.RequestToken, oauthVerifier string) (*flickr.OAuthToken, error)
}

type Authorize struct {
	secrets      Secrets
	Client       *flickr.FlickrClient
	authorizer   Authorizer
	requestToken *flickr.RequestToken
}

type flickrAuthorizer struct{}

func (a flickrAuthorizer) GetRequestToken(client *flickr.FlickrClient) (*flickr.RequestToken, error) {
	return flickr.GetRequestToken(client)
}

func (a flickrAuthorizer) GetAuthorizeUrl(client *flickr.FlickrClient, reqToken *flickr.RequestToken) (string, error) {
	return flickr.GetAuthorizeUrl(client, reqToken)
}

func (a flickrAuthorizer) GetAccessToken(client *flickr.FlickrClient, reqToken *flickr.RequestToken, oauthVerifier string) (*flickr.OAuthToken, error) {
	return flickr.GetAccessToken(client, reqToken, oauthVerifier)
}

func NewAuth(secrets Secrets) *Authorize {
	return &Authorize{
		secrets:    secrets,
		Client:     flickr.NewFlickrClient(secrets.ApiKey, secrets.ApiSecret),
		authorizer: flickrAuthorizer{},
	}
}

// func NeedsAuthentication(a AuthInfo) bool {
// 	return len(a.ApiKey) == 0 ||
// 		len(a.ApiSecret) == 0 ||
// 		len(a.OAuthToken) == 0 ||
// 		len(a.OAuthSecret) == 0
// }

func (a *Authorize) GetUrl() (string, error) {
	var err error
	a.requestToken, err = a.authorizer.GetRequestToken(a.Client)
	if err != nil {
		return "", fmt.Errorf("getting request token: %w", err)
	}

	url, err := a.authorizer.GetAuthorizeUrl(a.Client, a.requestToken)
	if err != nil {
		return "", fmt.Errorf("getting authorization URL: %s", err)
	}

	return url, nil
}

func (a *Authorize) GetAccessToken(confirmationCode string) error {
	accessTok, err := a.authorizer.GetAccessToken(a.Client, a.requestToken, confirmationCode)
	if err != nil {
		return fmt.Errorf("getting access token: %w", err)
	}
	a.Client.OAuthToken = accessTok.OAuthToken
	a.Client.OAuthTokenSecret = accessTok.OAuthTokenSecret
	return nil

}
