package main

import "fmt"

func main() {
	limit := 100
	sq := func(x int) int { return x * x }
	sumNatNum := 0
	sumSq := 0
	for i := 0; i <= limit; i++ {
		sumSq += sq(i)
		sumNatNum += i
	}
	fmt.Println(sq(sumNatNum) - sumSq)
}
