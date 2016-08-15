package hackerrank_algorithms_warmup

import (
	"fmt"
)

func main() {
	alicePoints := 0
	bobPoints := 0
	aliceTriplet, bobTriplet := ReadPoints()

	for i := 0; i < 3; i++ {
		if aliceTriplet[i] > bobTriplet[i] {
			alicePoints += 1
		} else if bobTriplet[i] > aliceTriplet[i] {
			bobPoints += 1
		}
	}
	fmt.Printf("%d %d", alicePoints, bobPoints)
}

func ReadPoints() ([]int, []int) {
	a := make([]int, 3)
	for i := range a {
		fmt.Scanf("%d", &a[i])
	}
	b := make([]int, 3)
	for i := range b {
		fmt.Scanf("%d", &b[i])
	}
	return a, b
}