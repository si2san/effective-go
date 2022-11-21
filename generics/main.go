package main

import "fmt"

func main() {
	ints := map[string]int64{
		"first":  34,
		"second": 34,
	}

	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	fmt.Printf("Non-Generic Sums: %v, and %v\n", SumInt(ints), SumFloats(floats))
	// go compiler includes type type inference, So can drop type arguments.
	fmt.Printf("Generic Sums:%v, and %v\n", SumIntsOrFloat(ints), SumIntsOrFloat(floats))
}

func SumInt(m map[string]int64) int64 {
	var s int64

	for _, v := range m {
		s += v
	}
	return s
}

func SumFloats(m map[string]float64) float64 {
	var s float64

	for _, v := range m {
		s += v
	}
	return s
}

// comparable constraint is predeclared in Go.
func SumIntsOrFloat[K comparable, V int64 | float64](m map[K]V) V {
	var s V

	for _, v := range m {
		s += v
	}
	return s
}
