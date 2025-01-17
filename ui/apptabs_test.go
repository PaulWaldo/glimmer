package ui

import (
	"testing"

	"fyne.io/fyne/v2/container"
	"github.com/stretchr/testify/assert"
)

func TestAppTabs_MakeUI(t *testing.T) {
	ui := makeUI()

	// Type assert to AppTabs to access its methods and fields
	appTabs := ui.Objects[0].(*container.AppTabs)

	// Assert that the container has two children (the tabs)
	assert.Equal(t, 2, len(appTabs.Items))

	// Assert that the tab labels are correct
	assert.Equal(t, "Contacts", appTabs.Items[0].Text)
	assert.Equal(t, "Groups", appTabs.Items[1].Text)
}
