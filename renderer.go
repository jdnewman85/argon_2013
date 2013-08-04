package argon

import (
	"log"
	"reflect"

	gl "github.com/chsc/gogl/gl43"
)

func init() {
	log.Println("renderer.go here")
}

type Attribute struct {
	Index      gl.Uint
	Size       gl.Int
	Kind       gl.Enum
	Normalized gl.Boolean
	Stride		gl.Sizei
	Offset     gl.Pointer
}

type Buffer struct {
	bo gl.Uint //Actual gl buffer object
}

type VertexBuffer struct {
	Buffer
	Offset gl.Intptr
	Stride gl.Intptr
}

type RenderData struct {
	ArrayData  gl.Pointer
	ArraySize  gl.Sizeiptr
	ElementNum gl.Sizei
}

type RendererBase struct {
	vao      gl.Uint
	vbo	VertexBuffer
	attributes    []Attribute
	defaultProgram Program
}

func NewRendererBase(renderAttributes []Attribute, defaultShaderPaths []string) *RendererBase {
	temp := new(RendererBase)
	temp.attributes = renderAttributes

	//Setup VAO and VBO
	gl.GenVertexArrays(1, &temp.vao)
	gl.GenBuffers(1, &temp.vbo.bo)

	//-Bind
	gl.BindVertexArray(temp.vao)
	gl.BindBuffer(gl.ARRAY_BUFFER, temp.vbo.bo)

	//-Attributes
	for _, t := range renderAttributes {
		gl.VertexAttribPointer(t.Index, t.Size, t.Kind, t.Normalized, t.Stride, t.Offset)
		gl.EnableVertexAttribArray(t.Index)
	}

	//-Unbind
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	gl.BindVertexArray(0)

	//Shader Program
	temp.defaultProgram, _ = CreateProgramFromFiles(defaultShaderPaths)
	//TODO ERROR on err here!

	temp.defaultProgram.Use()
	//-Uniforms //TODO Remove/Move/Change
	glUniformName := gl.GLString("inOrtho")
	defer gl.GLStringFree(glUniformName)
	inOrthoLoc := gl.GetUniformLocation(gl.Uint(temp.defaultProgram), glUniformName)
	orthoMat := MakeOrtho(Width, Height)
	gl.UniformMatrix4fv(inOrthoLoc, 1, 0, &orthoMat[0])

	//--Textures //TODO Remove/Move/Change
	texLoc := gl.GetUniformLocation(gl.Uint(temp.defaultProgram), gl.GLString("inTexture"))
	gl.Uniform1i(texLoc, 0)

	//TODO Error Handling/Reporting

	return temp
}

func (this *RendererBase) Render(elements RenderData, program Program) {

	//TODO Avoid unnessessary rebinds
	//Binds
	program.Use()
	gl.BindVertexArray(this.vao)
	gl.BindBuffer(gl.ARRAY_BUFFER, this.vbo.bo)

	//Update Buffer
	gl.BufferData(gl.ARRAY_BUFFER, elements.ArraySize, elements.ArrayData, gl.DYNAMIC_DRAW)

	//Draw
	gl.DrawArrays(gl.POINTS, 0, elements.ElementNum)

	//TODO Defer?
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	gl.BindVertexArray(0)
	gl.UseProgram(0)
}

//TEMP?
func (this *RendererBase) RenderBuffer(buffer gl.Uint) {
	//TODO Avoid unnessessary rebinds
	//Binds
	this.defaultProgram.Use()
	var tempVAO gl.Uint
	gl.GenVertexArrays(1, &tempVAO)
	gl.BindVertexArray(tempVAO)
	gl.BindBuffer(gl.ARRAY_BUFFER, buffer)

	//Update Buffer
	//	gl.BufferData(gl.ARRAY_BUFFER, elements.ArraySize, elements.ArrayData, gl.DYNAMIC_DRAW)

	//-Attributes
	for _, t := range this.attributes {
		gl.VertexAttribPointer(t.Index, t.Size, t.Kind, t.Normalized, t.Stride, t.Offset)
		gl.EnableVertexAttribArray(t.Index)
	}

	//Draw
	gl.DrawArrays(gl.POINTS, 0, 512*512)

	//TODO Defer?
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	gl.BindVertexArray(0)
	gl.UseProgram(0)
}

func (this *RendererBase) Draw(entity interface{}) {
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
	renderData := RenderData{gl.Pointer(entityPointer), gl.Sizeiptr(entitySize), gl.Sizei(numEntities)}
	this.Render(renderData, this.defaultProgram)
}

//TEMP?
func (this *RendererBase) DefaultProgram() Program {
	return this.defaultProgram
}
func (this *RendererBase) Vao() gl.Uint {
	return this.vao
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
