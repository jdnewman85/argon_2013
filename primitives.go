package argon

import (
	"log"

	gl "github.com/chsc/gogl/gl43"
)

func init() {
	log.Println("primitives.go here")
}

// Circle Primitive
var DefaultCircle Circle = Circle{0.0, 0.0, 20.0, 1.0, 1.0, 1.0, 1.0}

type Circle struct {
	X, Y                    gl.Float
	R                       gl.Float
	Red, Green, Blue, Alpha gl.Float
}

// Line Primitive
var DefaultLine Line = Line{0.0, 0.0, 100.0, 100.0, 10.0, 1.0, 1.0, 1.0, 1.0}

type Line struct {
	X1, Y1                  gl.Float
	X2, Y2                  gl.Float
	R                       gl.Float
	Red, Green, Blue, Alpha gl.Float
}
