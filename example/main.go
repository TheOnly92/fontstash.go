package main

import (
	"github.com/TheOnly92/fontstash.go/truetype"
	"fmt"
	"io/ioutil"
	"math"
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

	text := "Heljo World!"

	scale := font.ScaleForPixelHeight(15)
	ascent, _, _ := font.GetFontVMetrics()
	baseline := int(float64(ascent) * scale)

	var screen [1580]byte

	var xpos float64
	for ch, b := range text {
		xShift := xpos - math.Floor(xpos)
		advance, _ := font.GetCodepointHMetrics(int(b))
		x0, y0, x1, y1 := font.GetCodepointBitmapBoxSubpixel(int(b), scale, scale, xShift, 0)
		tmp := font.MakeCodepointBitmapSubpixel(screen[(baseline+y0)*79+int(xpos)+x0:], x1-x0, y1-y0, 79, scale, scale, xShift, 0, int(b))
		copy(screen[(baseline+y0)*79+int(xpos)+x0:], tmp)
		xpos += (float64(advance) * scale)
		if len(text) > ch+1 {
			xpos += scale * float64(font.GetCodepointKernAdvance(int(b), int(text[ch+1])))
		}
	}

	for j := 0; j < 20; j++ {
		for i := 0; i < 78; i++ {
			fmt.Printf("%s", string(" .:ioVM@"[screen[j*79+i]>>5]))
		}
		fmt.Printf("\n")
	}
}
