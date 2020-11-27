package main

import (
	"bytes"
)

type Line struct {
	data []byte
}

func (self Line) GetData() []byte {
	return self.data
}


type LexiconDocument struct {
	lines        []Line
}

func NewLexiconDocument() *LexiconDocument {
	newLexiconDocument := new(LexiconDocument)
	return newLexiconDocument
}

func (self LexiconDocument) GetLines() []Line {
	return self.lines
}

func (self LexiconDocument) GetLine(idx int) Line {
	return self.lines[idx]
}

func (self *LexiconDocument) AddLine(data []byte) {

	var newData []byte = data
	newData = bytes.TrimSuffix(newData, []byte{0x0D})
	newData = bytes.TrimSuffix(newData, []byte{0x0A})

	newLine := Line{data: newData}

	self.lines = append(self.lines, newLine)
}
