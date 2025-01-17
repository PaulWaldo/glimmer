package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func (at *apptabs) makeUI() *container.AppTabs {
	contactsTab := container.NewTabItem("Contacts", widget.NewLabel("Contact photos content will go here"))
	groupsTab := container.NewTabItem("Groups", widget.NewLabel("Group photos content will go here"))

	if at.AppTabs == nil {
		at.AppTabs = container.NewAppTabs(contactsTab, groupsTab)
	}

	return at.AppTabs
}
