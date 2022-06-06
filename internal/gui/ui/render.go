package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"github.com/jmillerv/daily/helpers"
	"github.com/jmillerv/daily/internal/gui/panels"
	"log"
)

const (
	preferenceCurrentPanel = "currentPanel"
	homePanel              = "home"
	aboutPanel             = "about"
	configurationPanel     = "configuration"
)

var issueLink = widget.NewHyperlink("issues", helpers.ParseURL("https://github.com/jmillerv/daily/issues"))
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

	//panel := container.NewBorder(container.NewVBox(title, widget.NewSeparator()), nil, nil, nil, content)
	if fyne.CurrentDevice().IsMobile() {
		w.SetContent(createNav(setPanel, false))
	} else {
		desktopTabs := container.NewHScroll(createTabs())
		w.SetContent(desktopTabs)
	}

	w.Resize(fyne.Size{Width: 800, Height: 560})
	w.ShowAndRun()
}
