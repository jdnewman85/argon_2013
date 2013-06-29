// Circle Primitive

package circle

import (
	"log"
	"unsafe"

	gl "github.com/chsc/gogl/gl33"

	"bitbucket.org/jdnewman/argon/renderer"
	"bitbucket.org/jdnewman/argon/shader"
)

func init() {
	log.Println("circle.go here")
}

var defaultCircle Circle = Circle{0.0, 0.0, 1.0, 1.0, 1.0, 1.0, 1.0}

var gRenderer *renderer.Renderer
var gShader *shader.Shader

type Circle struct {
	X, Y gl.Float
	R gl.Float
	Red, Green, Blue, Alpha gl.Float
}

func Init() {
	//Renderer
	//-Attributes
	//TODO Get a compiler error if these intermediate offset values are used directly
	xOffset := gl.Pointer(unsafe.Offsetof(defaultCircle.X))
	rOffset := gl.Pointer(unsafe.Offsetof(defaultCircle.R))
	redOffset := gl.Pointer(unsafe.Offsetof(defaultCircle.Red))
	sizeOfCircle := gl.Sizei(unsafe.Sizeof(defaultCircle))
	tempAttributes := []renderer.Attribute{
		renderer.Attribute{0, 2, gl.FLOAT, gl.FALSE, sizeOfCircle, xOffset},
		renderer.Attribute{1, 1, gl.FLOAT, gl.FALSE, sizeOfCircle, rOffset},
		renderer.Attribute{2, 4, gl.FLOAT, gl.FALSE, sizeOfCircle, redOffset},
	}
	//-Create
	gRenderer = renderer.Create(unsafe.Sizeof(defaultCircle), tempAttributes)

	//Shader
	gShader = shader.Create()
	gShader.LoadFromFile("./shaders/circle")
	gShader.Link()

	gShader.Use()
	//-Uniforms //TODO Remove/Move/Change
	inOrthoLoc := gl.GetUniformLocation(gShader.Program, gl.GLString("inOrtho"))
	orthoVec := makeOrtho(1280, 720) //TODO FINISH Replace with screen width/height
	gl.UniformMatrix4fv(inOrthoLoc, 1, 0, &orthoVec[0])

	//--Textures //TODO Remove/Move/Change
	texLoc := gl.GetUniformLocation(gShader.Program, gl.GLString("inTexture"))
	gl.Uniform1i(texLoc, 0)

	//TODO Errors
}

func Create() *Circle {
	temp := defaultCircle

	return &temp
}

func (this *Circle) Draw() {
	gRenderer.Draw(unsafe.Pointer(this), gShader)
}

//TODO REM
func makeOrtho(width, height int) [16]gl.Float {
	return [16]gl.Float {
		2.0 / gl.Float(width), 0.0, 0.0, 0.0,
		0.0, 2.0 / gl.Float(height), 0.0, 0.0,
		0.0, 0.0, -1.0, 0.0,
		-1.0, -1.0, 0.0, 1.0}
}
