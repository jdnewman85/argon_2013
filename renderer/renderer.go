
package renderer

import(
	"log"
	"reflect"

	gl "github.com/chsc/gogl/gl43"

	"bitbucket.org/jdnewman/argon"
	"bitbucket.org/jdnewman/argon/shader"
)

func init() {
	log.Println("renderer.go here")
}

type Attribute struct {
	Index gl.Uint
	Size gl.Int
	Kind gl.Enum
	Normalized gl.Boolean
	Stride gl.Sizei
	Offset gl.Pointer
}

type RenderData struct {
	ArrayData gl.Pointer
	ArraySize gl.Sizeiptr
	ElementNum gl.Sizei
}

type RendererBase struct {
	vao, vbo gl.Uint
	attributes []Attribute
	defaultShader *shader.Shader
}

type Renderer interface {
	Render(elements RenderData, aShader *shader.Shader)
}

func NewRendererBase(renderAttributes []Attribute, defaultShaderPaths []string) *RendererBase {
	temp := new(RendererBase)
	temp.attributes = renderAttributes


	//Setup VAO and VBO
	gl.GenVertexArrays(1, &temp.vao)
	gl.GenBuffers(1, &temp.vbo)

	//-Bind
	gl.BindVertexArray(temp.vao)
	gl.BindBuffer(gl.ARRAY_BUFFER, temp.vbo)

	//-Attributes
	for _,t := range renderAttributes {
		gl.VertexAttribPointer(t.Index, t.Size, t.Kind, t.Normalized, t.Stride, t.Offset)
		gl.EnableVertexAttribArray(t.Index)
	}

	//-Unbind
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	gl.BindVertexArray(0)


	//Shader
	temp.defaultShader, _ = shader.CreateFromFiles(defaultShaderPaths)
	//TODO ERROR on err here!

	temp.defaultShader.Use()
	//-Uniforms //TODO Remove/Move/Change
	inOrthoLoc := gl.GetUniformLocation(temp.defaultShader.Program, gl.GLString("inOrtho"))
	orthoVec := shader.MakeOrtho(argon.Width, argon.Height)
	gl.UniformMatrix4fv(inOrthoLoc, 1, 0, &orthoVec[0])

	//--Textures //TODO Remove/Move/Change
	texLoc := gl.GetUniformLocation(temp.defaultShader.Program, gl.GLString("inTexture"))
	gl.Uniform1i(texLoc, 0)

	//TODO Error Handling/Reporting


	return temp
}

func (this *RendererBase) Render(elements RenderData, aShader *shader.Shader) {

	//TODO Avoid unnessessary rebinds
	//Binds
	aShader.Use()
	gl.BindVertexArray(this.vao)
	gl.BindBuffer(gl.ARRAY_BUFFER, this.vbo)

	//Update Buffer
	gl.BufferData(gl.ARRAY_BUFFER, elements.ArraySize, elements.ArrayData, gl.DYNAMIC_DRAW)

	//Draw
	gl.DrawArrays(gl.POINTS, 0, elements.ElementNum)

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
		entitySize = reflect.TypeOf(entity).Elem().Size()*uintptr(numEntities)
	case reflect.Array:
		numEntities = reflect.ValueOf(entity).Len()
		tempSlice := reflect.MakeSlice(reflect.SliceOf(reflect.ValueOf(entity).Index(0).Type()), 0, numEntities)
		for i := 0; i < numEntities; i++ {
			tempSlice = reflect.Append(tempSlice, reflect.ValueOf(entity).Index(i))
		}
		entityPointer = tempSlice.Pointer()
		entitySize = reflect.TypeOf(entity).Elem().Size()*uintptr(numEntities)
	case reflect.Ptr:
		entityPointer = reflect.ValueOf(entity).Pointer()
		entitySize = reflect.TypeOf(entity).Elem().Size()*uintptr(numEntities)
	case reflect.Struct:
		tempSlice := reflect.MakeSlice(reflect.SliceOf(reflect.ValueOf(entity).Type()), 0, numEntities)
		tempSlice = reflect.Append(tempSlice, reflect.ValueOf(entity))
		entityPointer = tempSlice.Pointer()
		entitySize = reflect.TypeOf(entity).Size()*uintptr(numEntities)
	default:
		//TODO Better error stuffs
		log.Println("Renderer: Unhandled type: %s", interfaceType.String())
	}
	renderData := RenderData{gl.Pointer(entityPointer), gl.Sizeiptr(entitySize), gl.Sizei(numEntities)}
	this.Render(renderData, this.defaultShader)
}

//TODO Assert on dataSizes and such not matching multiples of stored correct value?
