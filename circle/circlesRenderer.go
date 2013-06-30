
package circle

import (
	"log"
	"unsafe"

	gl "github.com/chsc/gogl/gl33"

	"bitbucket.org/jdnewman/argon/renderer"
)

func init() {
	log.Println("circlesRenderer.go here")
}

func (this *Circles) Draw() {
	activeRenderer.Draw(this, activeShader)
}

func (this *Circles) RenderData() renderer.RenderData {
	//TODO Get rid of gl dependancies?
	circleArray := (*this).CircleArray()
	return renderer.RenderData{gl.Pointer(&circleArray[0]), gl.Sizeiptr(unsafe.Sizeof(defaultCircle)*uintptr(len(circleArray))), gl.Sizei(len(circleArray))}
}

