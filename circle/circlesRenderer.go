
package circle

import (
	"log"
	"unsafe"

	gl "github.com/chsc/gogl/gl33"
)

func init() {
	log.Println("circlesRenderer.go here")
}

func (this *Circles) Draw() {
	gRenderer.Draw(this, gShader)
}

func (this *Circles) DrawData() renderer.DrawData {
	//TODO Get rid of gl dependancies?
	circleArray := (*this).CircleArray()
	return renderer.DrawData{gl.Pointer(&circleArray[0]), gl.Sizeiptr(unsafe.Sizeof(defaultCircle)*uintptr(len(circleArray))), gl.Sizei(len(circleArray))}
}

