package main

import (
	"log"
	"os"
	"image"
	"image/color"
	"image/png"
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


func main() {

	/* Prepare Lexicon render system */
	mainFontStream, _ := os.Open("./fonts/VGA0.SFN")
	newFontReader := NewFontReader(mainFontStream)
	mainFont, _ := newFontReader.Read()
	mainFontStream.Close()

	log.Printf("Main Lexicon Font: glyph count = %d", mainFont.GetGlyphCount())

	/* Create A4 portrait paper */
	p := Paper{
		Name: "A4",
		Width: 210,
		Height: 297,
	}

	/* Initialize image with 300 DPI */
	var dpi int = 300
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

	/* Render rune */
	

	/* Save image */
	stream, err := os.Create("image.png")
	if err != nil {
		log.Fatal(err)
	}
	png.Encode(stream, img)
	stream.Close()

}
