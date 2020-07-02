package user_interface

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type SpriteInterface interface {
	Draw(t pixel.Target)
}

type UIInterface interface {
	SpriteInterface
	MouseMoving(LastPos, NewPos pixel.Vec)
	MouseTouching(Pos pixel.Vec, Button pixelgl.Button)
	MouseUnTouching(Pos pixel.Vec, Button pixelgl.Button)
	MouseDragging(TouchPos, NowPos pixel.Vec, Button pixelgl.Button)
}
