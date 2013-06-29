
package circle

import(
	"log"
	"unsafe"

	gl "github.com/chsc/gogl/gl33"

	"bitbucket.org/jdnewman/argon/renderer"
)

func init() {
	log.Println("circles.go here")
}

type Circles []Circle

func CreateCircles(num int) *Circles {
	temp := make(Circles, num)
	return &temp
}

func (this *Circles) CircleArray() []Circle {
	return []Circle(*this)
}

func (this *Circles) Draw() {
	gRenderer.Draw(this, gShader)
}

func (this *Circles) DrawData() renderer.DrawData {
	//TODO Get rid of gl dependancies?
	circleArray := (*this).CircleArray()
	return renderer.DrawData{gl.Pointer(&circleArray[0]), gl.Sizeiptr(unsafe.Sizeof(defaultCircle)*uintptr(len(circleArray))), gl.Sizei(len(circleArray))}
}

