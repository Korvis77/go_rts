package main

import (
	"image/color"
	_ "image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// 1. Packages & Import
// 2. Main Processor
// 2b. Listeners - just Mouse for now
// 3a. InputInterfaceProcessor
// 4. Interface Definition
// 5. *MouseInput Interface Methods

func MouseVectorListener() *MouseInput {

	var mouseCur rl.Vector2
	var draw bool = false
	var move bool = false

	if rl.IsMouseButtonReleased(rl.MouseButtonLeft) == true {
		move = true
	}

	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) == true {
		mouseFirst = rl.GetMousePosition()
		//fmt.Printf("mouseStart X: /%f  \n", mouseStart.X)
	}

	if rl.IsMouseButtonDown(rl.MouseButtonLeft) == true {
		draw = true
		mouseCur = rl.GetMousePosition()
	}

	returnedMouseInput := &MouseInput{
		vector1:  mouseFirst,
		vector2:  mouseCur,
		drawbool: draw,
		movebool: move,
	}

	return returnedMouseInput
}

func InputInterfaceProcessor(curInputInterface InputInterface) {
	inputActions := curInputInterface.BoolEval()

	for _, value := range inputActions {
		switch value {
		case "draw":
			curInputInterface.Draw()
		case "transform":
			curInputInterface.Transform()
		}
	}
}

type InputInterface interface {
	BoolEval() []string
	Draw()
	Transform()
}

type MouseInput struct {
	vector1  rl.Vector2
	vector2  rl.Vector2
	drawbool bool
	movebool bool
}

func (input *MouseInput) BoolEval() []string {
	var returnStringList = make([]string, 0)
	if input.drawbool == true {
		returnStringList = append(returnStringList, "draw")
	}
	if input.movebool == true {
		returnStringList = append(returnStringList, "transform")
	}
	return returnStringList
}

func (input *MouseInput) Draw() {
	NodeColor := color.RGBA{0, 0, 0, 255}
	// fmt.Printf("mouseStart X: %f  \n", input.vector1.X)
	// fmt.Printf("mouseCur X: %f  \n", input.vector2.X)
	DrawMousePath(input.vector1, input.vector2, NodeColor)
}

func (input *MouseInput) Transform() {

}
