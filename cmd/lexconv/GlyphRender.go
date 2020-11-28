package main

import (
	"image"
	"image/draw"
//	"image/color"
)

type GlyphRender struct {
	font *Font
}

func NewGlyphRender() *GlyphRender {
	newGlyphRender := new(GlyphRender)
	return newGlyphRender
}

func (self *GlyphRender) SetFont(font *Font) {
	self.font = font
}

func (self GlyphRender) Render(glyphIndex byte, x int, y int, page *LexiconDocumentPage) {
	glyph := self.font.GetGlyphByIndex(glyphIndex)
	glyphImage := glyph.GetImage()
	outRect := image.Rectangle{
		Min: image.Point{X: x, Y: y},
		Max: image.Point{X: x+8, Y: y+19},
	}
	out := page.GetImage()
	draw.Draw(out, outRect, glyphImage, image.ZP, draw.Over)
}
