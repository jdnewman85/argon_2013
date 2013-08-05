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
	LineRenderer   *argon.Renderer
)

func InitPrimitiveRenderers() {
	circleSources := []string{"./shaders/circle.vert", "./shaders/circle.geom", "./shaders/circle.frag"}
	sizeOfCircle := gl.Sizei(unsafe.Sizeof(argon.DefaultCircle))
	CircleRenderer = argon.CreateRenderer(CircleRenderAttributes(), circleSources, 0, sizeOfCircle)

	lineSources := []string{"./shaders/line.vert", "./shaders/line.geom", "./shaders/line.frag"}
	sizeOfLine := gl.Sizei(unsafe.Sizeof(argon.DefaultLine))
	LineRenderer = argon.CreateRenderer(LineRenderAttributes(), lineSources, 0, sizeOfLine)
}

func CircleRenderAttributes() []argon.Attribute {
	//TODO Compiler error if these intermediate offset values are used directly
	xOffset := gl.Uint(unsafe.Offsetof(argon.DefaultCircle.X))
	rOffset := gl.Uint(unsafe.Offsetof(argon.DefaultCircle.R))
	redOffset := gl.Uint(unsafe.Offsetof(argon.DefaultCircle.Red))
	tempAttributes := []argon.Attribute{
		argon.Attribute{0, 2, gl.FLOAT, gl.FALSE, xOffset},
		argon.Attribute{1, 1, gl.FLOAT, gl.FALSE, rOffset},
		argon.Attribute{2, 4, gl.FLOAT, gl.FALSE, redOffset},
	}

	return tempAttributes
}

func LineRenderAttributes() []argon.Attribute {
	//TODO Compiler error if these intermediate offset values are used directly
	x1Offset := gl.Uint(unsafe.Offsetof(argon.DefaultLine.X1))
	x2Offset := gl.Uint(unsafe.Offsetof(argon.DefaultLine.X2))
	rOffset := gl.Uint(unsafe.Offsetof(argon.DefaultLine.R))
	redOffset := gl.Uint(unsafe.Offsetof(argon.DefaultLine.Red))
	tempAttributes := []argon.Attribute{
		argon.Attribute{0, 2, gl.FLOAT, gl.FALSE, x1Offset},
		argon.Attribute{1, 2, gl.FLOAT, gl.FALSE, x2Offset},
		argon.Attribute{2, 1, gl.FLOAT, gl.FALSE, rOffset},
		argon.Attribute{3, 4, gl.FLOAT, gl.FALSE, redOffset},
	}

	return tempAttributes
}
