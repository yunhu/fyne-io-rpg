package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	app := app.New()

	w := app.NewWindow("Hello")
	w2 := app.NewWindow("win 2")
	w.SetContent(widget.NewVBox(
		widget.NewLabel("Hello Fyne!"),
		widget.NewButton("Quit", func() {
			w2.SetOnClosed(func() {
				w2.Hide()
			})
			w2.Resize(fyne.Size{100, 200})
			w2.Show()

		}),
	))

	w.Resize(fyne.Size{400, 600})
	w.Show()
	app.Run()
}
