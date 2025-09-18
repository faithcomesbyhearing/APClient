package main

import (
	"fyne.io/fyne/v2/app"
)

type Field interface {
	Load()
	Save()
	Clear()
}

func main() {
	myApp := app.NewWithID("org.fcbh.artificial_polyglot")
	PresentForm(myApp)
}
