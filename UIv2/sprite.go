package UIv2

import (
	"github.com/faiface/pixel"
)

type Sprite struct {
	Basic

	sprites   []*pixel.Sprite
	spriteNow int
}

//func (S *Sprite) _draw(t pixel.Target, drawPos pixel.Vec) {
//	drawPos = drawPos.Add(S.Pos)
//	S.draw(t, drawPos)
//	if S.Children.B() {
//		S.MapChildren(func(child basicInterface) {child._draw(t, drawPos)})
//	}
//}

func (S *Sprite) draw(t pixel.Target, drawPos pixel.Vec) {
	S.sprites[S.spriteNow].Draw(t, pixel.IM.Moved(drawPos))
}

//

func (G getter) Sprite(basic *Basic, sprites ...string) *Sprite {
	spr := make([]*pixel.Sprite, len(sprites))
	for i, s := range sprites {
		spr[i] = G.pixelSprite(s)
	}
	return &Sprite{Basic: *basic, sprites: spr}
}

func (G getter) StdSprite(Pos pixel.Vec, spriteType string, children ...basicInterface) *Sprite {
	spr := Get.spriteTypes(spriteType)
	return &Sprite{Basic: *G.Basic(Pos, spr[0].Frame().Size(), children...), sprites: spr}
}
