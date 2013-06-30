// Testbench for Argon Graphics Engine

package main

import (
	"fmt"

	"bitbucket.org/jdnewman/argon"
	"bitbucket.org/jdnewman/argon/circle"
)

func main() {
	fmt.Println("This is a testbench.")

	argon.Graphics(1280, 720, false)
	defer argon.EndGraphics()

	//Circles
	circle.Init()

//	myCircle0 := circle.Create()
//	myCircle1 := circle.Create()
//	myCircle2 := circle.Create()
//	myCircle3 := circle.Create()
//	myCircle0.X = 200.0
//	myCircle0.Y = 200.0
//	myCircle0.R = 25.0
//	myCircle0.Red = 1.0
//	myCircle0.Green = 0.0
//	myCircle0.Blue = 0.0
//	myCircle0.Alpha = 1.0
//	myCircle1.X = 200.0
//	myCircle1.Y = 500.0
//	myCircle1.R = 25.0
//	myCircle1.Red = 0.0
//	myCircle1.Green = 1.0
//	myCircle1.Blue = 0.0
//	myCircle1.Alpha = 1.0
//	myCircle2.X = 500.0
//	myCircle2.Y = 200.0
//	myCircle2.R = 25.0
//	myCircle2.Red = 0.0
//	myCircle2.Green = 0.0
//	myCircle2.Blue = 1.0
//	myCircle2.Alpha = 1.0
//	myCircle3.X = 500.0
//	myCircle3.Y = 500.0
//	myCircle3.R = 25.0
//	myCircle3.Red = 1.0
//	myCircle3.Green = 1.0
//	myCircle3.Blue = 1.0
//	myCircle3.Alpha = 1.0

	mahCircles := circle.CreateCircles(4)
	myCircles := mahCircles.CircleArray()
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

	for argon.WindowOpen() {
		argon.Cls()

		//Circles
//		myCircle0.Draw()
//		myCircle1.Draw()
//		myCircle2.Draw()
//		myCircle3.Draw()
		mahCircles.Draw()

		argon.Flip()
	}

}

