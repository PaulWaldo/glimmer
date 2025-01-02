package mocks

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"gopkg.in/masci/flickr.v3"
)

func NewFlickrMock(statusCode int, responseBody string) (*httptest.Server, *flickr.FlickrClient) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(statusCode)
		fmt.Fprint(w, responseBody)
	}))

	client := flickr.NewFlickrClient("api-key", "api-secret")
	client.BaseURL = mockServer.URL

	return mockServer, client
}
