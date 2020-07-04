package UIv2

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
)

type StandardValues struct {
	ButtonTouchFunc    func(button *Button)
	ButtonTouchButtons []pixelgl.Button
	SpriteSpriteTypes  []string

	TextSpriteFont  *text.Atlas
	TextSpriteText  string
	TextSpriteColor pixel.RGBA
}

type getter struct{}

var (
	Get    = getter{}
	StdVal = StandardValues{}

	Fonts       = map[string]map[float64]*text.Atlas{}
	Sprites     = map[string]*pixel.Sprite{}
	SpriteTypes = map[string][]string{}
)

func (G getter) pixelSprite(S string) *pixel.Sprite {
	spr, ok := Sprites[S]
	if !ok {
		panic("unknown sprite: " + S)
	}
	return spr
}

func (G getter) spriteTypes(S string) []*pixel.Sprite {
	sprTypes, ok := SpriteTypes[S]
	if !ok {
		panic("unknown sprite: " + S)
	}
	sprites := make([]*pixel.Sprite, len(sprTypes))
	for i, t := range sprTypes {
		sprites[i] = Sprites[S+"_"+t]
	}
	return sprites
}

func (G getter) Font(name string, size float64) *text.Atlas {
	return Fonts[name][size]
}
