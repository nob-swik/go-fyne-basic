package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")
	l := widget.NewLabel("Hello Fyne!")
	ne := widget.NewEntry()
	pe := widget.NewPasswordEntry()
	w.SetContent(
		widget.NewVBox(
			l,
			widget.NewForm(
				widget.NewFormItem("Name", ne),
				widget.NewFormItem("Pass", pe),
			),
			widget.NewButton("OK", func() {
				l.SetText(ne.Text + " & " + pe.Text)
			}),
		),
	)
	a.Settings().SetTheme(theme.LightTheme())
	w.ShowAndRun()
}
