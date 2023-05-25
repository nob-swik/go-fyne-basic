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
	r := widget.NewRadioGroup(
		[]string{"ONE", "TWO", "THREE"},
		func(s string) {
			if s == "" {
				l.SetText("not selected.")
			} else {
				l.SetText("selected: " + s)
			}
		})
		r.SetSelected("ONE")
	w.SetContent(
		widget.NewVBox(
			l, r,
		),
	)
	a.Settings().SetTheme(theme.LightTheme())
	w.ShowAndRun()
}
