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
		Width:  8 * 20,
		Height: 16 * 10,
	}

	w1.Resize(si)
	var can *canvas.Image

	dir := "D"
	//初始位置 放入通道里
	chPost := make(chan fyne.Position, 1)
	chPost <- fyne.Position{0, 0}
	//go func(c chan fyne.Position) {
	//	h := <-c
	//	h.Y += 1
	//	go func() {
	//		c <- h
	//	}()
	//}(chPost)
	hsize := fyne.NewSize(32, 48)
	go func() {
		i := 0
		for {
			t := time.Tick(30 * time.Millisecond)
			<-t
			num := i % 2
			var obj fyne.Resource
			obj = stepMap[dir][num]

			//初始化资源
			can = canvas.NewImageFromResource(obj)
			//widget.NewVBox()
			//要移动容器
			//fmt.Println("container", pos.X, pos.Y)
			container := fyne.NewContainer()
			container.Layout = layout.NewFixedGridLayout(hsize)
			//container.Layout = layout.NewMaxLayout()
			w1.SetContent(container)
			container.AddObject(can)
			go func() {
				conPos := <-chPost
				fmt.Println(conPos.X, conPos.Y)
				container.Move(fyne.Position{conPos.X, conPos.Y})
				p := container.Position()
				ndir, np := getDir(p, si, hsize, dir)
				fmt.Println(ndir, np)
				dir = ndir

				go func() {
					chPost <- np
				}()
			}()
			i++

		}
	}()

	w1.ShowAndRun()

}
func getDir(p fyne.Position, size fyne.Size, hsize fyne.Size, d string) (dir string, newP fyne.Position) {

	//当前位置是p
	//窗器大小是size
	//人物大小是hsize
	// x = 159 - 32 = 127,0 ~127
	// y = 159  - 48 = 111,0~111

	// D X 不变，y++
	// U x不变 y--

	// R Y不变 x++
	// L Y不变 x--

	nsize := size
	nsize.Width = size.Width - hsize.Width - 1
	nsize.Height = size.Height - hsize.Height - 1

	switch d {
	case "D":

		if p.Y < nsize.Height && p.X < nsize.Width {
			p.Y += 1
			return d, p
		}

		if p.Y == nsize.Height {
			if p.X == 0 {
				d = "R"
				p.X += 1
			} else if p.X == nsize.Width {
				d = "L"
				p.X -= 1
			} else {
				if d == "R" {
					p.X += 1
				}
				if d == "L" {
					p.X -= 1
				}
			}
		}

	case "U":

	case "R":
		if p.Y == nsize.Height && p.X == nsize.Width {
			d = "U"
			p.Y -= 1
		}
		if p.Y == 0 && p.X == nsize.Width {
			d = "L"
			p.X -= 1
		}

	case "L":

	}

	return d, p
}

//右转

func rgbGradient(x, y, w, h int) color.Color {
	g := int(float32(x) / float32(w) * float32(255))
	b := int(float32(y) / float32(h) * float32(255))

	return color.RGBA{uint8(255 - b), uint8(g), uint8(b), 0xff}
}
