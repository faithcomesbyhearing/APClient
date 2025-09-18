package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"gopkg.in/yaml.v3"
	"io"

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

	compare := NewCheckField("compare:", "html report of comparing two copies of text", true, &req.Compare.HTMLReport)
	fieldList = append(fieldList, compare)

	gFilter := NewIntField("gordon_filter:",
		"Enter 4 for filter, leave blank for no filter.",
		&req.Compare.GordonFilter)
	fieldList = append(fieldList, gFilter)

	loadButton := widget.NewButton("Load", func() {
		fileDialog := dialog.NewFileOpen(
			func(reader fyne.URIReadCloser, err error) {
				if err != nil {
					dialog.ShowError(err, myWindow)
					return
				}
				if reader == nil {
					return // User cancelled
				}
				defer reader.Close()

				filePath := reader.URI().Path()
				fmt.Printf("Loading: %s\n", filePath)

				// Example: Read as string
				content, err := io.ReadAll(reader)
				if err != nil {
					dialog.ShowError(err, myWindow)
					return
				}
				err = yaml.Unmarshal(content, &req)
				if err != nil {
					//return nil, fmt.Errorf("failed to parse YAML: %w", err)
				}
				for _, field := range fieldList {
					field.Load()
				}
			},
			myWindow,
		)
		fileDialog.SetFilter(storage.NewExtensionFileFilter([]string{".yaml", ".yml"}))
		listableURI, err := storage.ListerForURI(storage.NewFileURI("./"))
		if err == nil {
			fileDialog.SetLocation(listableURI)
		}
		fileDialog.Show()
	})
	clearButton := widget.NewButton("Clear", func() {
		for _, field := range fieldList {
			field.Clear()
		}
	})
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
		if req.Compare.HTMLReport {
			req.Compare.CompareSettings.LowerCase = true
			req.Compare.CompareSettings.RemovePromptChars = true
			req.Compare.CompareSettings.RemovePunctuation = true
			req.Compare.CompareSettings.DoubleQuotes.Remove = true
			req.Compare.CompareSettings.Apostrophe.Remove = true
			req.Compare.CompareSettings.Hyphen.Remove = true
			req.Compare.CompareSettings.DiacriticalMarks.NormalizeNFC = true
		}
		yamlData, err := yaml.Marshal(&req)
		if err != nil {
			fmt.Printf("Error marshaling YAML: %v\n", err)
			return
		}
		filename := req.DatasetName + ".yaml"
		err = os.WriteFile(filename, yamlData, 0644)
		if err != nil {
			fmt.Printf("Error writing file: %v\n", err)
			return
		}
		fmt.Println("Saved ", filename)
	}
	saveButton := widget.NewButton("Save", saveAction)
	runButton := widget.NewButton("Run", func() {
		saveAction()
		// CODE FOR AWS UPLOAD GOES HERE
	})
	quitButton := widget.NewButton("Quit", func() {
		app.Quit()
	})

	form := container.NewVBox(
		widget.NewSeparator(),
		container.NewGridWithColumns(5, loadButton, saveButton, clearButton, runButton, quitButton),
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
