// Testbench for Argon Graphics Engine

package main

import (
	"fmt"
	"reflect"

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
	myCircle0 := argon.NewCircle()
	myGraphics.DrawMap[reflect.TypeOf(myCircle0)] = circleRenderer
	myGraphics.DrawMap[reflect.TypeOf([]argon.Circle{*myCircle0})] = circleRenderer
	myGraphics.DrawMap[reflect.TypeOf([3]argon.Circle{*myCircle0})] = circleRenderer

	myCircle1 := argon.NewCircle()
	myCircle2 := argon.NewCircle()
	myCircle3 := argon.NewCircle()
	myCircle0.X = 200.0
	myCircle0.Y = 200.0
	myCircle0.R = 25.0
	myCircle0.Red = 1.0
	myCircle0.Green = 0.0
	myCircle0.Blue = 0.0
	myCircle0.Alpha = 1.0
	myCircle1.X = 200.0
	myCircle1.Y = 500.0
	myCircle1.R = 25.0
	myCircle1.Red = 0.0
	myCircle1.Green = 1.0
	myCircle1.Blue = 0.0
	myCircle1.Alpha = 1.0
	myCircle2.X = 500.0
	myCircle2.Y = 200.0
	myCircle2.R = 25.0
	myCircle2.Red = 0.0
	myCircle2.Green = 0.0
	myCircle2.Blue = 1.0
	myCircle2.Alpha = 1.0
	myCircle3.X = 500.0
	myCircle3.Y = 500.0
	myCircle3.R = 25.0
	myCircle3.Red = 1.0
	myCircle3.Green = 1.0
	myCircle3.Blue = 1.0
	myCircle3.Alpha = 1.0


	for !myGraphics.ShouldClose() {
		myGraphics.Cls()

		//Circles
//		circleRenderer.Draw(myCircle0)
//		circleRenderer.Draw(myCircle1)
//		circleRenderer.Draw(myCircle2)
//		circleRenderer.Draw(myCircle3)
		myGraphics.Draw([...]argon.Circle{*myCircle0, *myCircle1, *myCircle2})
//		myGraphics.Draw(myCircle1)
//		myGraphics.Draw(myCircle2)
		myGraphics.Draw(myCircle3)

		myGraphics.Flip()
	}

}

