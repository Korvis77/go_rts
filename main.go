package main

import (
	"fmt"
	"image/color"
	_ "image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// Global Varss

var mouseFirst rl.Vector2 // For mouse input

func main() {

	rl.InitWindow(800, 450, "raylib [core] example - basic window")
	defer rl.CloseWindow()

	unitBoundaryColor := color.RGBA{0, 0, 0, 255}
	boundaryNodeColor := color.RGBA{0, 255, 0, 255}
	var nodeTestText string = "testing testing"

	// StoreUserInputMap := map[string]bool{
	// 	"mouse_left":  false,
	// 	"mouse_right": false,
	// }

	UnitTemplate2x2 := Build2x2UnitTemplate()
	fmt.Printf("vector1 nested slices: %d  ", UnitTemplate2x2.vectorNodeSlice[0])

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		PlayerInputExecutor()

		rl.DrawText(nodeTestText, 190, 200, 20, rl.LightGray)
		//rl.DrawCircle(300, 300, 300, NodeColor)

		DrawUnit(UnitTemplate2x2, 300, 300, unitBoundaryColor, boundaryNodeColor)

		rl.EndDrawing()
	}
}

func ActiveTesting() {
	testlen := [][]int32{{1, 3, -3}, {2, -3, -3}}
	for index, value := range testlen {
		fmt.Printf("Index: %d  wow cool \n ", index)
		fmt.Printf("Value: %d  wow cool \n ", value)
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

func DrawUnit(
	unitTemplate *UnitTemplateStruct,
	xtransform int32,
	ytransform int32,
	unitBoundaryColor color.RGBA,
	coreNodeColor color.RGBA,
) {

	var coreNode1X int32
	var coreNode1Y int32
	var coreNode2X int32
	var coreNode2Y int32

	var sumBoundaryNodeOffsetX int32
	var sumBoundaryNodeOffsetY int32

	var boundaryNode1X int32
	var boundaryNode1Y int32
	// var boundaryNode2X int32
	// var boundaryNode2Y int32
	// var sumPrevBoundaryNodeOffsetX int32
	// var sumPrevBoundaryNodeOffsetY int32
	// var offsetCore1X int32
	// var offsetCore1Y int32
	// var offsetCore2X int32
	// var offsetCore2Y int32

	coreNodeCount := len(unitTemplate.coreNodeSlice)
	for index, _ := range unitTemplate.coreNodeSlice {
		if index == coreNodeCount-1 {
			coreNode1X = unitTemplate.coreNodeSlice[0].xOffset + xtransform
			coreNode1Y = unitTemplate.coreNodeSlice[0].yOffset + ytransform
			coreNode2X = unitTemplate.coreNodeSlice[index].xOffset + xtransform
			coreNode2Y = unitTemplate.coreNodeSlice[index].yOffset + ytransform
			rl.DrawCircle(coreNode2X, coreNode2Y, 5, coreNodeColor)
		} else {
			coreNode1X = unitTemplate.coreNodeSlice[index].xOffset + xtransform
			coreNode1Y = unitTemplate.coreNodeSlice[index].yOffset + ytransform
			// coreNode2X = unitTemplate.coreNodeSlice[index+1].xOffset + xtransform
			// coreNode2Y = unitTemplate.coreNodeSlice[index+1].yOffset + ytransform
		}
		rl.DrawCircle(coreNode1X, coreNode1Y, 5, coreNodeColor)
	}

	for _, value := range unitTemplate.vectorNodeSlice {

		parentList := value.parentCores
		for _, value2 := range parentList {
			sumBoundaryNodeOffsetX += value2[1]
			sumBoundaryNodeOffsetY += value2[2]
		}

		boundaryNode1X = sumBoundaryNodeOffsetX + xtransform
		boundaryNode1Y = sumBoundaryNodeOffsetY + ytransform
		// fmt.Printf("SumNodeX: %d  wow cool ", sumBoundaryNodeOffsetX)
		// fmt.Printf("SumNodeY: %d  wow cool \n ", sumBoundaryNodeOffsetY)
		rl.DrawCircle(boundaryNode1X, boundaryNode1Y, 3, unitBoundaryColor)
	}

	// Draw Node Cores
	// coreNodeCount := len(unitTemplate.coreNodeSlice)
	// for i := 0; i < coreNodeCount; i++ {
	// 	j := i + 1

	// 	if i == coreNodeCount-1 {
	// 		vec1xpos = unitTemplate.coreNodeSlice[0].xOffset + xtransform
	// 		vec1ypos = unitTemplate.coreNodeSlice[0].yOffset + ytransform
	// 		vec2xpos = unitTemplate.coreNodeSlice[i].xOffset + xtransform
	// 		vec2ypos = unitTemplate.coreNodeSlice[i].yOffset + ytransform
	// 	} else {
	// 		vec1xpos = unitTemplate.coreNodeSlice[i].xOffset + xtransform
	// 		vec1ypos = unitTemplate.coreNodeSlice[i].yOffset + ytransform
	// 		vec2xpos = unitTemplate.coreNodeSlice[j].xOffset + xtransform
	// 		vec2ypos = unitTemplate.coreNodeSlice[j].yOffset + ytransform
	// 	}
	// 	//rl.DrawLine(vec1xpos, vec1ypos, vec2xpos, vec2ypos, unitBoundaryColor)
	// 	rl.DrawCircle(vec1xpos, vec1ypos, 5, nodeCoreColor)
	// }

	// // Draw Unit Boundary
	// boundaryNodeCount := len(unitTemplate.vectorNodeSlice)
	// for i := 0; i < boundaryNodeCount; i++ {
	// 	j := i + 1

	// 	if i == coreNodeCount-1 {
	// 		vec1xpos = unitTemplate.vectorNodeSlice[0].xOffset + xtransform
	// 		vec1ypos = unitTemplate.vectorNodeSlice[0].yOffset + ytransform
	// 		vec2xpos = unitTemplate.vectorNodeSlice[i].xOffset + xtransform
	// 		vec2ypos = unitTemplate.coreNodeSlice[i].yOffset + ytransform
	// 	} else {
	// 		vec1xpos = unitTemplate.coreNodeSlice[i].xOffset + xtransform
	// 		vec1ypos = unitTemplate.coreNodeSlice[i].yOffset + ytransform
	// 		vec2xpos = unitTemplate.coreNodeSlice[j].xOffset + xtransform
	// 		vec2ypos = unitTemplate.coreNodeSlice[j].yOffset + ytransform
	// 	}
	// 	//rl.DrawLine(vec1xpos, vec1ypos, vec2xpos, vec2ypos, unitBoundaryColor)
	// 	rl.DrawCircle(vec1xpos, vec1ypos, 5, nodeCoreColor)
	// }

}

/* vector length, this is the hypotenuse,
   hypotenuse (vector length a.k.a. offset from x and y)
   cosine for x, sine for y,
   arctan(yoffset/xoffset)
   rotation (passed into the function)

   hypotenuse * cos/sin (arctan*(yoffset/xoffset) + rotation)
                         ^^^ solve once, becomes theta
*/
