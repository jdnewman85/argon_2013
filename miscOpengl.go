package argon

import (
	"log"

	gl "github.com/chsc/gogl/gl43"
)

func init() {
	log.Println("miscOpengl.go here")
}

const ArrayBuffer = gl.ARRAY_BUFFER

type Buffer gl.Uint
type Vao gl.Uint


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
//	//TODO Bind?
//	gl.BufferData(target, gl.Sizeiptr(size), gl.Pointer(data), usage)
//}

func (this Buffer) Data(target gl.Enum, size gl.Sizeiptr, data gl.Pointer, usage gl.Enum) {
	//TODO Bind?
	gl.BufferData(target, size, data, usage)
}

