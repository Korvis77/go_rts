package main

import (
	"image/color"
	_ "image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// Global Varss

var mouseFirst rl.Vector2 // For mouse input

func main() {

	rl.InitWindow(800, 450, "raylib [core] example - basic window")
	defer rl.CloseWindow()

	NodeColor := color.RGBA{0, 0, 0, 255}
	var nodeTestText string = "testing testing"

	// StoreUserInputMap := map[string]bool{
	// 	"mouse_left":  false,
	// 	"mouse_right": false,
	// }

	UnitTemplate2x2 := Build2x2UnitTemplate()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		PlayerInputExecutor()

		rl.DrawText(nodeTestText, 190, 200, 20, rl.LightGray)
		//rl.DrawRectangle(0, 0, 400, 200, NodeColor)
		//rl.DrawCircle(300, 300, 300, NodeColor)

		DrawUnitOutline(UnitTemplate2x2, 300, 300, NodeColor)

		rl.EndDrawing()
	}
}

// func PlayerInputListener(liveInputs []int32) {
// 	for i := 0; i < len(liveInputs); i++ {
// 		keyFound
// 	}
// }

func PlayerInputExecutor() {
	// var prevMouseStart rl.Vector2
	mouseInputCur := MouseVectorListener()
	InputInterfaceProcessor(mouseInputCur)
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

func MouseVectorListener() *MouseInput {

	// var firstMouseVector rl.Vector2
	// var secondMouseVector rl.Vector2
	// var returnVector rl.Vector2

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

	// if rl.IsMouseButtonDown(rl.MouseButtonLeft) == true {
	if rl.IsMouseButtonDown(rl.MouseButtonLeft) == true {
		draw = true
		mouseCur = rl.GetMousePosition()
		// fmt.Printf("mouseCur X: /%f  \n", mouseCur.X)
		//MouseVectorListener()
		//fmt.Printf("MouseButtonLeft int value: /%d  ", rl.MouseButtonLeft)
		//fmt.Printf("key pressed int: /%d   /n", rl.GetKeyPressed())
	} else {
		// mouseStart = rl.GetMousePosition()
		// }

	}

	returnedMouseInput := &MouseInput{
		vector1:  mouseFirst,
		vector2:  mouseCur,
		drawbool: draw,
		movebool: move,
	}

	return returnedMouseInput
}

// if rl.IsMouseButtonReleased(rl.MouseButtonLeft) == true {
// secondMouseVector = rl.GetMousePosition()

// normalizedMouseVector := rl.Vector2Subtract(firstMouseVector, secondMouseVector)
// normalizedMouseVector = rl.Vector2Normalize(normalizedMouseVector)

// lerpedVector := rl.Vector2Lerp(firstMouseVector, secondMouseVector, 1)
// //fmt.Printf("/n Vector X Result is: /%f  ", lerpedVector.X)
// //fmt.Printf("/n Vector Y Result is: /%f  ", lerpedVector.Y)

// v1 := rl.Vector2Normalize(firstMouseVector)
// v2 := rl.Vector2Normalize(secondMouseVector)
// vLineAngle := rl.Vector2LineAngle(v1, v2)
// fmt.Printf("/n Line Angle Float Solved is : /%f  ", vLineAngle)
// }

func DrawMousePath(v1 rl.Vector2, v2 rl.Vector2, color color.RGBA) {
	var v1x = int32(v1.X)
	var v1y = int32(v1.Y)
	var v2x = int32(v2.X)
	var v2y = int32(v2.Y)
	rl.DrawLine(v1x, v1y, v2x, v2y, color)
}

func DrawUnitOutline(unitTemplate *UnitTemplateStruct, xtransform int32, ytransform int32, nodeColor color.RGBA) {
	vectorCount := len(unitTemplate.coreNodeSlice)

	for i := 0; i < vectorCount; i++ {

		j := i + 1

		//Iterate on building position after applying sum of forces:

		if i == vectorCount-1 {
			vec1xpos := unitTemplate.coreNodeSlice[0].xOffset + xtransform
			vec1ypos := unitTemplate.coreNodeSlice[0].yOffset + ytransform
			vec2xpos := unitTemplate.coreNodeSlice[i].xOffset + xtransform
			vec2ypos := unitTemplate.coreNodeSlice[i].yOffset + ytransform
			rl.DrawLine(vec1xpos, vec1ypos, vec2xpos, vec2ypos, nodeColor)
		} else {
			vec1xpos := unitTemplate.coreNodeSlice[i].xOffset + xtransform
			vec1ypos := unitTemplate.coreNodeSlice[i].yOffset + ytransform
			vec2xpos := unitTemplate.coreNodeSlice[j].xOffset + xtransform
			vec2ypos := unitTemplate.coreNodeSlice[j].yOffset + ytransform
			// fmt.Printf("vec1xpos /%d :", vec1xpos)
			// fmt.Printf("vec1ypos /%d :", vec1ypos)
			// fmt.Printf("vec2xpos /%d :", vec2xpos)
			// fmt.Printf("vec2ypos /%d :", vec2ypos)
			rl.DrawLine(vec1xpos, vec1ypos, vec2xpos, vec2ypos, nodeColor)
		}

		// vectorToDraw := rl.Vector2{
		// 	X: vec2xpos - vec1xpos,
		// 	Y: vec2ypos - vec1ypos,
		// }

	}
}
