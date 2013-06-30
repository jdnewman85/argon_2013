
package circle

import(
	"log"
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

