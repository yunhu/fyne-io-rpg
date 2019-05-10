package main

import (
	"fmt"
	"fy/static"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"image/color"
	"time"
)

func makeCell() fyne.CanvasObject {
	cl := color.RGBA{128, 128, 128, 255}
	rect := canvas.NewRectangle(&cl)
	rect.SetMinSize(fyne.NewSize(1, 1))
	return rect
}
func makeBoxLayout() *fyne.Container {
	top := makeCell()
	bottom := makeCell()
	left := makeCell()
	right := makeCell()
	middle := widget.NewLabelWithStyle("xxxxxx center", fyne.TextAlignCenter, fyne.TextStyle{})

	borderLayout := layout.NewBorderLayout(top, bottom, left, right)
	return fyne.NewContainerWithLayout(borderLayout, top, bottom, left, right, middle)
}
func main() {

	app := app.New()
	//app.

	//app.Settings().SetTheme(theme.LightTheme())
	//
	//w := app.NewWindow("Hello")
	w1 := app.NewWindow("Canvas")
	si := fyne.Size{
		Width:  32 * 20,
		Height: 48 * 10,
	}

	w1.Resize(si)
	//can := canvas.NewImageFromImage(c.Image)
	var container *fyne.Container

	go func() {
		i := 1
		for i > 0 {
			t := time.Tick(100 * time.Millisecond)
			<-t
			num := i % 3
			var obj fyne.Resource

			switch num {
			case 0:
				obj = static.D1
			case 1:
				obj = static.D2
			case 2:
				obj = static.D3
			}

			container = fyne.NewContainer(
				canvas.NewImageFromResource(obj),
			)
			container.Move(fyne.Position{200, 50})
			fmt.Println(container.Position())
			container.Layout = layout.NewFixedGridLayout(fyne.NewSize(32, 48))
			//container.Layout = layout.NewVBoxLayout()
			i++
			container.Resize(fyne.Size{30, 30})
			fmt.Println(container.Size())
			w1.SetContent(container)
		}
	}()

	w1.ShowAndRun()

	//w.SetContent(widget.NewVBox(
	//
	//	widget.NewLabel("Hello Fyne!"),
	//	widget.NewButton("Quit", func() {
	//		app.Quit()
	//
	//	}),
	//
	//	widget.NewCheck("check", func(b bool) {
	//		if b {
	//			//app.Quit()
	//		}
	//	}),
	//	widget.NewButton("quit", func() {
	//		cnf := dialog.NewConfirm("ConfirmINfo", "Are you sure to exit?", func(b bool) {
	//			if b {
	//				w1.Show()
	//			}
	//		}, w)
	//		cnf.SetConfirmText("Yes")
	//		cnf.SetDismissText("NOOOO")
	//		cnf.Show()
	//	}),
	//	widget.NewTabContainer(
	//		widget.NewTabItem("Box", makeBoxLayout()),
	//	),
	//))
	//
	//si := fyne.Size{
	//	Width:  800,
	//	Height: 600,
	//}
	//w.Resize(si)
	//
	//w.ShowAndRun()

}
func rgbGradient(x, y, w, h int) color.Color {
	g := int(float32(x) / float32(w) * float32(255))
	b := int(float32(y) / float32(h) * float32(255))

	return color.RGBA{uint8(255 - b), uint8(g), uint8(b), 0xff}
}
