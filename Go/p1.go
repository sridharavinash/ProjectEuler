package main

import "fmt"

func multiples_sum(x, y, threshold int) int {
	sum := 0
	for i := x; i < threshold; i++ {
		if i%x == 0 || i%y == 0 {
			sum += i
		}
	}
	return sum
}

func main() {
	fmt.Println(multiples_sum(3, 5, 1000))
}
