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
}

type Authorize struct {
	secrets    Secrets
	client     *flickr.FlickrClient
	authorizer Authorizer
}

type flickrAuthorizer struct{}

func (a flickrAuthorizer) GetRequestToken(client *flickr.FlickrClient) (*flickr.RequestToken, error) {
	return flickr.GetRequestToken(client)
}

func (a flickrAuthorizer) GetAuthorizeUrl(client *flickr.FlickrClient, reqToken *flickr.RequestToken) (string, error) {
	return flickr.GetAuthorizeUrl(client, reqToken)
}

func NewAuth(secrets Secrets) *Authorize {
	return &Authorize{
		secrets:    secrets,
		client:     flickr.NewFlickrClient(secrets.ApiKey, secrets.ApiSecret),
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
	// a.client = flickr.NewFlickrClient(a.apiKey, a.apiSecret)
	requestTok, err := a.authorizer.GetRequestToken(a.client)
	if err != nil {
		return "", fmt.Errorf("getting request token: %w", err)
	}

	url, err := a.authorizer.GetAuthorizeUrl(a.client, requestTok)
	if err != nil {
		return "", fmt.Errorf("getting authorization URL: %s", err)
	}

	return url, nil
}
