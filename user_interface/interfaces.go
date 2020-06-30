package user_interface

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Sprite interface {
	Draw(t pixel.Target)
}

type UIInterface interface {
	Sprite
	MouseMoving(LastPos, NewPos pixel.Vec)
	MouseTouching(Pos pixel.Vec, Button pixelgl.Button)
	MouseUnTouching(Pos pixel.Vec, Button pixelgl.Button)
	MouseDragging(TouchPos, NowPos pixel.Vec, Button pixelgl.Button)
}
