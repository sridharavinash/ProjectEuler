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

func IncrementalSieve(num int) []int {
	arraysize := num - 1
	initarray := make([]int, arraysize)
	for i := 0; i < arraysize; i++ {
		initarray[i] = i + 2
	}

	k := 0
	start := initarray[k]

	for h := 0; h < int(math.Sqrt(float64(num))); h += 1 {
		for i := k; i < arraysize; i += start {
			if start == initarray[i] {
				continue
			}
			initarray[i] = 0
		}
		k += 1
		for initarray[k] == 0 {
			k += 1
		}
		start = initarray[k]
	}
	primes := make([]int, 0)
	for i := 0; i < arraysize; i++ {
		if initarray[i] != 0 {
			primes = append(primes, initarray[i])
		}
	}
	return primes
}
