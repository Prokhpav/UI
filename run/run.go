package run

import (
	UI "../UIv2"
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

	UI.Load.AllSprites("sprites")
	UI.Load.AllFonts("fonts", 24)

	UI.StdVal = UI.StandardValues{
		ButtonTouchFunc:    func(button *UI.Button) { fmt.Println(button.GetPos()) },
		ButtonTouchButtons: []pixelgl.Button{pixelgl.MouseButtonLeft},
		SpriteSpriteTypes:  []string{"button", "button_mini"},
		TextSpriteFont:     UI.Get.Font("Calibri", 24),
		TextSpriteText:     "",
		TextSpriteColor:    pixel.RGBA{A: 1},
	}

	Scene := UI.Get.Scene(WinSizeW, WinSizeH,
		//UI.Get.Basic(pixel.Vec{X: WinSizeW / 2., Y: WinSizeH / 2.}, pixel.Vec{Y: WinSizeH},
		//	UI.Get.StdButton(V(0, 100), 0, func(button *UI.Button) { fmt.Println(button.GetPos()) }, UI.Get.STS("button1")),
		//	UI.Get.StdButton(V(0, 30), 1, func(button *UI.Button) { fmt.Println(button.GetPos()) }, UI.Get.STS("b2")),
		//	UI.Get.StdButton(V(0, -40), 0, func(button *UI.Button) { fmt.Println(button.GetPos()) }, UI.Get.STS("button3")),
		//	UI.Get.StdButton(V(0, -110), 0, func(button *UI.Button) { fmt.Println(button.GetPos()) }, UI.Get.STS("button4")),
		//),
		UI.Get.ButtonConf(UI.ButtonConfig{PosY: 100}),
		UI.Get.ButtonConf(UI.ButtonConfig{PosY: 30}),
		UI.Get.ButtonConf(UI.ButtonConfig{PosY: -40}),
		UI.Get.ButtonConf(UI.ButtonConfig{PosType: "left-up"}),
	)

	var FMP, LMP, MP pixel.Vec
	T, LT := false, false

	for !win.Closed() {
		MP, LMP = win.MousePosition(), win.MousePreviousPosition()
		if win.JustPressed(pixelgl.MouseButtonLeft) {
			FMP = MP
		}
		T, LT = win.Pressed(pixelgl.MouseButtonLeft), T
		if MP != LMP {
			if !T {
				Scene.MouseMoving(LMP, MP)
			} else {
				Scene.MouseDragging(FMP, LMP, MP, pixelgl.MouseButtonLeft)
			}
		}
		if T != LT {
			if T {
				Scene.MouseTouching(MP, pixelgl.MouseButtonLeft)
			} else {
				Scene.MouseTouchEnding(MP, pixelgl.MouseButtonLeft)
			}
		}

		win.Clear(pixel.RGB(0, 0, 0))
		imd.Draw(win)
		Scene.Draw(win)
		imd.Clear()
		win.Update()

		select {
		case <-TickTimer:
		}
	}
}
