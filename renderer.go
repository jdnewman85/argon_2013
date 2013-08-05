package argon

import (
	"log"
	"reflect"

	gl "github.com/chsc/gogl/gl43"
)

func init() {
	log.Println("renderer.go here")
}

type Renderer struct {
	vao     Vao
	vbo     VertexBuffer
	program Program
}

func CreateRenderer(renderAttributes []Attribute, defaultShaderPaths []string, vbOffset gl.Intptr, vbStride gl.Sizei) *Renderer {
	temp := new(Renderer)

	//Setup VAO and VBO
	temp.vao = GenVao()
	temp.vbo = VertexBuffer{GenBuffer(), vbOffset, vbStride}

	//-Bind
	temp.vao.Bind()
	defer temp.vao.UnBind()
	temp.vbo.Bind()

	//-Attributes
	temp.vao.SetAttributes(renderAttributes)

	//Shader Program
	temp.program, _ = CreateProgramFromFiles(defaultShaderPaths)
	//TODO ERROR on err here!

	temp.program.Use()
	//-Uniforms //TODO Remove/Move/Change
	glUniformName := gl.GLString("inOrtho")
	defer gl.GLStringFree(glUniformName)
	inOrthoLoc := gl.GetUniformLocation(gl.Uint(temp.program), glUniformName)
	orthoMat := MakeOrtho(Width, Height)
	gl.UniformMatrix4fv(inOrthoLoc, 1, 0, &orthoMat[0])

	//--Textures //TODO Remove/Move/Change
	texLoc := gl.GetUniformLocation(gl.Uint(temp.program), gl.GLString("inTexture"))
	gl.Uniform1i(texLoc, 0)

	//TODO Error Handling/Reporting

	return temp
}

//----------------------------------------------------------------------------------This should take a buffer, and managing that buffer should be seperate?
func (this *Renderer) Render(data gl.Pointer, size gl.Sizeiptr, num gl.Sizei) {
	this.program.Use()
	defer this.program.Forgo()
	this.vao.Bind()
	defer this.vao.UnBind()

	//Update Buffer
	this.vbo.Data(ArrayBuffer, size, data, gl.DYNAMIC_DRAW)

	//Draw
	gl.DrawArrays(gl.POINTS, 0, num)
}

func (this *Renderer) Draw(entity interface{}) {
	//TODO: Cleanup this a bit
	var numEntities int = 1
	var entitySize uintptr = 0
	var entityPointer uintptr
	interfaceType := reflect.TypeOf(entity).Kind()
	switch interfaceType {
	case reflect.Slice:
		entityPointer = reflect.ValueOf(entity).Pointer()
		numEntities = reflect.ValueOf(entity).Len()
		entitySize = reflect.TypeOf(entity).Elem().Size() * uintptr(numEntities)
	case reflect.Array:
		numEntities = reflect.ValueOf(entity).Len()
		tempSlice := reflect.MakeSlice(reflect.SliceOf(reflect.ValueOf(entity).Index(0).Type()), 0, numEntities)
		for i := 0; i < numEntities; i++ {
			tempSlice = reflect.Append(tempSlice, reflect.ValueOf(entity).Index(i))
		}
		entityPointer = tempSlice.Pointer()
		entitySize = reflect.TypeOf(entity).Elem().Size() * uintptr(numEntities)
	case reflect.Ptr:
		entityPointer = reflect.ValueOf(entity).Pointer()
		entitySize = reflect.TypeOf(entity).Elem().Size() * uintptr(numEntities)
	case reflect.Struct:
		tempSlice := reflect.MakeSlice(reflect.SliceOf(reflect.ValueOf(entity).Type()), 0, numEntities)
		tempSlice = reflect.Append(tempSlice, reflect.ValueOf(entity))
		entityPointer = tempSlice.Pointer()
		entitySize = reflect.TypeOf(entity).Size() * uintptr(numEntities)
	default:
		//TODO Better error stuffs
		log.Println("Renderer: Unhandled type: %s", interfaceType.String())
	}
	this.Render(gl.Pointer(entityPointer), gl.Sizeiptr(entitySize), gl.Sizei(numEntities))
}

//TEMP?
func (this *Renderer) Program() Program {
	return this.program
}
func (this *Renderer) Vao() gl.Uint {
	return gl.Uint(this.vao)
}

//TODO Assert on dataSizes and such not matching multiples of stored correct value?

//TODO REM
func MakeOrtho(width, height int) [16]gl.Float {
	return [16]gl.Float{
		2.0 / gl.Float(width), 0.0, 0.0, 0.0,
		0.0, 2.0 / gl.Float(height), 0.0, 0.0,
		0.0, 0.0, -1.0, 0.0,
		-1.0, -1.0, 0.0, 1.0}
}
