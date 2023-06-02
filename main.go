package main

import (
	"fmt"
	"os"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/container"
	"fyne.io/fyne/dialog"
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
	a.Settings().SetTheme(theme.LightTheme())
	w := a.NewWindow("Text Editor by Fyne")
	sta := widget.NewLabel("status.")

	edit := widget.NewEntry()
	edit.MultiLine = true
	sc := container.NewScroll(edit)

	// new file function
	nf := func() {
		dialog.ShowConfirm("Alert", "Create new document?", func(b bool) {
			if b {
				edit.SetText("")
				sta.SetText("create new file.")
			}
		}, w)
	}

	// save file function
	sf := func() {
		sd := dialog.NewFileSave(func(uc fyne.URIWriteCloser, err error) {
			if err != nil {
				dialog.ShowError(err, w)
				return
			}
			if uc == nil {
				return
			}
			fmt.Println(uc.URI().String())
			err = os.WriteFile(uc.URI().String()[7:], []byte(edit.Text), os.ModePerm)
			if err != nil {
				dialog.ShowError(err, w)
				return
			}
			sta.SetText("Save file to '" + uc.URI().String()[7:] + "'.")
		}, w)
		sd.Show()
	}

	// open file function
	of := func() {
		od := dialog.NewFileOpen(func(uc fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, w)
				return
			}
			if uc == nil {
				return
			}
			cont, err := os.ReadFile(uc.URI().String()[7:])
			if err != nil {
				dialog.ShowError(err, w)
				return
			}
			edit.SetText(string(cont))
			sta.SetText("Open file from '" + uc.URI().String()[7:] + "'.")
		}, w)
		od.Show()
	}

	mm := fyne.NewMainMenu(
		fyne.NewMenu("File",
			fyne.NewMenuItem("New", func() {
				nf()
			}),
			fyne.NewMenuItem("Open...", func() {
				of()
			}),
			fyne.NewMenuItem("Save...", func() {
				sf()
			}),
			fyne.NewMenuItem("Quit", func() {
				a.Quit()
			}),
		),
		fyne.NewMenu("Edit",
			fyne.NewMenuItem("Cut", func() {
			}),
			fyne.NewMenuItem("Copy", func() {
			}),
			fyne.NewMenuItem("Paste", func() {
			}),
			fyne.NewMenuItem("Change Theme", func() {
			}),
		),
	)
	w.SetMainMenu(mm)

	w.SetContent(
		fyne.NewContainerWithLayout(
			layout.NewBorderLayout(
				nil, sta, nil, nil,
			),
			sc,
			sta,
		),
	)

	w.Resize(fyne.NewSize(800, 600))
	w.ShowAndRun()
}
