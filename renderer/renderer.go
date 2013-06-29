
package renderer

import(
	"log"
	"unsafe"

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

type Renderer struct {
	vao, vbo gl.Uint
	dataSize gl.Sizeiptr
	attributes []Attribute
}

func Create(dataSize uintptr, attributes []Attribute) *Renderer {
	temp := new(Renderer)

	temp.dataSize = gl.Sizeiptr(dataSize)
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

func (this *Renderer) Draw(data unsafe.Pointer, aShader *shader.Shader) {
	//TODO Avoid unnessessary rebinds
	//Binds
	aShader.Use()
	gl.BindVertexArray(this.vao)
	gl.BindBuffer(gl.ARRAY_BUFFER, this.vbo)

	//Update Buffer
	gl.BufferData(gl.ARRAY_BUFFER, this.dataSize, gl.Pointer(data), gl.DYNAMIC_DRAW)

	//Draw
	gl.DrawArrays(gl.POINTS, 0, gl.Sizei(1))

	//TODO Defer?
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	gl.BindVertexArray(0)
	gl.UseProgram(0)
}

