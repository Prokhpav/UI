package user_interface

import "C"
import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func (g getter) Slider(MinX, MinY, MaxX, MaxY float64, BGroundSprite string, button *Button) *Slider {
	return &Slider{
		Button:        button,
		MovingRect:    pixel.R(MinX, MinY, MaxX, MaxY),
		BGroundSprite: Sprites[BGroundSprite],
	}
}

func (g getter) EasySlider(centerX, centerY float64, BGroundSprite string, button *Button) *Slider {
	bgSpr := Sprites[BGroundSprite]
	min_ := pixel.V(centerX, centerY).Sub(bgSpr.Frame().Size().Scaled(0.5))
	return &Slider{
		Button:        button,
		MovingRect:    pixel.Rect{Min: min_, Max: min_.Add(bgSpr.Frame().Size())},
		BGroundSprite: bgSpr,
	}
}

type Slider struct {
	*Button
	MovingRect    pixel.Rect
	TouchPos      pixel.Vec
	BGroundSprite *pixel.Sprite
}

func (S *Slider) Draw(t pixel.Target) {
	S.BGroundSprite.Draw(t, pixel.IM.Moved(S.MovingRect.Center()))
	S.Button.Draw(t)
}

func (S *Slider) MouseTouching(Pos pixel.Vec, _ pixelgl.Button) {
	if S.status == 1 {
		S.status = 2
		S.TouchPos = Pos.Sub(S.Rect.Center())
	}
}

func (S *Slider) MouseUnTouching(_ pixel.Vec, _ pixelgl.Button) {
	if S.status == 2 {
		S.status = 1
	} else if S.status == 3 {
		S.status = 0
	}
}

func (S *Slider) MouseDragging(_, NowPos pixel.Vec, _ pixelgl.Button) {
	if S.status == 0 {
		if S.Rect.Contains(NowPos) {
			S.status = 1
		}
		return
	}
	if S.status == 1 {
		if !S.Rect.Contains(NowPos) {
			S.status = 0
		}
		return
	}
	size := S.Rect.Size()
	c := ToBounds(NowPos.Sub(S.TouchPos), S.MovingRect)
	S.Rect.Min = c.Sub(size.Scaled(0.5))
	S.Rect.Max = S.Rect.Min.Add(size)
	r := PosInRect(c, S.MovingRect)
	S.Button.Func(S.Button, r.X, r.Y)

	if S.Rect.Contains(NowPos) {
		S.status = 2
	} else {
		S.status = 3
	}
}
