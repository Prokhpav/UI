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
	GetSize() pixel.Vec
	SetPos(NewPosOrDelta pixel.Vec, IsDelta ...bool)
	GetChildren() *SpriteQueue

	MapChildren(f func(child basicInterface))
	setParent(parent basicInterface, index int)
	RemoveFromChildren()
	AddChild(child ...basicInterface)

	//_draw(t pixel.Target, drawPos pixel.Vec)
	//_mouseMoving(LastPos, NewPos pixel.Vec)
	//_mouseTouching(pos pixel.Vec, button pixelgl.Button)
	//_mouseDragging(FirstPos, LastPos, NewPos pixel.Vec, button ...pixelgl.Button)
	//_mouseTouchEnding(pos pixel.Vec, button pixelgl.Button)

	draw(t pixel.Target, drawPos pixel.Vec)
	mouseMoving(LastPos, NewPos pixel.Vec)
	mouseTouching(Pos pixel.Vec, button pixelgl.Button)
	mouseDragging(FirstPos, LastPos, NewPos pixel.Vec, button ...pixelgl.Button)
	mouseTouchEnding(Pos pixel.Vec, button pixelgl.Button)
}

type Basic struct {
	pos      pixel.Vec
	cPos     pixel.Vec
	posType  int
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

func (B *Basic) GetPos() pixel.Vec {
	return B.cPos
}
func (B *Basic) SetPos(NewPosOrDelta pixel.Vec, IsDelta ...bool) {
	if IsDelta != nil && IsDelta[0] {
		B.rect.Min = B.rect.Min.Add(NewPosOrDelta)
		B.rect.Max = B.rect.Max.Add(NewPosOrDelta)
		B.cPos = B.cPos.Add(NewPosOrDelta)
		B.pos = B.pos.Add(NewPosOrDelta)
	} else {
		d := NewPosOrDelta.Sub(B.pos)
		B.rect.Min = B.rect.Min.Add(d)
		B.rect.Max = B.rect.Max.Add(d)
		B.cPos = B.cPos.Add(d)
		B.pos = NewPosOrDelta
	}

}
func (B *Basic) GetSize() pixel.Vec {
	return B.rect.Size()
}
func (B *Basic) GetChildren() *SpriteQueue {
	return B.Children
}

func (B *Basic) MapChildren(f func(child basicInterface)) {
	for _, child := range B.Children.arr {
		f(child)
	}
}
func (B *Basic) setParent(parent basicInterface, index int) {
	B.cPos = getPosFromType2(B.posType, B.pos, B.rect.Size(), parent.GetSize())
	d := B.cPos.Sub(B.pos)
	B.rect.Min = B.rect.Min.Add(d)
	B.rect.Max = B.rect.Max.Add(d)
	B.index = index
	B.parent = parent
}
func (B *Basic) RemoveFromChildren() {
	if B.parent == nil {
		panic("Remove from nil parent")
	}
	B.parent.GetChildren().Pop(B.index)
	B.index = -1
	B.parent = nil
}
func (B *Basic) AddChild(children ...basicInterface) {
	i := len(B.Children.arr)
	for _, child := range children {
		child.setParent(B, i)
		i++
	}
	B.Children.Add(children...)
}

//func (S *Basic) _draw(t pixel.Target, drawPos pixel.Vec) {
//	drawPos = drawPos.Add(S.pos)
//	S.draw(t, drawPos)
//	if S.Children.B() {
//		S.MapChildren(func(child basicInterface) { child._draw(t, drawPos) })
//	}
//}
//func (S *Basic) _mouseMoving(LastPos, NewPos pixel.Vec) {
//	S.mouseMoving(LastPos, NewPos)
//	if S.Children.B() {
//		LastPos, NewPos = LastPos.Sub(S.pos), NewPos.Sub(S.pos)
//		S.MapChildren(func(child basicInterface) { child._mouseMoving(LastPos, NewPos) })
//	}
//}
//func (S *Basic) _mouseTouching(pos pixel.Vec, button pixelgl.Button) {
//	S.mouseTouching(pos, button)
//	if S.Children.B() {
//		pos = pos.Sub(S.pos)
//		S.MapChildren(func(child basicInterface) { child._mouseTouching(pos, button) })
//	}
//}
//func (S *Basic) _mouseDragging(FirstPos, LastPos, NewPos pixel.Vec, button ...pixelgl.Button) {
//	S.mouseDragging(FirstPos, LastPos, NewPos, button...)
//	if S.Children.B() {
//		FirstPos, LastPos, NewPos = FirstPos.Sub(S.pos), LastPos.Sub(S.pos), NewPos.Sub(S.pos)
//		S.MapChildren(func(child basicInterface) { child._mouseDragging(FirstPos, LastPos, NewPos, button...) })
//	}
//}
//func (S *Basic) _mouseTouchEnding(pos pixel.Vec, button pixelgl.Button) {
//	S.mouseTouchEnding(pos, button)
//	if S.Children.B() {
//		pos = pos.Sub(S.pos)
//		S.MapChildren(func(child basicInterface) { child._mouseTouchEnding(pos, button) })
//	}
//}
//
func (B *Basic) draw(_ pixel.Target, _ pixel.Vec)                     {}
func (B *Basic) mouseMoving(_, _ pixel.Vec)                           {}
func (B *Basic) mouseTouching(_ pixel.Vec, _ pixelgl.Button)          {}
func (B *Basic) mouseDragging(_, _, _ pixel.Vec, _ ...pixelgl.Button) {}
func (B *Basic) mouseTouchEnding(_ pixel.Vec, _ pixelgl.Button)       {}

//

func (G getter) Basic(Pos, Size pixel.Vec, children ...basicInterface) *Basic {
	Size = Size.Scaled(0.5)
	S := &Basic{
		pos:      Pos,
		rect:     pixel.Rect{Min: Pos.Sub(Size), Max: Pos.Add(Size)},
		Children: &SpriteQueue{arr: children},
	}
	for i, child := range children {
		child.setParent(S, i)
	}
	return S
}

//func getPosFromType(t string, pos, pSize pixel.Vec) pixel.Vec {
//	if i := strings.Index(t, "-"); i != -1 {
//		return getPosFromType(t[i+1:], getPosFromType(t[:i], pos, pSize), pSize)
//	}
//	if t == "center" {
//		return pixel.Vec{X: pos.X - pSize.X/2, Y: pos.Y - pSize.Y/2,}
//	}
//	if t == "right" {
//		return pixel.Vec{X: pos.X + pSize.X/2, Y: pos.Y}
//	}
//	if t == "left" {
//		return pixel.Vec{X: pos.X - pSize.X/2, Y: pos.Y}
//	}
//	if t == "up" {
//		return pixel.Vec{X: pos.X, Y: pos.Y + pSize.Y/2}
//	}
//	if t == "down" {
//		return pixel.Vec{X: pos.X, Y: pos.Y - pSize.Y/2}
//	}
//	return pos
//}

const _gPFTLLen int = 9

var _getPosFromTypeListNames = [_gPFTLLen]string{"center", "left-down", "left", "left-up", "up", "right-up", "right", "right-down", "down"}
var getPosFromTypeList = [_gPFTLLen][4]float64{
	{0.5, 0.5, 0, 0},   // center
	{0, 0, 0.5, 0.5},   // left-down
	{0, 0.5, 0.5, 0},   // left
	{0, 1, 0.5, -0.5},  // left-up
	{0.5, 1, 0, -0.5},  // up
	{1, 1, -0.5, -0.5}, // right-up
	{1, 0.5, -0.5, 0},  // right
	{1, 0, -0.5, -0.5}, // right-down
	{0.5, 0, 0, 0.5},   // down
}

func getPosFromType2(t int, pos, size, pSize pixel.Vec) pixel.Vec {
	l := getPosFromTypeList[t]
	return pixel.Vec{
		X: pos.X + pSize.X*l[0] + size.X*l[2],
		Y: pos.Y + pSize.Y*l[1] + size.Y*l[3],
	}
}

type BasicConfig struct {
	PosType      string
	Pos          pixel.Vec
	PosX, PosY   float64
	Size         pixel.Vec
	SizeX, SizeY float64
}

func (G getter) BasicConf(config BasicConfig, children ...basicInterface) *Basic {
	if config.Pos == pixel.ZV {
		config.Pos = pixel.Vec{X: config.PosX, Y: config.PosY}
	}
	if config.Size == pixel.ZV {
		config.Size = pixel.Vec{X: config.PosX, Y: config.PosY}
	}
	if config.PosType == "" {
		config.PosType = "center"
	}
	PosType := 0
	for i, TypeName := range _getPosFromTypeListNames {
		if config.PosType == TypeName {
			PosType = i
			break
		}
	}

	s := config.Size.Scaled(0.5)
	B := &Basic{
		pos:      config.Pos,
		posType:  PosType,
		rect:     pixel.Rect{Min: config.Pos.Sub(s), Max: config.Pos.Add(s)},
		Children: &SpriteQueue{arr: children},
	}
	for i, child := range children {
		child.setParent(B, i)
	}
	return B
}
