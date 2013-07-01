
package renderer

import (
	"log"
	"unsafe"

	gl "github.com/chsc/gogl/gl33"

	"bitbucket.org/jdnewman/argon"
)

func init() {
	log.Println("circleRenderer.go here")
}

type CircleRenderer struct {
	*RendererBase
}

func NewCircleRenderer() *CircleRenderer {
	circleRenderer := new(CircleRenderer)

	circleRenderer.RendererBase = NewRendererBase(CircleRenderAttributes())

	return circleRenderer
}

func CircleRenderAttributes() []Attribute {
	//TODO Compiler error if these intermediate offset values are used directly
	xOffset := gl.Pointer(unsafe.Offsetof(argon.DefaultCircle.X))
	rOffset := gl.Pointer(unsafe.Offsetof(argon.DefaultCircle.R))
	redOffset := gl.Pointer(unsafe.Offsetof(argon.DefaultCircle.Red))
	sizeOfCircle := gl.Sizei(unsafe.Sizeof(argon.DefaultCircle))
	tempAttributes := []Attribute{
		Attribute{0, 2, gl.FLOAT, gl.FALSE, sizeOfCircle, xOffset},
		Attribute{1, 1, gl.FLOAT, gl.FALSE, sizeOfCircle, rOffset},
		Attribute{2, 4, gl.FLOAT, gl.FALSE, sizeOfCircle, redOffset},
	}

	return tempAttributes
}

func (this *CircleRenderer) Draw(circle interface{}) {
	c := circle.(*argon.Circle)
	//If the renderer needed custom renderData, it could use it here
	//renderData := RenderData{gl.Pointer(c), gl.Sizeiptr(unsafe.Sizeof(argon.DefaultCircle)), 1}
	//this.Render(renderData, this.defaultShader)
	this.RendererBase.Draw(c)
}

//TODO Get rid of gl dependancies?
