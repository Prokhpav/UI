package user_interface

import (
	"github.com/faiface/pixel/text"
	_ "image/png"

	"github.com/faiface/pixel"
)

var (
	AllSprites = map[string]*pixel.Sprite{}
	AllFonts   = map[string]map[float64]*text.Atlas{}
)

func GetFont(name string, size float64) *text.Atlas {
	return AllFonts[name][size]
}
