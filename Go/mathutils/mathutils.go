package mathutils

import "fmt"
import "math"

func isSquare(n float64) bool {
	sqrt := math.Floor(math.Sqrt(n) + 0.5)
	return (sqrt*sqrt)-n == 0
}

func FermatFactor(num int) (float64, float64) {
	if num%2 == 0 {
		fmt.Println("Fermat Factorization only works on odd numbers")
		return -1, -1
	}
	numfloat := float64(num)
	a := math.Ceil((math.Sqrt(numfloat)))
	b := (a * a) - numfloat
	for isq := isSquare(b); !isq; isq = isSquare(b) {
		a += 1
		b = (a * a) - numfloat
	}

	return a - math.Sqrt(b), a + math.Sqrt(b)
}