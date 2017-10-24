package common

import (
	"image"

	"github.com/faiface/pixel"
)

type Spritesheet struct {
	subRect     []pixel.Rect
	picture     *image.Image
	pictureData *pixel.PictureData
}

func NewSpritesheetFromImage(im *image.Image, cellwidth, cellheight float64) Spritesheet {
	s := Spritesheet{}
	s.subRect = []pixel.Rect{}
	s.picture = im
	s.pictureData = pixel.PictureDataFromImage(*im)
	for x := s.pictureData.Bounds().Min.X; x < s.pictureData.Bounds().Max.X; x += cellwidth {
		for y := s.pictureData.Bounds().Min.Y; y < s.pictureData.Bounds().Max.Y; y += cellheight {
			s.subRect = append(s.subRect, pixel.R(x, y, x+cellwidth, y+cellheight))
		}
	}
	return s
}
func (s Spritesheet) GetSprite(i int) *pixel.Sprite {
	if i < 0 || i > len(s.subRect) {
		return nil
	}
	return pixel.NewSprite(s.pictureData, s.subRect[i])
}
