package main

import (
	"log"
	"os"
	"image"
	"image/color"
	"image/png"
	"bytes"
	"strconv"
)

type Paper struct {
	Name     string
	Width    int
	Height   int
}

func (self Paper) GetImageRect(dpi int) image.Rectangle {
	var widthNumerator float64 = float64(dpi * self.Width)
	var widthInPixel float64 = widthNumerator / float64(1.0 * 25.4)
	var heightNumerator float64 = float64(dpi * self.Height)
	var heightInPixel float64 = heightNumerator / float64(1.0 * 25.4)
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


func LoadFont(name string) *Font {
	mainFontStream, _ := os.Open(name)
	newFontReader := NewFontReader(mainFontStream)
	newFontReader.SetGlyphHeight(19)
	mainFont, _ := newFontReader.Read()
	log.Printf("Load Font %s: glyph count = %d", name, mainFont.GetGlyphCount())
	mainFontStream.Close()
	return mainFont
}

func main() {

	/* Prepare Lexicon render system */
	mainFont0 := LoadFont("./fonts/VGA0.SFN")
	mainFont1 := LoadFont("./fonts/VGA1.SFN")

	/* Glyph render */
	renderer := NewGlyphRender()

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

	/* Create render page */
	img := image.NewRGBA(imageRect)

	/* Render single pixel */
//	cyan := color.RGBA{100, 200, 200, 0xFF}

	/* Clear paper */
	for x := 0; x < imageRect.Max.X; x++ {
		for y := 0; y < imageRect.Max.Y; y++ {
			img.Set(x, y, color.White)
		}
	}

	/* Render process */
	var defaultBaseline = 19
	var currentBaseline = defaultBaseline
	var controlMode bool = false 
	var currentLine = 0
	var currentPos = 0
	for _, line := range document.GetLines() {

		/* Reset renderer */
		renderer.SetFont(mainFont0)
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
					renderer.SetFont(mainFont0)
				} else if ch == '1' {
					renderer.SetFont(mainFont1)
				}
				controlMode = false
			} else {
				if ch == 0xFF {
					controlMode = true
				} else {
					/* err := */ renderer.Render( ch, currentPos, currentLine, img )
					currentPos = currentPos + 8
				}
			}
		}
		currentLine = currentLine + currentBaseline

	}

	/* Save image */
	stream, err := os.Create("image.png")
	if err != nil {
		log.Fatal(err)
	}
	png.Encode(stream, img)
	stream.Close()

}
