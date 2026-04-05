package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct{}

// Read what Claude was talking about and then do a lot of studying of theory.
// Time to learn about how to decode an image from the image file's byte slice

// var (
// 	ebitenImage *ebiten.Image
// )

// func init() {
// 	// Decode an image from the image file's byte slice.
// 	img, _, err := image.Decode(bytes.NewReader(images.Ebiten_png))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	origEbitenImage := ebiten.NewImageFromImage(img)

// 	s := origEbitenImage.Bounds().Size()
// 	ebitenImage = ebiten.NewImage(s.X, s.Y)

// 	op := &ebiten.DrawImageOptions{}
// 	op.ColorScale.ScaleAlpha(1)
// 	ebitenImage.DrawImage(origEbitenImage, op)
// }

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {

	unitTemplate2x2 := Build2x2UnitTemplate()

	CountNodesUseTemplateInterface(unitTemplate2x2)

	//ConstructUnitTemplates()

	// unit1 := SpawnUnit(1, "2x2")
	// unit2 := SpawnUnit(2, "1NodeTriangle")

	// fmt.Printf("(%d)", unit1.unitID)
	// fmt.Printf("(%d)", unit2.unitID)

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}

	// unit1 := SpawnUnit(1, "2x2", 100, 100)

	// fmt.Printf("(%d)", unit1.unitID)

	//myUnit := NodeSequence("2x2", 1, 10, 20)

	// Print the unit-level fields
	// fmt.Printf("Unit ID: %d\n", myUnit.unitID)
	// fmt.Printf("Unit X: %d\n", myUnit.x)
	// fmt.Printf("Unit Y: %d\n", myUnit.y)

	// // Loop over the nodes and print each one
	// for _, node := range myUnit.Nodes {
	// 	fmt.Printf("  Node ID: %d\n", node.nodeID)
	// 	fmt.Printf("  Position: (%d, %d)\n", node.x, node.y)
	// 	fmt.Printf("  Velocity: (%d, %d)\n", node.xVel, node.yVel)
	// 	fmt.Printf("  Neighbors: %v\n", node.neighborNodes)
	// 	fmt.Println("  ---")
	// }

}

// type Unit struct {
// 	unitID       int
// 	unitTemplate string
// }

// func SpawnUnit2x2(unitIDPassed int, unitTemplatePassed string) Unit {

// 	unitTemplate1 := Construct2x2UnitTemplate(unitTemplatePassed)
// 	unit1 := Unit{
// 		unitID: unitIDPassed,
// 		unitTemplate: UnitTemplate{unitTemplate1}
// 	}
// 	return unit1
// }

// func SpawnUnit(
// 	unitID int,
// 	unitType string,
// 	x int,
// 	y int,
// ) Unit {
// 	switch unitType {
// 	case "2x2":
// 		GenerateNodes2x2(unitID, x, y)
// 	}
// 	return Unit{}
// }

// func GenerateNodes2x2(unitID int, passedx int, passedy int) {
// 	curx := float32(passedx)
// 	cury := float32(passedy)
// 	curx -= 5

// 	nodeVector1 := &NodeVector{
// 		parentnodeIDs: []int{1, 2},
// 		xvectoroffset: -5,
// 		yvectoroffset: 5,
// 	}

// 	node1 := &NodeCore{
// 		unitID:       unitID,
// 		nodeID:       1,
// 		x:            curx,
// 		y:            cury,
// 		childVectors: []NodeVector{*nodeVector1},
// 	}
// 	curx += 5

// 	fmt.Printf(" nodeVector1 parentNodes: (%d)\n", nodeVector1.parentnodeIDs)
// 	fmt.Printf(" node1 unitId, nodeID: (%d, %d)\n", node1.unitID, node1.nodeID)
// 	fmt.Printf(" node1 x, y: (%f, %f)\n", node1.x, node1.y)

// }

// func NodeSequence(nodeSequence string, unitID int, x int, y int) Unit {
// 	switch nodeSequence {
// 	case "2x2":
// 		return NodeSequence2x2(unitID, x, y)
// 	}
// 	return Unit{}
// }

// type Node struct {
// 	unitID        int
// 	nodeID        int
// 	x             int
// 	y             int
// 	xVel          int
// 	yVel          int
// 	neighborNodes []int
// }

// func NodeSequence2x2(passedUnitID int, passedx int, passedy int) Unit {

// 	var nodex int = passedx
// 	var nodey int = passedy
// 	node1 := Node{
// 		unitID:        1,
// 		nodeID:        1,
// 		x:             nodex,
// 		y:             nodey,
// 		xVel:          0,
// 		yVel:          0,
// 		neighborNodes: []int{2, 3, 4},
// 	}
// 	nodex -= 5
// 	node2 := Node{
// 		unitID:        1,
// 		nodeID:        1,
// 		x:             nodex,
// 		y:             nodey,
// 		xVel:          0,
// 		yVel:          0,
// 		neighborNodes: []int{2, 3, 4},
// 	}
// 	//offset x and y +- passed
// 	// make more nodes here and return lol
// 	return Unit{
// 		unitID: passedUnitID,
// 		x:      passedx,
// 		y:      passedy,
// 		Nodes:  []Node{node1, node2},
// 	}
// }

// func MoveUnit(xVel int, yVel int, rotation int) {
// 	for i := range 1 {
// 		fmt.Println(i)
// 	}
// }

/* vector length, this is the hypotenuse,
   hypotenuse (vector length a.k.a. offset from x and y)
   cosine for x, sine for y,
   arctan(yoffset/xoffset)
   rotation (passed into the function)

   hypotenuse * cos/sin (arctan*(yoffset/xoffset) + rotation)
                         ^^^ solve once, becomes theta
*/
