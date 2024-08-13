package arrays

import "fmt"

func Arrays() {
	var ints [3]int

	ints[0] = 1

	fmt.Printf("ints : %v\n", ints)

	// If you have an spare array, you can declare the positions you wanna set

	ints2 := [12]int{
		1, 5: 12, 8: 16,
	}

	fmt.Printf("Ints 2 = %v\n", ints2)

	ints3 := [...]int{10, 12, 123, 124}

	fmt.Printf("ints3 = %v, type = %T\n", ints3, ints3)

	// The built in function `len` takes the len of the array and returns it

	fmt.Printf("ints3 len = %d\n", len(ints3))

	// One of the main things about arrays in go is that go considers arrays of different sizes to be a different type this means that you can't assign an array of some size to another or make a function that works with arrays of all sizes
}
