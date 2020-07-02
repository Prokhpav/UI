package UIv2

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"strconv"
)

func StrToInt(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return num
}

type basicInterface interface {
	GetPos() pixel.Vec
	GetChildren() *SpriteQueue
	ChangePos(NewPos pixel.Vec)

	MapChildren(f func(child basicInterface))
	setParent(parent basicInterface, index int)
	RemoveFromChildren()
	AddChild(child ...basicInterface)

	//_draw(t pixel.Target, drawPos pixel.Vec)
	//_mouseMoving(LastPos, NewPos pixel.Vec)
	//_mouseTouching(Pos pixel.Vec, button pixelgl.Button)
	//_mouseDragging(FirstPos, LastPos, NewPos pixel.Vec, button ...pixelgl.Button)
	//_mouseTouchEnding(Pos pixel.Vec, button pixelgl.Button)

	draw(t pixel.Target, drawPos pixel.Vec)
	mouseMoving(LastPos, NewPos pixel.Vec)
	mouseTouching(Pos pixel.Vec, button pixelgl.Button)
	mouseDragging(FirstPos, LastPos, NewPos pixel.Vec, button ...pixelgl.Button)
	mouseTouchEnding(Pos pixel.Vec, button pixelgl.Button)
}

type Basic struct {
	Pos      pixel.Vec
	rect     pixel.Rect
	Children *SpriteQueue
	index    int
	parent   basicInterface
}

func draw(S basicInterface, t pixel.Target, drawPos pixel.Vec) {
	drawPos = drawPos.Add(S.GetPos())
	S.draw(t, drawPos)
	if S.GetChildren().B() {
		S.MapChildren(func(child basicInterface) { draw(child, t, drawPos) })
	}
}
func mouseMoving(S basicInterface, LastPos, NewPos pixel.Vec) {
	S.mouseMoving(LastPos, NewPos)
	if S.GetChildren().B() {
		LastPos, NewPos = LastPos.Sub(S.GetPos()), NewPos.Sub(S.GetPos())
		S.MapChildren(func(child basicInterface) { mouseMoving(child, LastPos, NewPos) })
	}
}
func mouseTouching(S basicInterface, Pos pixel.Vec, button pixelgl.Button) {
	S.mouseTouching(Pos, button)
	if S.GetChildren().B() {
		Pos = Pos.Sub(S.GetPos())
		S.MapChildren(func(child basicInterface) { mouseTouching(child, Pos, button) })
	}
}
func mouseDragging(S basicInterface, FirstPos, LastPos, NewPos pixel.Vec, button ...pixelgl.Button) {
	S.mouseDragging(FirstPos, LastPos, NewPos, button...)
	if S.GetChildren().B() {
		FirstPos, LastPos, NewPos = FirstPos.Sub(S.GetPos()), LastPos.Sub(S.GetPos()), NewPos.Sub(S.GetPos())
		S.MapChildren(func(child basicInterface) { mouseDragging(child, FirstPos, LastPos, NewPos, button...) })
	}
}
func mouseTouchEnding(S basicInterface, Pos pixel.Vec, button pixelgl.Button) {
	S.mouseTouchEnding(Pos, button)
	if S.GetChildren().B() {
		Pos = Pos.Sub(S.GetPos())
		S.MapChildren(func(child basicInterface) { mouseTouchEnding(child, Pos, button) })
	}
}

func (S *Basic) GetPos() pixel.Vec {
	return S.Pos
}
func (S *Basic) GetChildren() *SpriteQueue {
	return S.Children
}
func (S *Basic) ChangePos(NewPos pixel.Vec) {
	d := NewPos.Sub(S.Pos)
	S.rect.Min = S.rect.Min.Add(d)
	S.rect.Max = S.rect.Max.Add(d)
	S.Pos = NewPos
}

func (S *Basic) MapChildren(f func(child basicInterface)) {
	for _, child := range S.Children.arr {
		f(child)
	}
}
func (S *Basic) setParent(parent basicInterface, index int) {
	S.index = index
	S.parent = parent
}
func (S *Basic) RemoveFromChildren() {
	if S.parent == nil {
		panic("Remove from nil parent")
	}
	S.parent.GetChildren().Pop(S.index)
	S.index = -1
	S.parent = nil
}
func (S *Basic) AddChild(children ...basicInterface) {
	i := len(S.Children.arr)
	for _, child := range children {
		child.setParent(S, i)
		i++
	}
	S.Children.Add(children...)
}

//func (S *Basic) _draw(t pixel.Target, drawPos pixel.Vec) {
//	drawPos = drawPos.Add(S.Pos)
//	S.draw(t, drawPos)
//	if S.Children.B() {
//		S.MapChildren(func(child basicInterface) { child._draw(t, drawPos) })
//	}
//}
//func (S *Basic) _mouseMoving(LastPos, NewPos pixel.Vec) {
//	S.mouseMoving(LastPos, NewPos)
//	if S.Children.B() {
//		LastPos, NewPos = LastPos.Sub(S.Pos), NewPos.Sub(S.Pos)
//		S.MapChildren(func(child basicInterface) { child._mouseMoving(LastPos, NewPos) })
//	}
//}
//func (S *Basic) _mouseTouching(Pos pixel.Vec, button pixelgl.Button) {
//	S.mouseTouching(Pos, button)
//	if S.Children.B() {
//		Pos = Pos.Sub(S.Pos)
//		S.MapChildren(func(child basicInterface) { child._mouseTouching(Pos, button) })
//	}
//}
//func (S *Basic) _mouseDragging(FirstPos, LastPos, NewPos pixel.Vec, button ...pixelgl.Button) {
//	S.mouseDragging(FirstPos, LastPos, NewPos, button...)
//	if S.Children.B() {
//		FirstPos, LastPos, NewPos = FirstPos.Sub(S.Pos), LastPos.Sub(S.Pos), NewPos.Sub(S.Pos)
//		S.MapChildren(func(child basicInterface) { child._mouseDragging(FirstPos, LastPos, NewPos, button...) })
//	}
//}
//func (S *Basic) _mouseTouchEnding(Pos pixel.Vec, button pixelgl.Button) {
//	S.mouseTouchEnding(Pos, button)
//	if S.Children.B() {
//		Pos = Pos.Sub(S.Pos)
//		S.MapChildren(func(child basicInterface) { child._mouseTouchEnding(Pos, button) })
//	}
//}
//
func (S *Basic) draw(_ pixel.Target, _ pixel.Vec)                     {}
func (S *Basic) mouseMoving(_, _ pixel.Vec)                           {}
func (S *Basic) mouseTouching(_ pixel.Vec, _ pixelgl.Button)          {}
func (S *Basic) mouseDragging(_, _, _ pixel.Vec, _ ...pixelgl.Button) {}
func (S *Basic) mouseTouchEnding(_ pixel.Vec, _ pixelgl.Button)       {}

//

func (G getter) Basic(Pos, Size pixel.Vec, children ...basicInterface) *Basic {
	Size = Size.Scaled(0.5)
	S := &Basic{
		Pos:      Pos,
		rect:     pixel.Rect{Min: Pos.Sub(Size), Max: Pos.Add(Size)},
		Children: &SpriteQueue{arr: children},
	}
	for i, child := range children {
		child.setParent(S, i)
	}
	return S
}
