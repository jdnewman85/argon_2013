
package renderer

import(
	"log"
	"reflect"

	gl "github.com/chsc/gogl/gl33"

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
	entityPointer := reflect.ValueOf(entity).Pointer()
	entitySize := reflect.TypeOf(entity).Elem().Size()
	renderData := RenderData{gl.Pointer(entityPointer), gl.Sizeiptr(entitySize), 1}
	this.Render(renderData, this.defaultShader)
}

//TODO Assert on dataSizes and such not matching multiples of stored correct value?
