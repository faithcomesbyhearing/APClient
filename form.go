package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"strings"
)

func PresentForm(app fyne.App) {
	var result FormData
	myWindow := app.NewWindow("Artificial Polyglot Processing Request")
	myWindow.Resize(fyne.NewSize(650, 400))

	datasetRow, datasetName := createTextField("dataset_name:",
		"Enter a unique name for this dataset.")
	usernameRow, username := createTextField("username:",
		"Enter a unique name for yourself.")
	languageISORow, languageISO := createTextField("language_iso:",
		"Enter the ISO language code to be processed")
	altLanguageRow, altLanguage := createTextField("alt_language:",
		"This should be filled only if you wish to force the use of a specific language.")
	textDataRow, textData := createTextField("text_data:",
		"s3://{bucket}/{path}*.usx")
	audioDataRow, audioData := createTextField("audio_data:",
		"s3://{bucket}/{path}*.mp3")

	timestampsDesc := []string{"mms_align - most accurate, but uses lots of memory",
		"mms_fa_verse - slightly less accurate, but uses less memory",
		"no_timestamps"}
	timestampRow, timestamps := createRadioField("timestamps:", timestampsDesc, 0)

	trainingDesc := []string{"mms_adapter - training using MMS adapter module",
		"no_training"}
	trainingRow, training := createRadioField("training:", trainingDesc, 0)

	speechToTextDesc := []string{"mms_asr - speech to text using mms",
		"adapter_asr - speech to text using trained mms_adapter",
		"no_speech_to_text"}
	speechToTextRow, speechToText := createRadioField("speech_to_text:", speechToTextDesc, 1)

	compareDesc := []string{"compare - html report of compare", "no_compare"}
	compareRow, compare := createRadioField("compare:", compareDesc, 0)

	gFilterRow, gFilter := createTextField("gordon_filter:",
		"Enter 4 for filter, enter 0 for no filter.")

	loadButton := widget.NewButton("Load", func() {
		fmt.Println("Loading...")
	})
	saveButton := widget.NewButton("Save", func() {
		fmt.Println("Saving...")
	})
	clearButton := widget.NewButton("Clear", func() {
		fmt.Println("Clearing...")
	})
	runButton := widget.NewButton("Run", func() {
		result.DatasetName = datasetName.Text
		result.Username = username.Text
		result.LanguageISO = languageISO.Text
		result.AltLanguage = altLanguage.Text
		result.TextData = textData.Text
		result.AudioData = audioData.Text
		result.Timestamps = strings.Split(timestamps.Selected, " ")[0]
		result.Training = strings.Split(training.Selected, " ")[0]
		result.SpeechToText = strings.Split(speechToText.Selected, " ")[0]
		result.Compare = strings.Split(compare.Selected, " ")[0]
		result.GordonFilter = gFilter.Text
	})

	form := container.NewVBox(
		widget.NewSeparator(),
		container.NewGridWithColumns(4, loadButton, saveButton, clearButton, runButton),
		widget.NewSeparator(),
		datasetRow,
		usernameRow,
		languageISORow,
		altLanguageRow,
		textDataRow,
		audioDataRow,
		timestampRow,
		trainingRow,
		speechToTextRow,
		compareRow,
		gFilterRow,
	)
	scrollableForm := container.NewScroll(form)

	myWindow.SetContent(scrollableForm)
	myWindow.ShowAndRun()
}

func createTextField(label string, description string) (*fyne.Container, *widget.Entry) {
	entry := widget.NewEntry()
	entry.SetPlaceHolder(description)
	row := container.NewBorder(nil, nil,
		container.NewWithoutLayout(
			widget.NewLabel(label),
		), nil,
		entry,
	)
	return row, entry
}

func createRadioField(label string, descriptions []string, selected int) (*fyne.Container, *widget.RadioGroup) {
	group := widget.NewLabel(label)
	timestamps := widget.NewRadioGroup(descriptions, nil)
	timestamps.SetSelected(descriptions[selected])
	row := container.NewVBox(group, timestamps)
	return row, timestamps
}

/*
func persistData(datasetName *widget.Entry, username *widget.Entry,
	languageISO *widget.Entry, altLanguage *widget.Entry, textData *widget.Entry,
	audioData *widget.Entry, timestamps *widget.RadioGroup, training *widget.RadioGroup,
	speechToText *widget.RadioGroup, compare *widget.RadioGroup,
	gFilter *widget.Entry) FormData {
	var result FormData
	result.DatasetName = datasetName.Text
	result.Username = username.Text
	result.LanguageISO = languageISO.Text
	result.AltLanguage = altLanguage.Text
	result.TextData = textData.Text
	result.AudioData = audioData.Text
	result.Timestamps = strings.Split(timestamps.Selected, " ")[0]
	result.Training = strings.Split(training.Selected, " ")[0]
	result.SpeechToText = strings.Split(speechToText.Selected, " ")[0]
	result.Compare = strings.Split(compare.Selected, " ")[0]
	result.GordonFilter = gFilter.Text
	return result
}
*
*/
