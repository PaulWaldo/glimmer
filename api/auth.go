package api

import (
	"fmt"

	"gopkg.in/masci/flickr.v3"
)

type Authorizer interface {
	GetRequestToken(client *flickr.FlickrClient) (*flickr.RequestToken, error)
	GetAuthorizeUrl(client *flickr.FlickrClient, reqToken *flickr.RequestToken) (string, error)
	GetAccessToken(client *flickr.FlickrClient, reqToken *flickr.RequestToken, oauthVerifier string) (*flickr.OAuthToken, error)
}

type Secrets struct {
	ApiKey      string
	ApiSecret   string
	AccessToken string
	OAuthToken  string
	OAuthSecret string
}

type Authorization struct {
	Secrets      Secrets
	Client       *flickr.FlickrClient
	Authorizer   Authorizer
	RequestToken *flickr.RequestToken
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

func NewAuth(secrets Secrets) *Authorization {
	return &Authorization{
		Secrets:    secrets,
		Client:     flickr.NewFlickrClient(secrets.ApiKey, secrets.ApiSecret),
		Authorizer: flickrAuthorizer{},
	}
}

// func NeedsAuthentication(a AuthInfo) bool {
// 	return len(a.ApiKey) == 0 ||
// 		len(a.ApiSecret) == 0 ||
// 		len(a.OAuthToken) == 0 ||
// 		len(a.OAuthSecret) == 0
// }

func (a *Authorization) GetUrl() (string, error) {
	var err error
	a.RequestToken, err = a.Authorizer.GetRequestToken(a.Client)
	if err != nil {
		return "", fmt.Errorf("getting request token: %w", err)
	}

	url, err := a.Authorizer.GetAuthorizeUrl(a.Client, a.RequestToken)
	if err != nil {
		return "", fmt.Errorf("getting authorization URL: %s", err)
	}

	return url, nil
}

func (a *Authorization) GetAccessToken(confirmationCode string) error {
	accessTok, err := a.Authorizer.GetAccessToken(a.Client, a.RequestToken, confirmationCode)
	if err != nil {
		return fmt.Errorf("getting access token: %w", err)
	}
	a.Client.OAuthToken = accessTok.OAuthToken
	a.Client.OAuthTokenSecret = accessTok.OAuthTokenSecret
	return nil

}
