// Testbench for Argon Graphics Engine

package main

import (
	"fmt"
	"bitbucket.org/jdnewman/argon"
)

func main() {
	fmt.Println("This is a testbench.")

	argon.Graphics(1280, 720, false)
	defer argon.EndGraphics()

	argon.CircleInit()

	myCircle := argon.CreateCircle(3)
	myCircle[0].X = 200.0
	myCircle[0].Y = 200.0
	myCircle[0].R = 25.0
	myCircle[0].Red = 1.0
	myCircle[0].Green = 0.0
	myCircle[0].Blue = 0.0
	myCircle[0].Alpha = 1.0
	myCircle[1].X = 200.0
	myCircle[1].Y = 500.0
	myCircle[1].R = 25.0
	myCircle[1].Red = 0.0
	myCircle[1].Green = 1.0
	myCircle[1].Blue = 0.0
	myCircle[1].Alpha = 1.0
	myCircle[2].X = 500.0
	myCircle[2].Y = 200.0
	myCircle[2].R = 25.0
	myCircle[2].Red = 0.0
	myCircle[2].Green = 0.0
	myCircle[2].Blue = 1.0
	myCircle[2].Alpha = 1.0

	otherCircle := argon.CreateCircle(1)
	otherCircle[0].X = 500.0
	otherCircle[0].Y = 500.0
	otherCircle[0].R = 25.0
	otherCircle[0].Red = 1.0
	otherCircle[0].Green = 1.0
	otherCircle[0].Blue = 1.0
	otherCircle[0].Alpha = 1.0

	for argon.WindowOpen() {
		argon.Cls()
		myCircle.Draw()
		otherCircle.Draw()
		argon.Flip()
	}

}

