package main

import (
	"fmt"
)

// Key initialization actions include:
// 1. Constructing the unit templates
// 2. Creating/Loading in images into ebitimage?

// 1. Define struct:
type UnitTemplateStruct struct {
	unitID          int
	unitType        string
	coreNodeSlice   []CoreNode
	vectorNodeSlice []VectorNode
}

type CoreNode struct {
	coreID           int
	childVectorSlice []VectorNode
	xOffset          int32
	yOffset          int32
}

type VectorNode struct {
	vectorID        int
	parentCores     [][]int32
	parentCoreSlice []int
	xOffset         int32
	yOffset         int32
}

// 2. Build struct literal:
func Build2x2UnitTemplate() *UnitTemplateStruct {

	vector1 := VectorNode{
		vectorID: 1,
		parentCores: [][]int32{
			{1, -3, 3},
			{4, 3, 3},
		},
	}
	vector2 := VectorNode{
		vectorID: 2,
		parentCores: [][]int32{
			{1, -3, -3},
		},
	}
	vector3 := VectorNode{
		vectorID: 3,
		parentCores: [][]int32{
			{1, 3, -3},
			{2, -3, -3},
		},
	}
	vector4 := VectorNode{
		vectorID: 4,
		parentCores: [][]int32{
			{2, 3, -3},
		},
	}
	vector5 := VectorNode{
		vectorID: 5,
		parentCores: [][]int32{
			{2, 3, 3},
			{3, 3, -3},
		},
	}
	vector6 := VectorNode{
		vectorID: 6,
		parentCores: [][]int32{
			{3, 3, 3},
		},
	}
	vector7 := VectorNode{
		vectorID: 7,
		parentCores: [][]int32{
			{3, -3, 3},
			{4, -3, -3},
		},
	}
	vector8 := VectorNode{
		vectorID: 8,
		parentCores: [][]int32{
			{4, -3, 3},
		},
	}

	unitTemplate := &UnitTemplateStruct{
		unitID:   1,
		unitType: "2x2",
		coreNodeSlice: []CoreNode{
			{
				coreID:           1,
				childVectorSlice: []VectorNode{vector1, vector2, vector3},
				xOffset:          -25,
				yOffset:          -25,
			},
			{
				coreID:           2,
				childVectorSlice: []VectorNode{vector3, vector4, vector5},
				xOffset:          25,
				yOffset:          -25,
			},
			{
				coreID:           3,
				childVectorSlice: []VectorNode{vector5, vector6, vector7},
				xOffset:          25,
				yOffset:          25,
			},
			{
				coreID:           4,
				childVectorSlice: []VectorNode{vector7, vector8, vector1},
				xOffset:          -25,
				yOffset:          25,
			},
		},
		vectorNodeSlice: []VectorNode{vector1, vector2, vector3, vector4, vector5, vector6, vector7, vector8},
	}

	return unitTemplate
}

// 3. Interface for struct literal
type UnitTemplateInterface interface {
	NodeCount() int
}

// Now I can define different functions for different struct literals such as:

func (unitTemplateLiteral *UnitTemplateStruct) NodeCount() int {

	countOfNodes := len(unitTemplateLiteral.coreNodeSlice)

	fmt.Printf("length in NodeCount: %d \n", countOfNodes)
	return countOfNodes

}

func CountNodesUseTemplateInterface(unitTemplateExample UnitTemplateInterface) int {
	testint := unitTemplateExample.NodeCount()
	unitTemplateExample.NodeCount()
	fmt.Printf("length through function: %d \n", unitTemplateExample.NodeCount())

	if testint == 4 {
		fmt.Printf("this is a 2x2 \n")
	}
	if testint == 1 {
		fmt.Printf("this has 1 node \n")
	}

	return testint
}
