package user_interface

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/text"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"image"
	"io/ioutil"
	"os"
)

func GetFileName(name string) string {
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

func CloseFile(file *os.File) {
	err := file.Close()
	if err != nil {
		panic(err)
	}
}

func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer CloseFile(file)

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}

func LoadSprite(path string) {
	pic, err := loadPicture(path)
	if err != nil {
		panic(err)
	}
	sprite := pixel.NewSprite(pic, pic.Bounds())
	name := GetFileName(path)
	if _, ok := AllSprites[name]; ok {
		name = path
	}
	AllSprites[name] = sprite
}

func loadTTF(path string, size float64) (font.Face, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer CloseFile(file)

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

func LoadFont(path string, size float64) {
	face, err := loadTTF(path, size)
	if err != nil {
		panic(err)
	}

	name := GetFileName(path)
	if _, ok := AllFonts[name]; !ok {
		AllFonts[name] = map[float64]*text.Atlas{}
	}
	AllFonts[name][size] = text.NewAtlas(face, text.ASCII)
}

func LoadAllSprites(directoryPath string) {
	files, err := ioutil.ReadDir(directoryPath)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		filePath := directoryPath + "/" + file.Name()
		if file.IsDir() {
			LoadAllSprites(filePath)
		} else {
			LoadSprite(filePath)
		}
	}
}

func LoadAllFonts(directoryPath string, size float64) {
	files, err := ioutil.ReadDir(directoryPath)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		filePath := directoryPath + "/" + file.Name()
		if file.IsDir() {
			LoadAllFonts(filePath, size)
		} else {
			LoadFont(filePath, size)
		}
	}
}
