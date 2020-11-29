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
	newFontManager.LoadFont('2', "./fonts/VGA2.SFN")
	newFontManager.LoadFont('3', "./fonts/VGA3.SFN")
	newFontManager.LoadFont('4', "./fonts/VGA4.SFN")
	newFontManager.LoadFont('5', "./fonts/VGA5.SFN")
	newFontManager.LoadFont('6', "./fonts/VGA6.SFN")
	newFontManager.LoadFont('7', "./fonts/VGA7.SFN")
	newFontManager.LoadFont('8', "./fonts/VGA8.SFN")
	newFontManager.LoadFont('9', "./fonts/VGA9.SFN")
	newFontManager.LoadFont('A', "./fonts/VGA10.SFN")
	newFontManager.LoadFont('B', "./fonts/VGA11.SFN")

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
