package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func makeUI() *fyne.Container {
	contactsTab := container.NewTabItem("Contacts", widget.NewLabel("Contact photos content will go here")) // Placeholder content
	groupsTab := container.NewTabItem("Groups", widget.NewLabel("Group photos content will go here")) // Placeholder content

	appTabs := container.NewAppTabs(contactsTab, groupsTab)
	return &fyne.Container{Objects: []fyne.CanvasObject{appTabs}}
}
