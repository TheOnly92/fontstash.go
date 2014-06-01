package main

import (
	"./truetype"
	"image"
	"image/png"
	"io/ioutil"
	"os"
)

func main() {
	data, err := ioutil.ReadFile("ClearSans-Regular.ttf")
	if err != nil {
		panic(err)
	}

	font, err := truetype.InitFont(data, truetype.GetFontOffsetForIndex(data, 0))
	if err != nil {
		panic(err)
	}

	bW := 512 // bitmap width
	bH := 128 // bitmap height
	lH := 64  // line height

	scale := font.ScaleForPixelHeight(float64(lH))

	word := "how are you?"

	ascent, descent, _ := font.GetFontVMetrics()

	ascent = int(float64(ascent) * scale)
	descent = int(float64(descent) * scale)

	bitmap := make([]byte, bW*bH)

	x := 0
	for i, b := range word {
		cX1, cY1, cX2, cY2 := font.GetCodepointBitmapBox(int(b), scale, scale)

		y := ascent + cY1

		byteOffset := x + (y * bW)
		tmp := font.MakeCodepointBitmap(bitmap[byteOffset:], cX2-cX1, cY2-cY1, bW, scale, scale, int(b))
		copy(bitmap[byteOffset:], tmp)

		ax, _ := font.GetCodepointHMetrics(int(b))
		x += int(float64(ax) * scale)

		if len(word) > i+1 {
			kern := font.GetCodepointKernAdvance(int(b), int(word[i+1]))
			x += int(float64(kern) * scale)
		}
	}

	r := image.Rect(0, 0, bW, bH)
	g := image.NewGray(r)
	g.Pix = []uint8(bitmap)

	file, err := os.Create("test.png")
	if err != nil {
		panic(err)
	}
	png.Encode(file, g)
}
