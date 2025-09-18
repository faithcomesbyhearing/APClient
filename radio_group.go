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
	Value       *bool
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

func (r *RadioGroup) AddItem(label string, description string, value *bool) {
	item := RadioItem{Label: label, Description: description, Value: value}
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
	//r.Group.SetSelected(r.Descriptions[r.Default])
	//r.Entry.SetText(r.Text)
}

func (r RadioGroup) Save() {
	option := strings.Split(r.Group.Selected, " ")[0]
	for i := range r.Items {
		*r.Items[i].Value = r.Items[i].Label == option
	}
}

func (r RadioGroup) Clear() {
	r.Group.SetSelected(r.ItemLabels[r.Default])
	for i := range r.Items {
		*r.Items[i].Value = false
	}
	*r.Items[r.Default].Value = true
}
