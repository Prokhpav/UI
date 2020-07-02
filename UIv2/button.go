package UIv2

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Button struct {
	Sprite
	touchFunc    func(button *Button)
	touchButtons []pixelgl.Button
}

func (B *Button) containsButton(button pixelgl.Button) bool {
	for _, b := range B.touchButtons {
		if button == b {
			return true
		}
	}
	return false
}
func (B *Button) mouseMoving(_, NewPos pixel.Vec) {
	if B.rect.Contains(NewPos) {
		B.spriteNow = 1 // focused
	} else {
		B.spriteNow = 0 // not focused
	}
}
func (B *Button) mouseTouching(_ pixel.Vec, button pixelgl.Button) {
	if B.spriteNow == 1 {
		if B.containsButton(button) {
			B.spriteNow = 2 // touched
		}
	}
}
func (B *Button) mouseDragging(_, _, NewPos pixel.Vec, _ ...pixelgl.Button) {
	if B.spriteNow == 0 || B.spriteNow == 1 {
		if B.rect.Contains(NewPos) {
			B.spriteNow = 1 // focused
		} else {
			B.spriteNow = 0 // not focused
		}
	} else {
		if B.rect.Contains(NewPos) {
			B.spriteNow = 2 // touched
		} else {
			B.spriteNow = 3 // dragged
		}
	}
}
func (B *Button) mouseTouchEnding(_ pixel.Vec, button pixelgl.Button) {
	if !B.containsButton(button) {
		return
	}
	if B.spriteNow == 2 {
		B.spriteNow = 1
		B.touchFunc(B)
	} else if B.spriteNow == 3 {
		B.spriteNow = 0
	}
}

//

func (G getter) Button(sprite *Sprite, touchFunc func(button *Button), title *TextSprite, touchButtons ...pixelgl.Button) *Button {
	if title != nil {
		sprite.AddChild(title)
	}
	if touchButtons == nil {
		touchButtons = StdVal.ButtonTouchButtons
	}
	return &Button{
		Sprite:       *sprite,
		touchFunc:    touchFunc,
		touchButtons: touchButtons,
	}
}

func (G getter) StdButton(Pos pixel.Vec, stdSpriteTypeI int, touchFunc func(button *Button), title *TextSprite, children ...basicInterface) *Button {
	return G.Button(G.StdSprite(Pos, StdVal.ButtonSpriteTypes[stdSpriteTypeI], children...), touchFunc, title, StdVal.ButtonTouchButtons...)
}
