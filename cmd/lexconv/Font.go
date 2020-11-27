package main

type Font struct {
	glyphs []Glyph
}

func NewFont() *Font {
	newFont := new(Font)
	return newFont
}

func (self Font) GetGlyphCount() int {
	return len(self.glyphs)
}

func (self *Font) AddGlyph(g Glyph) {
	self.glyphs = append(self.glyphs, g)
}

func (self Font) GetGlyphByIndex(index byte) Glyph {
	return self.glyphs[index]
}
