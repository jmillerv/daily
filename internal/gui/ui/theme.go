package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var themeButton = widget.NewButtonWithIcon("theme", theme.ColorPaletteIcon(), changeTheme)

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
