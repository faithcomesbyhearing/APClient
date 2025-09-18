package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"strconv"
)

type IntField struct {
	Entry     *widget.Entry
	Container *fyne.Container
	Integer   *int
}

func NewIntField(label string, description string, integer *int) IntField {
	var r IntField
	r.Entry = widget.NewEntry()
	r.Entry.SetPlaceHolder(description)
	r.Container = container.NewBorder(nil, nil,
		container.NewWithoutLayout(
			widget.NewLabel(label),
		), nil, r.Entry)
	r.Integer = integer
	return r
}

func (r IntField) Load() {
	//r.Entry.
	//r.Entry.SetText(*r.Text)
}

func (r IntField) Save() {
	num, err := strconv.Atoi(r.Entry.Text)
	if err != nil {
		fmt.Println("must be a number")
		return
	}
	r.Integer = &num
}

func (r IntField) Clear() {
	*r.Integer = 0
	r.Entry.SetText("")
}
