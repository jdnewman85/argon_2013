// Testbench for Argon Graphics Engine

package main

import (
	"fmt"

	"bitbucket.org/jdnewman/argon"
	"bitbucket.org/jdnewman/argon/renderer"
)

func main() {
	fmt.Println("This is a testbench.")

	myGraphics, err := argon.NewGraphics(1280, 720, false)
	if err != nil {
		fmt.Println("ERROR: %s", err)
	}
	defer myGraphics.Destroy()

	//Circles
	circleRenderer := renderer.NewCircleRenderer()
	myCircle := make([]argon.Circle, 4)
	myGraphics.RegisterRenderer(myCircle, circleRenderer)

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
	myCircle[3].X = 500.0
	myCircle[3].Y = 500.0
	myCircle[3].R = 25.0
	myCircle[3].Red = 1.0
	myCircle[3].Green = 1.0
	myCircle[3].Blue = 1.0
	myCircle[3].Alpha = 1.0


	for !myGraphics.ShouldClose() {
		myGraphics.Cls()

		//Circles
		myGraphics.Draw(myCircle)

		myGraphics.Flip()
	}

}

