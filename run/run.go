package run

import (
	UI "../user_interface"
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"math/rand"
	"time"
)

const (
	WinSizeW = 800
	WinSizeH = 600
	FPS      = 60
)

var WinSize = pixel.Vec{X: WinSizeW, Y: WinSizeH}

func Run() {
	rand.Seed(time.Now().UnixNano())
	TickTimer := time.Tick(time.Second / FPS)
	win, err := pixelgl.NewWindow(pixelgl.WindowConfig{
		Title:  "UI test",
		Bounds: pixel.R(0, 0, float64(WinSizeW), float64(WinSizeH)),
	})
	if err != nil {
		panic(err)
	}
	imd := imdraw.New(nil)

	UI.LoadAllSprites("./sprites")
	UI.LoadAllFonts("./fonts", 36)

	UI.DefaultFont = UI.GetFont("Calibri", 36)

	Scene := UI.Scene{}
	ButtonSprites := [4]string{"button_default", "button_focused", "button_touched", "button_dragged"}
	ButtonFunc := func(B *UI.Button) { fmt.Println(B.Text.Text) }
	Scene.AddButton2(WinSize.Scaled(0.5).Add(pixel.V(0, 105.)), pixel.V(200, 80), ButtonSprites, UI.GetDefaultText("Touch"), ButtonFunc)
	Scene.AddButton2(WinSize.Scaled(0.5).Add(pixel.V(0, 35.0)), pixel.V(200, 80), ButtonSprites, UI.GetDefaultText("Button"), ButtonFunc)
	Scene.AddButton2(WinSize.Scaled(0.5).Add(pixel.V(0, -35.)), pixel.V(200, 80), ButtonSprites, UI.GetDefaultText("Shit"), ButtonFunc)
	Scene.AddButton2(WinSize.Scaled(0.5).Add(pixel.V(0, -105)), pixel.V(200, 80), ButtonSprites, UI.GetDefaultText("('v')"), ButtonFunc)

	T, LT := false, false
	for !win.Closed() {
		MP, LMP := win.MousePosition(), win.MousePreviousPosition()
		T, LT = win.Pressed(pixelgl.MouseButtonLeft), T
		if MP != LMP {
			if !T {
				for _, e := range Scene.UI {
					e.MouseMoving(LMP, MP)
				}
			} else {
				for _, e := range Scene.UI {
					e.MouseDragging(LMP, MP, pixelgl.MouseButtonLeft)
				}
			}
		}
		if T != LT {
			if T {
				for _, e := range Scene.UI {
					e.MouseTouching(MP, pixelgl.MouseButtonLeft)
				}
			} else {
				for _, e := range Scene.UI {
					e.MouseUnTouching(MP, pixelgl.MouseButtonLeft)
				}
			}
		}

		win.Clear(pixel.RGB(0, 0, 0))
		imd.Draw(win)

		for _, e := range Scene.UI {
			e.Draw(win)
		}

		imd.Clear()
		win.Update()

		select {
		case <-TickTimer:
		}
	}
}
