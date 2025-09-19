package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"strings"
)

type RadioItem struct {
	Label       string
	Description string
	IsBool      bool
	BoolValue   *bool
	IntValue    *int
}

func (r RadioItem) IsTrue() bool {
	if r.IsBool {
		return *r.BoolValue
	} else {
		return *r.IntValue != 0
	}
}
func (r RadioItem) SetBool(b bool) {
	if r.IsBool {
		*r.BoolValue = b
	} else if b {
		*r.IntValue = 1
	} else {
		*r.IntValue = 0
	}
}

type RadioGroup struct {
	Label      *widget.Label
	Group      *widget.RadioGroup
	Items      []RadioItem
	ItemLabels []string
	Default    int
	Container  *fyne.Container
}

func NewRadioGroup(label string) RadioGroup {
	var r RadioGroup
	r.Label = widget.NewLabel(label)
	r.Items = make([]RadioItem, 0, 10)
	return r
}

func (r *RadioGroup) AddItem(label string, description string, isBool bool,
	boolValue *bool, intValue *int) {
	item := RadioItem{Label: label, Description: description, IsBool: isBool,
		BoolValue: boolValue, IntValue: intValue}
	r.Items = append(r.Items, item)
}

func (r *RadioGroup) SetSelected(selected int) {
	r.Default = selected
	for _, item := range r.Items {
		r.ItemLabels = append(r.ItemLabels, item.Label+" - "+item.Description)
	}
	r.Group = widget.NewRadioGroup(r.ItemLabels, nil)
	r.Group.SetSelected(r.ItemLabels[selected])
	r.Container = container.NewVBox(r.Label, r.Group)
}

func (r RadioGroup) Load() {
	count := 0
	for i := range r.Items {
		if r.Items[i].IsTrue() {
			r.Group.SetSelected(r.ItemLabels[i])
			count++
		}
	}
	if count == 0 { // Set Defaults
		r.Group.SetSelected(r.ItemLabels[len(r.Items)-1])
	}
}

func (r RadioGroup) Save() {
	option := strings.Split(r.Group.Selected, " ")[0]
	for i := range r.Items {
		if r.Items[i].Label == option {
			r.Items[i].SetBool(true)
		} else {
			r.Items[i].SetBool(false)
		}
	}
}

func (r RadioGroup) Clear() {
	r.Group.SetSelected(r.ItemLabels[r.Default])
	for i := range r.Items {
		r.Items[i].SetBool(false)
	}
	r.Items[r.Default].SetBool(true)
}
