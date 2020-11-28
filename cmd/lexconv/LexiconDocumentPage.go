package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
)

type LexiconDocumentPage struct {
	img       draw.Image
	name      string
	imageRect image.Rectangle
}

func NewLexiconDocumentPage(imageRect image.Rectangle) *LexiconDocumentPage {

	newLexiconDocumentPage := new(LexiconDocumentPage)

	newLexiconDocumentPage.imageRect = imageRect
	newLexiconDocumentPage.name = "image.png"

	newLexiconDocumentPage.Reset()

	return newLexiconDocumentPage

}

func (self *LexiconDocumentPage) Reset() {

	/* Create render page */
	self.img = image.NewRGBA(self.imageRect)

	/* Render single pixel */
	//	cyan := color.RGBA{100, 200, 200, 0xFF}

	/* Clear paper */
	for x := 0; x < self.imageRect.Max.X; x++ {
		for y := 0; y < self.imageRect.Max.Y; y++ {
			self.img.Set(x, y, color.White)
		}
	}

}

func (self *LexiconDocumentPage) SetName(name string) {
	self.name = name
}

func (self LexiconDocumentPage) Save() {

	/* Save image */
	stream, err := os.Create(self.name)
	if err != nil {
		log.Fatal(err)
	}
	err1 := png.Encode(stream, self.img)
	if err1 != nil {
		log.Fatal(err1)
	}
	stream.Close()

}

func (self LexiconDocumentPage) GetImage() draw.Image {
	return self.img
}
