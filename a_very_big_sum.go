package hackerrank_algorithms_warmup

import (
"strings"
"strconv"
"bufio"
"os"
"fmt"
)


func main() {
	r := bufio.NewScanner(os.Stdin)
	r.Scan()
	var sum int64 = 0
	for r.Scan() {
		numbers := strings.Split(r.Text(), " ")
		for _, n := range numbers {
			x, _ :=strconv.Atoi(n)
			sum += int64(x)
		}
	}
	fmt.Println(sum)
}

