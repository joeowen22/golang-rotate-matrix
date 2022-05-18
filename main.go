package main

import "fmt"

func main() {
	fmt.Println("Matrix rotator")
	fmt.Println(StartRotation([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {1, 2, 3}}, 5))

}

func StartRotation(matrix [][]int, rotations int) [][]int {

	switch rotations % 4 {
	case 1:
		return Rotate(matrix, 1)
	case 2:
		return Rotate(matrix, 2)
	case 3:
		return Rotate(matrix, 3)
	default:
		return matrix
	}
}

func Rotate(matrix [][]int, times int) [][]int {
	if times == 0 {
		return matrix
	}
	newMatrix := FlipMatrix(matrix)

	newMatrix = TransposeMatrix(newMatrix)
	return Rotate(newMatrix, times-1)

}

func FlipMatrix(matrix [][]int) [][]int {
	newMatrix := matrix

	for i, j := 0, len(newMatrix)-1; i < j; i, j = i+1, j-1 {
		newMatrix[i], newMatrix[j] = newMatrix[j], newMatrix[i]
	}

	return newMatrix
}

func TransposeMatrix(matrix [][]int) [][]int {
	newMatrix := make([][]int, 0)
	for col := 0; col < len(matrix[0]); col++ {
		newRow := make([]int, 0)
		for row := 0; row < len(matrix); row++ {
			newRow = append(newRow, matrix[row][col])
		}
		newMatrix = append(newMatrix, newRow)
	}
	return newMatrix
}
