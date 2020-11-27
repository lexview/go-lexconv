package main

import (
	"log"
	"io"
)

type FontReader struct {
	stream       io.Reader
	glyphHeight  int
}

func NewFontReader(stream io.Reader) *FontReader {
	newFontReader := new(FontReader)
	newFontReader.stream = stream
	return newFontReader
}

func (self *FontReader) SetGlyphHeight(height int) {
	self.glyphHeight = height
}

func (self FontReader) Read() (*Font, error) {

	newFont := new(Font)

	/* Step 1. Get glyph height */
//	if (self.glyphHeight == 0) {
//		// Step 1. Get stream size
//		// Step 2. Stream size divide on 256
//	}

	/* Step 2. Reading glyph */
	for {
		data := make([]byte, 19)
		n, err := self.stream.Read(data)
		if err != nil {
			log.Printf("err = %+v", err)
			break
		}

		newGlyph := NewGlyph()
		newGlyph.Parse(data)
		newFont.AddGlyph(*newGlyph)

		log.Printf("size = %d vec = %+v", n, data)
	}

	return newFont, nil
}
