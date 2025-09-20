package main

import (
	"errors"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"strconv"
)

type IntField struct {
	Entry     *widget.Entry
	Container *fyne.Container
	Label     string
	Integer   *int
	Window    fyne.Window
}

func NewIntField(label string, description string, integer *int, window fyne.Window) IntField {
	var r IntField
	r.Entry = widget.NewEntry()
	r.Label = label
	r.Entry.SetPlaceHolder(description)
	r.Container = container.NewBorder(nil, nil,
		container.NewWithoutLayout(
			widget.NewLabel(label),
		), nil, r.Entry)
	r.Integer = integer
	return r
}

func (r IntField) Load() {
	r.Entry.SetText(strconv.Itoa(*r.Integer))
}

func (r IntField) Save() {
	if r.Entry.Text == "" {
		*r.Integer = 0
	} else {
		num, err := strconv.Atoi(r.Entry.Text)
		if err != nil {
			dialog.ShowError(errors.New(r.Label+" must be number or blank"), r.Window)
			return
		}
		*r.Integer = num
	}
}

func (r IntField) Clear() {
	*r.Integer = 0
	r.Entry.SetText("")
}
