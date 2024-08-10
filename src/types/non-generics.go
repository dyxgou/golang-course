package types

import "fmt"

func sumInts(m map[string]int) int {
	counter := 0

	for _, value := range m {
		counter += value
	}

	return counter
}

func sumFloats(m map[string]float64) float64 {
	counter := 0.0

	for _, value := range m {
		counter += value
	}

	return counter
}

// if we wanted to  make a new sum function we gotta create a new implementation even thought the fucntioins just change in their type parameters

func NonGenerics() {
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

	n := sumInts(ints)
	f := sumFloats(floats)

	fmt.Printf("ints = %d, floats = %f", n, f)
}
