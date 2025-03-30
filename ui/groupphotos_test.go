package ui

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/PaulWaldo/glimmer/api"
	"github.com/stretchr/testify/assert" // Added for assertions
	"gopkg.in/masci/flickr.v3"
	"gopkg.in/masci/flickr.v3/groups"
)

func TestSetGroups(t *testing.T) {
	p := &groupPhotosUI{
		ma:       &myApp{}, // Initialize myApp
		cardByID: make(map[string]*fyne.CanvasObject),
	}
	testGroups := []groups.Group{
		{Nsid: "1", Name: "Group 1"},
		{Nsid: "2", Name: "Group 2"},
	}

	p.setGroups(testGroups)

	assert.Equal(t, len(testGroups), len(p.groupCards))
	for _, group := range testGroups {
		card := (*p.cardByID[group.Nsid]).(*GroupCard) // Type assertion
		assert.NotNil(t, card)
		assert.Equal(t, group.Name, card.Title)
	}
}

// TestNewGroupPhotoCard tests the creation of a photo card for group photos
func TestNewGroupPhotoCard(t *testing.T) {
	// Create a mock transport
	transport := &mockTransport{
		responses: make(map[string]mockResponse),
	}

	// Add mock response for photos.getInfo (needed for loadImage)
	photoInfoResponse := `<?xml version="1.0" encoding="utf-8" ?>
		<rsp stat="ok">
			<photo id="12345" secret="secret123" server="server123" farm="1" title="Test Photo">
				<owner nsid="owner123" username="testuser" />
			</photo>
		</rsp>`
	transport.responses["flickr.photos.getInfo"] = mockResponse{statusCode: 200, body: photoInfoResponse}

	// Create client with mock transport
	client := flickr.GetTestClient()
	client.HTTPClient = &http.Client{Transport: transport}

	photo := api.Photo{
		ID:       "12345",
		Owner:    "owner123",
		Secret:   "secret123",
		Server:   "server123",
		Username: "testuser",
		Title:    "Test Photo",
	}

	// Create a group photo card
	photoCard := NewGroupPhotoCard(photo, client)

	// Initially, the content should be a progress bar
	_, isProgress := photoCard.Content.(*widget.ProgressBarInfinite)
	assert.True(t, isProgress, "Initial content should be a progress bar")

	// We expect the image to be loaded, which would change the content from a progress bar
	assert.Eventually(t, func() bool {
		// Check if the content has changed from a progress bar to an image
		_, stillProgress := photoCard.Content.(*widget.ProgressBarInfinite)
		return !stillProgress
	}, 3*time.Second, 100*time.Millisecond, "Image should be loaded within timeout")
}

// mockTransport and mockResponse types for testing
type mockTransport struct {
	responses map[string]mockResponse
}

type mockResponse struct {
	statusCode int
	body       string
}

func (t *mockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if err := req.ParseForm(); err != nil {
		return nil, err
	}

	method := req.FormValue("method")

	response, ok := t.responses[method]
	if !ok {
		return nil, fmt.Errorf("no mock response found for method %q", method)
	}

	return &http.Response{
		StatusCode: response.statusCode,
		Body:       io.NopCloser(strings.NewReader(response.body)),
		Header:     make(http.Header),
	}, nil
}
