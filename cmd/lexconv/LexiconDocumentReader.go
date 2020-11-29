package main

import (
	"bytes"
	"io"
	"io/ioutil"
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

	content, err1 := ioutil.ReadAll(self.reader)
	if err1 != nil {
		return nil, err1
	}

	rows := bytes.Split(content, []byte{0x0A})
	for _, row := range rows {
		newLexiconDocument.AddLine(row)
	}

	return newLexiconDocument, nil
}
