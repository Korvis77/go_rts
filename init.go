package main

// Key initialization actions include:
// 1. Constructing the unit templates
// 2. Creating/Loading in images into ebitimage?

type UnitTemplate struct {
	unitType  string
	xStarting float32
	yStarting float32
	nodeSlice []NodeCore
}

type NodeCore struct {
	nodeID       int
	x            float32
	y            float32
	xVel         float32
	yVel         float32
	childVectors []NodeVector
}

type NodeVector struct {
	parentnodeIDs []int
	xvectoroffset float32
	yvectoroffset float32
	xsummedforces float32
	ysummedforces float32
}

func ConstructUnitTemplates() {

	//First Vector at 9 o'clock
	unitType := "2x2"
	Construct2x2UnitTemplate(unitType)

	//repeat above pattern
}

func Construct2x2UnitTemplate(unitTypePassed string) UnitTemplate {

	nodeVector1 := NodeVector{[]int{1, 4}, -2.5, -2.5, 0, 0}

	node1 := ConstructNodeCore(1, 0, 0, []NodeVector{nodeVector1})
	// node2 := ConstructNodeCore(2, 5, 0, []int{3, 4, 5})
	// node3 := ConstructNodeCore(3, 5, 5, []int{5, 6, 7})
	// node4 := ConstructNodeCore(4, 0, 5, []int{7, 8, 1})

	//constructedNodeSlice := []NodeCore{node1}

	unit := UnitTemplate{
		unitType:  "2x2",
		xStarting: 0,
		yStarting: 0,
		nodeSlice: []NodeCore{node1}}

	return unit
}

func ConstructNodeCore(
	nodeID int,
	x float32,
	y float32,
	childVectors []NodeVector) NodeCore {

	node := NodeCore{nodeID, x, y, 0, 0, childVectors}

	return node
}
