package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")
	w.SetContent(
		fyne.NewContainerWithLayout(
			layout.NewGridLayout(3),
			widget.NewButton("one", nil),
			widget.NewButton("two", nil),
			widget.NewButton("three", nil),
			widget.NewButton("four", nil),
			layout.NewSpacer(),
			widget.NewButton("five", nil),
		),
	)
	a.Settings().SetTheme(theme.LightTheme())
	w.ShowAndRun()
}
