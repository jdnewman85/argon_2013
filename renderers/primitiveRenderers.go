package renderers

import (
	"log"
	"unsafe"

	gl "github.com/chsc/gogl/gl43"

	"bitbucket.org/jdnewman/argon"
)

func init() {
	log.Println("renderers.go here")
}

var (
	CircleRenderer *argon.Renderer
	LineRenderer *argon.Renderer
)

func InitPrimitiveRenderers() {
	circleSources := []string{"./shaders/circle.vert", "./shaders/circle.geom", "./shaders/circle.frag"}
	CircleRenderer = argon.CreateRenderer(CircleRenderAttributes(), circleSources)

	lineSources := []string{"./shaders/line.vert", "./shaders/line.geom", "./shaders/line.frag"}
	LineRenderer = argon.CreateRenderer(LineRenderAttributes(), lineSources)
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

func LineRenderAttributes() []argon.Attribute {
	//TODO Compiler error if these intermediate offset values are used directly
	x1Offset := gl.Pointer(unsafe.Offsetof(argon.DefaultLine.X1))
	x2Offset := gl.Pointer(unsafe.Offsetof(argon.DefaultLine.X2))
	rOffset := gl.Pointer(unsafe.Offsetof(argon.DefaultLine.R))
	redOffset := gl.Pointer(unsafe.Offsetof(argon.DefaultLine.Red))
	sizeOfElement := gl.Sizei(unsafe.Sizeof(argon.DefaultLine))
	tempAttributes := []argon.Attribute{
		argon.Attribute{0, 2, gl.FLOAT, gl.FALSE, sizeOfElement, x1Offset},
		argon.Attribute{1, 2, gl.FLOAT, gl.FALSE, sizeOfElement, x2Offset},
		argon.Attribute{2, 1, gl.FLOAT, gl.FALSE, sizeOfElement, rOffset},
		argon.Attribute{3, 4, gl.FLOAT, gl.FALSE, sizeOfElement, redOffset},
	}

	return tempAttributes
}

