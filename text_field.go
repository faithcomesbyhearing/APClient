package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type TextField struct {
	Entry     *widget.Entry
	Container *fyne.Container
	Text      *string
}

func NewTextField(label string, description string, text *string) TextField {
	var f TextField
	f.Entry = widget.NewEntry()
	f.Entry.SetPlaceHolder(description)
	f.Container = container.NewBorder(nil, nil,
		container.NewWithoutLayout(
			widget.NewLabel(label),
		), nil, f.Entry)
	f.Text = text
	return f
}

func (r TextField) Load() {
	r.Entry.SetText(*r.Text)
}

func (r TextField) Save() {
	*r.Text = r.Entry.Text
}

func (r TextField) Clear() {
	*r.Text = ""
	r.Entry.SetText("")
}
