package main

import (
	"log"
	"runtime"

	"github.com/DualGo/gl/v4.1-core/gl"
	"github.com/DualGo/glfont"
	"github.com/DualGo/glfw/v3.1/glfw"
)

const windowWidth = 800
const windowHeight = 600

func init() {
	runtime.LockOSThread()
}

func main() {

	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.Resizable, glfw.True)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, _ := glfw.CreateWindow(int(windowWidth), int(windowHeight), "glfontExample", nil, nil)

	window.MakeContextCurrent()
	glfw.SwapInterval(1)

	if err := gl.Init(); err != nil {
		panic(err)
	}

	//load font (fontfile, font scale, window width, window height
	font, err := glfont.LoadFont("Robot.ttf", int32(52), windowWidth, windowHeight)
	if err != nil {
		log.Panicf("LoadFont: %v", err)
	}

	gl.ClearColor(0.0, 0.0, 0.0, 0.0)

	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		//set color and draw text
		//font.SetColor(1.0, 1.0, 1.0, 1.0)
		//r,g,b,a font color
		font.Printf(100, 100, 1.0, "Lorem ipsum dolor sit amet, consectetur adipiscing elit.") //x,y,scale,string,printf args

		window.SwapBuffers()
		glfw.PollEvents()

	}
}
