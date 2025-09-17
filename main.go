package main

import (
	"fyne.io/fyne/v2/app"
)

type FormData struct {
	DatasetName  string `json:"dataset_name"`
	Username     string `json:"username"`
	LanguageISO  string `json:"language_iso"`
	AltLanguage  string `json:"alt_language"`
	TextData     string `json:"text_data"`
	AudioData    string `json:"audio_data"`
	Timestamps   string `json:"timestamps"`
	Training     string `json:"training"`
	SpeechToText string `json:"speech_to_text"`
	Compare      string `json:"compare"`
	GordonFilter string `json:"gordon_filter"`
}

func main() {
	myApp := app.New()
	PresentForm(myApp)
	//fmt.Println(formData)
}
