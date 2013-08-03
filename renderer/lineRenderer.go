
package renderer

import (
	"log"
	"unsafe"

	gl "github.com/chsc/gogl/gl33"

	"bitbucket.org/jdnewman/argon"
)

func init() {
	log.Println("lineRenderer.go here")
}

type LineRenderer struct {
	*RendererBase
}

func NewLineRenderer() *LineRenderer {
	//Load Default Shaders
	sources := []string{"./shaders/line.vert", "./shaders/line.geom", "./shaders/line.frag"}
	rendererBase := NewRendererBase(LineRenderAttributes(), sources)

	return &LineRenderer{rendererBase}
}

func LineRenderAttributes() []Attribute {
	//TODO Compiler error if these intermediate offset values are used directly
	x1Offset := gl.Pointer(unsafe.Offsetof(argon.DefaultLine.X1))
	x2Offset := gl.Pointer(unsafe.Offsetof(argon.DefaultLine.X2))
	rOffset := gl.Pointer(unsafe.Offsetof(argon.DefaultLine.R))
	redOffset := gl.Pointer(unsafe.Offsetof(argon.DefaultLine.Red))
	sizeOfElement := gl.Sizei(unsafe.Sizeof(argon.DefaultLine))
	tempAttributes := []Attribute{
		Attribute{0, 2, gl.FLOAT, gl.FALSE, sizeOfElement, x1Offset},
		Attribute{1, 2, gl.FLOAT, gl.FALSE, sizeOfElement, x2Offset},
		Attribute{2, 1, gl.FLOAT, gl.FALSE, sizeOfElement, rOffset},
		Attribute{3, 4, gl.FLOAT, gl.FALSE, sizeOfElement, redOffset},
	}

	return tempAttributes
}

//func (this *LineRenderer) Draw(line interface{}) {
	//l := line.(*argon.Line)
	//If the renderer needed custom renderData, it could use it here
	//renderData := RenderData{gl.Pointer(c), gl.Sizeiptr(unsafe.Sizeof(argon.DefaultLine)), 1}
	//this.Render(renderData, this.defaultShader)
//	this.RendererBase.Draw(line)
//}

//TODO Get rid of gl dependancies?
