package common

import (
	"image"
	_ "image/png"
	"os"

	"github.com/faiface/pixel"
)

type ImageAssetManager struct {
	assetMap map[string]image.Image
}

func NewImageAssetManager() ImageAssetManager {
	i := ImageAssetManager{}
	i.assetMap = map[string]image.Image{}
	return i
}
func (im *ImageAssetManager) LoadPictureData(path string) (*pixel.PictureData, error) {
	i, err := im.Load(path)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(i), nil
}
func (im *ImageAssetManager) Load(path string) (image.Image, error) {
	if _, ok := im.assetMap[path]; !ok {
		file, err := os.Open(path)
		if err != nil {
			return nil, err
		}
		defer file.Close()
		img, _, err := image.Decode(file)
		if err != nil {
			return nil, err
		}
		im.assetMap[path] = img
	}
	return im.assetMap[path], nil
}
func (im ImageAssetManager) IsLoaded(path string) bool {
	_, ok := im.assetMap[path]
	return ok
}
func (im *ImageAssetManager) Unload(path string) {
	delete(im.assetMap, path)
}

/*type SpriteSheetAssetManager struct {
	assetMap map[string]*Spritesheet
}

func (im *SpriteSheetAssetManager) Load(path string) (*Spritesheet, error) {

}
func (im SpriteSheetAssetManager) IsLoaded(path string) bool {
	_, ok := im.assetMap[path]
	return ok
}
func (im *SpriteSheetAssetManager) Unload(path string) {
	delete(im.assetMap, path)
}*/
