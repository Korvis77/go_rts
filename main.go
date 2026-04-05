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
		//rl.DrawCircle(300, 300, 300, NodeColor)

		DrawUnitOutline(UnitTemplate2x2, 300, 300, NodeColor)

		rl.EndDrawing()
	}
}

func PlayerInputExecutor() {
	// var prevMouseStart rl.Vector2
	mouseInputCur := MouseVectorListener()
	InputInterfaceProcessor(mouseInputCur)
}

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
			rl.DrawLine(vec1xpos, vec1ypos, vec2xpos, vec2ypos, nodeColor)
		}
	}
}

/* vector length, this is the hypotenuse,
   hypotenuse (vector length a.k.a. offset from x and y)
   cosine for x, sine for y,
   arctan(yoffset/xoffset)
   rotation (passed into the function)

   hypotenuse * cos/sin (arctan*(yoffset/xoffset) + rotation)
                         ^^^ solve once, becomes theta
*/
