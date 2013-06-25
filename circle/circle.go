// Circle Primitive

package circle

import (
	gl "github.com/chsc/gogl/gl33"
	"log"
	"bitbucket.org/jdnewman/argon/shader"
	"unsafe"
)

func init() {
	log.Println("circle.go here")
}

//TODO Make this our default value, so it doubles as default and global for getting offsets and such
var dummyCircle Circle

//TODO Move
var vao, vbo gl.Uint
var gShader shader.Shader

type Circle struct {
	X, Y gl.Float
	R gl.Float
	Red, Green, Blue, Alpha gl.Float
}

func Init() {
	//Setup VAO and VBO
	gl.GenVertexArrays(1, &vao)
	gl.GenBuffers(1, &vbo)

	//-Bind
	gl.BindVertexArray(vao)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)

	//-Attributes
	gl.VertexAttribPointer(0, 2, gl.FLOAT, gl.FALSE, gl.Sizei(unsafe.Sizeof(dummyCircle)), gl.Pointer(unsafe.Offsetof(dummyCircle.X)))
	gl.VertexAttribPointer(1, 1, gl.FLOAT, gl.FALSE, gl.Sizei(unsafe.Sizeof(dummyCircle)), gl.Pointer(unsafe.Offsetof(dummyCircle.R)))
	gl.VertexAttribPointer(2, 4, gl.FLOAT, gl.FALSE, gl.Sizei(unsafe.Sizeof(dummyCircle)), gl.Pointer(unsafe.Offsetof(dummyCircle.Red)))

	gl.EnableVertexAttribArray(0)
	gl.EnableVertexAttribArray(1)
	gl.EnableVertexAttribArray(2)

	//-Unbind
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	gl.BindVertexArray(0)

	//Shader
	gShader.LoadFromFile("./shaders/circle")

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
	temp := new(Circle)

	temp.X = 0.0;
	temp.Y = 0.0;
	temp.R = 1.0;
	temp.Red = 1.0;
	temp.Green = 1.0;
	temp.Blue = 1.0;
	temp.Alpha = 1.0;

	return temp
}

func (this *Circle) Draw() {
	//TODO Avoid unnessessary rebinds
	//Binds
	gShader.Use()
	gl.BindVertexArray(vao)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)

	//Update Buffer
	//gl.BufferData(gl.ARRAY_BUFFER, gl.Sizeiptr(unsafe.Sizeof(dummyCircle)), gl.Pointer(this), gl.DYNAMIC_DRAW)
	gl.BufferData(gl.ARRAY_BUFFER, gl.Sizeiptr(7*4), gl.Pointer(this), gl.DYNAMIC_DRAW)

	//Draw
	gl.DrawArrays(gl.POINTS, 0, gl.Sizei(1))

	//TODO Defer?
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	gl.BindVertexArray(0)
	gl.UseProgram(0)
}

//TODO REM
func makeOrtho(width, height int) [16]gl.Float {
	return [16]gl.Float {
		2.0 / gl.Float(width), 0.0, 0.0, 0.0,
		0.0, 2.0 / gl.Float(height), 0.0, 0.0,
		0.0, 0.0, -1.0, 0.0,
		-1.0, -1.0, 0.0, 1.0}
}
