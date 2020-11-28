package main

import (
	"log"
	"os"
)

type FontManager struct {
	fonts map[byte]Font
}

func NewFontManager() *FontManager {
	newFontManager := new(FontManager)
	newFontManager.fonts = make(map[byte]Font)
	return newFontManager
}

func (self *FontManager) LoadFont(index byte, name string) error {

	mainFontStream, _ := os.Open(name)
	newFontReader := NewFontReader(mainFontStream)
	newFontReader.SetGlyphHeight(19)
	mainFont, _ := newFontReader.Read()
	log.Printf("Load Font %s: glyph count = %d", name, mainFont.GetGlyphCount())
	mainFontStream.Close()

	/* Register */
	self.fonts[index] = *mainFont

	return nil
}

func (self *FontManager) GetFont(index byte) *Font {
	var newFont Font = self.fonts[index]
	return &newFont
}
