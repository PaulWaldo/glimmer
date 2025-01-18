package mocks

import (
	"fmt"
	"gopkg.in/masci/flickr.v3"
	"net/http"
	"net/http/httptest"
)

func NewFlickrMock(statusCode int, responseBody string) (*httptest.Server, *flickr.FlickrClient) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(statusCode)
		fmt.Fprint(w, responseBody)
	}))

	client := flickr.NewFlickrClient("api-key", "api-secret")
	client.Args.Set("base_url", mockServer.URL)

	return mockServer, client
}
