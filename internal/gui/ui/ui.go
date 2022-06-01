package ui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/jmillerv/daily/internal/gui/panels"
	"log"
)

const (
	preferenceCurrentPanel = "currentPanel"
)

var themeButton = widget.NewButtonWithIcon("theme", theme.ColorPaletteIcon(), changeTheme)
var themeBool = binding.NewBool()
var topWindow fyne.Window

func Render() {
	a := app.NewWithID("com.jmillerv.daily")
	w := a.NewWindow("daily - a cross-platform stand up app")

	err := themeBool.Set(false)
	if err != nil {
		log.Fatal("error setting theme bool")
	}
	topWindow = w

	w.SetMaster()

	content := container.NewMax()
	title := widget.NewLabel("Component name")
	setPanel := func(p panels.Panel) {
		if fyne.CurrentDevice().IsMobile() {
			child := a.NewWindow(p.Title)
			topWindow = child
			child.SetContent(p.View(topWindow))
			child.Show()
			child.SetOnClosed(func() {
				topWindow = w
			})
			return
		}
		title.SetText(p.Title)

		content.Objects = []fyne.CanvasObject{p.View(w)}
		content.Refresh()
	}

	panel := container.NewBorder(container.NewVBox(title, widget.NewSeparator()), nil, nil, nil, content)
	if fyne.CurrentDevice().IsMobile() {
		w.SetContent(createNav(setPanel, false))
	} else {
		split := container.NewHSplit(createNav(setPanel, true), panel)
		split.Offset = 0.2
		w.SetContent(split)
	}

	w.Resize(fyne.Size{Width: 800, Height: 560})
	w.ShowAndRun()
}

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

	themes := container.New(layout.NewGridLayout(1),
		themeButton,
	)
	return container.NewBorder(nil, themes, nil, nil, tree)
}

func changeTheme() {
	a := fyne.CurrentApp()
	b, _ := themeBool.Get()
	if !b {
		a.Settings().SetTheme(theme.LightTheme())
		_ = themeBool.Set(true)
		return
	}
	if b {
		a.Settings().SetTheme(theme.DarkTheme())
		_ = themeBool.Set(false)
		return
	}
}
