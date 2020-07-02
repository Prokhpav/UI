package user_interface

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func (g getter) ButtonSplit(center, size pixel.Vec, spriteDefault, spriteFocused, spriteTouched, spriteDragged string, txt *TextSprite, touchFunc func(B *Button, v ...float64)) *Button {
	if txt != nil {
		txt.Pos = txt.Pos.Add(center)
	}
	return &Button{
		Rect:    pixel.Rect{Min: center.Sub(size.Scaled(0.5)), Max: center.Add(size.Scaled(0.5))},
		Sprites: [4]*pixel.Sprite{Sprites[spriteDefault], Sprites[spriteFocused], Sprites[spriteTouched], Sprites[spriteDragged]},
		status:  0,
		Text:    txt,
		Func:    touchFunc,
	}
}

func (g getter) Button(center, size pixel.Vec, sprites [4]string, txt *TextSprite, touchFunc func(B *Button, v ...float64)) *Button {
	return g.ButtonSplit(center, size, sprites[0], sprites[1], sprites[2], sprites[3], txt, touchFunc)
}

func (g getter) EasyButton(centerX, centerY float64, spriteType string, txt *TextSprite, touchFunc func(B *Button, v ...float64)) *Button {
	return g.ButtonSplit(pixel.V(centerX, centerY), Sprites[Get.SpriteType(spriteType, 0)].Frame().Size(), Get.SpriteType(spriteType, 0), Get.SpriteType(spriteType, 1), Get.SpriteType(spriteType, 2), Get.SpriteType(spriteType, 3), txt, touchFunc)
}

type Button struct {
	Rect    pixel.Rect
	Sprites [4]*pixel.Sprite
	status  int // 0 - not focused, 1 - focused, 2 - touched, 3 - dragged
	Text    *TextSprite
	Func    func(B *Button, v ...float64)
	//DefaultSpr *pixel.SpriteInterface
	//FocusedSpr *pixel.SpriteInterface
	//TouchedSpr *pixel.SpriteInterface
	//DraggedSpr *pixel.SpriteInterface
}

func (B *Button) Draw(t pixel.Target) {
	B.Sprites[B.status].Draw(t, pixel.IM.Moved(B.Rect.Center()))
	if B.Text != nil {
		B.Text.Draw(t)
	}
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
