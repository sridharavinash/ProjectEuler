package main

import "fmt"
import "math"
import "strconv"

func isPalindrome(num int) bool {
	n := strconv.Itoa(num)
	last := len(n) - 1
	for i := 0; i <= len(n)/2; i++ {
		if n[i] != n[last-i] {
			return false
		}
	}
	return true
}

func main() {
	noOfdigits := 3
	max := 0
	count := 0
	start := int(math.Pow10(noOfdigits) - 1)
	end := int(math.Pow10(noOfdigits - 1))
	for i := start; i >= end; i-- {
		if i*start < max {
			continue
		}
		for j := start; j >= end; j-- {
			prod := i * j
			if isPalindrome(prod) {
				count += 1
				fmt.Println(prod)
				if prod > max {
					max = prod
				}
			}
		}
	}
	fmt.Println(max, count)
}
