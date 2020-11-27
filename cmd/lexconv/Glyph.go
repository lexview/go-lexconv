package main

import (
	"image"
	"image/color"
)

type Glyph struct {
	img *image.RGBA
}

func NewGlyph() *Glyph {
	newGlyph := new(Glyph)
	return newGlyph
}

func (self *Glyph) Parse(data []byte) {

	img := image.NewRGBA(image.Rectangle{
		Min: image.Point{X: 0, Y: 0},
		Max: image.Point{X: 8, Y: 19},
	})

	for idx, row := range data {
		for bit := 0; bit < 8; bit++ {
			var mask byte = 0x80 >> bit
			if row & mask == mask {
				img.Set(bit, idx, color.Black)
			}
		}
	}

	self.img = img

}

func (self Glyph) GetImage() *image.RGBA {
	return self.img
}
