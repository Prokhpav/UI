package user_interface

import (
	"github.com/faiface/pixel/text"
	_ "image/png"
	"strconv"

	"github.com/faiface/pixel"
)

type getter struct{}

var (
	Sprites     = map[string]*pixel.Sprite{}
	Fonts       = map[string]map[float64]*text.Atlas{}
	SpriteTypes = map[string][]string{}
	Get         = getter{}
)

func (g getter) Font(name string, size float64) *text.Atlas {
	return Fonts[name][size]
}

func (g getter) SpriteType(SpriteType string, Type int) string {
	return SpriteType + "_" + SpriteTypes[SpriteType][Type]
}

func Cmp(a, min, max float64) float64 {
	if a < min {
		return min
	}
	if a > max {
		return max
	}
	return a
}

func ToBounds(pos pixel.Vec, bounds pixel.Rect) pixel.Vec {
	return pixel.V(Cmp(pos.X, bounds.Min.X, bounds.Max.X), Cmp(pos.Y, bounds.Min.Y, bounds.Max.Y))
}

func StrToInt(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return num
}

func PosInRect(pos pixel.Vec, rect pixel.Rect) pixel.Vec {
	p, s := pos.Sub(rect.Min), rect.Size()
	r := pixel.V(0, 0)
	if s.X != 0 {
		r.X = p.X / s.X
	}
	if s.Y != 0 {
		r.Y = p.Y / s.Y
	}
	return r
}
