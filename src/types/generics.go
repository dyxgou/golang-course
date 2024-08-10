package types

import "fmt"

type number interface {
	int | float64 | float32
}

func sumIntsOrFloats[K comparable, V number](m map[K]V) V {
	var counter V

	for _, value := range m {
		counter += value
	}

	return counter
}

func Generics() {
	fmt.Println("Generics")
	ints := map[string]int{
		"number1": 1,
		"number2": 2,
		"number3": 3,
		"number4": 4,
		"number5": 5,
		"number6": 6,
	}

	floats := map[string]float64{
		"float1":  1.0,
		"float2":  2.1,
		"number3": 3.4,
		"number4": 4.4,
		"number5": 5.4,
		"number6": 6.4,
	}

	n := sumIntsOrFloats(ints)
	f := sumIntsOrFloats(floats)

	// Inferred types
	// n := sumIntsOrFloats[string, int](ints)
	// f := sumIntsOrFloats[string, float64](floats)

	fmt.Printf("ints = %d, floats = %f\n", n, f)
}
