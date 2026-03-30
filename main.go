package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct{}

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
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}

	myUnit := NodeSequence("2x2", 1, 10, 20)

	// Print the unit-level fields
	fmt.Printf("Unit ID: %d\n", myUnit.unitID)
	fmt.Printf("Unit X: %d\n", myUnit.x)
	fmt.Printf("Unit Y: %d\n", myUnit.y)

	// Loop over the nodes and print each one
	for _, node := range myUnit.Nodes {
		fmt.Printf("  Node ID: %d\n", node.nodeID)
		fmt.Printf("  Position: (%d, %d)\n", node.x, node.y)
		fmt.Printf("  Velocity: (%d, %d)\n", node.xVel, node.yVel)
		fmt.Printf("  Neighbors: %v\n", node.neighborNodes)
		fmt.Println("  ---")
	}

}

type Node struct {
	unitID        int
	nodeID        int
	x             int
	y             int
	xVel          int
	yVel          int
	neighborNodes []int
}

type Unit struct {
	unitID int
	x      int
	y      int
	Nodes  []Node
}

func NodeSequence(nodeSequence string, unitID int, x int, y int) Unit {
	switch nodeSequence {
	case "2x2":
		return NodeSequence2x2(unitID, x, y)
	}
	return Unit{}
}

func NodeSequence2x2(passedUnitID int, passedx int, passedy int) Unit {
	//node1slice := []int{2, 3, 4}
	node1 := Node{
		unitID:        1,
		nodeID:        1,
		x:             passedx,
		y:             passedy,
		xVel:          0,
		yVel:          0,
		neighborNodes: []int{2, 3, 4},
	}
	//node1.neighborNodes = node1slice

	return Unit{
		unitID: passedUnitID,
		x:      passedx,
		y:      passedy,
		Nodes:  []Node{node1},
	}
}
