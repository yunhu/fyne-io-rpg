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

	w1 := app.NewWindow("Canvas")
	si := fyne.Size{
		Width:  32 * 20,
		Height: 48 * 10,
	}

	w1.Resize(si)
	dir := "D"
	go func() {
		i := 1
		for i > 0 {
			t := time.Tick(90 * time.Millisecond)
			<-t
			num := i % 2
			var obj fyne.Resource
			obj = stepMap[dir][num]

			//初始化资源
			can := canvas.NewImageFromResource(obj)
			//初始化容器
			container := fyne.NewContainer()
			container.Resize(si)
			//容器加到windows里
			w1.SetContent(container)
			//加布局，设置资源大小
			container.Layout = layout.NewFixedGridLayout(fyne.NewSize(32, 48))
			//
			p := can.Position()
			p.Add(fyne.Position{p.X, i + 122})
			container.AddObject(can)
			can.Move(fyne.Position{p.X, i + 1})
			pos := container.Position()
			//container.Move(fyne.Position{pos.X, i + 1})
			fmt.Println("container", pos.X, pos.Y)
			fmt.Println("obj pos is ", p.X, p.Y)
			fmt.Println("container size is", container.Size())
			fmt.Println("obj size is ", can.Size())
			i++
		}
	}()

	w1.ShowAndRun()

}

//右转

func rgbGradient(x, y, w, h int) color.Color {
	g := int(float32(x) / float32(w) * float32(255))
	b := int(float32(y) / float32(h) * float32(255))

	return color.RGBA{uint8(255 - b), uint8(g), uint8(b), 0xff}
}
