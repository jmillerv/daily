package panels

import (
	"fyne.io/fyne/v2/driver/desktop"
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

type questionEntry struct {
	entry *widget.Entry
}

func (q *questionEntry) TypedShortcut(s fyne.Shortcut) {
	if _, ok := s.(*desktop.CustomShortcut); !ok {
		log.Println("shortcut test")
		q.entry.TypedShortcut(s)
		return
	}
	log.Println("tab shortcut")
}

// keys
var whatYouDidKey string
var whatYouWillKey string
var whatBlocksYouKey string

// question entries
var whatYouDidQuestion *questionEntry
var whatYouWillQuestion *questionEntry
var whatBlocksYouQuestion *questionEntry

var questions *fyne.Container

// answer bindings
var whatYouDidAnswer binding.String
var whatYouWillAnswer binding.String
var whatBlocksYouAnswer binding.String

// labels
var whatYouDidLabel *canvas.Text
var whatYouWillLabel *canvas.Text
var whatBlocksYouLabel *canvas.Text

var ctrlTab desktop.CustomShortcut

// date
var CurrentDate string
var Today *widget.Label

func homeScreen(_ fyne.Window) fyne.CanvasObject {
	app := fyne.CurrentApp()
	ctrlTab = desktop.CustomShortcut{KeyName: fyne.KeyTab, Modifier: desktop.ControlModifier}
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
	saveButton := createSaveButton()
	clearButton := createClearButton()
	buttons := container.NewGridWithColumns(2, clearButton, saveButton)

	// Layout
	questionRows := container.NewGridWithRows(3, whatYouDidQuestion.entry, whatYouWillQuestion.entry, whatBlocksYouQuestion.entry)
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

func newQuestionEntry() *questionEntry {
	newEntry := widget.NewMultiLineEntry()
	newEntry.TypedShortcut(&ctrlTab)
	return &questionEntry{entry: newEntry}
}

// createSaveButton returns a button that sets the answers.
func createSaveButton() *widget.Button {
	return widget.NewButtonWithIcon("Save", theme.DocumentSaveIcon(), func() {
		_ = whatYouDidAnswer.Set(whatYouDidQuestion.entry.Text)
		_ = whatYouWillAnswer.Set(whatYouWillQuestion.entry.Text)
		_ = whatBlocksYouAnswer.Set(whatBlocksYouQuestion.entry.Text)
		log.Println("saving stand up")
	})
}

// createClearButton returns the button that runs the clear() function and resets the answers.
func createClearButton() *widget.Button {
	return widget.NewButtonWithIcon("Clear", theme.ContentClearIcon(), func() {
		clear()
		// TODO Find a better way to implement this
		_ = whatYouDidAnswer.Set("")
		_ = whatYouWillAnswer.Set("")
		_ = whatBlocksYouAnswer.Set("")
		log.Println("clearing stand up")
	})
}

// setKeys creates the keys for fyne's key/value store
func setKeys() {
	whatYouDidKey = CurrentDate + wydKey
	whatYouWillKey = CurrentDate + wywKey
	whatBlocksYouKey = CurrentDate + wbyKey
}

// setLabels
func setLabels() {
	whatYouDidLabel = canvas.NewText(whatYouDidQuestion.entry.Text, theme.FocusColor())
	whatYouDidLabel.Alignment = fyne.TextAlignCenter
	whatYouWillLabel = canvas.NewText(whatYouWillQuestion.entry.Text, theme.FocusColor())
	whatYouWillLabel.Alignment = fyne.TextAlignCenter
	whatBlocksYouLabel = canvas.NewText(whatBlocksYouQuestion.entry.Text, theme.FocusColor())
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
	whatYouDidQuestion = newQuestionEntry()
	whatYouDidQuestion.entry.SetPlaceHolder(whatYouDid)

	whatYouWillQuestion = newQuestionEntry()
	whatYouWillQuestion.entry.SetPlaceHolder(whatYouWill)

	whatBlocksYouQuestion = newQuestionEntry()
	whatBlocksYouQuestion.entry.SetPlaceHolder(whatBlocksYou)

}

// clear removes the values for the keys associated with the day's standup
func clear() {
	whatYouDidQuestion.entry.Text = ""
	whatYouWillQuestion.entry.Text = ""
	whatBlocksYouQuestion.entry.Text = ""
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
