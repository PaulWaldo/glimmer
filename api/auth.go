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

// type Secrets struct {
// 	ApiKey      string
// 	ApiSecret   string
// 	AccessToken string
// 	OAuthToken  string
// 	OAuthSecret string
// }

type Authorization struct {
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

func NewAuthorizer() *Authorization {
	return &Authorization{
		Authorizer: flickrAuthorizer{},
	}
}

func (a *Authorization) GetUrl(client *flickr.FlickrClient) (string, error) {
	var err error
	a.RequestToken, err = a.Authorizer.GetRequestToken(client)
	if err != nil {
		return "", fmt.Errorf("getting request token: %w", err)
	}

	url, err := a.Authorizer.GetAuthorizeUrl(client, a.RequestToken)
	if err != nil {
		return "", fmt.Errorf("getting authorization URL: %s", err)
	}

	return url, nil
}

func (a *Authorization) GetAccessToken(client *flickr.FlickrClient, confirmationCode string) error {
	accessTok, err := a.Authorizer.GetAccessToken(client, a.RequestToken, confirmationCode)
	if err != nil {
		return fmt.Errorf("getting access token: %w", err)
	}
	client.OAuthToken = accessTok.OAuthToken
	client.OAuthTokenSecret = accessTok.OAuthTokenSecret
	return nil

}
