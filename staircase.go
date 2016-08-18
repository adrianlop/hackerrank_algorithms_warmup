package hackerrank_algorithms_warmup

import (
	"fmt"
	"strings"
)

func main() {
	var size int
	fmt.Scanf("%d", &size)
	for i := 1; i <= size; i++ {
		fmt.Printf("%s%s\n",
			strings.Repeat(" ", size-i),
			strings.Repeat("#", i))
	}
}
