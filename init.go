package main

// Key initialization actions include:
// 1. Constructing the unit templates
// 2. Creating/Loading in images into ebitimage?

const unitTemplateCount = 2

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

func UnitTemplateBuilder() {
	for i := 0; i == unitTemplateCount; i++ {
		unitTemplateSlice := make([]unitTemplate, 0, unitTemplateCount)
	}		
	curUnitTemplate := ConstructUnit("2x2")
	append(unitTemplateSlice, curUnitTemplate)
	curUnitTemplate := ConstructUnit("1NodeTriangle")
	append(unitTemplateSlice, curUnitTemplate)
}

func ConstructNodeCore(
	nodeID int,
	x float32,
	y float32,
	childVectors []NodeVector) NodeCore {

	node := NodeCore{nodeID, x, y, 0, 0, childVectors}

	return node
}

func ConstructUnit(unitTemplateString string) UnitTemplate {

	switch unitTemplateString {
	case "2x2":
		unitTemplateReturned := Construct2x2UnitTemplate("2x2")
	case "1NodeTriangle":
		unitTemplateReturned := Construct1NodeTriangleUnitTemplate("1NodeTriangle")
	}

	return unitTemplate
}

func ConstructUnitTemplates() {

	Construct2x2UnitTemplate("2x2")
	Construct1NodeTriangleUnitTemplate("1NodeTriangle")

}

func Construct2x2UnitTemplate(unitTypePassed string) UnitTemplate {

	//First Vector at 9 o'clock, everything goes clockwise starting at 9
	//Offset is based off 1st vector in list FOR NOW
	nodeVector1 := NodeVector{[]int{1, 4}, -2.5, -2.5, 0, 0}
	nodeVector2 := NodeVector{[]int{1}, -2.5, 2.5, 0, 0}
	nodeVector3 := NodeVector{[]int{1, 2}, 2.5, 2.5, 0, 0}
	nodeVector4 := NodeVector{[]int{2}, 2.5, 2.5, 0, 0}
	nodeVector5 := NodeVector{[]int{2, 3}, 2.5, -2.5, 0, 0}
	nodeVector6 := NodeVector{[]int{3}, 2.5, 2.5, 0, 0}
	nodeVector7 := NodeVector{[]int{3, 4}, -2.5, 2.5, 0, 0}
	nodeVector8 := NodeVector{[]int{4}, -2.5, 2.5, 0, 0}

	node1 := ConstructNodeCore(1, 0, 0, []NodeVector{nodeVector1, nodeVector2, nodeVector3})
	node2 := ConstructNodeCore(2, 0, 0, []NodeVector{nodeVector3, nodeVector4, nodeVector5})
	node3 := ConstructNodeCore(3, 0, 0, []NodeVector{nodeVector5, nodeVector6, nodeVector7})
	node4 := ConstructNodeCore(4, 0, 0, []NodeVector{nodeVector7, nodeVector8, nodeVector1})

	unit := UnitTemplate{
		unitType:  "2x2",
		xStarting: 0,
		yStarting: 0,
		nodeSlice: []NodeCore{node1, node2, node3, node4}
	}

	return unit
}

func Construct1NodeTriangleUnitTemplate(unitTypePassed string) UnitTemplate {

	//First Vector at 9 o'clock, everything goes clockwise starting at 9
	//Offset is based off 1st vector in list FOR NOW
	nodeVector1 := NodeVector{[]int{1}, -3, -1.5, 0, 0}
	nodeVector2 := NodeVector{[]int{1}, 3, 0, 0, 0}
	nodeVector3 := NodeVector{[]int{1}, -3, 1.5, 0, 0}

	node1 := ConstructNodeCore(1, 0, 0, []NodeVector{nodeVector1, nodeVector2, nodeVector3})
	unit := UnitTemplate{
		unitType:  "1NodeTriangle",
		xStarting: 0,
		yStarting: 0,
		nodeSlice: []NodeCore{node1}
	}

	return unit
}
