package panels

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/jmillerv/daily/helpers"
)

func aboutScreen(_ fyne.Window) fyne.CanvasObject {
	return container.NewCenter(
		container.NewVBox(
			container.NewHBox(
				widget.NewHyperlink("readme", helpers.ParseURL("https://github.com/jmillerv/daily#readme")),
				widget.NewLabel("-"),
				widget.NewHyperlink("report issue", helpers.ParseURL("https://github.com/jmillerv/daily/issues")))))
}
