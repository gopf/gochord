package main

import (
    "os"
    "strconv"
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
const fontSize = 24
const dotSize = 6

func main() {
    loadFont()

    for _, arg := range os.Args[1:] {
        chord := parseChord(arg)
        stringCount, fretCount, fretOffset := getFretInfo(chord)

        dest := image.NewRGBA(image.Rect(0, 0, leftMargin + rightMargin + spacing * stringCount, topMargin + bottomMargin + fretCount * spacing))
        gc := draw2dimg.NewGraphicContext(dest)

        drawChord(gc, chord, stringCount, fretCount, fretOffset)

        draw2dimg.SaveToPngFile(chord.name + ".png", dest)
    }
}

func loadFont() {
    font, err := truetype.Parse(goregular.TTF)
	if err == nil {
        draw2d.RegisterFont(draw2d.FontData{Name: "luxi"}, font)
    }
}

func setFillColor(gc draw2d.GraphicContext, isRoot bool, defaultColor uint8) {
    c := defaultColor
    if isRoot {
        c = 0
    }
    gc.SetFillColor(color.RGBA{c, c, c, 0xff})
}

func drawChord(gc draw2d.GraphicContext, chord chordInfo, stringCount int, fretCount int, fretOffset int) {
	gc.SetFillColor(color.RGBA{0xff, 0xff, 0xff, 0xff})
	gc.SetStrokeColor(color.RGBA{0x00, 0x00, 0x00, 0xff})
	gc.SetLineWidth(2)

    drawBox(gc, stringCount, fretCount)
    if fretOffset > 0 {
        drawText(gc, fontSize * 2 / 3, leftMargin + stringCount * spacing + dotSize + 4, topMargin + spacing + fontSize * 1 / 3, strconv.Itoa(fretOffset + 1))
    } else {
        drawNut(gc, stringCount)
    }

    drawFrets(gc, stringCount, fretCount)
    drawStrings(gc, stringCount, fretCount)
    drawText(gc, fontSize, leftMargin, 5 + fontSize, chord.name)

    for i, v := range chord.strings {
        x := float64(leftMargin + i * spacing)
        for _, fret := range v.frets {
            if fret == 0 {
                setFillColor(gc, v.rootFret == fret, 0xff)
                draw2dkit.Circle(gc, x, float64(topMargin - dotSize - 4), dotSize)
                gc.FillStroke()
            } else if fret > 0 {
                setFillColor(gc, v.rootFret == fret, 0xc0)
                draw2dkit.Circle(gc, x, float64(topMargin + (fret - fretOffset) * spacing - dotSize - 4), dotSize)
                gc.FillStroke()
            } else {
                delta := dotSize * 0.7
                gc.MoveTo(x - delta, topMargin - 4)
                gc.LineTo(float64(x + delta), float64(topMargin - 4 - 2 * delta))
                gc.Stroke()
                gc.MoveTo(x + delta, topMargin - 4)
                gc.LineTo(float64(x - delta), float64(topMargin - 4 - 2 * delta))
                gc.Stroke()
            }
        }
    }
}

func getFretInfo(chord chordInfo) (int, int, int) {
    var min, max int = 999, -1
    stringCount := len(chord.strings) - 1
    for _, v := range chord.strings {
        for _, fret := range v.frets {
            if fret >= 0 {
                if fret > max {
                    max = fret
                }
                if fret < min {
                    min = fret
                }
            }
        }
    }
    if max - min >= 4 {
        if (min > 0) {
            return stringCount, max - min + 1, min - 1
        } else {
            return stringCount, max - min, min
        }
    } else {
        return stringCount, 4, min - 1
    }
}

func drawText(gc draw2d.GraphicContext, fontSize int, x int, y int, text string) {
    gc.SetFontSize(float64(fontSize))
	gc.SetFillColor(color.RGBA{0x00, 0x00, 0x00, 0xff})
	gc.FillStringAt(text, float64(x), float64(y))
}

func drawBox(gc draw2d.GraphicContext, stringCount int, fretCount int) {
	gc.MoveTo(leftMargin, topMargin)
	gc.LineTo(float64(leftMargin + stringCount * spacing), float64(topMargin))
	gc.LineTo(float64(leftMargin + stringCount * spacing), float64(topMargin + fretCount * spacing))
	gc.LineTo(leftMargin, float64(topMargin + fretCount * spacing))
	gc.Close()
	gc.FillStroke()
}

func drawNut(gc draw2d.GraphicContext, stringCount int) {
    gc.MoveTo(float64(leftMargin), float64(topMargin + 4))
    gc.LineTo(float64(leftMargin + spacing * stringCount), float64(topMargin + 4))
    gc.Stroke()
}

func drawFrets(gc draw2d.GraphicContext, stringCount int, fretCount int) {
    for i:=1 ; i < fretCount ; i++ {
        gc.MoveTo(float64(leftMargin), float64(topMargin + i * spacing))
        gc.LineTo(float64(leftMargin + spacing * stringCount), float64(topMargin + i * spacing))
        gc.Stroke()
    }
}

func drawStrings(gc draw2d.GraphicContext, stringCount int, fretCount int) {
    for i:=1 ; i < stringCount ; i++ {
        gc.MoveTo(float64(leftMargin + spacing * i), float64(topMargin))
        gc.LineTo(float64(leftMargin + spacing * i), float64(topMargin + fretCount * spacing))
        gc.Stroke()
    }
}
