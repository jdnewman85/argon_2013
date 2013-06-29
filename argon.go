// Argon Graphics Engine

package argon

import (
	"fmt"
	"log"
	"os"

	gl "github.com/chsc/gogl/gl33"
	"github.com/jteeuwen/glfw"
)

func init() {
	log.Println("argon.go here")
}

func Graphics(aWidth, aHeight int, aFullscreen bool) error {

	//glfw
	if err := glfw.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "glfw: %s\n", err)
		return err
	}

	//-hints
	glfw.OpenWindowHint(glfw.WindowNoResize, 1)

	//Fullscreen param
	fullscreen := glfw.Windowed
	if  aFullscreen {
		fullscreen = glfw.Fullscreen
	}

	if err := glfw.OpenWindow(aWidth, aHeight, 0, 0, 0, 0, 0, 0, fullscreen); err != nil {
		fmt.Fprintf(os.Stderr, "glfw: %s\n", err)
		return err
	}

	glfw.SetSwapInterval(1)

	//opengl
	if err := gl.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "gl: %s\n", err)
		return err
	}

	//-initial state
	gl.Viewport(0, 0, gl.Sizei(aWidth), gl.Sizei(aHeight))
	gl.ClearColor(0.4, 0.4, 0.4, 1.0)
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	//-log version
	glMajor, glMinor, glRev := glfw.GLVersion()
	log.Printf("%d %d %d", glMajor, glMinor, glRev)
	log.Println(gl.GoStringUb(gl.GetString(gl.VERSION)))

	return nil
}

func EndGraphics() {
	//TODO Free all images, and stuffs?
	glfw.CloseWindow()
	glfw.Terminate()
}

func Cls() {
	//TODO Other buffers
	gl.Clear(gl.COLOR_BUFFER_BIT)
}

func Flip() {
	//TODO Timing management and vsync?
	glfw.SwapBuffers()
}

func WindowOpen() bool {
	if glfw.WindowParam(glfw.Opened) == gl.TRUE {
		return true
	}
	return false
}

