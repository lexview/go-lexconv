package main

import (
	"bytes"
	"fmt"
	"image"
	"log"
	"strconv"
)

type LexiconDocumentRender struct {
	document    *LexiconDocument
	fontManager *FontManager
	renderer    *GlyphRender
	imageRect   image.Rectangle
}

func NewLexiconDocumentRender() *LexiconDocumentRender {
	newLexiconDocumentRender := new(LexiconDocumentRender)
	newLexiconDocumentRender.renderer = NewGlyphRender()
	return newLexiconDocumentRender
}

func (self *LexiconDocumentRender) SetDocument(document *LexiconDocument) {
	self.document = document
}

func (self *LexiconDocumentRender) SetFontManager(fontManager *FontManager) {
	self.fontManager = fontManager
}

func (self *LexiconDocumentRender) SetRectangle(imageRect image.Rectangle) {
	self.imageRect = imageRect
}

func (self *LexiconDocumentRender) Render() {

	var pageIndex int = 1
	var page *LexiconDocumentPage

	page = NewLexiconDocumentPage(self.imageRect)

	var defaultBaseline = 19
	var currentBaseline = defaultBaseline
	var controlMode bool = false
	var currentLine = 0
	var currentPos = 0
	var underline bool = false

	for _, line := range self.document.GetLines() {

		if currentLine + 32 > self.imageRect.Max.Y {
			//
			page.SetName(fmt.Sprintf("page_%04d.png", pageIndex))
			page.Save()
			//
			page = NewLexiconDocumentPage(self.imageRect)
			//
			currentLine = 0
			pageIndex += 1
		}

		/* Reset renderer */
		self.renderer.SetFont(self.fontManager.GetFont('0'))
		underline = false
		data := line.GetData()
		currentPos = 0

		/* Change baseline size */
		if bytes.HasPrefix(data, []byte{0xFF, 0xE8}) {
			newMultiply := string(data[2:])
			value, _ := strconv.ParseFloat(newMultiply, 64)
			log.Printf("baseline = %+v", value)
			currentBaseline = int(19 * value)
			continue
		}

		for _, ch := range data {
			if controlMode {
				if ch == '0' {
					self.renderer.SetFont(self.fontManager.GetFont('0'))
				} else if ch == '1' {
					self.renderer.SetFont(self.fontManager.GetFont('1'))
				} else if ch == '2' {
					self.renderer.SetFont(self.fontManager.GetFont('2'))
				} else if ch == '3' {
					self.renderer.SetFont(self.fontManager.GetFont('3'))
				} else if ch == '4' {
					self.renderer.SetFont(self.fontManager.GetFont('4'))
				} else if ch == '5' {
					self.renderer.SetFont(self.fontManager.GetFont('5'))
				} else if ch == '6' {
					self.renderer.SetFont(self.fontManager.GetFont('6'))
				} else if ch == '7' {
					self.renderer.SetFont(self.fontManager.GetFont('7'))
				} else if ch == '8' {
					self.renderer.SetFont(self.fontManager.GetFont('8'))
				} else if ch == '9' {
					self.renderer.SetFont(self.fontManager.GetFont('9'))
				} else if ch == 'A' {
					self.renderer.SetFont(self.fontManager.GetFont('A'))
				} else if ch == 'B' {
					self.renderer.SetFont(self.fontManager.GetFont('B'))
				} else if ch == '_' {
					underline = true
				} else if ch == '.' {
					underline = false
				} else {
					log.Printf("warn: wrong directive %c", ch)
				}

				controlMode = false

			} else {
				if ch == 0xFF {
					controlMode = true
				} else {
					/* err := */ self.renderer.Render( ch, currentPos, currentLine, page )
					if underline {
						self.renderer.Render( '_', currentPos, currentLine + 2, page )
					}
					currentPos = currentPos + 8
				}
			}
		}
		currentLine = currentLine + currentBaseline

	}

	/**/
	page.SetName(fmt.Sprintf("page_%04d.png", pageIndex))
	page.Save()

}
