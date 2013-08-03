// Sprite Primitive

package argon

import (
	"log"

	gl "github.com/chsc/gogl/gl33"
)

func init() {
	log.Println("sprite.go here")
}

var DefaultSprite Sprite = Sprite{0.0, 0.0, 32.0, 32.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 0.0, 0}

type Sprite struct {
	X1, Y1 gl.Float
	Width, Height gl.Float
	Red, Green, Blue, Alpha gl.Float
	ScaleX, ScaleY gl.Float
	Angle gl.Float
	texture gl.Uint //TODO Needs to be my own object
}

