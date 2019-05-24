package main

import (
	"fy/static"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"image/color"
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
	var stepMap = make(map[string]map[int]*fyne.StaticResource)
	var subUStempMap = make(map[int]*fyne.StaticResource)
	var subDStempMap = make(map[int]*fyne.StaticResource)
	var subLStempMap = make(map[int]*fyne.StaticResource)
	var subRStempMap = make(map[int]*fyne.StaticResource)
	subUStempMap[0] = static.U1
	subUStempMap[1] = static.U2
	subUStempMap[2] = static.U3
	stepMap["U"] = subUStempMap
	stepMap["D"] = subDStempMap
	stepMap["L"] = subLStempMap
	stepMap["R"] = subRStempMap

	subDStempMap[0] = static.D1
	subDStempMap[1] = static.D2
	subDStempMap[2] = static.D3

	subLStempMap[0] = static.L1
	subLStempMap[1] = static.L2
	subLStempMap[2] = static.L3

	subRStempMap[0] = static.R1
	subRStempMap[1] = static.R2
	subRStempMap[2] = static.R3

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

	var obj fyne.Resource
	obj = stepMap["D"][0]
	can := canvas.NewImageFromResource(obj)
	container := fyne.NewContainer()
	container.Layout = layout.NewFixedGridLayout(fyne.NewSize(32, 48))
	container.AddObject(can)
	w1.SetContent(container)
	container.Move(fyne.Position{200, 200})

	w1.Resize(si)
	w1.ShowAndRun()

}

//右转

func rgbGradient(x, y, w, h int) color.Color {
	g := int(float32(x) / float32(w) * float32(255))
	b := int(float32(y) / float32(h) * float32(255))

	return color.RGBA{uint8(255 - b), uint8(g), uint8(b), 0xff}
}
