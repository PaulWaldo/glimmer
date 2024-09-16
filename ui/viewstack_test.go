package ui

import (
	"fmt"
	"testing"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/test"
	"fyne.io/fyne/v2/widget"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPush(t *testing.T) {
	app := test.NewApp()
	win := app.NewWindow("Test")
	vs := NewViewStack(win)

	const numViews = 3
	expectedViews := make([]*fyne.Container, numViews)
	for i := range numViews {
		l := widget.NewLabel(fmt.Sprintf("View %d", i))
		expectedViews[i] = container.NewStack(l)
		vs.Push(expectedViews[i])
	}

	// Check that the views are in the correct order
	for i, v := range vs.elements {
		if v != expectedViews[i] {
			assert.Equal(t, expectedViews[i], v)
		}
	}

	// Check that the current view is the last one pushed
	assert.Equal(t, expectedViews[numViews-1], win.Content())
}
func TestPop(t *testing.T) {
	app := test.NewApp()
	win := app.NewWindow("Test")
	vs := NewViewStack(win)

	const numViews = 3
	expectedViews := make([]*fyne.Container, numViews)
	for i := range numViews {
		l := widget.NewLabel(fmt.Sprintf("View %d", i))
		expectedViews[i] = container.NewStack(l)
		vs.Push(expectedViews[i])
	}

	// Check that views are in correct order by Popping them
	for i := numViews - 1; i >= 0; i-- {
		view, ok := win.Content().(*fyne.Container)
		require.True(t, ok)
		label := view.Objects[0].(*widget.Label)
		require.Equal(t, expectedViews[i].Objects[0].(*widget.Label).Text, label.Text)
		vs.Pop()
	}
}

func TestFirst(t *testing.T) {
	app := test.NewApp()
	win := app.NewWindow("Test")
	vs := NewViewStack(win)

	const numViews = 3
	expectedViews := make([]*fyne.Container, numViews)
	for i := range numViews {
		l := widget.NewLabel(fmt.Sprintf("View %d", i))
		expectedViews[i] = container.NewStack(l)
		vs.Push(expectedViews[i])
	}

	vs.First()
	view, ok := win.Content().(*fyne.Container)
	require.True(t, ok)
	label := view.Objects[0].(*widget.Label)
	require.Equal(t, expectedViews[0].Objects[0].(*widget.Label).Text, label.Text)
}

func TestViewStack_IsEmpty(t *testing.T) {
	app := test.NewApp()
	win := app.NewWindow("Test")
	vs := NewViewStack(win)
	assert.True(t, vs.IsEmpty())
	// Push a view
	vs.Push(container.NewStack())
	assert.False(t, vs.IsEmpty())
	// Pop the view
	vs.Pop()
	// Check that it is empty again
	assert.True(t, vs.IsEmpty())
}
