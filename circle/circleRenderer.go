
package circle

import (
	"log"
	"unsafe"

	gl "github.com/chsc/gogl/gl33"

	"bitbucket.org/jdnewman/argon"
	"bitbucket.org/jdnewman/argon/renderer"
	"bitbucket.org/jdnewman/argon/shader"
)

func init() {
	log.Println("circleRenderer.go here")
}

const(
	defaultShaderPath = "./shaders/circle"
)

type Renderer struct {
	*renderer.RendererBase
	defaultShader *shader.Shader
}

func NewRenderer() *Renderer {
	temp := new(Renderer)
	//RendererBase
	temp.RendererBase = renderer.CreateBase(RenderAttributes())

	//Shader
	temp.defaultShader = shader.Create()
	temp.defaultShader.LoadFromFile(defaultShaderPath)
	temp.defaultShader.Link()

	temp.defaultShader.Use()
	//-Uniforms //TODO Remove/Move/Change
	inOrthoLoc := gl.GetUniformLocation(temp.defaultShader.Program, gl.GLString("inOrtho"))
	orthoVec := shader.MakeOrtho(argon.Width, argon.Height)
	gl.UniformMatrix4fv(inOrthoLoc, 1, 0, &orthoVec[0])

	//--Textures //TODO Remove/Move/Change
	texLoc := gl.GetUniformLocation(temp.defaultShader.Program, gl.GLString("inTexture"))
	gl.Uniform1i(texLoc, 0)

	//TODO Errors

	return temp
}

func RenderAttributes() []renderer.Attribute {
	//TODO Compiler error if these intermediate offset values are used directly
	xOffset := gl.Pointer(unsafe.Offsetof(defaultCircle.X))
	rOffset := gl.Pointer(unsafe.Offsetof(defaultCircle.R))
	redOffset := gl.Pointer(unsafe.Offsetof(defaultCircle.Red))
	sizeOfCircle := gl.Sizei(unsafe.Sizeof(defaultCircle))
	tempAttributes := []renderer.Attribute{
		renderer.Attribute{0, 2, gl.FLOAT, gl.FALSE, sizeOfCircle, xOffset},
		renderer.Attribute{1, 1, gl.FLOAT, gl.FALSE, sizeOfCircle, rOffset},
		renderer.Attribute{2, 4, gl.FLOAT, gl.FALSE, sizeOfCircle, redOffset},
	}

	return tempAttributes
}

func (this *Renderer) Draw(circle interface{}) {
	//TODO May be possible to get the pointer, size and such without the assert?
	c := circle.(*Circle)
	renderData := renderer.RenderData{gl.Pointer(c), gl.Sizeiptr(unsafe.Sizeof(defaultCircle)), 1}
	this.Render(renderData, this.defaultShader)
}

//TODO Get rid of gl dependancies?
