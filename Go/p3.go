package main

import "fmt"
import "math"
import "./mathutils"

func main() {
	num := 600851475143
	threshold := int(math.Sqrt(float64(num))) + 1
	primes := mathutils.PrimesSieve(num, threshold)
	max_factor := 0
	for _, v := range primes {
		if num%v == 0 {
			if v > max_factor {
				max_factor = v
			}
		}
	}
	fmt.Println(max_factor)
}
