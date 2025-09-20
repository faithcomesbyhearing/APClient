package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"github.com/faithcomesbyhearing/fcbh-dataset-io/decode_yaml/request"
	"gopkg.in/yaml.v3"
	"os"
)

type ClientConfig struct {
	QueueBucket     string                  `yaml:"queue_bucket"`
	IsNew           bool                    `yaml:"is_new"`
	NotifyOk        []string                `yaml:"notify_ok"`
	NotifyErr       []string                `yaml:"notify_err"`
	MMSAdapter      request.MMSAdapter      `yaml:"mms_adapter"`
	CompareSettings request.CompareSettings `yaml:"compare_settings"`
}

func loadConfig(window fyne.Window) ClientConfig {
	var config ClientConfig
	filename := "APClient.yaml"
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		config = createConfig(filename, window)
	} else {
		config = readConfig(filename, window)
	}
	return config
}

func createConfig(filename string, window fyne.Window) ClientConfig {
	var config ClientConfig
	config.QueueBucket = "dataset-queue"
	config.IsNew = true
	config.NotifyOk = []string{"jbarndt@fcbhmail.org",
		"ezornes@fcbhmail.org",
		"gfiddes@fcbhmail.org",
		"edomschot@fcbhmail.org"}
	config.NotifyErr = []string{"jbarndt@fcbhmail.org",
		"gary@shortsands.com",
		"ezornes@fcbhmail.org"}
	config.MMSAdapter.NumEpochs = 16
	config.MMSAdapter.BatchMB = 4
	config.MMSAdapter.LearningRate = 1e-03
	config.MMSAdapter.WarmupPct = 12
	config.MMSAdapter.GradNormMax = 0.4
	config.CompareSettings.LowerCase = true
	config.CompareSettings.RemovePromptChars = true
	config.CompareSettings.RemovePunctuation = true
	config.CompareSettings.DoubleQuotes.Remove = true
	config.CompareSettings.Apostrophe.Remove = true
	config.CompareSettings.Hyphen.Remove = true
	config.CompareSettings.DiacriticalMarks.NormalizeNFC = true
	yamlData, err := yaml.Marshal(&config)
	if err != nil {
		dialog.ShowError(err, window)
		return config
	}
	err = os.WriteFile(filename, yamlData, 0644)
	if err != nil {
		dialog.ShowError(err, window)
		return config
	}
	return config
}

func readConfig(filename string, window fyne.Window) ClientConfig {
	var config ClientConfig
	content, err := os.ReadFile(filename)
	if err != nil {
		dialog.ShowError(err, window)
		return config
	}
	err = yaml.Unmarshal(content, &config)
	if err != nil {
		dialog.ShowError(err, window)
	}
	return config
}
