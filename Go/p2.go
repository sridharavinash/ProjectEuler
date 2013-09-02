package main

import "fmt"

func fib_even_sum(n int) int {
	prev_sum := 0
	sum := 1
	evensum := 0
	for i := 0; i < n; i++ {
		prev_sum, sum = sum, prev_sum+sum
		if sum > 4000000 {
			break
		}
		if sum%2 == 0 {
			evensum += sum
			fmt.Println(sum)
		}

	}
	return evensum
}

func main() {
	fmt.Printf("Total Even Fib Sum = %d", fib_even_sum(100))
}
