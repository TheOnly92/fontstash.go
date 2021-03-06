package main

import (
	"github.com/TheOnly92/fontstash.go/truetype"
	"github.com/go-gl/gl"
	glfw "github.com/go-gl/glfw3"
	"io/ioutil"
	"runtime"
)

func main() {
	runtime.LockOSThread()

	if !glfw.Init() {
		panic("Can't init glfw!")
	}
	defer glfw.Terminate()

	window, err := glfw.CreateWindow(800, 600, "fontstash example", nil, nil)
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()
	glfw.SwapInterval(1)
	gl.Init()

	data, err := ioutil.ReadFile("ClearSans-Regular.ttf")
	if err != nil {
		panic(err)
	}

	gl.Enable(gl.TEXTURE_2D)

	tmpBitmap := make([]byte, 512*512)
	cdata, err, _, tmpBitmap := truetype.BakeFontBitmap(data, 0, 32, tmpBitmap, 512, 512, 32, 96)
	ftex := gl.GenTexture()
	ftex.Bind(gl.TEXTURE_2D)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.ALPHA, 512, 512, 0, gl.ALPHA, gl.UNSIGNED_BYTE, tmpBitmap)

	gl.ClearColor(0.3, 0.3, 0.32, 1.)

	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		gl.MatrixMode(gl.PROJECTION)
		gl.LoadIdentity()
		gl.Ortho(0, 800, 600, 0, 0, 1)
		gl.MatrixMode(gl.MODELVIEW)
		gl.LoadIdentity()
		gl.Disable(gl.DEPTH_TEST)
		gl.Color4ub(255, 255, 255, 255)
		gl.Enable(gl.BLEND)
		gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

		my_print(100, 100, "The quick brown fox jumps over the fence", ftex, cdata)

		window.SwapBuffers()
		glfw.PollEvents()
	}
}

func my_print(x, y float64, text string, ftex gl.Texture, cdata []*truetype.BakedChar) {
	gl.Enable(gl.TEXTURE_2D)
	ftex.Bind(gl.TEXTURE_2D)
	gl.Begin(gl.QUADS)
	for _, b := range text {
		if int(b) >= 32 && int(b) < 128 {
			var q *truetype.AlignedQuad
			x, q = truetype.GetBakedQuad(cdata, 512, 512, int(b)-32, x, y, true)
			gl.TexCoord2f(q.S0, q.T0)
			gl.Vertex2f(q.X0, q.Y0)
			gl.TexCoord2f(q.S1, q.T0)
			gl.Vertex2f(q.X1, q.Y0)
			gl.TexCoord2f(q.S1, q.T1)
			gl.Vertex2f(q.X1, q.Y1)
			gl.TexCoord2f(q.S0, q.T1)
			gl.Vertex2f(q.X0, q.Y1)
		}
	}
	gl.End()
}
