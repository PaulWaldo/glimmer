package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/PaulWaldo/glimmer/api"
	"gopkg.in/masci/flickr.v3/photos"
)

type PhotoCard struct {
	widget.BaseWidget
	PhotoInfo *photos.PhotoInfoResponse
	Title     string
	UserName  string
	// Title     *widget.Label
	// UserName  *widget.Label
	// image     *canvas.Image
}

func NewPhotoCard(title string, userName string, info *photos.PhotoInfoResponse) *PhotoCard {
	card := &PhotoCard{
		Title:     title,
		UserName:  userName,
		PhotoInfo: info,
	}
	// card := &PhotoCard{
	// 	Title:     widget.NewLabel(title),
	// 	UserName:  widget.NewLabel(userName),
	// 	PhotoInfo: info,
	// }
	card.ExtendBaseWidget(card)
	// card.title.Importance = widget.HighImportance
	// card.Title.Alignment = fyne.TextAlignLeading
	// card.Title.TextStyle = fyne.TextStyle{Bold: true}
	// card.Title.Wrapping = fyne.TextWrapWord
	// err := card.loadImage()
	// if err != nil {
	// 	return nil, err
	// }
	// return card, nil
	return card
}

// var _ *photoCardRenderer = *fyne.WidgetRenderer(nil)

type photoCardRenderer struct {
	photoCard *PhotoCard
	Title     *widget.Label
	UserName  *widget.Label
	image     *canvas.Image
}

func (r *photoCardRenderer) loadImage() error {
	photoUrl := fmt.Sprintf("https://live.staticflickr.com/%s/%s_%s_%s.jpg", r.photoCard.PhotoInfo.Photo.Server, r.photoCard.PhotoInfo.Photo.Id, r.photoCard.PhotoInfo.Photo.Secret, "z")
	uri, err := storage.ParseURI(photoUrl)
	if err != nil {
		return fmt.Errorf("parsing photoCard image url %q: %w", err, err)
	}
	fmt.Println("Downloading ", uri)
	r.image = canvas.NewImageFromURI(uri)
	r.image.Resize(fyne.NewSize(300, 300))
	r.image.Refresh()
	return nil
}

func (r *photoCardRenderer) MinSize() fyne.Size {
	return fyne.NewSize(
		r.image.MinSize().Width+r.Title.MinSize().Width+r.UserName.MinSize().Width,
		r.image.MinSize().Height+r.Title.MinSize().Height+r.UserName.MinSize().Height,
	)
}

func (r *photoCardRenderer) Layout(s fyne.Size) {
	padding := r.photoCard.Theme().Size(theme.SizeNamePadding)
	cellSize := r.MinSize()
	midX := cellSize.Width / 2
	imageSize := r.image.MinSize()
	r.image.Move(fyne.NewPos(midX-imageSize.Width/2, padding))
	nextY :=

		r.Title.Move(fyne.NewPos(50, 50))
	r.UserName.Move(fyne.NewPos(100, 100))

}

func (r *photoCardRenderer) Destroy() {
}

func (r *photoCardRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{r.image, r.Title, r.UserName}
}

func (r *photoCardRenderer) Refresh() {
	err := r.loadImage()
	if err != nil {
		fyne.LogError("refreshing", err)
	}
}

func (c *PhotoCard) CreateRenderer() fyne.WidgetRenderer {
	title := widget.NewLabel(c.Title)
	userName := widget.NewLabel(c.UserName)
	// card.title.Importance = widget.HighImportance
	title.Alignment = fyne.TextAlignLeading
	title.TextStyle = fyne.TextStyle{Bold: true}
	title.Wrapping = fyne.TextWrapWord
	// err := card.loadImage()
	// if err != nil {
	// 	return nil, err
	// }
	r := &photoCardRenderer{
		photoCard: c,
		Title:     title,
		UserName:  userName,
	}
	err := r.loadImage()
	if err != nil {
		fyne.LogError("refreshing", err)
	}
	return r
}

// func (c *PhotoCard) makeUI() *fyne.Container {
// 	return container.NewVBox(c.image, c.Title, c.UserName)
// }

type contactPhotos struct {
	ma         *myApp
	container  *fyne.Container
	title      *widget.Label
	photoList  *fyne.Container
	photos     []api.Photo
	photoCards []fyne.CanvasObject
}

func (p *contactPhotos) makeUI() *fyne.Container {
	p.title = widget.NewLabel("Contact Photos")
	// m := image.NewRGBA(image.Rect(0, 0, 640, 640))
	// for x := range 640 {
	// 	for y := range 640 {
	// 		m.Set(x, y, color.RGBA{255, 0, 0, 255})
	// 	}
	// }
	// placeholderImage := canvas.NewImageFromImage(m)

	// Create cards for each photo
	p.photoCards = make([]fyne.CanvasObject, len(p.photos))
	for i, photo := range p.photos {
		info, err := photos.GetInfo(p.ma.client, photo.Id, photo.Secret)
		if err != nil {
			fyne.LogError("Failed to get photo info", err)
			continue
		}
		// photoUrl := fmt.Sprintf("https://live.staticflickr.com/%s/%s_%s_%s.jpg", info.Photo.Server, info.Photo.Id, info.Photo.Secret, "z")
		// uri, err := storage.ParseURI(photoUrl)
		// if err != nil {
		// 	fyne.LogError("parsing url", err)
		// 	continue
		// }
		// fmt.Println("Downloading ", uri)
		// c := canvas.NewImageFromURI(uri)
		card := NewPhotoCard(photo.Title, photo.Username, info)
		// card := newTapCard(photo.Title, photo.Username, nil, func() {
		// 	pv := &photoView{ma: p.ma, photo: photo}
		// 	cont, err := pv.makeUI()
		// 	if err != nil {
		// 		fyne.LogError("parsing url", err)
		// 		return
		// 	}
		// 	p.ma.vs.Push(cont)
		// })
		// card.Content = c
		// c.FillMode = canvas.ImageFillContain

		p.photoCards[i] = card
	}

	gw := container.NewGridWrap(fyne.NewSize(500, 500), p.photoCards...)
	scrollingGrid := container.NewScroll(gw)

	p.container = container.NewStack(
		container.NewStack(),
		container.NewBorder(p.title, nil, nil, nil, scrollingGrid),
	)
	// p.container = container.NewBorder(p.title, nil, nil, nil, scrollingGrid)
	return p.container
}

// type tapCard struct {
// 	*widget.Card
// 	tap func()
// }

// func newTapCard(title, subtitle string, content fyne.CanvasObject, fn func()) *tapCard {
// 	i := &tapCard{tap: fn}
// 	i.Card = widget.NewCard(title, subtitle, content)
// 	i.ExtendBaseWidget(i)
// 	return i
// }

// func (t *tapCard) Tapped(e *fyne.PointEvent) {
// 	if t.tap == nil {
// 		return
// 	}
// 	t.tap()
// }
