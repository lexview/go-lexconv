package main

import (
	"io"
)

type FontReader struct {
}

func NewFontReader(io.Reader) *FontReader {
	newFontReader := new(FontReader)
	return newFontReader
}

func (self FontReader) Read() (*Font, error) {
	newFont := new(Font)
	return newFont, nil
}
