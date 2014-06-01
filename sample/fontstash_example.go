package main

import (
	"github.com/go-gl/gl"
	glfw "github.com/go-gl/glfw3"
	//"os"
	fontstash "./fontstash.go"
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

	stash := fontstash.Create(512, 512)

	clearSansRegular, err := stash.AddFont("ClearSans-Regular.ttf")
	if err != nil {
		panic(err)
	}

	clearSansItalic, err := stash.AddFont("ClearSans-Italic.ttf")
	if err != nil {
		panic(err)
	}

	clearSansBold, err := stash.AddFont("ClearSans-Bold.ttf")
	if err != nil {
		panic(err)
	}

	droidJapanese, err := stash.AddFont("DroidSansJapanese.ttf")
	if err != nil {
		panic(err)
	}

	gl.ClearColor(0.3, 0.3, 0.32, 1.)

	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		gl.MatrixMode(gl.PROJECTION)
		gl.LoadIdentity()
		gl.Ortho(0, 800, 0, 600, -1, 1)
		gl.MatrixMode(gl.MODELVIEW)
		gl.LoadIdentity()
		gl.Disable(gl.DEPTH_TEST)
		gl.Color4ub(255, 255, 255, 255)
		gl.Enable(gl.BLEND)
		gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

		gl.Disable(gl.TEXTURE_2D)
		gl.Begin(gl.QUADS)
		gl.Vertex2i(0, -5)
		gl.Vertex2i(5, -5)
		gl.Vertex2i(5, -11)
		gl.Vertex2i(0, -11)
		gl.End()

		sx := float64(100)
		sy := float64(250)

		stash.BeginDraw()

		dx := sx
		dy := sy
		dx = stash.DrawText(clearSansRegular, 24, dx, dy, "The quick ")
		dx = stash.DrawText(clearSansItalic, 48, dx, dy, "brown ")
		dx = stash.DrawText(clearSansRegular, 24, dx, dy, "fox ")
		_, _, lh := stash.VMetrics(clearSansItalic, 24)
		dx = sx
		dy -= lh * 1.2
		dx = stash.DrawText(clearSansItalic, 24, dx, dy, "jumps over ")
		dx = stash.DrawText(clearSansBold, 24, dx, dy, "the lazy ")
		dx = stash.DrawText(clearSansRegular, 24, dx, dy, "dog.")
		dx = sx
		dy -= lh * 1.2
		dx = stash.DrawText(clearSansRegular, 12, dx, dy, "Now is the time for all good men to come to the aid of the party.")
		_, _, lh = stash.VMetrics(clearSansItalic, 12)
		dx = sx
		dy -= lh * 1.2 * 2
		dx = stash.DrawText(clearSansItalic, 18, dx, dy, "Ég get etið gler án þess að meiða mig.")
		_, _, lh = stash.VMetrics(clearSansItalic, 18)
		dx = sx
		dy -= lh * 1.2
		stash.DrawText(droidJapanese, 18, dx, dy, "どこかに置き忘れた、サングラスと打ち明け話。")

		stash.EndDraw()
		gl.Enable(gl.DEPTH_TEST)

		window.SwapBuffers()
		glfw.PollEvents()
	}
}
