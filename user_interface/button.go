package user_interface

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Button struct {
	Rect    pixel.Rect
	Sprites [4]*pixel.Sprite
	status  int // 0 - not focused, 1 - focused, 2 - touched, 3 - dragged
	Text    *TextSprite
	Func    func(B *Button)
	//DefaultSpr *pixel.Sprite
	//FocusedSpr *pixel.Sprite
	//TouchedSpr *pixel.Sprite
	//DraggedSpr *pixel.Sprite
}

func (B *Button) Draw(t pixel.Target) {
	B.Sprites[B.status].Draw(t, pixel.IM.Moved(B.Rect.Center()))
	B.Text.Draw(t)
}

func (B *Button) MouseMoving(_, NewPos pixel.Vec) {
	if B.Rect.Contains(NewPos) {
		B.status = 1
	} else {
		B.status = 0
	}
}

func (B *Button) MouseTouching(_ pixel.Vec, _ pixelgl.Button) {
	if B.status == 1 {
		B.status = 2
	}
}

func (B *Button) MouseUnTouching(_ pixel.Vec, _ pixelgl.Button) {
	if B.status == 2 {
		B.status = 1
		B.Func(B)
	} else if B.status == 3 {
		B.status = 0
	}
}

func (B *Button) MouseDragging(_, NowPos pixel.Vec, _ pixelgl.Button) {
	if B.status == 0 {
		if B.Rect.Contains(NowPos) {
			B.status = 1
		}
	} else if B.status == 1 {
		if !B.Rect.Contains(NowPos) {
			B.status = 0
		}
	} else if B.status == 2 {
		if !B.Rect.Contains(NowPos) {
			B.status = 3
		}
	} else if B.status == 3 {
		if B.Rect.Contains(NowPos) {
			B.status = 2
		}
	}
}
