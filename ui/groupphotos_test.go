package ui

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"mime"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"
	"testing"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
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

// minimalJPEG represents a minimal valid 1x1 JPEG image
var minimalJPEG = []byte{
	0xff, 0xd8, 0xff, 0xe0, 0x00, 0x10, 0x4a, 0x46, 0x49, 0x46, 0x00, 0x01, 0x01, 0x00, 0x00, 0x01,
	0x00, 0x01, 0x00, 0x00, 0xff, 0xdb, 0x00, 0x43, 0x00, 0x08, 0x06, 0x06, 0x07, 0x06, 0x05, 0x08,
	0x07, 0x07, 0x07, 0x09, 0x09, 0x08, 0x0a, 0x0c, 0x14, 0x0d, 0x0c, 0x0b, 0x0b, 0x0c, 0x19, 0x12,
	0x13, 0x0f, 0x14, 0x1d, 0x1a, 0x1f, 0x1e, 0x1d, 0x1a, 0x1c, 0x1c, 0x20, 0x24, 0x2e, 0x27, 0x20,
	0x22, 0x2c, 0x23, 0x1c, 0x1c, 0x28, 0x37, 0x29, 0x2c, 0x30, 0x31, 0x34, 0x34, 0x34, 0x1f, 0x27,
	0x39, 0x3d, 0x38, 0x32, 0x3c, 0x2e, 0x33, 0x34, 0x32, 0xff, 0xc0, 0x00, 0x0b, 0x08, 0x00, 0x01,
	0x00, 0x01, 0x01, 0x01, 0x11, 0x00, 0xff, 0xc4, 0x00, 0x14, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x09, 0xff, 0xc4, 0x00, 0x14,
	0x10, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0xff, 0xda, 0x00, 0x08, 0x01, 0x01, 0x00, 0x00, 0x3f, 0x00, 0x7f, 0x00, 0xff, 0xd9,
}

func simpleJPEGImage() []byte {
	simpleImage := image.NewGray(image.Rect(0, 0, 1, 1))
	// Encode the grayscale image to JPEG
	var buf bytes.Buffer                       // Use a buffer to store the encoded JPEG data
	err := jpeg.Encode(&buf, simpleImage, nil) // Use nil for default JPEG options
	if err != nil {
		fyne.LogError("Failed to encode JPEG", err)
		return []byte{}
	}

	// buf.Bytes() now contains the JPEG encoded image data
	return buf.Bytes()
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
	transport.responses["flickr.photos.getInfo"] = mockResponse{statusCode: 200, body: photoInfoResponse, isImage: false}

	// Add mock response for the image URL
	imageURL := "https://live.staticflickr.com/server123/12345_secret123_z.jpg"
	transport.responses[imageURL] = mockResponse{
		statusCode: 200,
		body:       string(simpleJPEGImage()),
		isImage:    true,
	}

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

	// Create a mock version of func canvas.NewImageFromURI(uri fyne.URI) *canvas.Image
	mockNewImageFromURI := func(uri fyne.URI) *canvas.Image {
		assert.Equal(t, "https://live.staticflickr.com/server123/12345_secret123_z.jpg", uri.String())
		return canvas.NewImageFromReader(strings.NewReader(string(simpleJPEGImage())), "fakeImage.jpg")
	}
	NewImageFromURI = mockNewImageFromURI

	// Wait for the image to load (which should happen quickly with the mock response)
	assert.Eventually(t, func() bool {
		_, isProgress := photoCard.Content.(*widget.ProgressBarInfinite)
		return !isProgress // The content should no longer be a progress bar
	}, 2*time.Second, 10*time.Millisecond) // Adjust timeout as needed

	// Assert that the content is now an image
	assert.IsType(t, &canvas.Image{}, photoCard.Content)

	// We expect the image to be loaded, which would change the content from a progress bar
	assert.Eventually(t, func() bool {
		// Check if the content has changed from a progress bar to an image
		_, stillProgress := photoCard.Content.(*widget.ProgressBarInfinite)
		return !stillProgress
	}, 500*time.Second, 100*time.Millisecond, "Image should be loaded within timeout")
}

// mockTransport and mockResponse types for testing
type mockTransport struct {
	responses map[string]mockResponse
}

type mockResponse struct {
	statusCode int
	body       string
	isImage    bool
}

func (t *mockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Handle GET requests (for image downloads)
	if req.Method == "GET" {
		urlStr := req.URL.String()
		if resp, ok := t.responses[urlStr]; ok {
			header := make(http.Header)
			if resp.isImage {
				header.Set("Content-Type", "image/jpeg")
			}
			return &http.Response{
				StatusCode: resp.statusCode,
				Body:       io.NopCloser(strings.NewReader(resp.body)),
				Header:     header,
			}, nil
		}
		return nil, fmt.Errorf("no mock response for GET request: %s", urlStr)
	}

	// Handle POST requests (for Flickr API)
	if req.Method == "POST" {
		if err := req.ParseForm(); err != nil {
			return nil, err
		}

		method := req.FormValue("method")
		groupID := req.FormValue("group_id")

		if req.Method == "POST" && method == "" && groupID == "" {
			mediaType, params, err := mime.ParseMediaType(req.Header.Get("Content-Type"))
			if err != nil {
				return nil, fmt.Errorf("parsing media type: %w", err)
			}
			if strings.HasPrefix(mediaType, "multipart/") {
				mr := multipart.NewReader(req.Body, params["boundary"])
				values := make(map[string][]string)
				for {
					p, err := mr.NextPart()
					if err == io.EOF {
						break
					}
					if err != nil {
						return nil, fmt.Errorf("reading next part: %w", err)
					}
					sl, err := io.ReadAll(p)
					if err != nil {
						return nil, fmt.Errorf("reading all from part: %w", err)
					}
					values[p.FormName()] = append(values[p.FormName()], string(sl))
				}
				req.Form = url.Values(values)
				method = req.FormValue("method")
				groupID = req.FormValue("group_id")
			}
		}

		if method == "" && groupID == "" {
			err := req.ParseMultipartForm(1024 * 1024) // Adjust limit as needed
			if err != nil {
				return nil, err
			}

			method = req.FormValue("method")
			groupID = req.FormValue("group_id")

			if method == "" || groupID == "" {
				return nil, fmt.Errorf("method or group_id not found in multipart form")
			}
		}

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

	return nil, fmt.Errorf("unsupported HTTP method: %s", req.Method)
}
