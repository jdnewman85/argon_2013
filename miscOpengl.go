package argon

import (
	"log"

	gl "github.com/chsc/gogl/gl43"
)

func init() {
	log.Println("miscOpengl.go here")
}

const ArrayBuffer = gl.ARRAY_BUFFER

type Attribute struct {
	Index      gl.Uint
	Size       gl.Int
	Kind       gl.Enum
	Normalized gl.Boolean
	Offset     gl.Uint
}

type Buffer gl.Uint
type Vao gl.Uint

type VertexBuffer struct {
	Buffer
	Offset gl.Intptr
	Stride gl.Sizei
}

//func (this VertexBuffer) Bind(index gl.Uint) {
//	gl.BindVertexBuffer(index, gl.Uint(this.Buffer), this.Offset, this.Stride)
//}//TODO Other Buffer bind indicies
//TODO ??? Should this be a vao method that takes a VertexBuffer as a parameters?
//TODO		I guess it depends on if the VAO binds? idk
func (this VertexBuffer) BindVertexBuffer() {
	gl.BindVertexBuffer(0, gl.Uint(this.Buffer), this.Offset, this.Stride)
}

func GenVao() Vao {
	var vao gl.Uint
	gl.GenVertexArrays(1, &vao)
	return Vao(vao)
}

func GenBuffer() Buffer {
	var buffer gl.Uint
	gl.GenBuffers(1, &buffer)
	return Buffer(buffer)
}

func (this Vao) Bind() {
	gl.BindVertexArray(gl.Uint(this))
}

func (this Buffer) Bind(target gl.Enum) {
	gl.BindBuffer(target, gl.Uint(this))
}

func (this Vao) UnBind() {
	gl.BindVertexArray(0)
}

//TODO Would be nice to not know the target here...?
func (this Buffer) UnBind(target gl.Enum) {
	gl.BindBuffer(target, 0)
}

//func (this Buffer) Data(target gl.Enum, size int, data uintptr, usage gl.Enum) {
//	gl.BufferData(target, gl.Sizeiptr(size), gl.Pointer(data), usage)
//}

func (this Buffer) Data(target gl.Enum, size gl.Sizeiptr, data gl.Pointer, usage gl.Enum) {
	gl.BufferData(target, size, data, usage)
}

func (this Vao) SetAttributes(attributes []Attribute) {
	//Attributes
	for _, t := range attributes {
		gl.VertexAttribFormat(t.Index, t.Size, t.Kind, t.Normalized, t.Offset)
		gl.VertexAttribBinding(t.Index, 0) //TODO Other Buffer bind indicies
		gl.EnableVertexAttribArray(t.Index)
	}
}

func (this Vao) Draw(num gl.Sizei, program Program) {
	program.Use()
	defer program.Forgo()
	this.Bind()
	defer this.UnBind()

	gl.DrawArrays(gl.POINTS, 0, num)
}

