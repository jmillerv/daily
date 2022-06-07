package panels

import (
	"log"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

const (
	whatYouDid        = "What did you do yesterday?"
	wydKey            = "_WYD"
	whatYouWill       = "What will you do Today?"
	wywKey            = "_WYW"
	whatBlocksYou     = "Is anything blocking your progress?"
	wbyKey            = "_WBY"
	defaultDateFormat = "01-02-2006"
)

// keys
var whatYouDidKey string
var whatYouWillKey string
var whatBlocksYouKey string

// question entries
var whatYouDidQuestion *widget.Entry
var whatYouWillQuestion *widget.Entry
var whatBlocksYouQuestion *widget.Entry

var questions *fyne.Container

// answer bindings
var whatYouDidAnswer binding.String
var whatYouWillAnswer binding.String
var whatBlocksYouAnswer binding.String

// labels
var whatYouDidLabel *canvas.Text
var whatYouWillLabel *canvas.Text
var whatBlocksYouLabel *canvas.Text

// date
var CurrentDate string
var Today *widget.Label

func homeScreen(_ fyne.Window) fyne.CanvasObject {
	app := fyne.CurrentApp()

	// keys
	setKeys()

	// Question elements
	setQuestions()

	// labels
	setLabels()

	// Bindings & Data
	// Set Date
	CurrentDate = time.Now().Format("01-02-2006")
	Today = widget.NewLabel(CurrentDate)
	Today.Alignment = fyne.TextAlignCenter

	// answers
	setAnswers(app)

	// Button Elements
	saveButton := widget.NewButtonWithIcon("Save", theme.DocumentSaveIcon(), func() {
		_ = whatYouDidAnswer.Set(whatYouDidQuestion.Text)
		_ = whatYouWillAnswer.Set(whatYouWillQuestion.Text)
		_ = whatBlocksYouAnswer.Set(whatBlocksYouQuestion.Text)
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
	questions = container.NewGridWithColumns(1, questionRows)
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
		Today,
		answers,
		buttons,
	)
}

// setKeys creates the keys for fyne's key/value store
func setKeys() {
	whatYouDidKey = CurrentDate + wydKey
	whatYouWillKey = CurrentDate + wywKey
	whatBlocksYouKey = CurrentDate + wbyKey
}

// setLabels
func setLabels() {
	whatYouDidLabel = canvas.NewText(whatYouDidQuestion.Text, theme.FocusColor())
	whatYouDidLabel.Alignment = fyne.TextAlignCenter
	whatYouWillLabel = canvas.NewText(whatYouWillQuestion.Text, theme.FocusColor())
	whatYouWillLabel.Alignment = fyne.TextAlignCenter
	whatBlocksYouLabel = canvas.NewText(whatBlocksYouQuestion.Text, theme.FocusColor())
	whatBlocksYouLabel.Alignment = fyne.TextAlignCenter

}

// setAnswers()
func setAnswers(app fyne.App) {
	whatYouDidAnswer = binding.BindPreferenceString(whatYouDidKey, app.Preferences())
	whatYouDidAnswer.AddListener(
		binding.NewDataListener(func() {
			whatYouDidLabel.Text, _ = whatYouDidAnswer.Get()
			whatYouDidLabel.Refresh()
		}))

	whatYouWillAnswer = binding.BindPreferenceString(whatYouWillKey, app.Preferences())
	whatYouWillAnswer.AddListener(
		binding.NewDataListener(func() {
			whatYouWillLabel.Text, _ = whatYouWillAnswer.Get()
			whatYouWillLabel.Refresh()
		}))

	whatBlocksYouAnswer = binding.BindPreferenceString(whatBlocksYouKey, app.Preferences())
	whatBlocksYouAnswer.AddListener(
		binding.NewDataListener(func() {
			whatBlocksYouLabel.Text, _ = whatBlocksYouAnswer.Get()
			whatBlocksYouLabel.Refresh()
		}))

}

// setQuestions puts the question into a base state with placeholder suggestions
func setQuestions() {
	whatYouDidQuestion = widget.NewMultiLineEntry()
	whatYouDidQuestion.SetPlaceHolder(whatYouDid)

	whatYouWillQuestion = widget.NewMultiLineEntry()
	whatYouWillQuestion.SetPlaceHolder(whatYouWill)

	whatBlocksYouQuestion = widget.NewMultiLineEntry()
	whatBlocksYouQuestion.SetPlaceHolder(whatBlocksYou)

}

// clear removes the values for the keys associated with the day's standup
func clear() {
	whatYouDidQuestion.Text = ""
	whatYouWillQuestion.Text = ""
	whatBlocksYouQuestion.Text = ""
	app := fyne.CurrentApp()
	app.Preferences().RemoveValue(whatYouDidKey)
	app.Preferences().RemoveValue(whatYouWillKey)
	app.Preferences().RemoveValue(whatBlocksYouKey)
	questions.Refresh()
}

func UpdateDate(setDate string) {
	now := time.Now().Format("01-02-2006")
	if setDate != now {
		CurrentDate = now
	}
}
