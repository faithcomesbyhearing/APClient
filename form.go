package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"gopkg.in/yaml.v3"

	//"github.com/faithcomesbyhearing/fcbh-dataset-io/decode_yaml"
	"github.com/faithcomesbyhearing/fcbh-dataset-io/decode_yaml/request"
	"os"
)

type RequestHelper struct {
	Train_MMSAdapter  bool
	Compare_NoCompare bool
}

func PresentForm(app fyne.App) {
	myWindow := app.NewWindow("Artificial Polyglot Processing Request")
	myWindow.Resize(fyne.NewSize(650, 400))

	var req request.Request
	var help RequestHelper
	var fieldList []Field
	datasetName := NewTextField("dataset_name:",
		"Enter a unique name for this dataset.", &req.DatasetName)
	fieldList = append(fieldList, datasetName)
	username := NewTextField("username:",
		"Enter a unique name for yourself.", &req.Username)
	fieldList = append(fieldList, username)
	languageISO := NewTextField("language_iso:",
		"Enter the ISO language code to be processed", &req.LanguageISO)
	fieldList = append(fieldList, languageISO)
	altLanguage := NewTextField("alt_language:",
		"This should be filled only if you wish to force the use of a specific language.",
		&req.AltLanguage)
	fieldList = append(fieldList, altLanguage)
	textData := NewTextField("text_data:",
		"s3://{bucket}/{path}*.usx", &req.TextData.AWSS3)
	fieldList = append(fieldList, textData)
	audioData := NewTextField("audio_data:",
		"s3://{bucket}/{path}*.mp3", &req.AudioData.AWSS3)
	fieldList = append(fieldList, audioData)

	timestamps := NewRadioGroup("timestamps:")
	timestamps.AddItem("mms_align", "most accurate, but uses lots of memory",
		&req.Timestamps.MMSAlign)
	timestamps.AddItem("mms_fa_verse", "slightly less accurate, but uses less memory",
		&req.Timestamps.MMSFAVerse)
	timestamps.AddItem("no_timestamps", "", &req.Timestamps.NoTimestamps)
	timestamps.SetSelected(0)
	fieldList = append(fieldList, timestamps)

	training := NewRadioGroup("training:")
	training.AddItem("mms_adapter", "training using MMS adapter module",
		&help.Train_MMSAdapter)
	training.AddItem("no_training", "", &req.Training.NoTraining)
	training.SetSelected(0)
	fieldList = append(fieldList, training)

	speechToText := NewRadioGroup("speech_to_text:")
	speechToText.AddItem("mms_asr", "speech to text using mms",
		&req.SpeechToText.MMS)
	speechToText.AddItem("adapter_asr", "speech to text using trained mms_adapter",
		&req.SpeechToText.MMSAdapter)
	speechToText.AddItem("no_speech_to_text", "", &req.SpeechToText.NoSpeechToText)
	speechToText.SetSelected(1)
	fieldList = append(fieldList, speechToText)

	compare := NewRadioGroup("compare:")
	compare.AddItem("compare", "html report of compare",
		&req.Compare.HTMLReport)
	compare.AddItem("no_compare", "", &help.Compare_NoCompare)
	compare.SetSelected(0)
	fieldList = append(fieldList, compare)

	gFilter := NewIntField("gordon_filter:",
		"Enter 4 for filter, leave blank for no filter.",
		&req.Compare.GordonFilter)
	fieldList = append(fieldList, gFilter)

	loadButton := widget.NewButton("Load", func() {
		fmt.Println("Loading...")
	})
	//saveButton := widget.NewButton("Save", func() {
	saveAction := func() {
		for _, field := range fieldList {
			field.Save()
		}
		req = templateA(req) // Add defaults from template
		if help.Train_MMSAdapter {
			req.Training.MMSAdapter.NumEpochs = 16
			req.Training.MMSAdapter.BatchMB = 4
			req.Training.MMSAdapter.LearningRate = 1e-03
			req.Training.MMSAdapter.WarmupPct = 12
			req.Training.MMSAdapter.GradNormMax = 0.4
		}
		yamlData, err := yaml.Marshal(&req)
		if err != nil {
			fmt.Printf("Error marshaling YAML: %v\n", err)
			return
		}
		filename := req.Username + "_" + req.DatasetName + ".yaml"
		err = os.WriteFile(filename, yamlData, 0644)
		if err != nil {
			fmt.Printf("Error writing file: %v\n", err)
			return
		}
		fmt.Println("Saved ", filename)
	}
	saveButton := widget.NewButton("Save", saveAction)
	clearButton := widget.NewButton("Clear", func() {
		for _, field := range fieldList {
			field.Clear()
		}
	})
	runButton := widget.NewButton("Run", func() {
		saveAction()
		// CODE FOR AWS UPLOAD GOES HERE
	})

	form := container.NewVBox(
		widget.NewSeparator(),
		container.NewGridWithColumns(4, loadButton, saveButton, clearButton, runButton),
		widget.NewSeparator(),
		datasetName.Container,
		username.Container,
		languageISO.Container,
		altLanguage.Container,
		textData.Container,
		audioData.Container,
		timestamps.Container,
		training.Container,
		speechToText.Container,
		compare.Container,
		gFilter.Container,
	)
	scrollableForm := container.NewScroll(form)

	myWindow.SetContent(scrollableForm)
	myWindow.ShowAndRun()
}
