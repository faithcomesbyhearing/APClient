package main

import (
	"context"
	"errors"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"github.com/faithcomesbyhearing/fcbh-dataset-io/decode_yaml"
	"github.com/faithcomesbyhearing/fcbh-dataset-io/decode_yaml/request"
	"gopkg.in/yaml.v3"
	"io"
	"os"
)

func PresentForm(app fyne.App) {
	myWindow := app.NewWindow("Artificial Polyglot Processing Request")
	myWindow.Resize(fyne.NewSize(650, 400))
	config := loadConfig(myWindow)

	var req request.Request
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
		true, &req.Timestamps.MMSAlign, nil)
	timestamps.AddItem("mms_fa_verse", "slightly less accurate, but uses less memory",
		true, &req.Timestamps.MMSFAVerse, nil)
	timestamps.AddItem("no_timestamps", "",
		true, &req.Timestamps.NoTimestamps, nil)
	timestamps.SetSelected(0)
	fieldList = append(fieldList, timestamps)

	training := NewRadioGroup("training:")
	training.AddItem("mms_adapter", "training using MMS adapter module",
		false, nil, &req.Training.MMSAdapter.NumEpochs)
	training.AddItem("no_training", "",
		true, &req.Training.NoTraining, nil)
	training.SetSelected(0)
	fieldList = append(fieldList, training)

	speechToText := NewRadioGroup("speech_to_text:")
	speechToText.AddItem("mms_asr", "speech to text using mms",
		true, &req.SpeechToText.MMS, nil)
	speechToText.AddItem("adapter_asr", "speech to text using trained mms_adapter",
		true, &req.SpeechToText.MMSAdapter, nil)
	speechToText.AddItem("no_speech_to_text", "",
		true, &req.SpeechToText.NoSpeechToText, nil)
	speechToText.SetSelected(1)
	fieldList = append(fieldList, speechToText)

	compare := NewCheckField("compare:", "html report of comparing two copies of text", true, &req.Compare.HTMLReport)
	fieldList = append(fieldList, compare)

	gFilter := NewIntField("gordon_filter:",
		"Enter 4 for filter, or zero for no filter.",
		&req.Compare.GordonFilter, myWindow)
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
				content, err := io.ReadAll(reader)
				if err != nil {
					dialog.ShowError(err, myWindow)
					return
				}
				req = request.Request{}
				err = yaml.Unmarshal(content, &req)
				if err != nil {
					dialog.ShowError(err, myWindow)
					return
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
		req = request.Request{}
		for _, field := range fieldList {
			field.Clear()
		}
	})
	createAction := func() []byte {
		for _, field := range fieldList {
			field.Save()
		}
		req.IsNew = config.IsNew
		req.NotifyOk = config.NotifyOk
		req.NotifyErr = config.NotifyErr
		if req.Training.MMSAdapter.NumEpochs != 0 {
			req.Training.MMSAdapter = config.MMSAdapter
		}
		if req.Compare.HTMLReport {
			req.Compare.CompareSettings = config.CompareSettings
		}
		yamlData, err := yaml.Marshal(&req)
		if err != nil {
			dialog.ShowError(err, myWindow)
			return []byte{}
		}
		decoder := decode_yaml.NewRequestDecoder(context.Background())
		_, status := decoder.Process(yamlData)
		if status != nil {
			dialog.ShowError(errors.New(status.Message), myWindow)
			return []byte{}
		}
		return yamlData
	}
	saveButton := widget.NewButton("Save", func() {
		yamlData := createAction()
		if len(yamlData) > 0 {
			_ = saveFile(req, yamlData, myWindow)
		}
	})
	runButton := widget.NewButton("Run", func() {
		yamlData := createAction()
		if len(yamlData) > 0 {
			filename := saveFile(req, yamlData, myWindow)
			if len(filename) > 0 {
				enqueueYaml(filename, yamlData, config.QueueBucket, myWindow)
			}
		}
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

func saveFile(req request.Request, yamlData []byte, window fyne.Window) string {
	filename := req.DatasetName + ".yaml"
	err := os.WriteFile(filename, yamlData, 0644)
	if err != nil {
		dialog.ShowError(err, window)
		return ""
	}
	return filename
}
