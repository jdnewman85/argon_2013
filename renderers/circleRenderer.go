package renderers

import (
	"log"
	"unsafe"

	gl "github.com/chsc/gogl/gl43"

	"bitbucket.org/jdnewman/argon"
)

func init() {
	log.Println("circleRenderer.go here")
}

type CircleRenderer struct {
	*argon.RendererBase
}

func NewCircleRenderer() *CircleRenderer {
	//Load Default Shaders
	sources := []string{"./shaders/circle.vert", "./shaders/circle.geom", "./shaders/circle.frag"}
	rendererBase := argon.NewRendererBase(CircleRenderAttributes(), sources)

	return &CircleRenderer{rendererBase}
}

func CircleRenderAttributes() []argon.Attribute {
	//TODO Compiler error if these intermediate offset values are used directly
	xOffset := gl.Pointer(unsafe.Offsetof(argon.DefaultCircle.X))
	rOffset := gl.Pointer(unsafe.Offsetof(argon.DefaultCircle.R))
	redOffset := gl.Pointer(unsafe.Offsetof(argon.DefaultCircle.Red))
	sizeOfElement := gl.Sizei(unsafe.Sizeof(argon.DefaultCircle))
	tempAttributes := []argon.Attribute{
		argon.Attribute{0, 2, gl.FLOAT, gl.FALSE, sizeOfElement, xOffset},
		argon.Attribute{1, 1, gl.FLOAT, gl.FALSE, sizeOfElement, rOffset},
		argon.Attribute{2, 4, gl.FLOAT, gl.FALSE, sizeOfElement, redOffset},
	}

	return tempAttributes
}

//func (this *CircleRenderer) Draw(circle interface{}) {
//c := circle.(*argon.Circle)
//If the renderer needed custom renderData, it could use it here
//renderData := RenderData{gl.Pointer(c), gl.Sizeiptr(unsafe.Sizeof(argon.DefaultCircle)), 1}
//this.Render(renderData, this.defaultShader)
//	this.RendererBase.Draw(circle)
//}

//TODO Get rid of gl dependancies?
