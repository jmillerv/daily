package ui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/jmillerv/daily/internal/gui/panels"
)

func createNav(setPanel func(panel panels.Panel), loadPrevious bool) fyne.CanvasObject {
	a := fyne.CurrentApp()
	tree := &widget.Tree{
		ChildUIDs: func(uid string) []string {
			return panels.PanelIndex[uid]
		},
		CreateNode: func(branch bool) fyne.CanvasObject {
			return widget.NewLabel("Collection Widgets")
		},
		IsBranch: func(uid string) bool {
			children, ok := panels.PanelIndex[uid]
			return ok && len(children) > 0
		},
		UpdateNode: func(uid string, branch bool, obj fyne.CanvasObject) {
			p, ok := panels.Panels[uid]
			if !ok {
				fyne.LogError(fmt.Sprintf("Missing panel %s", uid), nil)
			}
			obj.(*widget.Label).SetText(p.Title)
		},
		OnSelected: func(uid string) {
			if p, ok := panels.Panels[uid]; ok {
				a.Preferences().SetString(preferenceCurrentPanel, uid)
				setPanel(p)
			}
		},
	}

	if loadPrevious {
		currentPref := a.Preferences().StringWithFallback(preferenceCurrentPanel, "home")
		tree.Select(currentPref)
	}
	issueCenter := container.NewCenter(issueLink)
	themes := container.New(layout.NewGridLayout(1),
		issueCenter,
		themeButton,
	)
	return container.NewBorder(nil, themes, nil, nil, tree)
}
