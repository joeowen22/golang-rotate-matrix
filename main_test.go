package main

import (
	"fmt"
	"testing"
)

var oneRotation [][]int = [][]int{{1, 7, 4, 1}, {2, 8, 5, 2}, {3, 9, 6, 3}}
var twoRotation [][]int = [][]int{{3, 2, 1}, {9, 8, 7}, {6, 5, 4}, {3, 2, 1}}
var threeRotation [][]int = [][]int{{3, 6, 9, 3}, {2, 5, 8, 2}, {1, 4, 7, 1}}

func TestOneRotation(t *testing.T) {
	output := StartRotation(getStartingMatrix(), 1)
	fmt.Println(output)
	compareMatrices(oneRotation, output, t)
}

func TestTwoRotation(t *testing.T) {
	output := StartRotation(getStartingMatrix(), 2)
	fmt.Println(output)

	compareMatrices(twoRotation, output, t)
}

func TestThreeRotation(t *testing.T) {
	output := StartRotation(getStartingMatrix(), 3)

	compareMatrices(threeRotation, output, t)
}

func TestNoRotation(t *testing.T) {
	output := StartRotation(getStartingMatrix(), 4)

	compareMatrices(getStartingMatrix(), output, t)
}

func testNegativeNumber(t *testing.T) {
	output := StartRotation(getStartingMatrix(), -3)

	compareMatrices(oneRotation, output, t)
}

func compareMatrices(expected [][]int, actual [][]int, t *testing.T) {
	for i := 0; i < len(actual); i++ {
		for j := 0; j < len(actual[i]); j++ {
			if actual[i][j] != expected[i][j] {
				t.Errorf("[%v][%v]: Expected %v but was %v", i, j, expected[i][j], actual[i][j])
			}
		}
	}
}

func getStartingMatrix() [][]int {
	return [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {1, 2, 3}}
}
