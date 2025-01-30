package ui

import (
	"testing"

	"fyne.io/fyne/v2/app"
)

func TestTwoTabsExist(t *testing.T) {
	ma := &myApp{} // Initialize myApp as needed for your tests
	ma.app = app.NewWithID(AppID)
	ma.prefs = NewPreferences(ma.app)
	ma.window = ma.app.NewWindow("Glimmer")

	// Initialize apptabs and assign it to ma.at
	ma.at = &apptabs{ma: ma}
	ma.at.makeUI()

	tabItems := ma.at.appTabs.Items
	expectedTitles := []string{"Contacts", "Groups"}

	if len(tabItems) != len(expectedTitles) {
		t.Fatalf("Expected %d tabs, but got %d", len(expectedTitles), len(tabItems))
	}

	for i, title := range expectedTitles {
		if tabItems[i].Text != title {
			t.Errorf("Expected tab %d to be '%s', but got '%s'", i, title, tabItems[i].Text)
		}
	}
}
