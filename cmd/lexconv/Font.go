package main

type Font struct {
}

func NewFont() *Font {
	newFont := new(Font)
	return newFont
}

func (self Font) GetGlyphCount() int {
	return 0
}
