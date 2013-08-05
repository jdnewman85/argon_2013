package argon

import (
	"log"
	"reflect"

	gl "github.com/chsc/gogl/gl43"
)

func init() {
	log.Println("renderer.go here")
}

type RenderData struct {
	ArrayData  gl.Pointer
	ArraySize  gl.Sizeiptr
	ElementNum gl.Sizei
}

type Renderer struct {
	vao    Vao
	vbo VertexBuffer
	attributes    []Attribute
	defaultProgram Program
}

//------------------------------------------------------------------------------------------The majority of this is VAO/Attribute setup, which should be seperated, switched to new system
//------------------------------------------------------------------------------------------The remainder can maybe be summed up in a renderable interface or something?
func CreateRenderer(renderAttributes []Attribute, defaultShaderPaths []string, vbOffset gl.Intptr, vbStride gl.Sizei) *Renderer {
	temp := new(Renderer)
	temp.attributes = renderAttributes

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

//----------------------------------------------------------------------------------This should take a buffer, and managing that buffer should be seperate?
func (this *Renderer) Render(elements RenderData, program Program) {

	//TODO Avoid unnessessary rebinds
	//Binds
	program.Use()
	defer program.Forgo()
	this.vao.Bind()
	defer this.vao.UnBind()

	//Update Buffer
	this.vbo.Data(ArrayBuffer, elements.ArraySize, elements.ArrayData, gl.DYNAMIC_DRAW)

	//Draw
	gl.DrawArrays(gl.POINTS, 0, elements.ElementNum)
}

//TEMP?
func (this *Renderer) RenderBuffer(buffer Buffer) {
	//Shader
	this.defaultProgram.Use()
	defer this.defaultProgram.Forgo()
	tempVAO := GenVao()

	//Bind vao
	tempVAO.Bind()
	defer tempVAO.UnBind()

	//Attributes
	tempVAO.SetAttributes(this.attributes)

	//Draw
	gl.DrawArrays(gl.POINTS, 0, 512*512)

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
	renderData := RenderData{gl.Pointer(entityPointer), gl.Sizeiptr(entitySize), gl.Sizei(numEntities)}
	this.Render(renderData, this.defaultProgram)
}

//TEMP?
func (this *Renderer) DefaultProgram() Program {
	return this.defaultProgram
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
