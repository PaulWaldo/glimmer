package ui

import (
	"testing"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"github.com/stretchr/testify/assert"
)

func TestAppTabs_MakeUI(t *testing.T) {
	ui := makeUI()

	// Assert that the returned UI element is an *fyne.Container (AppTabs)
	assert.IsType(t, &fyne.Container{}, ui)

	// Type assert to AppTabs to access its methods
	appTabs := ui.(*fyne.Container).Objects[0].(*container.AppTabs)

    // Assert that the container has two children (the tabs)
	assert.Equal(t, 2, len(appTabs.Items))


	// Assert that the tab labels are correct
	assert.Equal(t, "Contacts", appTabs.Items[0].Text)
	assert.Equal(t, "Groups", appTabs.Items[1].Text)
}
