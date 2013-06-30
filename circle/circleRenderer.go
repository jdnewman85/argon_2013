
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

var(
	defaultRenderer *renderer.RendererBase
	defaultShader *shader.Shader
	activeRenderer *renderer.RendererBase
	activeShader *shader.Shader
)

func DefaultRenderer() *renderer.RendererBase {
	return defaultRenderer
}

func DefaultShader() *shader.Shader {
	return defaultShader
}

func ActiveRenderer() *renderer.RendererBase {
	return activeRenderer
}

func ActiveShader() *shader.Shader {
	return activeShader
}

func setRenderer(r *renderer.RendererBase) {
	activeRenderer = r
}

func setShader(s *shader.Shader) {
	activeShader = s
}

func RendererInit() {
	//RendererBase
	defaultRenderer = renderer.CreateBase(RenderAttributes())
	activeRenderer = defaultRenderer

	//Shader
	defaultShader = shader.Create()
	activeShader = defaultShader
	defaultShader.LoadFromFile(defaultShaderPath)
	defaultShader.Link()

	defaultShader.Use()
	//-Uniforms //TODO Remove/Move/Change
	inOrthoLoc := gl.GetUniformLocation(defaultShader.Program, gl.GLString("inOrtho"))
	orthoVec := shader.MakeOrtho(argon.Width(), argon.Height())
	gl.UniformMatrix4fv(inOrthoLoc, 1, 0, &orthoVec[0])

	//--Textures //TODO Remove/Move/Change
	texLoc := gl.GetUniformLocation(defaultShader.Program, gl.GLString("inTexture"))
	gl.Uniform1i(texLoc, 0)

	//TODO Errors
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

func (this *Circle) RenderData() renderer.RenderData {
	return renderer.RenderData{gl.Pointer(this), gl.Sizeiptr(unsafe.Sizeof(defaultCircle)), 1}
}

func (this *Circle) Draw() {
	activeRenderer.Draw(this, activeShader)
}

//TODO Get rid of gl dependancies?
