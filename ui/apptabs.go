package ui

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type apptabs struct {
	ma      *myApp
	appTabs *container.AppTabs
}

func (at *apptabs) makeUI() {
	contactsTab := container.NewTabItem("Contacts", widget.NewLabel("Contact photos content will go here"))
	groupsTab := container.NewTabItem("Groups", widget.NewLabel("Group photos content will go here"))

	at.appTabs = container.NewAppTabs(contactsTab, groupsTab)
}
