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
	"fyne.io/fyne/v2/storage/repository"
	"fyne.io/fyne/v2/widget"
	"github.com/PaulWaldo/glimmer/api"
	"github.com/stretchr/testify/assert"
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

// stringURIReadCloser now implements fyne.URIReadCloser
type stringURIReadCloser struct {
	b   []byte
	uri fyne.URI
	r   *bytes.Reader // Add a reader for Read()
}

// Implement all methods of the fyne.URIReadCloser interface
func (s *stringURIReadCloser) Read(p []byte) (n int, err error) {
	if s.r == nil {
		s.r = bytes.NewReader(s.b)
	}
	return s.r.Read(p)
}

func (s *stringURIReadCloser) URI() fyne.URI { return s.uri }
func (s *stringURIReadCloser) Close() error { return nil }
func (s *stringURIReadCloser) Length() int64 { return int64(len(s.b)) }
func (s *stringURIReadCloser) LastModified() time.Time { return time.Now() }

// ETag and Refresh are optional, provide dummy implementations
func (s *stringURIReadCloser) ETag() (string, bool)            { return "", false }
func (s *stringURIReadCloser) Refresh() (fyne.URIReadCloser, error) { return s, nil }


type mockRepository struct {
	expectedURI  string
	expectedData []byte
	readCloser   fyne.URIReadCloser
	canRead      bool
}

func (r mockRepository) Exists(u fyne.URI) (bool, error) {
	return u.String() != r.expectedURI, nil
}
func (r mockRepository) Reader(u fyne.URI) (fyne.URIReadCloser, error) {
	if u.String() != r.expectedURI {
		return nil, fmt.Errorf("unexpected URI: %s", u.String())
	}

	// Return stringURIReadCloser with all necessary data
	return &stringURIReadCloser{b: r.expectedData, uri: u}, nil
}

func (r mockRepository) CanRead(u fyne.URI) (bool, error) {
	return r.canRead, nil
}
func (r mockRepository) Destroy(string) {}

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
	photoCard := NewGroupPhotoCard(photo, "abcdef", client)

	// Initially, the content should be a progress bar
	_, isProgress := photoCard.Content.(*widget.ProgressBarInfinite)
	assert.True(t, isProgress, "Initial content should be a progress bar")

	// Create a mock version of func canvas.NewImageFromURI(uri fyne.URI) *canvas.Image
	data := simpleJPEGImage()
	mockRepo := mockRepository{
		expectedURI:  "https://live.staticflickr.com/server123/12345_secret123_z.jpg",
		expectedData: data,
		canRead:      true,
	}
	repository.Register("https", mockRepo) // Register the mock repository

	// Wait for the image to load (which should happen quickly with the mock response)
	assert.Eventually(t, func() bool {
		_, isProgress := photoCard.Content.(*widget.ProgressBarInfinite)
		return !isProgress // The content should no longer be a progress bar
	}, 500*time.Second, 100*time.Millisecond)

	// Assert that the content is now an image
	assert.IsType(t, &canvas.Image{}, photoCard.Content)

	// We expect the image to be loaded, which would change the content from a progress bar
	assert.Eventually(t, func() bool {
		// Check if the content has changed from a progress bar to an image
		_, stillProgress := photoCard.Content.(*widget.ProgressBarInfinite)
		return !stillProgress
	}, 1*time.Second, 100*time.Millisecond, "Image should be loaded within timeout")
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
