package UIv2

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/text"
)

type TextSprite struct {
	*Basic
	text   string
	color  pixel.RGBA
	Sprite *text.Text
}

func (T *TextSprite) Text() string {
	return T.text
}

func (T *TextSprite) SetText(txt string) {
	T.text = txt
	T.Sprite.Clear()
	_, err := T.Sprite.WriteString(txt)
	if err != nil {
		panic(err)
	}
	T.rect = T.Sprite.Bounds().Moved(T.pos)
}

func (T *TextSprite) draw(t pixel.Target, drawPos pixel.Vec) {
	T.Sprite.Draw(t, pixel.IM.Moved(drawPos.Sub(T.Sprite.Bounds().Center())))
}

//

func (G getter) TextSprite(basic *Basic, text_, fontName string, fontSize float64, color pixel.RGBA) *TextSprite {
	font := G.Font(fontName, fontSize)
	T := &TextSprite{
		Basic:  basic,
		Sprite: text.New(pixel.V(0, 0), font),
	}
	T.Sprite.Color = color
	T.SetText(text_)
	return T
}

func (G getter) StdTextSprite(Pos pixel.Vec, text_ string, children ...basicInterface) *TextSprite {
	if text_ == "" {
		text_ = StdVal.TextSpriteText
	}
	T := &TextSprite{
		Basic:  Get.Basic(Pos, pixel.ZV, children...),
		Sprite: text.New(pixel.V(0, 0), StdVal.TextSpriteFont),
	}
	T.Sprite.Color = StdVal.TextSpriteColor
	T.SetText(text_)
	return T
}

func (G getter) STS(text_ string) *TextSprite {
	return G.StdTextSprite(pixel.ZV, text_)
}
