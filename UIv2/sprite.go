package UIv2

import (
	"github.com/faiface/pixel"
)

type Sprite struct {
	*Basic

	sprites   []*pixel.Sprite
	spriteNow int
}

//func (S *Sprite) _draw(t pixel.Target, drawPos pixel.Vec) {
//	drawPos = drawPos.Add(S.pos)
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
	return &Sprite{Basic: basic, sprites: spr}
}

func (G getter) StdSprite(Pos pixel.Vec, spriteType string, children ...basicInterface) *Sprite {
	spr := Get.spriteTypes(spriteType)
	return &Sprite{Basic: G.Basic(Pos, spr[0].Frame().Size(), children...), sprites: spr}
}

//

type SpriteConfig struct {
	PosType       string
	Pos           pixel.Vec
	PosX, PosY    float64
	Size          pixel.Vec
	SizeX, SizeY  float64
	Sprites       []*pixel.Sprite
	SpriteType    string
	StdSpriteType int
	SpriteNow     int
}

func (G getter) SpriteConf(config SpriteConfig, children ...basicInterface) *Sprite {
	if config.SpriteType != "" {
		config.Sprites = append(config.Sprites, G.spriteTypes(config.SpriteType)...)
	} else if config.Sprites == nil {
		config.Sprites = G.spriteTypes(StdVal.SpriteSpriteTypes[config.StdSpriteType])
	}
	if config.Size == pixel.ZV && config.SizeX == 0 && config.SizeY == 0 {
		config.Size = config.Sprites[0].Frame().Size()
	}
	return &Sprite{
		Basic: G.BasicConf(BasicConfig{
			PosType: config.PosType,
			Pos:     config.Pos,
			PosX:    config.PosX,
			PosY:    config.PosY,
			Size:    config.Size,
			SizeX:   config.SizeX,
			SizeY:   config.SizeY,
		}, children...),
		sprites:   config.Sprites,
		spriteNow: config.SpriteNow,
	}
}
