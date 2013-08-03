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
	myCircles := make([]argon.Circle, 5)
	myGraphics.RegisterRenderer(myCircles, circleRenderer)

	myCircles[0].X = 200.0
	myCircles[0].Y = 200.0
	myCircles[0].R = 25.0
	myCircles[0].Red = 1.0
	myCircles[0].Green = 0.0
	myCircles[0].Blue = 0.0
	myCircles[0].Alpha = 1.0
	myCircles[1].X = 200.0
	myCircles[1].Y = 500.0
	myCircles[1].R = 25.0
	myCircles[1].Red = 0.0
	myCircles[1].Green = 1.0
	myCircles[1].Blue = 0.0
	myCircles[1].Alpha = 1.0
	myCircles[2].X = 500.0
	myCircles[2].Y = 200.0
	myCircles[2].R = 25.0
	myCircles[2].Red = 0.0
	myCircles[2].Green = 0.0
	myCircles[2].Blue = 1.0
	myCircles[2].Alpha = 1.0
	myCircles[3].X = 500.0
	myCircles[3].Y = 500.0
	myCircles[3].R = 25.0
	myCircles[3].Red = 1.0
	myCircles[3].Green = 1.0
	myCircles[3].Blue = 1.0
	myCircles[3].Alpha = 1.0
	myCircles[4] = argon.DefaultCircle

	//Lines
	lineRenderer := renderer.NewLineRenderer()
	myLines := make([]argon.Line, 5)
	myGraphics.RegisterRenderer(myLines, lineRenderer)

	myLines[0].X1 = 200.0
	myLines[0].Y1 = 200.0
	myLines[0].X2 = 250.0
	myLines[0].Y2 = 250.0
	myLines[0].R = 10.0
	myLines[0].Red = 1.0
	myLines[0].Green = 0.0
	myLines[0].Blue = 0.0
	myLines[0].Alpha = 1.0
	myLines[1].X1 = 200.0
	myLines[1].Y1 = 500.0
	myLines[1].X2 = 250.0
	myLines[1].Y2 = 250.0
	myLines[1].R = 10.0
	myLines[1].Red = 0.0
	myLines[1].Green = 1.0
	myLines[1].Blue = 0.0
	myLines[1].Alpha = 1.0
	myLines[2].X1 = 500.0
	myLines[2].Y1 = 200.0
	myLines[2].X2 = 250.0
	myLines[2].Y2 = 250.0
	myLines[2].R = 10.0
	myLines[2].Red = 0.0
	myLines[2].Green = 0.0
	myLines[2].Blue = 1.0
	myLines[2].Alpha = 1.0
	myLines[3].X1 = 500.0
	myLines[3].Y1 = 500.0
	myLines[3].X2 = 250.0
	myLines[3].Y2 = 250.0
	myLines[3].R = 10.0
	myLines[3].Red = 1.0
	myLines[3].Green = 1.0
	myLines[3].Blue = 1.0
	myLines[3].Alpha = 1.0
	myLines[4] = argon.DefaultLine

	for !myGraphics.ShouldClose() {
		myGraphics.Cls()

		//Circles
		myGraphics.Draw(myCircles)
		myGraphics.Draw(myLines)

		myGraphics.Flip()
	}

}

