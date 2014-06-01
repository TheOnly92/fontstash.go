package main

import (
    "fmt"
    "github.com/TheOnly92/fontstash.go/truetype"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("ClearSans-Regular.ttf")
	if err != nil {
		panic(err)
	}

	tmpBitmap := make([]byte, 512*512)
	cdata, err, _, tmpBitmap := truetype.BakeFontBitmap(data, 0, 32, tmpBitmap, 512, 512, 32, 96)
	var x, y float64
	b := 'b'
	x, q := truetype.GetBakedQuad(cdata, 512, 512, int(b)-32, x, y, true)

	fmt.Println(q)
}
