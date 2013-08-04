// Sprite Primitive

package argon

import (
	"log"

	gl "github.com/chsc/gogl/gl43"
)

func init() {
	log.Println("texture.go here")
}

type Texture struct {
	texture       gl.Uint
	Width, Height gl.Int
}
