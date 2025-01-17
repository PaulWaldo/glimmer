package ui

import (
	"testing"

	"fyne.io/fyne/v2/container"
	"github.com/stretchr/testify/assert"
)

func TestAppTabs_MakeUI(t *testing.T) {
	ma := &myApp{}
	at := &apptabs{ma: ma}
	ui := at.makeUI()

	// This assertion will fail because makeUI doesn't use at.appTabs yet
	ui := at.makeUI()

	assert.NotNil(t, at.appTabs, "at.appTabs should not be nil after makeUI")

	assert.Equal(t, 2, len(ui.Items))

	assert.Equal(t, "Contacts", ui.Items[0].Text)
	assert.Equal(t, "Groups", ui.Items[1].Text)
}
