package ui

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type apptabs struct {
	ma      *myApp
	appTabs *container.AppTabs // Use a named field
}

func (at *apptabs) makeUI() *container.AppTabs {
	contactsTab := container.NewTabItem("Contacts", widget.NewLabel("Contact photos content will go here"))
	groupsTab := container.NewTabItem("Groups", widget.NewLabel("Group photos content will go here"))

	if at.appTabs == nil { // Access the named field
		at.appTabs = container.NewAppTabs(contactsTab, groupsTab)
	}

	return at.appTabs // Return the named field
}
