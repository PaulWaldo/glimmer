package ui

import (
	"testing"

	"fyne.io/fyne/v2/container"
	"github.com/stretchr/testify/assert"
)

func TestAppTabs_MakeUI(t *testing.T) {
	ma := &myApp{}
	at := &apptabs{ma: ma}
	at.makeUI() // Call makeUI to initialize at.appTabs

	assert.NotNil(t, at.appTabs, "at.appTabs should not be nil after makeUI")

	assert.Equal(t, 2, len(at.appTabs.Items))

	assert.Equal(t, "Contacts", at.appTabs.Items[0].Text)
	assert.Equal(t, "Groups", at.appTabs.Items[1].Text)
}
