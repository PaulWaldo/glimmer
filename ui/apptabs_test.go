package ui

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppTabs_MakeUI(t *testing.T) {
	ma := &myApp{} // You'll need to initialize myApp appropriately for the test
	at := &apptabs{ma: ma}
	ui := at.makeUI()

	// This assertion will fail because makeUI doesn't use at.appTabs yet
	assert.NotNil(t, at.appTabs, "at.appTabs should not be nil after makeUI") // Access the named field

	appTabs := ui.Objects[0].(*container.AppTabs)

	assert.Equal(t, 2, len(appTabs.Items))

	assert.Equal(t, "Contacts", appTabs.Items[0].Text)
	assert.Equal(t, "Groups", appTabs.Items[1].Text)
}
