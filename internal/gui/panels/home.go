package panels

import (
	"log"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

const (
	whatYouDid    = "What did you do yesterday?"
	wydKey        = "_WYD"
	whatYouWill   = "What will you do today?"
	wywKey        = "_WYW"
	whatBlocksYou = "Is anything blocking your progress?"
	wbyKey        = "_WBY"
)

var whatYouDidKey string
var whatYouWillKey string
var whatBlocksYouKey string

func homeScreen(_ fyne.Window) fyne.CanvasObject {
	app := fyne.CurrentApp()

	// Question elements
	whatYouDidQuestion := widget.NewMultiLineEntry()
	whatYouDidQuestion.SetPlaceHolder(whatYouDid)

	whatYouWillQuestion := widget.NewMultiLineEntry()
	whatYouWillQuestion.SetPlaceHolder(whatYouWill)

	whatBlocksYouQuestion := widget.NewMultiLineEntry()
	whatBlocksYouQuestion.SetPlaceHolder(whatBlocksYou)

	// labels
	whatYouDidLabel := canvas.NewText(whatYouDidQuestion.Text, theme.FocusColor())
	whatYouDidLabel.Alignment = fyne.TextAlignCenter
	whatYouWillLabel := canvas.NewText(whatYouWillQuestion.Text, theme.FocusColor())
	whatYouWillLabel.Alignment = fyne.TextAlignCenter
	whatBlocksYouLabel := canvas.NewText(whatBlocksYouQuestion.Text, theme.FocusColor())
	whatBlocksYouLabel.Alignment = fyne.TextAlignCenter

	// Bindings & Data
	currentDate := time.Now().Format("01-02-2006")
	today := widget.NewLabel(currentDate)
	today.Alignment = fyne.TextAlignCenter

	whatYouDidKey = currentDate + wydKey
	whatYouDidAnswer := binding.BindPreferenceString(whatYouDidKey, app.Preferences())
	whatYouDidAnswer.AddListener(
		binding.NewDataListener(func() {
			whatYouDidLabel.Text, _ = whatYouDidAnswer.Get()
			whatYouDidLabel.Refresh()
		}))

	whatYouWillKey = currentDate + wywKey
	whatYouWillAnswer := binding.BindPreferenceString(whatYouWillKey, app.Preferences())
	whatYouWillAnswer.AddListener(
		binding.NewDataListener(func() {
			whatYouWillLabel.Text, _ = whatYouWillAnswer.Get()
			whatYouWillLabel.Refresh()
		}))

	whatBlocksYouKey = currentDate + wbyKey
	whatBlocksYouAnswer := binding.BindPreferenceString(whatBlocksYouKey, app.Preferences())
	whatBlocksYouAnswer.AddListener(
		binding.NewDataListener(func() {
			whatBlocksYouLabel.Text, _ = whatBlocksYouAnswer.Get()
			whatBlocksYouLabel.Refresh()
		}))

	// Button Elements
	saveButton := widget.NewButtonWithIcon("Save", theme.DocumentSaveIcon(), func() {
		_ = whatYouDidAnswer.Set(strings.Replace(whatYouDidQuestion.Text, "\n", " ", -1))
		_ = whatYouWillAnswer.Set(strings.Replace(whatYouWillQuestion.Text, "\n", " ", -1))
		_ = whatBlocksYouAnswer.Set(strings.Replace(whatBlocksYouQuestion.Text, "\n", " ", -1))
		log.Println("saving stand up")
	})
	clearButton := widget.NewButtonWithIcon("Clear", theme.ContentClearIcon(), func() {
		clear()
		// TODO Find a better way to implement this
		_ = whatYouDidAnswer.Set("")
		_ = whatYouWillAnswer.Set("")
		_ = whatBlocksYouAnswer.Set("")
		log.Println("clearing stand up")
	})
	buttons := container.NewGridWithColumns(2, clearButton, saveButton)

	// Layout

	questionRows := container.NewGridWithRows(3, whatYouDidQuestion, whatYouWillQuestion, whatBlocksYouQuestion)
	questions := container.NewGridWithColumns(1, questionRows)
	answers := container.NewCenter(container.NewVBox(
		widget.NewLabel(whatYouDid),
		whatYouDidLabel,
		widget.NewLabel(whatYouWill),
		whatYouWillLabel,
		widget.NewLabel(whatBlocksYou),
		whatBlocksYouLabel,
	))

	return container.NewVBox(
		questions,
		buttons,
		today,
		answers,
	)
}

func clear() {
	app := fyne.CurrentApp()
	app.Preferences().RemoveValue(whatYouDidKey)
	app.Preferences().RemoveValue(whatYouWillKey)
	app.Preferences().RemoveValue(whatBlocksYouKey)
}
