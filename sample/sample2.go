package main

import (
	"./truetype"
	"fmt"
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

	var c int
	if len(os.Args) > 1 {
		c = int(os.Args[1][0])
	} else {
		c = int('a')
	}

	bitmap, w, h := font.GetCodepointBitmap(0, font.ScaleForPixelHeight(20), c, 0, 0)

	for j := 0; j < h; j++ {
		for i := 0; i < w; i++ {
			fmt.Printf("%s", string(" .:ioVM@"[bitmap[j*w+i]>>5]))
		}
		fmt.Printf("\n")
	}
}
