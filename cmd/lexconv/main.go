package main

import (
	"image"
	"log"
	"os"
)

type Paper struct {
	Name     string
	Width    int
	Height   int
}

func (self Paper) GetImageRect(dpi int) image.Rectangle {
	var widthNumerator float64 = float64(dpi * self.Width)
	var widthInPixel float64 = widthNumerator / (1.0 * 25.4)
	var heightNumerator float64 = float64(dpi * self.Height)
	var heightInPixel float64 = heightNumerator / (1.0 * 25.4)
	return image.Rectangle{
		Min: image.Point{
			X: 0,
			Y: 0,
		},
		Max: image.Point{
			X: int(widthInPixel),
			Y: int(heightInPixel),
		},
	}
}




func main() {

	/* Prepare Lexicon render system */
	newFontManager := NewFontManager()
	newFontManager.LoadFont('0', "./fonts/VGA0.SFN")
	newFontManager.LoadFont('1', "./fonts/VGA1.SFN")

	/* Document reader */
	mainDocumentStream, _ := os.Open("./os-01.lex")
	documentReader := NewLexiconDocumentReader(mainDocumentStream)
	document, _ := documentReader.Read()
	mainDocumentStream.Close()

	/* Create A4 portrait paper */
	p := Paper{
		Name: "A4",
		Width: 210,
		Height: 297,
	}

	/* Initialize image with 300 DPI */
	var dpi int = 75
	imageRect := p.GetImageRect(dpi)

	log.Printf("  %s    %d x %d mm ( %d x %d )", p.Name, p.Width, p.Height, imageRect.Max.X, imageRect.Max.Y)


	/* Render process */
	documentRender := NewLexiconDocumentRender()
	documentRender.SetDocument(document)
	documentRender.SetFontManager(newFontManager)
	documentRender.SetRectangle(imageRect)
	documentRender.Render()


}
