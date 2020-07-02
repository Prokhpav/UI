package user_interface

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/text"
)

var (
	DefaultPos       = pixel.Vec{X: 0, Y: 0}
	DefaultFont      *text.Atlas
	DefaultTextColor = pixel.RGBA{A: 1}
)

func (g getter) TextSprite(pos pixel.Vec, txt string, font *text.Atlas, color pixel.RGBA) *TextSprite {
	T := &TextSprite{Pos: pos, Text: txt, font: font, Sprite: text.New(pixel.V(0, 0), font)}
	T.Sprite.Color = color
	T.SetText(txt)
	return T
}

func (g getter) DefaultText(txt string) *TextSprite {
	return g.TextSprite(DefaultPos, txt, DefaultFont, DefaultTextColor)
}

type TextSprite struct {
	Pos    pixel.Vec
	Text   string
	font   *text.Atlas
	Sprite *text.Text
}

func (T *TextSprite) SetText(txt string) {
	T.Text = txt
	T.Sprite.Clear()
	_, err := T.Sprite.WriteString(txt)
	if err != nil {
		panic(err)
	}
}

func (T *TextSprite) Draw(t pixel.Target) {
	T.Sprite.Draw(t, pixel.IM.Moved(T.Pos.Sub(T.Sprite.Bounds().Center())))
}
