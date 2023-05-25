package main

import (
	"strconv"

	"fyne.io/fyne/app"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")
	l := widget.NewLabel("Hello Fyne!")
	s := widget.NewSlider(0.0, 100.)
	s.SetValue(20)
	b := widget.NewButton("Check", func ()  {
		l.SetText("value: " + strconv.Itoa(int(s.Value)))
	})
	w.SetContent(
		widget.NewVBox(
			l, s, b,
		),
	)
	a.Settings().SetTheme(theme.LightTheme())
	w.ShowAndRun()
}
