package main

import "fmt"
import "./mathutils"

func main() {
	inarray := make([]int, 0)
	limit := 20
	for i := 2; i <= limit; i++ {
		inarray = append(inarray, i)
	}
	x := mathutils.Lcm(inarray)
	fmt.Println(x)
}
