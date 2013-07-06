// Circle Primitive

package argon

import (
	"log"

	gl "github.com/chsc/gogl/gl33"
)

func init() {
	log.Println("circle.go here")
}

var DefaultCircle Circle = Circle{0.0, 0.0, 1.0, 1.0, 1.0, 1.0, 1.0}

type Circle struct {
	X, Y gl.Float
	R gl.Float
	Red, Green, Blue, Alpha gl.Float
}

