package main

import "fmt"
import "./mathutils"

func main() {
	index := 10001
	numOfPrimes := 0
	threshold := 1.0
	for numOfPrimes < index {
		numOfPrimes = mathutils.NoOfPrimes(threshold)
		threshold += 1
	}
	fmt.Printf("To find number %d prime finding primes upto %d\n", index, int(threshold))
	primes := mathutils.PrimesSieve(int(threshold))
	fmt.Printf("number %d prime is %d", index, primes[index-1])
}
