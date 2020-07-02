package UIv2

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/text"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"image"
	_ "image/png"
	"io/ioutil"
	"os"
	"strings"
)

type loader struct{}

var Load = loader{}

func (L loader) getFileName(name string) string {
	dot := len(name) - 1
	for dot >= 0 && name[dot] != '.' {
		dot--
	}
	slh := dot - 1
	for slh >= 0 && name[slh] != '/' {
		slh--
	}
	return name[slh+1 : dot]
}

func (L loader) closeFile(file *os.File) {
	err := file.Close()
	if err != nil {
		panic(err)
	}
}

func (L loader) picture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer L.closeFile(file)
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}

func (L loader) Sprite(path string) {
	pic, err := L.picture(path)
	if err != nil {
		panic(err)
	}
	sprite := pixel.NewSprite(pic, pic.Bounds())
	name := L.getFileName(path)
	if _, ok := Sprites[name]; ok {
		name = path
	}
	Sprites[name] = sprite
}

func (L loader) ttf(path string, size float64) (font.Face, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer L.closeFile(file)

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	font_, err := truetype.Parse(bytes)
	if err != nil {
		return nil, err
	}

	return truetype.NewFace(font_, &truetype.Options{
		Size:              size,
		GlyphCacheEntries: 1,
	}), nil
}

func (L loader) Font(path string, size float64) {
	face, err := L.ttf(path, size)
	if err != nil {
		panic(err)
	}

	name := L.getFileName(path)
	if _, ok := Fonts[name]; !ok {
		Fonts[name] = map[float64]*text.Atlas{}
	}
	Fonts[name][size] = text.NewAtlas(face, text.ASCII)
}

func (L loader) AllSprites(directoryPath string) {
	files, err := ioutil.ReadDir(directoryPath)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		filePath := directoryPath + "/" + file.Name()
		if file.IsDir() {
			L.AllSprites(filePath)
		} else if strings.HasSuffix(filePath, ".png") {
			L.ComposeSprite(filePath)
		}
	}
}

func (L loader) AllFonts(directoryPath string, size float64) {
	files, err := ioutil.ReadDir(directoryPath)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		filePath := directoryPath + "/" + file.Name()
		if file.IsDir() {
			L.AllFonts(filePath, size)
		} else {
			L.Font(filePath, size)
		}
	}
}

func (L loader) ComposeSprite(path string) {
	pic, err := L.picture(path)
	if err != nil {
		panic(err)
	}

	dot := len(path) - 1
	for dot >= 0 && path[dot] != '.' {
		dot--
	}
	infoPath := path[:dot] + ".txt"
	infoStr, err := ioutil.ReadFile(infoPath)
	if err != nil { // sprite is not composed
		sprite := pixel.NewSprite(pic, pic.Bounds())
		name := L.getFileName(path)
		Sprites[name] = sprite
		return
	}

	infoList := strings.Split(strings.ReplaceAll(string(infoStr), "\r", ""), "\n")
	bounds := [4]float64{}
	for _, info := range infoList {
		nameList := strings.Split(info, ": ")
		List := strings.Split(nameList[1], ", ")
		for j := 0; j < 4; j++ {
			bounds[j] = float64(StrToInt(List[j]))
		}
		sprite := pixel.NewSprite(pic, pixel.R(bounds[0], bounds[1], bounds[2], bounds[3]))
		name := L.getFileName(path)
		SpriteTypes[name] = append(SpriteTypes[name], nameList[0])
		Sprites[name+"_"+nameList[0]] = sprite
	}
}
