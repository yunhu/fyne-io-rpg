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

	w1.Resize(si)
	//can := canvas.NewImageFromImage(c.Image)
	//var container *fyne.Container
	dir := "D"
	go func() {
		i := 1
		for i > 0 {
			t := time.Tick(10 * time.Millisecond)
			<-t
			num := i % 2
			var obj fyne.Resource
			obj = stepMap[dir][num]

			can := canvas.NewImageFromResource(obj)
			//can.Move(fyne.Position{100, 200})
			//
			//lay := layout.NewGridLayout(2)
			//if i > 1 {
			//	pos := container.Position()
			//	container.Move(fyne.Position{pos.X, i + 1})
			//	fmt.Println(pos.X, pos.Y, i)
			//}
			//container = fyne.NewContainerWithLayout(lay, can)
			//背景
			//先画第一张
			//第二张到m=ove前都要隐藏
			container := fyne.NewContainer()
			w1.SetContent(container)
			container.Layout = layout.NewFixedGridLayout(fyne.NewSize(32, 48))
			container.AddObject(can)

			pos := container.Position()

			//如果大于等于height 那就转
			if i+1 >= si.Height-48 {
				dir = "R"
				container.Move(fyne.Position{pos.X + i, si.Height - 48})
			} else {

				container.Move(fyne.Position{pos.X, i + 1})
			}
			fmt.Println(pos.X, pos.Y)
			////container.Show()
			//
			////container.Hide()
			////container.Hide()
			//w1.Show()
			//po := container.Position()
			//fmt.Println("----")
			i++
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

//右转

func rgbGradient(x, y, w, h int) color.Color {
	g := int(float32(x) / float32(w) * float32(255))
	b := int(float32(y) / float32(h) * float32(255))

	return color.RGBA{uint8(255 - b), uint8(g), uint8(b), 0xff}
}
