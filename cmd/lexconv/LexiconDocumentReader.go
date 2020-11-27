package main

import (
	"io"
	"bytes"
)

type LexiconDocumentReader struct {
	reader        io.Reader
}

func NewLexiconDocumentReader(reader io.Reader) *LexiconDocumentReader {
	newLexiconDocumentReader := new(LexiconDocumentReader)
	newLexiconDocumentReader.reader = reader
	return newLexiconDocumentReader
}

func (self *LexiconDocumentReader) Read() (*LexiconDocument, error) {

	newLexiconDocument := NewLexiconDocument()

	data := make([]byte, 8192)

	self.reader.Read(data)

	rows := bytes.Split(data, []byte{0x0A})
	for _, row := range rows {
		newLexiconDocument.AddLine(row)
	}

	return newLexiconDocument, nil
}
