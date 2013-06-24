// Circle Primitive

package argon

import (
	gl "github.com/chsc/gogl/gl33"
	"log"
	"unsafe"
)

func init() {
	log.Println("circle.go here")
}

//TODO Must be a more elegant way
//TODO Possibly by completely seperating the primitives from their materials?
var circleMaterial Material
var dummyCircle Circle

type Circle struct {
	//TODO Native types instead?
	X, Y gl.Float
	R gl.Float
	Red, Green, Blue, Alpha gl.Float
}

type Circles []Circle

func CircleInit() {
	//Material
	circleMaterial.LoadShaders("./shader/circle")
	circleMaterial.Bind()

	//Attributes
	gl.VertexAttribPointer(0, 2, gl.FLOAT, gl.FALSE, gl.Sizei(unsafe.Sizeof(dummyCircle)), gl.Pointer(unsafe.Offsetof(dummyCircle.X)))
	gl.VertexAttribPointer(1, 1, gl.FLOAT, gl.FALSE, gl.Sizei(unsafe.Sizeof(dummyCircle)), gl.Pointer(unsafe.Offsetof(dummyCircle.R)))
	gl.VertexAttribPointer(2, 4, gl.FLOAT, gl.FALSE, gl.Sizei(unsafe.Sizeof(dummyCircle)), gl.Pointer(unsafe.Offsetof(dummyCircle.Red)))

	gl.EnableVertexAttribArray(0)
	gl.EnableVertexAttribArray(1)
	gl.EnableVertexAttribArray(2)

	//Uniforms //TODO May Remove/Move/Change
	inOrthoLoc := gl.GetUniformLocation(circleMaterial.shaderProgram, gl.GLString("inOrtho"))
	orthoVec := makeOrtho(1280, 720) //TODO FINISH Replace with screen width/height
	gl.UniformMatrix4fv(inOrthoLoc, 1, 0, &orthoVec[0])

	//-Textures //TODO May Remove/Move/Change
	texLoc := gl.GetUniformLocation(circleMaterial.shaderProgram, gl.GLString("inTexture"))
	gl.Uniform1i(texLoc, 0)

	circleMaterial.UnBind()
	//TODO Errors

}

func CreateCircle(num int) Circles {
	temp := make([]Circle, num)

	for i := range temp {
		temp[i].X = 0.0;
		temp[i].Y = 0.0;
		temp[i].R = 1.0;
		temp[i].Red = 1.0;
		temp[i].Green = 1.0;
		temp[i].Blue = 1.0;
		temp[i].Alpha = 1.0;
	}

	return temp
}

func (this Circles) Draw() {
	//TODO Avoid unnessessary rebinds
	circleMaterial.Draw(len(this), int(unsafe.Sizeof(dummyCircle)), gl.Pointer(&this[0]))
}

//TODO REM
func makeOrtho(width, height int) [16]gl.Float {
	return [16]gl.Float {
		2.0 / gl.Float(width), 0.0, 0.0, 0.0,
		0.0, 2.0 / gl.Float(height), 0.0, 0.0,
		0.0, 0.0, -1.0, 0.0,
		-1.0, -1.0, 0.0, 1.0}
}
