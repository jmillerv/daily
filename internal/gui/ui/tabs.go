package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"github.com/jmillerv/daily/internal/gui/panels"
)

// Tabbed Navigation
func createTabs() fyne.CanvasObject {
	// HOME TAB
	homeTabPane := panels.Panels[homePanel]

	// ABOUT TAB
	aboutTabPane := panels.Panels[aboutPanel]

	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("StandUp", theme.HomeIcon(), homeTabPane.View(nil)),
		container.NewTabItemWithIcon("About", theme.InfoIcon(), aboutTabPane.View(nil)),
	)
	return container.NewVBox(tabs, themeButton)
}
