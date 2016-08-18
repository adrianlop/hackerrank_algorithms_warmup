package hackerrank_algorithms_warmup

import (
	"fmt"
)

func main() {
	var size, positives, negatives, zeroes int
	fmt.Scanf("%d", &size)
	ints := readArray(size)

	for _, num := range ints {
		if num > 0 {
			positives++
		} else if num < 0 {
			negatives++
		} else {
			zeroes++
		}
	}
	fmt.Printf("%06f\n%06f\n%06f",
		float32(positives)/float32(size),
		float32(negatives)/float32(size),
		float32(zeroes)/float32(size),
	)
}

func readArray(size int) ([]int) {
	array := make([]int, size)
	for x := 0; x < size; x++ {
			fmt.Scanf("%d", &array[x])
	}
	return array
}