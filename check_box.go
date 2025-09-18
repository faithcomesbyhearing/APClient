package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type CheckField struct {
	CheckBox  *widget.Check
	Container *fyne.Container
	Default   bool
	Value     *bool
}

func NewCheckField(label string, description string, checked bool, value *bool) CheckField {
	var f CheckField
	f.CheckBox = widget.NewCheck(label, func(checked bool) {
		if f.Value != nil {
			*f.Value = checked
		}
	})
	f.Container = container.NewHBox(
		f.CheckBox,
		widget.NewLabel(description),
	)
	f.Value = value
	f.Default = checked
	f.CheckBox.SetChecked(checked)
	return f
}

func (r CheckField) Load() {
	r.CheckBox.SetChecked(*r.Value)
}

func (r CheckField) Save() {
	*r.Value = r.CheckBox.Checked
}

func (r CheckField) Clear() {
	*r.Value = r.Default
	r.CheckBox.SetChecked(r.Default)
}
