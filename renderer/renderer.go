
package renderer

import(
	"log"
	//"unsafe"

	gl "github.com/chsc/gogl/gl33"

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

type Renderable interface {
	//TODO Include getting of Renderer, default shader?
	RenderData() RenderData
}

type RendererBase struct {
	vao, vbo gl.Uint
	attributes []Attribute
}

func CreateBase(attributes []Attribute) *RendererBase {
	temp := new(RendererBase)

	temp.attributes = attributes

	//Setup VAO and VBO
	gl.GenVertexArrays(1, &temp.vao)
	gl.GenBuffers(1, &temp.vbo)

	//-Bind
	gl.BindVertexArray(temp.vao)
	gl.BindBuffer(gl.ARRAY_BUFFER, temp.vbo)

	//-Attributes
	for _,t := range attributes {
		gl.VertexAttribPointer(t.Index, t.Size, t.Kind, t.Normalized, t.Stride, t.Offset)
		gl.EnableVertexAttribArray(t.Index)
	}

	//-Unbind
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	gl.BindVertexArray(0)

	return temp
}

func (this *RendererBase) Draw(elements Renderable, aShader *shader.Shader) {

	//Get Drawable data
	drawData := elements.RenderData()

	//TODO Avoid unnessessary rebinds
	//Binds
	aShader.Use()
	gl.BindVertexArray(this.vao)
	gl.BindBuffer(gl.ARRAY_BUFFER, this.vbo)

	//Update Buffer
	gl.BufferData(gl.ARRAY_BUFFER, drawData.ArraySize, drawData.ArrayData, gl.DYNAMIC_DRAW)

	//Draw
	gl.DrawArrays(gl.POINTS, 0, drawData.ElementNum)

	//TODO Defer?
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	gl.BindVertexArray(0)
	gl.UseProgram(0)
}

//TODO Assert on dataSizes and such not matching multiples of stored correct value?
