package main

import (
    "os"
    "github.com/llgcode/draw2d"
    "github.com/llgcode/draw2d/draw2dimg"
    "github.com/llgcode/draw2d/draw2dkit"
    "github.com/golang/freetype/truetype"
    "golang.org/x/image/font/gofont/goregular"
    "image"
	"image/color"
)

const leftMargin = 10
const rightMargin = 40
const topMargin = 60
const bottomMargin = 10
const spacing = 30
const fretCount = 4
const fontSize = 24
const dotSize = 6

func main() {
    loadFont()

    for _, arg := range os.Args[1:] {
        chord := parseChord(arg)
        stringCount := len(chord.strings) - 1

        dest := image.NewRGBA(image.Rect(0, 0, leftMargin + rightMargin + spacing * stringCount, topMargin + bottomMargin + fretCount * spacing))
        gc := draw2dimg.NewGraphicContext(dest)

        drawChord(gc, chord, stringCount)

        draw2dimg.SaveToPngFile(chord.name + ".png", dest)
    }
}

func loadFont() {
    font, err := truetype.Parse(goregular.TTF)
	if err == nil {
        draw2d.RegisterFont(draw2d.FontData{Name: "luxi"}, font)
    }
}

func drawChord(gc draw2d.GraphicContext, chord chordInfo, stringCount int) {
	gc.SetFillColor(color.RGBA{0xff, 0xff, 0xff, 0xff})
	gc.SetStrokeColor(color.RGBA{0x00, 0x00, 0x00, 0xff})
	gc.SetLineWidth(2)

    drawBox(gc, stringCount)
    drawNut(gc, stringCount)
    drawFrets(gc, stringCount)
    drawStrings(gc, stringCount)
    drawText(gc, leftMargin, 5 + fontSize, chord.name)

    for i, v := range chord.strings {
        if v.mainFret > 0 {
            gc.SetFillColor(color.RGBA{0xc0, 0xc0, 0xc0, 0xff})
        } else {
            gc.SetFillColor(color.RGBA{0xff, 0xff, 0xff, 0xff})
        }
        draw2dkit.Circle(gc, float64(leftMargin + i * spacing), float64(topMargin + v.mainFret * spacing - dotSize - 4), dotSize)
        gc.FillStroke()
    }
}

func drawText(gc draw2d.GraphicContext, x int, y int, text string) {
    gc.SetFontSize(fontSize)
	gc.SetFillColor(color.RGBA{0x00, 0x00, 0x00, 0xff})
	gc.FillStringAt(text, float64(x), float64(y))
}

func drawBox(gc draw2d.GraphicContext, stringCount int) {
	gc.MoveTo(leftMargin, topMargin) // should always be called first for a new path
	gc.LineTo(float64(leftMargin + stringCount * spacing), float64(topMargin))
	gc.LineTo(float64(leftMargin + stringCount * spacing), float64(topMargin + fretCount * spacing))
	gc.LineTo(leftMargin, topMargin + fretCount * spacing)
	gc.Close()
	gc.FillStroke()
}

func drawNut(gc draw2d.GraphicContext, stringCount int) {
    gc.MoveTo(float64(leftMargin), float64(topMargin + 4))
    gc.LineTo(float64(leftMargin + spacing * stringCount), float64(topMargin + 4))
    gc.Stroke()
}

func drawFrets(gc draw2d.GraphicContext, stringCount int) {
    for i:=1 ; i < fretCount ; i++ {
        gc.MoveTo(float64(leftMargin), float64(topMargin + i * spacing))
        gc.LineTo(float64(leftMargin + spacing * stringCount), float64(topMargin + i * spacing))
        gc.Stroke()
    }
}

func drawStrings(gc draw2d.GraphicContext, stringCount int) {
    for i:=1 ; i < stringCount ; i++ {
        gc.MoveTo(float64(leftMargin + spacing * i), float64(topMargin))
        gc.LineTo(float64(leftMargin + spacing * i), float64(topMargin + fretCount * spacing))
        gc.Stroke()
    }
}
