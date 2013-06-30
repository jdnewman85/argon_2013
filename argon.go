// Argon Graphics Engine

package argon

import (
	"fmt"
	"log"
	"os"

	gl "github.com/chsc/gogl/gl33"
	"github.com/jteeuwen/glfw"
)

type Graphics struct {
	width, height int
}

var(
	//TODO TEMP Remove when ortho uniform becomes external to renderer
	Width, Height int
)

func init() {
	log.Println("argon.go here")
}

func NewGraphics(width, height int, fullscreen bool) (*Graphics, error) {

	//glfw
	if err := glfw.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "glfw: %s\n", err)
		return nil,err
	}

	//-hints
	glfw.OpenWindowHint(glfw.WindowNoResize, 1)

	//Fullscreen param
	glfwFullscreen := glfw.Windowed
	if  fullscreen {
		glfwFullscreen = glfw.Fullscreen
	}

	if err := glfw.OpenWindow(width, height, 0, 0, 0, 0, 0, 0, glfwFullscreen); err != nil {
		fmt.Fprintf(os.Stderr, "glfw: %s\n", err)
		return nil,err
	}

	glfw.SetSwapInterval(1)

	//opengl
	if err := gl.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "gl: %s\n", err)
		return nil, err
	}

	//-initial state
	gl.Viewport(0, 0, gl.Sizei(width), gl.Sizei(height))
	gl.ClearColor(0.4, 0.4, 0.4, 1.0)
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	//-log version
	glMajor, glMinor, glRev := glfw.GLVersion()
	log.Printf("%d %d %d", glMajor, glMinor, glRev)
	log.Println(gl.GoStringUb(gl.GetString(gl.VERSION)))

	//TODO TEMP Remove when ortho uniform becomes external to renderer
	Width, Height = width, height
	return &Graphics{width: width, height: height}, nil
}

func (this *Graphics) Destroy() {
	//TODO TEMP Remove when ortho uniform becomes external to renderer
	Width, Height = 0, 0
	this.width, this.height = 0, 0

	//TODO Free all images, and stuffs?
	glfw.CloseWindow()
	glfw.Terminate()
}

func (this *Graphics) Cls() {
	//TODO Other buffers
	gl.Clear(gl.COLOR_BUFFER_BIT)
}

func (this *Graphics) Flip() {
	//TODO Timing management and vsync?
	glfw.SwapBuffers()
}

func (this *Graphics) Open() bool {
	if glfw.WindowParam(glfw.Opened) == gl.TRUE {
		return true
	}
	return false
}

func (this *Graphics) Width() int {
	return this.width
}

func (this *Graphics) Height() int {
	return this.height
}
