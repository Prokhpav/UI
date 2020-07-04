package UIv2

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Scene struct {
	*Basic
}

func (S *Scene) Draw(t pixel.Target) {
	draw(S, t, pixel.V(0, 0))
}

func (S *Scene) MouseMoving(LastPos, NewPos pixel.Vec) {
	mouseMoving(S, LastPos, NewPos)
}
func (S *Scene) MouseTouching(Pos pixel.Vec, button pixelgl.Button) {
	mouseTouching(S, Pos, button)
}
func (S *Scene) MouseDragging(FirstPos, LastPos, NewPos pixel.Vec, button ...pixelgl.Button) {
	mouseDragging(S, FirstPos, LastPos, NewPos, button...)
}
func (S *Scene) MouseTouchEnding(Pos pixel.Vec, button pixelgl.Button) {
	mouseTouchEnding(S, Pos, button)
}

//

func (G getter) Scene(WinSizeW, WinSizeH float64, children ...basicInterface) *Scene {
	S := &Scene{Basic: G.Basic(pixel.ZV, pixel.Vec{X: WinSizeW, Y: WinSizeH}, children...)}
	return S
}
