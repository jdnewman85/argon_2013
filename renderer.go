package argon

import (
	"log"
	"reflect"

	gl "github.com/chsc/gogl/gl43"
)

func init() {
	log.Println("renderer.go here")
}

type BasicRenderer struct {
	Vao
	Vbo     VertexBuffer
	Program
}

type Renderer interface {
	Draw(interface{})
}

func CreateBasicRenderer(renderAttributes []Attribute, defaultShaderPaths []string, vbOffset gl.Intptr, vbStride gl.Sizei) *BasicRenderer {
	temp := new(BasicRenderer)

	//Setup VAO and VBO
	temp.Vao = GenVao()
	temp.Vbo = VertexBuffer{GenBuffer(), vbOffset, vbStride}

	//-Bind
	temp.Vao.Bind()
	defer temp.Vao.UnBind()
	temp.Vbo.Bind()

	//-Attributes
	temp.Vao.SetAttributes(renderAttributes)

	//Shader Program
	temp.Program, _ = CreateProgramFromFiles(defaultShaderPaths)
	//TODO ERROR on err here!

	temp.Program.Use()
	//-Uniforms //TODO Remove/Move/Change
	glUniformName := gl.GLString("inOrtho")
	defer gl.GLStringFree(glUniformName)
	inOrthoLoc := gl.GetUniformLocation(gl.Uint(temp.Program), glUniformName)
	orthoMat := MakeOrtho(Width, Height)
	gl.UniformMatrix4fv(inOrthoLoc, 1, 0, &orthoMat[0])

	//--Textures //TODO Remove/Move/Change
	texLoc := gl.GetUniformLocation(gl.Uint(temp.Program), gl.GLString("inTexture"))
	gl.Uniform1i(texLoc, 0)

	//TODO Error Handling/Reporting

	return temp
}

func (this *BasicRenderer) Draw(entity interface{}) {
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
	//	this.Render(gl.Pointer(entityPointer), gl.Sizeiptr(entitySize), gl.Sizei(numEntities))

	//Update Buffer
	this.Vbo.Bind()
	defer this.Vbo.UnBind()
	this.Vbo.Data(ArrayBuffer, gl.Sizeiptr(entitySize), gl.Pointer(entityPointer), gl.DYNAMIC_DRAW)

	//Draw
	this.Vao.Draw(gl.Sizei(numEntities), this.Program)
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
