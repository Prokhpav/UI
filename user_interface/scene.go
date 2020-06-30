package user_interface

import (
	"github.com/faiface/pixel"
)

type Scene struct {
	UI []UIInterface
}

func (S *Scene) Add(a ...UIInterface) {
	S.UI = append(S.UI, a...)
}

func (S *Scene) AddButton(center, size pixel.Vec, spriteDefault, spriteFocused, spriteTouched, spriteDragged string, txt *TextSprite, touchFunc func(B *Button)) {
	txt.Pos = txt.Pos.Add(center)
	S.Add(&Button{
		Rect:    pixel.Rect{Min: center.Sub(size.Scaled(0.5)), Max: center.Add(size.Scaled(0.5))},
		Sprites: [4]*pixel.Sprite{AllSprites[spriteDefault], AllSprites[spriteFocused], AllSprites[spriteTouched], AllSprites[spriteDragged]},
		status:  0,
		Text:    txt,
		Func:    touchFunc,
	})
}

func (S *Scene) AddButton2(center, size pixel.Vec, sprites [4]string, txt *TextSprite, touchFunc func(B *Button)) {
	S.AddButton(center, size, sprites[0], sprites[1], sprites[2], sprites[3], txt, touchFunc)
}
