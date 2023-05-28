package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

type MyEntry struct {
	widget.Entry
	entered func(e *MyEntry)
}

func NewMyEntry(f func(e *MyEntry)) *MyEntry {
	e := &MyEntry{}
	e.ExtendBaseWidget(e)
	e.entered = f
	return e
}

func (e *MyEntry) KeyDown(key *fyne.KeyEvent) {
	switch key.Name {
	case fyne.KeyReturn, fyne.KeyEnter:
		e.entered(e)
	default:
		e.Entry.KeyDown(key)
	}
}

func main() {
	a := app.New()
	w := a.NewWindow("Hello")
	l := widget.NewLabel("Hello Fyne!")
	tb := widget.NewToolbar(
		widget.NewToolbarAction(theme.HomeIcon(), func() {
			l.SetText("Select Home Icon!")
		}),
		widget.NewToolbarAction(theme.InfoIcon(), func() {
			l.SetText("Select Information Icon!")
		}),
	)
	mm := fyne.NewMainMenu(
		fyne.NewMenu("File",
			fyne.NewMenuItem("New", func() {
				l.SetText("select New menu item.")
			}),
			fyne.NewMenuItem("Quit", func() {
				a.Quit()
			}),
		),
	)
	w.SetMainMenu(mm)
	e := NewMyEntry(func(e *MyEntry) {
		s := e.Text
		e.SetText("")
		l.SetText("you type '" + s + "'.")
	})
	w.SetContent(
		fyne.NewContainerWithLayout(
			layout.NewBorderLayout(
				nil, tb, nil, nil,
			),
			widget.NewVBox(
				l,e,
			),
			tb,
		),
	)
	a.Settings().SetTheme(theme.LightTheme())
	w.Resize(fyne.NewSize(800, 600))
	w.ShowAndRun()
}
