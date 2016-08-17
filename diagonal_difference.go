package hackerrank_algorithms_warmup

import (
	"fmt"
)

func main() {
	var size int
	fmt.Scanf("%d", &size)
	matrix := readMatrix(size)

	diff := diagonalSum(matrix)-secondaryDiagonalSum(matrix)
	if diff < 0 {
		diff = -diff
	}
	fmt.Printf("%d", diff)
}

func readMatrix(size int) ([][]int) {
	matrix := make([][]int, size)
	for x := 0; x < size; x++ {
		matrix[x] = make([]int, size)
		for y := 0; y < size; y++ {
			fmt.Scanf("%d", &matrix[x][y])
		}
	}
	return matrix
}

func diagonalSum(matrix [][]int) int {
	sum := 0
	for i := 0; i < len(matrix); i++ {
		sum += matrix[i][i]
	}
	return sum
}

func secondaryDiagonalSum(matrix [][]int) int {
	sum := 0
	for i := 0; i < len(matrix); i++ {
		sum += matrix[i][len(matrix)-i-1]
	}
	return sum
}