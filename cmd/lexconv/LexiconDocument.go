package main

type LexiconDocument struct {
	lines        [][]byte
}

func NewLexiconDocument() *LexiconDocument {
	newLexiconDocument := new(LexiconDocument)
	return newLexiconDocument
}

func (self LexiconDocument) GetLines() [][]byte {
	return self.lines
}

func (self LexiconDocument) GetLine(idx int) []byte {
	return self.lines[idx]
}

func (self *LexiconDocument) AddLine(line []byte) {
	self.lines = append(self.lines, line)
}
