package panels

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/jmillerv/daily/internal/gui/ui"
)

var ThemeButton = widget.NewButtonWithIcon("theme", theme.ColorPaletteIcon(), changeTheme)

func changeTheme() {
	a := fyne.CurrentApp()
	b, _ := ui.themeBool.Get()
	if !b {
		a.Settings().SetTheme(theme.LightTheme())
		_ = ui.themeBool.Set(true)
		return
	}
	if b {
		a.Settings().SetTheme(theme.DarkTheme())
		_ = ui.themeBool.Set(false)
		return
	}
}
